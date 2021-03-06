package cmd

import (
	"context"
	"encoding/base64"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	cfg_helper "github.com/nayotta/metathings/pkg/common/config"
	service "github.com/nayotta/metathings/pkg/plugin/mosquitto/service"
	storage "github.com/nayotta/metathings/pkg/plugin/mosquitto/storage"
)

type MosquittoPluginOption struct {
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
	MosquittoStorage              map[string]interface{}
	Webhook                       cmd_contrib.WebhookOption
}

func NewMosquittoPluginOption() *MosquittoPluginOption {
	return &MosquittoPluginOption{
		ServiceBaseOption: cmd_contrib.CreateServiceBaseOption(),
	}
}

var (
	mosquitto_plugin_opt *MosquittoPluginOption
)

func init_mosquitto_storage(opt *MosquittoPluginOption) {
	mms := map[string]interface{}{}
	vms := cmd_helper.GetFromStage().Sub("mosquitto_storage")
	for _, key := range vms.AllKeys() {
		mms[key] = vms.Get(key)
	}
	opt.MosquittoStorage = mms
}

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

			init_mosquitto_storage(opt_t)

			mosquitto_plugin_opt = opt_t
			mosquitto_plugin_opt.SetServiceName("mosquitto-plugin")
			mosquitto_plugin_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: cmd_helper.Run("mosquitto-plugin", runMosquittoPlugin),
	}
)

func GetMosquittoPluginOptions() (
	*MosquittoPluginOption,
	cmd_contrib.ServiceOptioner,
	cmd_contrib.ListenOptioner,
	cmd_contrib.LoggerOptioner,
	cmd_contrib.ServiceEndpointsOptioner,
	cmd_contrib.OpentracingOptioner,
) {
	return mosquitto_plugin_opt,
		mosquitto_plugin_opt,
		mosquitto_plugin_opt,
		mosquitto_plugin_opt,
		mosquitto_plugin_opt,
		mosquitto_plugin_opt
}

func NewMosquittoPluginStorage(
	opt *MosquittoPluginOption,
	logger log.FieldLogger,
) (storage.Storage, error) {
	drv, args, err := cfg_helper.ParseConfigOption("driver", opt.MosquittoStorage, "logger", logger)
	if err != nil {
		return nil, err
	}

	return storage.NewStorage(drv, args...)
}

func NewMosquittoPluginServiceOption(opt *MosquittoPluginOption) *service.MosquittoPluginServiceOption {
	o := &service.MosquittoPluginServiceOption{}
	o.Webhook.Secret = base64.StdEncoding.EncodeToString([]byte(opt.Webhook.GetSecret()))

	return o
}

type NewMosquittoPluginRoutingParams struct {
	fx.In

	Router  *mux.Router
	Service *service.MosquittoPluginService
	Tracer  opentracing.Tracer `name:"opentracing_tracer" optional:"true"`
}

func NewMosquittoPluginRouting(p NewMosquittoPluginRoutingParams) error {
	p.Router.HandleFunc("/webhook", p.Service.WebhookHandler)

	return nil
}

func runMosquittoPlugin() error {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			GetMosquittoPluginOptions,
			mux.NewRouter,
			cmd_contrib.NewLogger("mosquitto-plugin"),
			cmd_contrib.NewListener,
			cmd_contrib.NewOpentracing,
			NewMosquittoPluginStorage,
			NewMosquittoPluginServiceOption,
			service.NewMosquittoPluginService,
		),
		fx.Invoke(
			NewMosquittoPluginRouting,
			cmd_contrib.NewHttpServer,
		),
	)

	if err := app.Start(context.Background()); err != nil {
		return err
	}
	go func() {
		defer app.Stop(context.Background())

		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM)
		<-ch
	}()

	<-app.Done()
	if err := app.Err(); err != nil {
		return err
	}

	return nil
}

func init() {
	mosquitto_plugin_opt = NewMosquittoPluginOption()

	flags := mosquittoPluginCmd.Flags()

	flags.StringVarP(mosquitto_plugin_opt.GetListenP(), "listen", "l", "127.0.0.1:21883", "Metathings Mosquitto Plugin Service listening address")

	pluginCmd.AddCommand(mosquittoPluginCmd)
}
