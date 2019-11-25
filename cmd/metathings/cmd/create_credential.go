package cmd

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

type CreateCredentialOption struct {
	cmd_contrib.ClientBaseOption `mapstructure:",squash"`
	Name                         string
	Token                        string
	Domain                       string
}

func NewCreateCredentialOption() *CreateCredentialOption {
	return &CreateCredentialOption{
		ClientBaseOption: cmd_contrib.CreateClientBaseOption(),
	}
}

var (
	create_credential_opt *CreateCredentialOption
)

var (
	createCredentialCmd = &cobra.Command{
		Use:   "create",
		Short: "Create Credential",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if base_opt.Config == "" {
				create_credential_opt.BaseOption = *base_opt
			} else {
				cmd_helper.UnmarshalConfig(create_credential_opt)
				base_opt = &create_credential_opt.BaseOption
			}
			if create_credential_opt.Token == "" {
				create_credential_opt.Token = cmd_helper.GetTokenFromEnv()
			}
		}),
		Run: cmd_helper.Run("create credential", create_credential),
	}
)

func GetCreateCredentialOptions() (
	*CreateCredentialOption,
	cmd_contrib.ServiceEndpointsOptioner,
	cmd_contrib.TransportCredentialOptioner,
	cmd_contrib.LoggerOptioner,
) {
	return create_credential_opt,
		create_credential_opt,
		create_credential_opt,
		create_credential_opt
}

func create_credential() error {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			GetCreateCredentialOptions,
			cmd_contrib.NewLogger("create_credential"),
			cmd_contrib.NewClientTransportCredentials,
			cmd_contrib.NewClientFactory,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, opt *CreateCredentialOption, cli_fty *client_helper.ClientFactory) {
				lc.Append(fx.Hook{
					OnStart: func(context.Context) error {
						cli, cfn, err := cli_fty.NewIdentityd2ServiceClient()
						if err != nil {
							return err
						}
						defer cfn()

						ctx := context_helper.WithToken(context.TODO(), create_credential_opt.Token)
						show_entity_res, err := cli.ShowEntity(ctx, &empty.Empty{})
						if err != nil {
							return err
						}
						ent := show_entity_res.GetEntity()
						cred_name := create_credential_opt.Name
						ent_id := ent.GetId()
						cred_dom_id := create_credential_opt.Domain
						if cred_dom_id == "" {
							doms := ent.GetDomains()
							if len(doms) == 0 {
								return errors.New("please bind entity and domain first")
							}
							cred_dom_id = doms[0].GetId()
						}

						create_credential_req := &pb.CreateCredentialRequest{
							Credential: &pb.OpCredential{
								Name: &wrappers.StringValue{Value: cred_name},
								Domain: &pb.OpDomain{
									Id: &wrappers.StringValue{Value: cred_dom_id},
								},
								Entity: &pb.OpEntity{
									Id: &wrappers.StringValue{Value: ent_id},
								},
							},
						}
						create_credential_res, err := cli.CreateCredential(ctx, create_credential_req)
						if err != nil {
							return err
						}

						cred := create_credential_res.GetCredential()
						data := map[string]interface{}{
							"credential": map[string]interface{}{
								"domain": cred_dom_id,
								"id":     cred.GetId(),
								"secret": cred.GetSecret(),
							},
						}

						if err = cmd_contrib.ProcessOutput(opt, data); err != nil {
							return err
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
	create_credential_opt = NewCreateCredentialOption()

	flags := createCredentialCmd.Flags()

	flags.StringVar(&create_credential_opt.Name, "name", "", "Credential Name")
	flags.StringVar(&create_credential_opt.Token, "token", "", "Token")
	flags.StringVar(&create_credential_opt.Domain, "domain", "", "Credential Domain")
	flags.StringVarP(&create_credential_opt.Output, "output", "o", "shell", "Output Format")

	credentialCmd.AddCommand(createCredentialCmd)
}
