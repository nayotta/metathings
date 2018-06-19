package cmd

import (
	"context"
	"errors"
	"fmt"

	gpb "github.com/golang/protobuf/ptypes/wrappers"
	"github.com/nayotta/viper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	pb "github.com/nayotta/metathings/pkg/proto/identityd"
)

type _tokenIssueOptions struct {
	_rootOptions              // include application_credential
	UserId                    string
	Username                  string
	Password                  string
	UserDomainId              string
	UserDomainName            string
	DomainId                  string
	ProjectId                 string
	ApplicationCredentialName string
	Env                       bool
}

var (
	token_issue_opts *_tokenIssueOptions
)

var (
	tokenCmd = &cobra.Command{
		Use:   "token",
		Short: "Token Toolkits",
	}

	tokenIssueCmd = &cobra.Command{
		Use:   "issue",
		Short: "Issue Token",
		PreRun: defaultPreRunHooks(func() {
			if root_opts.Config == "" {
				return
			}

			cmd_helper.UnmarshalConfig(token_issue_opts)
			root_opts = &core_agentd_opts._rootOptions
			core_agentd_opts.Stage = cmd_helper.GetStageFromEnv()
		}),
		Run: func(cmd *cobra.Command, args []string) {
			if err := issueToken(); err != nil {
				log.Fatalf("failed to issue token: %v", err)
			}
		},
	}
)

func parseIssueTokenRequest() (*pb.IssueTokenRequest, error) {
	req := &pb.IssueTokenRequest{}
	if token_issue_opts.Password != "" {
		req.Method = pb.AUTH_METHOD_PASSWORD
		payload := &pb.PasswordPayload{}

		payload.Password = &gpb.StringValue{Value: token_issue_opts.Password}
		if token_issue_opts.UserId != "" {
			payload.Id = &gpb.StringValue{Value: token_issue_opts.UserId}
		} else if token_issue_opts.Username != "" {
			payload.Username = &gpb.StringValue{Value: token_issue_opts.Username}
			if token_issue_opts.UserDomainId != "" {
				payload.DomainId = &gpb.StringValue{Value: token_issue_opts.UserDomainId}
			} else if token_issue_opts.UserDomainName != "" {
				payload.DomainName = &gpb.StringValue{Value: token_issue_opts.UserDomainName}
			} else {
				return nil, errors.New("required domain id or name when issue token by username")
			}
		}
		if token_issue_opts.DomainId != "" || token_issue_opts.ProjectId != "" {
			payload.Scope = &pb.TokenScope{}
		}

		if token_issue_opts.DomainId != "" {
			payload.Scope.DomainId = &gpb.StringValue{Value: token_issue_opts.DomainId}
		} else if token_issue_opts.ProjectId != "" {
			payload.Scope.ProjectId = &gpb.StringValue{Value: token_issue_opts.ProjectId}
		}

		req.Payload = &pb.IssueTokenRequest_Password{payload}
	} else if token_issue_opts.Token != "" {
		req.Method = pb.AUTH_METHOD_TOKEN
		payload := &pb.TokenPayload{}

		payload.TokenId = &gpb.StringValue{Value: token_issue_opts.Token}

		req.Payload = &pb.IssueTokenRequest_Token{payload}
	} else if root_opts.ApplicationCredential.Secret != "" {
		req.Method = pb.AUTH_METHOD_APPLICATION_CREDENTIAL
		payload := &pb.ApplicationCredentialPayload{}

		payload.Secret = &gpb.StringValue{Value: root_opts.ApplicationCredential.Secret}
		if root_opts.ApplicationCredential.Id != "" {
			payload.Id = &gpb.StringValue{Value: root_opts.ApplicationCredential.Id}
		} else if token_issue_opts.ApplicationCredentialName != "" {
			payload.Name = &gpb.StringValue{Value: token_issue_opts.ApplicationCredentialName}
			if token_issue_opts.DomainId != "" {
				payload.DomainId = &gpb.StringValue{Value: token_issue_opts.DomainId}
			} else {
				return nil, errors.New("required domain id when issue token by application credential")
			}
		}

		req.Payload = &pb.IssueTokenRequest_ApplicationCredential{payload}
	} else {
		return nil, errors.New("required password or token or secret")
	}
	return req, nil
}

func issueToken() error {
	req, err := parseIssueTokenRequest()
	if err != nil {
		return err
	}

	ctx := context.Background()
	cli, cfn, err := getClientFactory().NewIdentityServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	var header metadata.MD
	res, err := cli.IssueToken(ctx, req, grpc.Header(&header))
	if err != nil {
		return err
	}

	token_str := header["authorization"][0]
	token_str = token_str[3:len(token_str)]
	if token_issue_opts.Env {
		fmt.Printf(`export MT_TOKEN=%v
# Run this command to configure your shell
# eval $(metathings token issue ... --env)
`, token_str)
	} else {
		log.WithFields(log.Fields{
			"user_id": res.Token.User.Id,
			"user":    res.Token.User.Name,
			"token":   token_str,
		}).Info("issue token")
	}

	return nil
}

func init() {
	token_issue_opts = &_tokenIssueOptions{}

	tokenIssueCmd.Flags().StringVar(&token_issue_opts.UserId, "user-id", "", "User ID")
	viper.BindPFlag("user-id", tokenIssueCmd.Flags().Lookup("user-id"))

	tokenIssueCmd.Flags().StringVar(&token_issue_opts.Username, "username", "", "User Name")
	viper.BindPFlag("username", tokenIssueCmd.Flags().Lookup("username"))

	tokenIssueCmd.Flags().StringVar(&token_issue_opts.Password, "password", "", "User Password")
	viper.BindPFlag("password", tokenIssueCmd.Flags().Lookup("password"))

	tokenIssueCmd.Flags().StringVar(&token_issue_opts.UserDomainId, "user-domain-id", "", "User Domain ID")
	viper.BindPFlag("user-domain-id", tokenIssueCmd.Flags().Lookup("user-domain-id"))

	tokenIssueCmd.Flags().StringVar(&token_issue_opts.UserDomainName, "user-domain-name", "", "User Domain Name")
	viper.BindPFlag("user-domain-name", tokenIssueCmd.Flags().Lookup("user-domain-name"))

	tokenIssueCmd.Flags().StringVar(&token_issue_opts.DomainId, "domain-id", "", "Scope Domain ID")
	viper.BindPFlag("domain-id", tokenIssueCmd.Flags().Lookup("domain-id"))

	tokenIssueCmd.Flags().StringVar(&token_issue_opts.ProjectId, "project-id", "", "Scope Project ID")
	viper.BindPFlag("project-id", tokenIssueCmd.Flags().Lookup("project-id"))

	tokenIssueCmd.Flags().StringVar(&token_issue_opts.ApplicationCredentialName, "application-credential-name", "", "Application Credential Name")
	viper.BindPFlag("application-credential-name", tokenIssueCmd.Flags().Lookup("application-credential-name"))

	tokenIssueCmd.Flags().BoolVar(&token_issue_opts.Env, "env", false, "Output as shell script for setup shell environment")

	tokenCmd.AddCommand(tokenIssueCmd)
	RootCmd.AddCommand(tokenCmd)
}
