package cmd

import (
	"context"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "github.com/bigdatagz/metathings/pkg/proto/identity"
	gpb "github.com/golang/protobuf/ptypes/wrappers"
)

var (
	token_issue_opts struct {
		user_id     string
		username    string
		password    string
		domain_id   string
		domain_name string
		project_id  string

		token string

		app_cred_id   string
		app_cred_name string
		secret        string

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
			if err := issueToken(); err != nil {
				log.Fatalf("failed to issue token: %v", err)
			}
		},
	}
)

func parseIssueTokenRequest() (*pb.IssueTokenRequest, error) {
	opts := token_issue_opts
	req := &pb.IssueTokenRequest{}
	if opts.password != "" {
		req.Method = pb.AUTH_METHOD_PASSWORD
		payload := &pb.PasswordPayload{}

		payload.Password = &gpb.StringValue{opts.password}
		if opts.user_id != "" {
			payload.Id = &gpb.StringValue{opts.user_id}
		} else if opts.username != "" {
			payload.Username = &gpb.StringValue{opts.username}
			if opts.domain_id != "" {
				payload.DomainId = &gpb.StringValue{opts.domain_id}
			} else if opts.domain_name != "" {
				payload.DomainName = &gpb.StringValue{opts.domain_name}
			} else {
				return nil, errors.New("required domain id or name when issue token by username")
			}
		}
		if opts.project_id != "" {
			payload.Scope = &pb.TokenScope{
				ProjectId: &gpb.StringValue{opts.project_id},
			}
		}

		req.Payload = &pb.IssueTokenRequest_Password{payload}
	} else if opts.token != "" {
		req.Method = pb.AUTH_METHOD_TOKEN
		payload := &pb.TokenPayload{}

		payload.TokenId = &gpb.StringValue{opts.token}

		req.Payload = &pb.IssueTokenRequest_Token{payload}
	} else if opts.secret != "" {
		req.Method = pb.AUTH_METHOD_APPLICATION_CREDENTIAL
		payload := &pb.ApplicationCredentialPayload{}

		payload.Secret = &gpb.StringValue{opts.secret}
		if opts.app_cred_id != "" {
			payload.Id = &gpb.StringValue{opts.app_cred_id}
		} else if opts.app_cred_name != "" {
			payload.Name = &gpb.StringValue{opts.app_cred_name}
			if opts.domain_id != "" {
				payload.DomainId = &gpb.StringValue{opts.domain_id}
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
	tokenIssueCmd.Flags().StringVar(&token_issue_opts.username, "mt-username", "", "User Name")
	tokenIssueCmd.Flags().StringVar(&token_issue_opts.password, "mt-password", "", "User Password")
	tokenIssueCmd.Flags().StringVar(&token_issue_opts.domain_id, "mt-domain-id", "", "User Domain ID")
	tokenIssueCmd.Flags().StringVar(&token_issue_opts.domain_name, "mt-domain-name", "", "User Domain Name")
	tokenIssueCmd.Flags().StringVar(&token_issue_opts.project_id, "mt-project-id", "", "User Project ID")
	tokenIssueCmd.Flags().StringVar(&token_issue_opts.token, "mt-token", "", "User Token")
	tokenIssueCmd.Flags().StringVar(&token_issue_opts.app_cred_id, "mt-application-credential-id", "", "Application Credential ID")
	tokenIssueCmd.Flags().StringVar(&token_issue_opts.app_cred_name, "mt-application-credential-name", "", "Application Credential Name")
	tokenIssueCmd.Flags().StringVar(&token_issue_opts.secret, "mt-secret", "", "Application Credential Secret")

	tokenIssueCmd.Flags().BoolVar(&token_issue_opts.env, "env", false, "Output as shell script for setup shell environment")

	tokenCmd.AddCommand(tokenIssueCmd)
}
