package cmd

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	connection "github.com/nayotta/metathings/pkg/deviced/connection"
	service "github.com/nayotta/metathings/pkg/deviced/service"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	policy "github.com/nayotta/metathings/pkg/identityd2/policy"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type DevicedOption struct {
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
	ConnectionCenter              struct {
		Storage map[string]interface{}
		Bridge  map[string]interface{}
	}
}

func NewDevicedOption() *DevicedOption {
	return &DevicedOption{
		ServiceBaseOption: cmd_contrib.CreateServiceBaseOption(),
	}
}

var (
	deviced_opt *DevicedOption
)

func init_connection_center(opt *DevicedOption) {
	init_connection_center_storage(opt)
	init_connection_center_bridge(opt)
}

func init_connection_center_storage(opt *DevicedOption) {
	mccs := map[string]interface{}{}
	vccs := cmd_helper.GetFromStage().Sub("connection_center").Sub("storage")
	for _, key := range vccs.AllKeys() {
		mccs[key] = vccs.Get(key)
	}
	opt.ConnectionCenter.Storage = mccs
}

func init_connection_center_bridge(opt *DevicedOption) {
	mccb := map[string]interface{}{}
	vccb := cmd_helper.GetFromStage().Sub("connection_center").Sub("bridge")
	for _, key := range vccb.AllKeys() {
		mccb[key] = vccb.Get(key)
	}
	opt.ConnectionCenter.Bridge = mccb
}

var (
	devicedCmd = &cobra.Command{
		Use:   "deviced",
		Short: "Device Service Daemon",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if base_opt.Config == "" {
				return
			}

			opt_t := NewDevicedOption()
			cmd_helper.UnmarshalConfig(opt_t)
			base_opt = &opt_t.BaseOption

			init_cmd_option(opt_t, deviced_opt)
			init_connection_center(opt_t)

			deviced_opt = opt_t
			deviced_opt.SetServiceName("deviced")
			deviced_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: cmd_helper.Run("deviced", runDeviced),
	}
)

func GetDevicedOptions() (
	*DevicedOption,
	cmd_contrib.ListenOptioner,
	cmd_contrib.TransportCredentialOptioner,
	cmd_contrib.StorageOptioner,
	cmd_contrib.LoggerOptioner,
	cmd_contrib.ServiceEndpointsOptioner,
	cmd_contrib.CredentialOptioner,
) {
	return deviced_opt,
		deviced_opt,
		deviced_opt,
		deviced_opt,
		deviced_opt,
		deviced_opt,
		deviced_opt
}

func NewDevicedStorage(opt cmd_contrib.StorageOptioner, logger log.FieldLogger) (storage.Storage, error) {
	return storage.NewStorage(opt.GetDriver(), opt.GetUri(), "logger", logger)
}

func parse_connection_center_option(x map[string]interface{}) (string, []interface{}, error) {
	var key string
	var val interface{}
	var name string
	var ok bool

	y := []interface{}{}

	if val, ok = x["name"]; !ok {
		return "", nil, ErrInvalidArgument
	}

	if name, ok = val.(string); !ok {
		return "", nil, ErrInvalidArgument
	}

	for key, val = range x {
		if key == "name" {
			continue
		}

		y = append(y, key, val)
	}

	return name, y, nil
}

func NewConnectionCenter(opt *DevicedOption, logger log.FieldLogger) (connection.ConnectionCenter, error) {
	var name string
	var args []interface{}
	var err error
	var conn_stor connection.Storage
	var conn_brfty connection.BridgeFactory
	var cc connection.ConnectionCenter

	if name, args, err = parse_connection_center_option(opt.ConnectionCenter.Storage); err != nil {
		return nil, err
	}

	if conn_stor, err = connection.NewStorage(name, args...); err != nil {
		return nil, err
	}

	if name, args, err = parse_connection_center_option(opt.ConnectionCenter.Bridge); err != nil {
		return nil, err
	}

	if conn_brfty, err = connection.NewBridgeFactory(name, args...); err != nil {
		return nil, err
	}

	if cc, err = connection.NewConnectionCenter(conn_brfty, conn_stor, logger); err != nil {
		return nil, err
	}

	return cc, nil
}

func NewMetathingsDevicedServiceOption(opt *DevicedOption) *service.MetathingsDevicedServiceOption {
	return &service.MetathingsDevicedServiceOption{}
}

func runDeviced() error {
	app := fx.New(
		fx.Provide(
			GetDevicedOptions,
			cmd_contrib.NewTransportCredentials,
			cmd_contrib.NewLogger("deviced"),
			cmd_contrib.NewListener,
			cmd_contrib.NewGrpcServer,
			cmd_contrib.NewClientFactory,
			cmd_contrib.NewTokener,
			token_helper.NewTokenValidator,
			NewConnectionCenter,
			NewDevicedStorage,
			NewMetathingsDevicedServiceOption,
			policy.NewEnforcer,
			service.NewMetathingsDevicedService,
		),
		fx.Invoke(
			pb.RegisterDevicedServiceServer,
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

func init() {
	deviced_opt = NewDevicedOption()

	flags := devicedCmd.Flags()

	flags.StringVarP(deviced_opt.GetListenP(), "listen", "l", "127.0.0.1:5001", "MetaThings Device Service listening address")
	flags.StringVar(deviced_opt.GetStorage().GetDriverP(), "storage-driver", "sqlite3", "MetaThtings Device Service Storage Driver")
	flags.StringVar(deviced_opt.GetStorage().GetUriP(), "storage-uri", "", "MetaThings Deviced Service Storage URI")
	flags.StringVar(deviced_opt.GetCertFileP(), "cert-file", "certs/server.crt", "MetaThings Device Service Credential File")
	flags.StringVar(deviced_opt.GetKeyFileP(), "key-file", "certs/server.key", "MetaThings Device Service Key File")

	RootCmd.AddCommand(devicedCmd)
}
