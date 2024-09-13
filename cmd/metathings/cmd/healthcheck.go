package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"google.golang.org/protobuf/types/known/emptypb"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
)

type HealthcheckOption struct {
	cmd_contrib.ClientBaseOption `mapstructure:",squash"`
	Policyd                      int
	Identityd2                   int
	Tagd                         int
	Deviced                      int
}

func NewHealthcheckOption() *HealthcheckOption {
	return &HealthcheckOption{
		ClientBaseOption: cmd_contrib.CreateClientBaseOption(),
	}
}

var (
	healthcheck_opt *HealthcheckOption

	healthcheckCmd = &cobra.Command{
		Use:   "healthcheck",
		Short: "Service health check",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if base_opt.Config == "" {
				healthcheck_opt.BaseOption = *base_opt
				return
			}

			cmd_helper.UnmarshalConfig(healthcheck_opt)
			base_opt = &healthcheck_opt.BaseOption

			healthcheck_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: cmd_helper.Run("healthcheck", run_healthcheck),
	}
)

func GetHealthcheckOptions() (
	*HealthcheckOption,
	cmd_contrib.ServiceEndpointsOptioner,
	cmd_contrib.LoggerOptioner,
) {
	return healthcheck_opt,
		cmd_contrib.NewServiceEndpointsOptionWithTransportCredentialOption(healthcheck_opt, healthcheck_opt),
		healthcheck_opt
}

func run_healthcheck() error {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			GetHealthcheckOptions,
			cmd_contrib.NewLogger("healthcheck"),
			cmd_contrib.NewClientFactory,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, opt *HealthcheckOption, cli_fty *client_helper.ClientFactory) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						if opt.Policyd > 0 {
							cli, cfn, err := cli_fty.NewPolicydServiceClient()
							if err != nil {
								return err
							}
							defer cfn()
							_, err = cli.Healthz(ctx, &emptypb.Empty{})
							if err != nil {
								return err
							}
						}

						if opt.Identityd2 > 0 {
							cli, cfn, err := cli_fty.NewIdentityd2ServiceClient()
							if err != nil {
								return err
							}
							defer cfn()
							_, err = cli.Healthz(ctx, &emptypb.Empty{})
							if err != nil {
								return err
							}
						}

						if opt.Tagd > 0 {
							cli, cfn, err := cli_fty.NewTagdServiceClient()
							if err != nil {
								return err
							}
							defer cfn()
							_, err = cli.Healthz(ctx, &emptypb.Empty{})
							if err != nil {
								return err
							}
						}

						if opt.Deviced > 0 {
							cli, cfn, err := cli_fty.NewDevicedServiceClient()
							if err != nil {
								return err
							}
							defer cfn()
							_, err = cli.Healthz(ctx, &emptypb.Empty{})
							if err != nil {
								return err
							}
						}

						fmt.Println("OK")
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
	healthcheck_opt = NewHealthcheckOption()

	flags := healthcheckCmd.Flags()

	flags.CountVar(&healthcheck_opt.Policyd, "policyd", "health check policyd")
	flags.CountVar(&healthcheck_opt.Identityd2, "identityd2", "health check identityd2")
	flags.CountVar(&healthcheck_opt.Tagd, "tagd", "health check tagd")
	flags.CountVar(&healthcheck_opt.Deviced, "deviced", "health check deviced")

	RootCmd.AddCommand(healthcheckCmd)
}
