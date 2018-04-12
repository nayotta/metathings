package cmd

import (
	"context"
	"errors"
	"fmt"

	gpb "github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "github.com/bigdatagz/metathings/pkg/proto/identity"
)

var (
	token_issue_opts struct {
		user_id     string
		username    string
		password    string
		domain_id   string
		domain_name string
		project_id  string

		application_credential_id   string
		application_credential_name string
		secret                      string

		env bool
	}
)

var (
	tokenCmd = &cobra.Command{
		Use:   "token",
		Short: "Token Toolkits",
	}

	tokenIssueCmd = &cobra.Command{
		Use:   "issue",
		Short: "Issue Token",
		Run: func(cmd *cobra.Command, args []string) {
			initialize()
			if err := issueToken(); err != nil {
				log.Fatalf("failed to issue token: %v", err)
			}
		},
	}
)

func parseIssueTokenRequest() (*pb.IssueTokenRequest, error) {
	req := &pb.IssueTokenRequest{}
	if V("password") != "" {
		req.Method = pb.AUTH_METHOD_PASSWORD
		payload := &pb.PasswordPayload{}

		payload.Password = &gpb.StringValue{V("password")}
		if V("user_id") != "" {
			payload.Id = &gpb.StringValue{V("user_id")}
		} else if V("username") != "" {
			payload.Username = &gpb.StringValue{V("username")}
			if V("domain_id") != "" {
				payload.DomainId = &gpb.StringValue{V("domain_id")}
			} else if V("domain_name") != "" {
				payload.DomainName = &gpb.StringValue{V("domain_name")}
			} else {
				return nil, errors.New("required domain id or name when issue token by username")
			}
		}
		if V("project_id") != "" {
			payload.Scope = &pb.TokenScope{
				ProjectId: &gpb.StringValue{V("project_id")},
			}
		}

		req.Payload = &pb.IssueTokenRequest_Password{payload}
	} else if V("token") != "" {
		req.Method = pb.AUTH_METHOD_TOKEN
		payload := &pb.TokenPayload{}

		payload.TokenId = &gpb.StringValue{V("token")}

		req.Payload = &pb.IssueTokenRequest_Token{payload}
	} else if V("secret") != "" {
		req.Method = pb.AUTH_METHOD_APPLICATION_CREDENTIAL
		payload := &pb.ApplicationCredentialPayload{}

		payload.Secret = &gpb.StringValue{V("secret")}
		if V("application_credential_id") != "" {
			payload.Id = &gpb.StringValue{V("application_credential_id")}
		} else if V("application_credential_name") != "" {
			payload.Name = &gpb.StringValue{V("application_credential_name")}
			if V("domain_id") != "" {
				payload.DomainId = &gpb.StringValue{V("domain_id")}
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
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(root_opts.addr, opts...)
	if err != nil {
		return err
	}
	defer conn.Close()
	cli := pb.NewIdentityServiceClient(conn)

	var header metadata.MD
	res, err := cli.IssueToken(ctx, req, grpc.Header(&header))
	if err != nil {
		return err
	}

	token_str := header["x-subject-token"][0]
	if token_issue_opts.env {
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
	tokenIssueCmd.Flags().StringVar(&token_issue_opts.user_id, "mt-user-id", "", "User ID")
	viper.BindPFlag(A("USER_ID"), tokenIssueCmd.Flags().Lookup("mt-user-id"))
	tokenIssueCmd.Flags().StringVar(&token_issue_opts.username, "mt-username", "", "User Name")
	viper.BindPFlag(A("USERNAME"), tokenIssueCmd.Flags().Lookup("mt-username"))
	tokenIssueCmd.Flags().StringVar(&token_issue_opts.password, "mt-password", "", "User Password")
	viper.BindPFlag(A("PASSWORD"), tokenIssueCmd.Flags().Lookup("mt-password"))
	tokenIssueCmd.Flags().StringVar(&token_issue_opts.domain_id, "mt-domain-id", "", "User Domain ID")
	viper.BindPFlag(A("DOMAIN_ID"), tokenIssueCmd.Flags().Lookup("mt-domain-id"))
	tokenIssueCmd.Flags().StringVar(&token_issue_opts.domain_name, "mt-domain-name", "", "User Domain Name")
	viper.BindPFlag(A("DOMAIN_NAME"), tokenIssueCmd.Flags().Lookup("mt-domain-name"))
	tokenIssueCmd.Flags().StringVar(&token_issue_opts.project_id, "mt-project-id", "", "User Project ID")
	viper.BindPFlag(A("PROJECT_ID"), tokenIssueCmd.Flags().Lookup("mt-project-id"))
	tokenIssueCmd.Flags().StringVar(&token_issue_opts.application_credential_id, "mt-application-credential-id", "", "Application Credential ID")
	viper.BindPFlag(A("APPLICATION_CREDENTIAL_ID"), tokenIssueCmd.Flags().Lookup("mt-application-credential-id"))
	tokenIssueCmd.Flags().StringVar(&token_issue_opts.application_credential_name, "mt-application-credential-name", "", "Application Credential Name")
	viper.BindPFlag(A("APPLICATION_CREDENTIAL_NAME"), tokenIssueCmd.Flags().Lookup("mt-application-credential-name"))
	tokenIssueCmd.Flags().StringVar(&token_issue_opts.secret, "mt-secret", "", "Application Credential Secret")
	viper.BindPFlag(A("SECRET"), tokenIssueCmd.Flags().Lookup("mt-secret"))
	tokenIssueCmd.Flags().BoolVar(&token_issue_opts.env, "env", false, "Output as shell script for setup shell environment")

	tokenCmd.AddCommand(tokenIssueCmd)
	RootCmd.AddCommand(tokenCmd)
}
