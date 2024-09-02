package cmd

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/PeerXu/option-go"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	cfg_helper "github.com/nayotta/metathings/pkg/common/config"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	redis_helper "github.com/nayotta/metathings/pkg/common/redis"
	service "github.com/nayotta/metathings/pkg/plugin/vernemq/service"
	storage "github.com/nayotta/metathings/pkg/plugin/vernemq/storage"
)

type VernemqPluginOption struct {
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
	VernemqStorage                map[string]any
	Webhook                       cmd_contrib.WebhookOption
}

func NewVernemqPluginOption() *VernemqPluginOption {
	return &VernemqPluginOption{
		ServiceBaseOption: cmd_contrib.CreateServiceBaseOption(),
	}
}

var (
	vernemq_plugin_opt *VernemqPluginOption
)

func init_vernemq_storage(opt *VernemqPluginOption) {
	mvs := make(map[string]any)
	vvs := cmd_helper.GetFromStage().Sub("vernemq_storage")
	for _, key := range vvs.AllKeys() {
		mvs[key] = vvs.Get(key)
	}
	opt.VernemqStorage = mvs
}

var (
	vernemqPluginCmd = &cobra.Command{
		Use:   "vernemq",
		Short: "Metathings Service Vernemq Plugin",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if base_opt.Config == "" {
				return
			}

			opt_t := NewVernemqPluginOption()
			cmd_helper.UnmarshalConfig(opt_t)
			base_opt = &opt_t.BaseOption

			init_vernemq_storage(opt_t)

			vernemq_plugin_opt = opt_t
			vernemq_plugin_opt.SetServiceName("vernemq-plugin")
			vernemq_plugin_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: cmd_helper.Run("vernemq-plugin", runVernemqPlugin),
	}
)

func GetVernemqPluginOptions() (
	*VernemqPluginOption,
	cmd_contrib.ServiceOptioner,
	cmd_contrib.ListenOptioner,
	cmd_contrib.LoggerOptioner,
	cmd_contrib.ServiceEndpointsOptioner,
	cmd_contrib.OpentracingOptioner,
) {

	return vernemq_plugin_opt,
		vernemq_plugin_opt,
		vernemq_plugin_opt,
		vernemq_plugin_opt,
		vernemq_plugin_opt,
		vernemq_plugin_opt
}

func NewVernemqPluginStorage(opt *VernemqPluginOption, logger logrus.FieldLogger) (storage.Storage, error) {
	drv, args, err := cfg_helper.ParseConfigOption("driver", opt.VernemqStorage)
	if err != nil {
		return nil, err
	}

	var opts []option.ApplyOption
	switch drv {
	case "redis":
		var host, username, password string
		var port, db int
		if err = opt_helper.Setopt(map[string]func(string, any) error{
			"host":     opt_helper.ToString(&host),
			"port":     opt_helper.ToInt(&port),
			"username": opt_helper.ToString(&username),
			"password": opt_helper.ToString(&password),
			"db":       opt_helper.ToInt(&db),
		})(args...); err != nil {
			return nil, err
		}
		addr := net.JoinHostPort(host, cast.ToString(port))
		client := redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:    []string{addr},
			Username: username,
			Password: password,
			DB:       db,
		})
		opts = append(opts, redis_helper.WithRedisClient(client), log_helper.WithLogger(logger))
	default:
		return nil, ErrUnsupportedStorageDriverFn(drv)
	}

	return storage.New(drv, opts...)
}

func NewVernemqPluginService(opt *VernemqPluginOption, logger logrus.FieldLogger, s storage.Storage) (*service.VernemqPluginService, error) {
	return service.NewVernemqPluginService(
		service.WithWebhookSecret(opt.Webhook.Secret),
		log_helper.WithLogger(logger),
		storage.WithStorage(s),
	)
}

type NewVernemqPluginRoutingParams struct {
	fx.In

	Router  *mux.Router
	Service *service.VernemqPluginService
	Tracer  opentracing.Tracer `name:"opentracing_tracer" option:"true"`
}

func NewVernemqPluginRouting(p NewVernemqPluginRoutingParams) error {
	p.Router.HandleFunc("/webhook", p.Service.WebhookHandler)

	return nil
}

func runVernemqPlugin() error {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			GetVernemqPluginOptions,
			mux.NewRouter,
			cmd_contrib.NewLogger("vernemq-plugin"),
			cmd_contrib.NewListener,
			cmd_contrib.NewOpentracing,
			NewVernemqPluginStorage,
			NewVernemqPluginService,
		),
		fx.Invoke(
			NewVernemqPluginRouting,
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
	vernemq_plugin_opt = NewVernemqPluginOption()

	flags := vernemqPluginCmd.Flags()

	flags.StringVarP(vernemq_plugin_opt.GetListenP(), "listen", "l", "127.0.0.1:21884", "Metathings VerneMQ Plugin Service listening address")

	pluginCmd.AddCommand(vernemqPluginCmd)
}
