package metathings_plugin_evaluator_cmd

import (
	"strings"

	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	cfg_helper "github.com/nayotta/metathings/pkg/common/config"
	constant_helper "github.com/nayotta/metathings/pkg/common/constant"
	evaluatord_storage "github.com/nayotta/metathings/pkg/evaluatord/storage"
	service "github.com/nayotta/metathings/pkg/plugin/evaluator/service"
	dssdk "github.com/nayotta/metathings/sdk/data_storage"
	dsdk "github.com/nayotta/metathings/sdk/deviced"
)

type EvaluatorPluginOption struct {
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
	Evaluator                     struct {
		Endpoint string
	}
	DataStorage   map[string]interface{}
	SimpleStorage map[string]interface{}
	TaskStorage   map[string]interface{}
	Caller        map[string]interface{}
}

func NewEvaluatorPluginOption() *EvaluatorPluginOption {
	return &EvaluatorPluginOption{
		ServiceBaseOption: cmd_contrib.CreateServiceBaseOption(),
	}
}

func LoadEvaluatorPluginOption(path string) func() (*EvaluatorPluginOption, error) {
	return func() (*EvaluatorPluginOption, error) {
		viper.AutomaticEnv()
		viper.SetEnvPrefix(constant_helper.PREFIX_METATHINGS_PLUGIN)
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
		viper.BindEnv("stage")

		viper.SetConfigFile(path)
		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}

		opt := NewEvaluatorPluginOption()
		cmd_helper.UnmarshalConfig(&opt)

		cmd_helper.InitManyStringMapFromConfigWithStage([]cmd_helper.InitManyOption{
			{&opt.TaskStorage, "task_storage"},
		})

		opt.SetServiceName("evaluator-plugin")

		cmd_helper.InitManyStringMapFromConfigWithStage([]cmd_helper.InitManyOption{
			{&opt.DataStorage, "data_storage"},
			{&opt.SimpleStorage, "simple_storage"},
			{&opt.Caller, "caller"},
		})

		return opt, nil
	}
}

func GetEvaluatorPluginOptions(opt *EvaluatorPluginOption) (
	cmd_contrib.LoggerOptioner,
	cmd_contrib.ServiceOptioner,
	cmd_contrib.CredentialOptioner,
	cmd_contrib.ServiceEndpointsOptioner,
	cmd_contrib.OpentracingOptioner,
) {
	return opt,
		opt,
		opt,
		opt,
		opt
}

type NewEvaluatorPluginOptionParams struct {
	fx.In

	Tracer opentracing.Tracer `name:"opentracing_tracer"`
	Option *EvaluatorPluginOption
}

func NewEvaluatorPluginServiceOption(p NewEvaluatorPluginOptionParams) (*service.EvaluatorPluginServiceOption, error) {
	opt := service.NewEvaluatorPluginServiceOption()

	opt.Evaluator.Endpoint = p.Option.Evaluator.Endpoint
	opt.IsTraced = p.Tracer != nil

	return opt, nil
}

func NewDataStorage(o *EvaluatorPluginOption, logger log.FieldLogger) (dssdk.DataStorage, error) {
	name, args, err := cfg_helper.ParseConfigOption("name", o.DataStorage, "logger", logger)
	if err != nil {
		return nil, err
	}

	ds, err := dssdk.NewDataStorage(name, args...)
	if err != nil {
		return nil, err
	}

	return ds, nil
}

func NewSimpleStorage(o *EvaluatorPluginOption, cli_fty *client_helper.ClientFactory, logger log.FieldLogger) (dsdk.SimpleStorage, error) {
	name, args, err := cfg_helper.ParseConfigOption("name", o.SimpleStorage, "logger", logger)
	if err != nil {
		return nil, err
	}

	switch name {
	case "default":
		args = append(args, "client_factory", cli_fty)
	}

	ss, err := dsdk.NewSimpleStorage(name, args...)
	if err != nil {
		return nil, err
	}

	return ss, nil
}

type NewTaskStorageParams struct {
	fx.In

	Option *EvaluatorPluginOption
	Logger log.FieldLogger
	Tracer opentracing.Tracer `name:"opentracing_tracer" optional:"true"`
}

func NewTaskStorage(p NewTaskStorageParams) (evaluatord_storage.TaskStorage, error) {
	var drv string
	var args []interface{}
	var err error

	if drv, args, err = cfg_helper.ParseConfigOption("name", p.Option.TaskStorage, "logger", p.Logger, "tracer", p.Tracer); err != nil {
		return nil, err
	}

	return evaluatord_storage.NewTaskStorage(drv, args...)

}

func NewCaller(o *EvaluatorPluginOption, cli_fty *client_helper.ClientFactory, logger log.FieldLogger) (dsdk.Caller, error) {
	name, args, err := cfg_helper.ParseConfigOption("name", o.Caller, "logger", logger)
	if err != nil {
		return nil, err
	}

	switch name {
	case "default":
		args = append(args, "client_factory", cli_fty)
	}

	c, err := dsdk.NewCaller(name, args...)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func NewEvaluatorPluginService(cfg string) (*service.EvaluatorPluginService, error) {
	var srv *service.EvaluatorPluginService

	c := dig.New()
	c.Provide(LoadEvaluatorPluginOption(cfg))
	c.Provide(GetEvaluatorPluginOptions)
	c.Provide(cmd_contrib.NewLogger("evaluator-plugin"))
	c.Provide(cmd_contrib.NewTokener)
	c.Provide(cmd_contrib.NewOpentracing)
	c.Provide(cmd_contrib.NewClientFactory)
	c.Provide(NewDataStorage)
	c.Provide(NewSimpleStorage)
	c.Provide(NewTaskStorage)
	c.Provide(NewCaller)
	c.Provide(NewEvaluatorPluginServiceOption)
	c.Provide(service.NewEvaluatorPluginService)
	if err := c.Invoke(func(evltr_plg_srv *service.EvaluatorPluginService) {
		srv = evltr_plg_srv
	}); err != nil {
		return nil, err
	}

	return srv, nil
}
