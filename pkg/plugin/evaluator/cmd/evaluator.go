package metathings_plugin_evaluator_cmd

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.uber.org/dig"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	cfg_helper "github.com/nayotta/metathings/pkg/common/config"
	constant_helper "github.com/nayotta/metathings/pkg/common/constant"
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
}

func NewEvaluatorPluginOption() *EvaluatorPluginOption {
	return &EvaluatorPluginOption{
		ServiceBaseOption: cmd_contrib.CreateServiceBaseOption(),
	}
}

func init_string_map_from_config_with_stage(dst *map[string]interface{}, key string) {
	sm := make(map[string]interface{})
	vm := cmd_helper.GetFromStage().Sub(key)
	for _, k := range vm.AllKeys() {
		sm[k] = vm.Get(k)
	}
	*dst = sm
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

		init_string_map_from_config_with_stage(&opt.DataStorage, "data_storage")
		init_string_map_from_config_with_stage(&opt.SimpleStorage, "simple_storage")

		return opt, nil
	}
}

func GetEvaluatorPluginOptions(opt *EvaluatorPluginOption) (
	cmd_contrib.LoggerOptioner,
	cmd_contrib.CredentialOptioner,
	cmd_contrib.ServiceEndpointsOptioner,
	cmd_contrib.OpentracingOptioner,
) {
	return opt,
		opt,
		opt,
		opt
}

func NewEvaluatorPluginServiceOption(o *EvaluatorPluginOption) (*service.EvaluatorPluginServiceOption, error) {
	opt := service.NewEvaluatorPluginServiceOption()

	opt.Evaluator.Endpoint = o.Evaluator.Endpoint

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
	c.Provide(NewEvaluatorPluginServiceOption)
	c.Provide(service.NewEvaluatorPluginService)
	if err := c.Invoke(func(evltr_plg_srv *service.EvaluatorPluginService) {
		srv = evltr_plg_srv
	}); err != nil {
		return nil, err
	}

	return srv, nil
}
