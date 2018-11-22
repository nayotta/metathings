package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	service "github.com/nayotta/metathings/pkg/device/service"
)

type RunDeviceOption struct {
	cmd_contrib.ClientBaseOption `mapstructure:",squash"`
}

func NewRunDeviceOption() *RunDeviceOption {
	return &RunDeviceOption{
		ClientBaseOption: cmd_contrib.CreateClientBaseOption(),
	}
}

var (
	run_device_opt *RunDeviceOption
)

var (
	runDeviceCmd = &cobra.Command{
		Use:   "run",
		Short: "Run Device Service Daemon",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if base_opt.Config == "" {
				run_device_opt.BaseOption = *base_opt
				return
			}

			cmd_helper.UnmarshalConfig(run_device_opt)
			base_opt = &run_device_opt.BaseOption

			run_device_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: cmd_helper.Run("run device", run_device),
	}
)

func GetRunDeviceOptions() (
	*RunDeviceOption,
	cmd_contrib.LoggerOptioner,
	cmd_contrib.ServiceEndpointsOptioner,
	cmd_contrib.CredentialOptioner,
) {
	return run_device_opt,
		run_device_opt,
		run_device_opt,
		run_device_opt
}

func NewMetathingsDeviceServiceOption(opt *RunDeviceOption) *service.MetathingsDeviceServiceOption {
	return &service.MetathingsDeviceServiceOption{}
}

func run_device() error {
	var err error

	app := fx.New(
		fx.Provide(
			GetRunDeviceOptions,
			cmd_contrib.NewLogger("device"),
			cmd_contrib.NewClientFactory,
			cmd_contrib.NewTokener,
			NewMetathingsDeviceServiceOption,
			service.NewMetathingsDeviceService,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, srv service.MetathingsDeviceService) error {
				lc.Append(fx.Hook{
					OnStart: func(context.Context) error {
						if err = srv.Start(); err != nil {
							return err
						}
						return nil
					},
				})

				return nil
			},
		),
	)

	if err = app.Start(context.Background()); err != nil {
		return err
	}

	<-app.Done()
	if err = app.Err(); err != nil {
		return err
	}

	return nil
}

func init() {
	run_device_opt = NewRunDeviceOption()

	deviceCmd.AddCommand(runDeviceCmd)
}
