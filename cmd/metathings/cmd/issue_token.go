package cmd

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	pb "github.com/nayotta/metathings/proto/identityd2"
)

type IssueTokenOption struct {
	cmd_contrib.ClientBaseOption `mapstructure:",squash"`
	DomainId                     string
	Username                     string
	Password                     string
	CredentialId                 string
	CredentialSecret             string
	Timestamp                    int64
	Nonce                        int64
	Token                        string

	Env bool
}

func NewIssueTokenOption() *IssueTokenOption {
	return &IssueTokenOption{
		ClientBaseOption: cmd_contrib.CreateClientBaseOption(),
	}
}

var (
	issue_token_opt *IssueTokenOption
)

var (
	issueTokenCmd = &cobra.Command{
		Use:   "issue",
		Short: "Issue Token",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if base_opt.Config == "" {
				issue_token_opt.BaseOption = *base_opt
				return
			}

			cmd_helper.UnmarshalConfig(issue_token_opt)
			base_opt = &issue_token_opt.BaseOption

			issue_token_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: cmd_helper.Run("issue token", issue_token),
	}
)

func GetIssueTokenOptions() (
	*IssueTokenOption,
	cmd_contrib.ServiceEndpointsOptioner,
	cmd_contrib.LoggerOptioner,
) {
	return issue_token_opt,
		cmd_contrib.NewServiceEndpointsOptionWithTransportCredentialOption(issue_token_opt, issue_token_opt),
		issue_token_opt
}

func issue_token_by_password(opt *IssueTokenOption, cli pb.IdentitydServiceClient) (*pb.Token, error) {
	req := &pb.IssueTokenByPasswordRequest{
		Entity: &pb.OpEntity{
			Name:     &wrappers.StringValue{Value: opt.Username},
			Password: &wrappers.StringValue{Value: opt.Password},
			Domains: []*pb.OpDomain{
				&pb.OpDomain{
					Id: &wrappers.StringValue{Value: opt.DomainId},
				},
			},
		},
	}

	res, err := cli.IssueTokenByPassword(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res.Token, nil
}

func issue_token_by_credential(opt *IssueTokenOption, cli pb.IdentitydServiceClient) (*pb.Token, error) {
	var ts time.Time
	if opt.Timestamp == 0 {
		ts = time.Now()
	} else {
		ts = time.Unix(0, opt.Timestamp)
	}
	pb_ts := pb_helper.FromTime(ts)

	if opt.Nonce == 0 {
		opt.Nonce = rand.Int63()
	}

	hmac := passwd_helper.MustParseHmac(opt.CredentialSecret, opt.CredentialId, ts, opt.Nonce)

	req := &pb.IssueTokenByCredentialRequest{
		Credential: &pb.OpCredential{
			Id: &wrappers.StringValue{Value: opt.CredentialId},
		},
		Timestamp: &pb_ts,
		Nonce:     &wrappers.Int64Value{Value: opt.Nonce},
		Hmac:      &wrappers.StringValue{Value: hmac},
	}

	res, err := cli.IssueTokenByCredential(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res.Token, err
}

func issue_token_by_token(opt *IssueTokenOption, cli pb.IdentitydServiceClient) (*pb.Token, error) {
	req := &pb.IssueTokenByTokenRequest{
		Token: &pb.OpToken{
			Text: &wrappers.StringValue{Value: opt.Token},
		},
	}

	res, err := cli.IssueTokenByToken(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res.Token, err
}

func issue_token() error {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			GetIssueTokenOptions,
			cmd_contrib.NewLogger("issue_token"),
			cmd_contrib.NewClientFactory,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, opt *IssueTokenOption, cli_fty *client_helper.ClientFactory) {
				lc.Append(fx.Hook{
					OnStart: func(context.Context) error {
						var token *pb.Token

						cli, cfn, err := cli_fty.NewIdentityd2ServiceClient()
						if err != nil {
							return err
						}
						defer cfn()

						if opt.Username != "" && opt.Password != "" {
							token, err = issue_token_by_password(opt, cli)
						} else if opt.CredentialId != "" && opt.CredentialSecret != "" {
							token, err = issue_token_by_credential(opt, cli)
						} else if opt.Token != "" {
							token, err = issue_token_by_token(opt, cli)
						}
						if err != nil {
							return err
						}

						tkn_txt_str := token.Text
						if opt.Env {
							fmt.Printf(`export MT_TOKEN=%v
# Run this command to configure your shell
# eval $(metathings token issue ...)
`, tkn_txt_str)
						} else {
							log.WithFields(log.Fields{
								"token": tkn_txt_str,
							}).Info("issue token")
						}

						return nil
					},
				})
			},
		),
	)

	if err := app.Start(context.Background()); err != nil {
		return err
	}
	if err := app.Err(); err != nil {
		return err
	}

	return nil
}

func init() {
	issue_token_opt = NewIssueTokenOption()

	flags := issueTokenCmd.Flags()

	flags.StringVar(&issue_token_opt.DomainId, "domain-id", "", "Domain ID")
	flags.StringVar(&issue_token_opt.Username, "username", "", "Username for issue token by password mode")
	flags.StringVar(&issue_token_opt.Password, "password", "", "Password for issue token by password mode")
	flags.StringVar(&issue_token_opt.Token, "token", "", "Token for issue token by token mode")
	flags.StringVar(&issue_token_opt.CredentialId, "credential-id", "", "Credential ID for issue token by credential mode")
	flags.StringVar(&issue_token_opt.CredentialSecret, "credential-secret", "", "Credential Secret for issue token by credential mode")
	flags.Int64Var(&issue_token_opt.Timestamp, "timestamp", 0, "Timestamp for issue token by credential mode")
	flags.Int64Var(&issue_token_opt.Nonce, "nonce", 0, "Nonce for issue token by credential mode")
	flags.BoolVar(&issue_token_opt.Env, "env", false, "Output as shell script for setup shell environment")

	tokenCmd.AddCommand(issueTokenCmd)
}
