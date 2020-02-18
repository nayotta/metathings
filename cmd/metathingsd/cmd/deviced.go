package cmd

import (
	"context"

	"github.com/opentracing/opentracing-go"
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
	evaluatord_sdk "github.com/nayotta/metathings/sdk/evaluatord"
)

type DevicedOption struct {
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
	SessionStorage                map[string]interface{}
	SimpleStorage                 map[string]interface{}
	ConnectionCenter              struct {
		Storage map[string]interface{}
		Bridge  map[string]interface{}
	}
	Flow         map[string]interface{}
	FlowSet      map[string]interface{}
	DataLauncher map[string]interface{}
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

func init_flow(opt *DevicedOption) {
	mf := map[string]interface{}{}
	vf := cmd_helper.GetFromStage().Sub("flow")
	for _, key := range vf.AllKeys() {
		mf[key] = vf.Get(key)
	}
	opt.Flow = mf
}

func init_flow_set(opt *DevicedOption) {
	mfs := map[string]interface{}{}
	vfs := cmd_helper.GetFromStage().Sub("flow_set")
	for _, key := range vfs.AllKeys() {
		mfs[key] = vfs.Get(key)
	}
	opt.FlowSet = mfs
}

func init_data_launcher(opt *DevicedOption) {
	mdl := map[string]interface{}{}
	vdl := cmd_helper.GetFromStage().Sub("data_launcher")
	for _, key := range vdl.AllKeys() {
		mdl[key] = vdl.Get(key)
	}
	opt.DataLauncher = mdl
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
			init_flow(opt_t)
			init_flow_set(opt_t)
			init_data_launcher(opt_t)

			deviced_opt = opt_t
			deviced_opt.SetServiceName("deviced")
			deviced_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: cmd_helper.Run("deviced", runDeviced),
	}
)

func GetDevicedOptions() (
	*DevicedOption,
	cmd_contrib.ServiceOptioner,
	cmd_contrib.ListenOptioner,
	cmd_contrib.TransportCredentialOptioner,
	cmd_contrib.StorageOptioner,
	cmd_contrib.LoggerOptioner,
	cmd_contrib.ServiceEndpointsOptioner,
	cmd_contrib.CredentialOptioner,
	cmd_contrib.OpentracingOptioner,
) {
	return deviced_opt,
		deviced_opt,
		deviced_opt,
		deviced_opt,
		deviced_opt,
		deviced_opt,
		deviced_opt,
		deviced_opt,
		deviced_opt
}

type NewDevicedDataLauncherParams struct {
	fx.In

	Option *DevicedOption
	Logger log.FieldLogger
	Tracer opentracing.Tracer `name:"opentracing_tracer" optional:"true"`
}

func NewDevicedDataLauncher(p NewDevicedDataLauncherParams) (evaluatord_sdk.DataLauncher, error) {
	var name string
	var args []interface{}
	var err error

	if name, args, err = cfg_helper.ParseConfigOption("name", p.Option.DataLauncher, "logger", p.Logger); err != nil {
		return nil, err
	}

	return evaluatord_sdk.NewDataLauncher(name, args...)
}

type NewDevicedStorageParams struct {
	fx.In

	Option cmd_contrib.StorageOptioner
	Logger log.FieldLogger
	Tracer opentracing.Tracer `name:"opentracing_tracer" optional:"true"`
}

func NewDevicedStorage(p NewDevicedStorageParams) (storage.Storage, error) {
	return storage.NewStorage(p.Option.GetDriver(), p.Option.GetUri(), "logger", p.Logger, "tracer", p.Tracer)
}

func NewConnectionCenter(opt *DevicedOption, sess_stor session_storage.SessionStorage, logger log.FieldLogger) (connection.ConnectionCenter, error) {
	var name string
	var args []interface{}
	var err error
	var conn_stor connection.Storage
	var conn_brfty connection.BridgeFactory
	var cc connection.ConnectionCenter

	if name, args, err = cfg_helper.ParseConfigOption("name", opt.ConnectionCenter.Storage, "logger", logger); err != nil {
		return nil, err
	}

	if conn_stor, err = connection.NewStorage(name, args...); err != nil {
		return nil, err
	}

	if name, args, err = cfg_helper.ParseConfigOption("name", opt.ConnectionCenter.Bridge, "logger", logger); err != nil {
		return nil, err
	}

	if conn_brfty, err = connection.NewBridgeFactory(name, args...); err != nil {
		return nil, err
	}

	if cc, err = connection.NewConnectionCenter(conn_brfty, conn_stor, sess_stor, logger); err != nil {
		return nil, err
	}

	return cc, nil
}

