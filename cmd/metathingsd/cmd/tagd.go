package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	policy "github.com/nayotta/metathings/pkg/identityd2/policy"
	pb "github.com/nayotta/metathings/pkg/proto/tagd"
	service "github.com/nayotta/metathings/pkg/tagd/service"
)

type TagdOption struct {
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
	Storage                       map[string]interface{}
}

func NewTagdOption() *TagdOption {
	return &TagdOption{
		ServiceBaseOption: cmd_contrib.CreateServiceBaseOption(),
	}
}

var (
	tagd_opt *TagdOption
)

var (
	tagdCmd = &cobra.Command{
		Use:   "tagd",
		Short: "Tag Service Daemon",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if base_opt.Config == "" {
				return
			}

			opt_t := NewTagdOption()
			cmd_helper.UnmarshalConfig(opt_t)
			base_opt = &opt_t.BaseOption

			tagd_opt = opt_t
			tagd_opt.SetServiceName("tagd")
			tagd_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: cmd_helper.Run("tagd", runTagd),
	}
)

func GetTagdOptions() (
	*TagdOption,
	cmd_contrib.ListenOptioner,
	cmd_contrib.TransportCredentialOptioner,
	cmd_contrib.LoggerOptioner,
	cmd_contrib.ServiceEndpointsOptioner,
	cmd_contrib.CredentialOptioner,
) {

	return tagd_opt,
		tagd_opt,
		tagd_opt,
		tagd_opt,
		tagd_opt,
		tagd_opt
}

func runTagd() error {
	app := fx.New(
		fx.Provide(
			GetTagdOptions,
			cmd_contrib.NewTransportCredentials,
			cmd_contrib.NewLogger("tagd"),
			cmd_contrib.NewListener,
			cmd_contrib.NewGrpcServer,
			cmd_contrib.NewClientFactory,
			cmd_contrib.NewTokener,
			token_helper.NewTokenValidator,
			policy.NewEnforcer,
			service.NewMetathingsTagdService,
		),
		fx.Invoke(
			pb.RegisterTagdServiceServer,
		),
	)

	if err := app.Start(context.TODO()); err != nil {
		return err
	}
	defer app.Stop(context.TODO())

	<-app.Done()
	if err := app.Err(); err != nil {
		return err
	}

	return nil
}

func init() {
	tagd_opt = NewTagdOption()

	flags := tagdCmd.Flags()

	flags.StringVarP(tagd_opt.GetListenP(), "listen", "l", "127.0.0.1:5002", "Metathings Tag Service listening address")
	flags.StringVar(tagd_opt.GetCertFileP(), "cert-file", "certs/server.crt", "Metathings Tag Service Credential File")
	flags.StringVar(tagd_opt.GetKeyFileP(), "key-file", "certs/server.key", "Metathings Tag Service Key File")

	RootCmd.AddCommand(tagdCmd)
}
