package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	cfg_helper "github.com/nayotta/metathings/pkg/common/config"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	service "github.com/nayotta/metathings/pkg/device_cloud/service"
	storage "github.com/nayotta/metathings/pkg/device_cloud/storage"
)

type DeviceCloudOption struct {
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
	Storage                       map[string]interface{}
	Connection                    map[string]interface{}
}

func NewDeviceCloudOption() *DeviceCloudOption {
	return &DeviceCloudOption{
		ServiceBaseOption: cmd_contrib.CreateServiceBaseOption(),
	}
}

var (
	device_cloud_opt *DeviceCloudOption
)

func init_device_cloud_storage(opt *DeviceCloudOption) {
	ms := map[string]interface{}{}
	vs := cmd_helper.GetFromStage().Sub("storage")
	for _, key := range vs.AllKeys() {
		switch key {
		case "driver":
			ms[key] = vs.GetString(key)
		case "db":
			ms[key] = vs.GetInt(key)
		default:
			ms[key] = vs.Get(key)
		}
	}
	opt.Storage = ms
}

var (
	deviceCloudCmd = &cobra.Command{
		Use:   "device_cloud",
		Short: "Device Cloud Service Daemon",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if base_opt.Config == "" {
				return
			}

			opt_t := NewDeviceCloudOption()
			cmd_helper.UnmarshalConfig(opt_t)
			base_opt = &opt_t.BaseOption

			init_device_cloud_storage(opt_t)

			device_cloud_opt = opt_t
			device_cloud_opt.SetServiceName("device_cloud")
			device_cloud_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: cmd_helper.Run("device_cloud", runDeviceCloud),
	}
)

func GetDeviceCloudOptions() (
	*DeviceCloudOption,
	cmd_contrib.ServiceOptioner,
	cmd_contrib.ListenOptioner,
	cmd_contrib.LoggerOptioner,
	cmd_contrib.ServiceEndpointsOptioner,
	cmd_contrib.CredentialOptioner,
	cmd_contrib.TransportCredentialOptioner,
	cmd_contrib.OpentracingOptioner,
) {
	return device_cloud_opt,
		device_cloud_opt,
		device_cloud_opt,
		device_cloud_opt,
		device_cloud_opt,
		device_cloud_opt,
		device_cloud_opt,
		device_cloud_opt
}

func NewDeviceCloudStorage(opt *DeviceCloudOption, logger log.FieldLogger) (storage.Storage, error) {
	var drv string
	var args []interface{}
	var err error

	if drv, args, err = cfg_helper.ParseConfigOption("driver", opt.Storage, "logger", logger); err != nil {
		return nil, err
	}

	return storage.NewStorage(drv, args...)
}

func NewMetathingsDeviceCloudServiceOption(opt *DeviceCloudOption) *service.MetathingsDeviceCloudServiceOption {
	dc_opt := &service.MetathingsDeviceCloudServiceOption{}
	dc_opt.Session.Id = id_helper.NewId()
	dc_opt.Connection = opt.Connection
	dc_opt.Credential.Id = opt.GetCredentialId()
	dc_opt.Credential.Secret = opt.GetCredentialSecret()
	return dc_opt
}

type NewDeviceCloudRoutingParams struct {
	fx.In

	Router  *mux.Router
	Service *service.MetathingsDeviceCloudService
}

func NewDeviceCloudRouting(p NewDeviceCloudRoutingParams) error {
	sr := p.Router.PathPrefix("/v1/device_cloud").Subrouter()
	sr.HandleFunc("/actions/heartbeat", p.Service.Heartbeat).Methods("POST")
	sr.HandleFunc("/actions/issue_module_token", p.Service.IssueModuleToken).Methods("POST")
	sr.HandleFunc("/actions/show_module", p.Service.ShowModule).Methods("POST")
	sr.HandleFunc("/actions/push_frame_to_flow", p.Service.PushFrameToFlow).Methods("POST")
	sr.HandleFunc("/actions/show_module_firmware_descriptor", p.Service.ShowModuleFirmwareDescriptor).Methods("POST")

	return nil
}

func runDeviceCloud() error {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			GetDeviceCloudOptions,
			mux.NewRouter,
			cmd_contrib.NewServerTransportCredentials,
			cmd_contrib.NewLogger("device_cloud"),
			cmd_contrib.NewListener,
			cmd_contrib.NewOpentracing,
			cmd_contrib.NewClientFactory,
			cmd_contrib.NewNoExpireTokener,
			token_helper.NewTokenValidator,
			NewDeviceCloudStorage,
			NewMetathingsDeviceCloudServiceOption,
			service.NewMetathingsDeviceCloudService,
		),
		fx.Invoke(
			NewDeviceCloudRouting,
			cmd_contrib.NewHttpServer,
		),
	)

	if err := app.Start(context.TODO()); err != nil {
		return err
	}
	go func() {
		defer app.Stop(context.TODO())

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
	device_cloud_opt = NewDeviceCloudOption()

	flags := deviceCloudCmd.Flags()

	flags.StringVarP(device_cloud_opt.GetListenP(), "listen", "l", "127.0.0.1:5003", "MetaThings Device Cloud Service listening address")
	flags.StringVar(device_cloud_opt.GetCertFileP(), "cert-file", "certs/server.crt", "MetaThings Device Cloud Service Credential File")
	flags.StringVar(device_cloud_opt.GetKeyFileP(), "key-file", "certs/server.key", "MetaThings Device Cloud Service Key File")

	RootCmd.AddCommand(deviceCloudCmd)
}