func NewSessionStorage(opt *DevicedOption, logger log.FieldLogger) (session_storage.SessionStorage, error) {
	drv, args, err := cfg_helper.ParseConfigOption("driver", opt.SessionStorage, "logger", logger)
	if err != nil {
		return nil, err
	}

	return session_storage.NewSessionStorage(drv, args...)
}

func NewSimpleStorage(opt *DevicedOption, logger log.FieldLogger) (simple_storage.SimpleStorage, error) {
	name, args, err := cfg_helper.ParseConfigOption("name", opt.SimpleStorage, "logger", logger)
	if err != nil {
		return nil, err
	}

	return simple_storage.NewSimpleStorage(name, args...)
}

func NewFlowFactory(opt *DevicedOption, logger log.FieldLogger) (flow.FlowFactory, error) {
	name, args, err := cfg_helper.ParseConfigOption("driver", opt.Flow, "logger", logger)
	if err != nil {
		return nil, err
	}

	return flow.NewFlowFactory(name, args...)
}

func NewFlowSetFactory(opt *DevicedOption, logger log.FieldLogger) (flow.FlowSetFactory, error) {
	name, args, err := cfg_helper.ParseConfigOption("driver", opt.FlowSet, "logger", logger)
	if err != nil {
		return nil, err
	}

	return flow.NewFlowSetFactory(name, args...)
}

func NewMetathingsDevicedServiceOption(opt *DevicedOption) *service.MetathingsDevicedServiceOption {
	o := &service.MetathingsDevicedServiceOption{}

	o.Methods.PutObjectStreaming.Timeout = 1200
	if to, ok := opt.SimpleStorage["timeout"]; ok {
		if toi, ok := to.(int); ok {
			o.Methods.PutObjectStreaming.Timeout = int64(toi)
		}
	}

	o.Methods.PutObjectStreaming.ChunkSize = 256 * 1024
	if cs, ok := opt.SimpleStorage["chunk_size"]; ok {
		if csi, ok := cs.(int); ok {
			o.Methods.PutObjectStreaming.ChunkSize = int64(csi)
		}
	}

	o.Methods.PutObjectStreaming.ChunkPerRequest = 4
	if cpr, ok := opt.SimpleStorage["chunk_per_request"]; ok {
		if cpri, ok := cpr.(int); ok {
			o.Methods.PutObjectStreaming.ChunkPerRequest = cpri
		}
	}

	o.Methods.PutObjectStreaming.PullRequestRetry = 10
	if prr, ok := opt.SimpleStorage["pull_request_retry"]; ok {
		if prri, ok := prr.(int); ok {
			o.Methods.PutObjectStreaming.PullRequestRetry = prri
		}
	}

	o.Methods.PutObjectStreaming.PullRequestTimeout = 12
	if prt, ok := opt.SimpleStorage["pull_request_timeout"]; ok {
		if prti, ok := prt.(int); ok {
			o.Methods.PutObjectStreaming.PullRequestTimeout = int64(prti)
		}
	}

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
			cmd_contrib.NewOpentracing,
			cmd_contrib.NewGrpcServer,
			cmd_contrib.NewClientFactory,
			cmd_contrib.NewNoExpireTokener,
			NewDevicedDataLauncher,
			token_helper.NewTokenValidator,
			NewSessionStorage,
			NewSimpleStorage,
			NewConnectionCenter,
			NewDevicedStorage,
			NewMetathingsDevicedServiceOption,
			authorizer.NewAuthorizer,
			cmd_contrib.NewValidator,
			NewFlowFactory,
			NewFlowSetFactory,
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
