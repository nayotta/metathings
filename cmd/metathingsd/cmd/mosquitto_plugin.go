package cmd

import (
	"context"
	"net/http"

	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	service "github.com/nayotta/metathings/pkg/plugin/mosquitto/service"
)

type MosquittoPluginOption struct {
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
}

func NewMosquittoPluginOption() *MosquittoPluginOption {
	return &MosquittoPluginOption{
		ServiceBaseOption: cmd_contrib.CreateServiceBaseOption(),
	}
}

var (
	mosquitto_plugin_opt *MosquittoPluginOption
)

var (
	mosquittoPluginCmd = &cobra.Command{
		Use:   "mosquitto",
		Short: "Metathings Service Mosquitto Plugin",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if base_opt.Config == "" {
				return
			}

			opt_t := NewMosquittoPluginOption()
			cmd_helper.UnmarshalConfig(opt_t)
			base_opt = &opt_t.BaseOption

			mosquitto_plugin_opt = opt_t
			mosquitto_plugin_opt.SetServiceName("mosquitto-plugin")
			mosquitto_plugin_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: cmd_helper.Run("mosquitto-plugin", runMosquittoPlugin),
	}
)

func GetMosquittoPluginOptions() (
	*MosquittoPluginOption,
	cmd_contrib.ListenOptioner,
	cmd_contrib.LoggerOptioner,
	cmd_contrib.ServiceEndpointsOptioner,
) {
	return mosquitto_plugin_opt,
		mosquitto_plugin_opt,
		mosquitto_plugin_opt,
		mosquitto_plugin_opt
}

func runMosquittoPlugin() error {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			GetMosquittoPluginOptions,
			cmd_contrib.NewLogger("mosquitto-plugin"),
			cmd_contrib.NewHttpServer,
			cmd_contrib.NewClientFactory,
			service.NewMosquittoPluginService,
		),
		fx.Invoke(
			func(s *http.Server, srv *service.MosquittoPluginService) {
				mux := http.NewServeMux()
				mux.HandleFunc("/webhook", srv.WebhookHandler)
				s.Handler = mux
			},
		),
	)

	if err := app.Start(context.Background()); err != nil {
		return err
	}
	defer app.Stop(context.Background())

	<-app.Done()
	if err := app.Err(); err != nil {
		return err
	}

	return nil
}
