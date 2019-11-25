package cmd

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	cfg_helper "github.com/nayotta/metathings/pkg/common/config"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	connection "github.com/nayotta/metathings/pkg/deviced/connection"
	flow "github.com/nayotta/metathings/pkg/deviced/flow"
	service "github.com/nayotta/metathings/pkg/deviced/service"
	session_storage "github.com/nayotta/metathings/pkg/deviced/session_storage"
	simple_storage "github.com/nayotta/metathings/pkg/deviced/simple_storage"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	authorizer "github.com/nayotta/metathings/pkg/identityd2/authorizer"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type DevicedOption struct {
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
	SessionStorage                map[string]interface{}
	SimpleStorage                 map[string]interface{}
	ConnectionCenter              struct {
		Storage map[string]interface{}
		Bridge  map[string]interface{}
	}
	Flow struct {
		Mongo struct {
			Uri      string
			Database string
		}
		Redis struct {
			Addr     string
			DB       int
			Password string
		}
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

func init_session_storage(opt *DevicedOption) {
	mss := map[string]interface{}{}
	vss := cmd_helper.GetFromStage().Sub("session_storage")
	for _, key := range vss.AllKeys() {
		mss[key] = vss.Get(key)
	}
	opt.SessionStorage = mss
}

func init_simple_storage(opt *DevicedOption) {
	mss := map[string]interface{}{}
	vss := cmd_helper.GetFromStage().Sub("simple_storage")
	for _, key := range vss.AllKeys() {
		mss[key] = vss.Get(key)
	}
	opt.SimpleStorage = mss
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

			init_service_cmd_option(opt_t, deviced_opt)
			init_session_storage(opt_t)
			init_simple_storage(opt_t)
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

func NewConnectionCenter(opt *DevicedOption, sess_stor session_storage.SessionStorage, logger log.FieldLogger) (connection.ConnectionCenter, error) {
	var name string
	var args []interface{}
	var err error
	var conn_stor connection.Storage
	var conn_brfty connection.BridgeFactory
	var cc connection.ConnectionCenter

	if name, args, err = parse_connection_center_option(opt.ConnectionCenter.Storage); err != nil {
		return nil, err
	}
	args = append(args, "logger", logger)

	if conn_stor, err = connection.NewStorage(name, args...); err != nil {
		return nil, err
	}

	if name, args, err = parse_connection_center_option(opt.ConnectionCenter.Bridge); err != nil {
		return nil, err
	}
	args = append(args, "logger", logger)

	if conn_brfty, err = connection.NewBridgeFactory(name, args...); err != nil {
		return nil, err
	}

	if cc, err = connection.NewConnectionCenter(conn_brfty, conn_stor, sess_stor, logger); err != nil {
		return nil, err
	}

	return cc, nil
}

func NewSessionStorage(opt *DevicedOption, logger log.FieldLogger) (session_storage.SessionStorage, error) {
	drv, args, err := cfg_helper.ParseConfigOption("driver", opt.SessionStorage)
	if err != nil {
		return nil, err
	}
	args = append(args, "logger", logger)

	sess_stor, err := session_storage.NewSessionStorage(drv, args...)
	if err != nil {
		return nil, err
	}

	return sess_stor, nil
}

func NewSimpleStorage(opt *DevicedOption, logger log.FieldLogger) (simple_storage.SimpleStorage, error) {
	name, args, err := cfg_helper.ParseConfigOption("name", opt.SimpleStorage)
	if err != nil {
		return nil, err
	}
	args = append(args, "logger", logger)

	simp_stor, err := simple_storage.NewSimpleStorage(name, args...)
	if err != nil {
		return nil, err
	}

	return simp_stor, nil
}

func NewMetathingsDevicedServiceOption(opt *DevicedOption) *service.MetathingsDevicedServiceOption {
	o := &service.MetathingsDevicedServiceOption{}
	return o
}

func runDeviced() error {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			GetDevicedOptions,
			cmd_contrib.NewServerTransportCredentials,
			cmd_contrib.NewLogger("deviced"),
			cmd_contrib.NewListener,
			cmd_contrib.NewGrpcServer,
			cmd_contrib.NewClientFactory,
			cmd_contrib.NewNoExpireTokener,
			token_helper.NewTokenValidator,
			NewSessionStorage,
			NewSimpleStorage,
			NewConnectionCenter,
			NewDevicedStorage,
			NewMetathingsDevicedServiceOption,
			authorizer.NewAuthorizer,
			cmd_contrib.NewValidator,
			func(opt *DevicedOption, logger log.FieldLogger) (flow.FlowFactory, error) {
				return flow.NewFlowFactory(
					"default",
					"redis_stream_addr", opt.Flow.Redis.Addr,
					"redis_stream_db", opt.Flow.Redis.DB,
					"redis_stream_password", opt.Flow.Redis.Password,
					"mongo_uri", opt.Flow.Mongo.Uri,
					"mongo_database", opt.Flow.Mongo.Database,
					"logger", logger,
				)
			},
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
