package cmd

import (
	"context"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	fx_helper "github.com/nayotta/metathings/pkg/common/fx"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	version_helper "github.com/nayotta/metathings/pkg/common/version"
	service "github.com/nayotta/metathings/pkg/device/service"
	pb "github.com/nayotta/metathings/pkg/proto/device"
)

type RunDeviceOption struct {
	cmd_contrib.ClientBaseOption `mapstructure:",squash"`
	BinarySynchronizer           map[string]interface{}
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

			cmd_helper.InitManyStringMapFromConfigWithStage([]cmd_helper.InitManyOption{
				{Dst: &run_device_opt.BinarySynchronizer, Key: "binary_synchronizer"},
			})

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
	cmd_contrib.ListenOptioner,
	cmd_contrib.TransportCredentialOptioner,
	cmd_contrib.BinarySynchronizerOption,
) {
	return run_device_opt,
		run_device_opt,
		run_device_opt,
		run_device_opt,
		run_device_opt,
		run_device_opt,
		cmd_contrib.BinarySynchronizerOption{
			Option: run_device_opt.BinarySynchronizer,
		}
}

func NewMetathingsDeviceServiceOption(opt *RunDeviceOption) *service.MetathingsDeviceServiceOption {
	return &service.MetathingsDeviceServiceOption{
		ModuleAliveTimeout:   83 * time.Second,
		HeartbeatInterval:    17 * time.Second,
		HeartbeatMaxRetry:    3,
		MinReconnectInterval: 7 * time.Second,
		MaxReconnectInterval: 137 * time.Second,
		PingInterval:         29 * time.Second,
	}
}

func run_device() error {
	var err error
	var app *fx.App

	app = fx.New(
		fx.NopLogger,
		fx.Provide(
			fx_helper.NewFxAppGetter(&app),
			GetRunDeviceOptions,
			cmd_contrib.NewLogger("device"),
			cmd_contrib.NewServerTransportCredentials,
			cmd_contrib.NewClientFactory,
			cmd_contrib.NewNoExpireTokener,
			cmd_contrib.NewBinarySynchronizer,
			version_helper.NewVersioner(version_str),
			token_helper.NewTokenValidator,
			NewMetathingsDeviceServiceOption,
			service.NewMetathingsDeviceService,
			func(x service.MetathingsDeviceService) pb.DeviceServiceServer {
				return x
			},
			cmd_contrib.NewListener,
			cmd_contrib.NewGrpcServer,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, srv service.MetathingsDeviceService, logger log.FieldLogger) error {
				lc.Append(fx.Hook{
					OnStart: func(context.Context) error {
						if err = srv.Start(); err != nil {
							return err
						}
						return nil
					},
					OnStop: func(context.Context) error {
						defer os.Exit(1)
						logger.Infof("receive exit signal")
						return nil
					},
				})

				return nil
			},
			pb.RegisterDeviceServiceServer,
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
