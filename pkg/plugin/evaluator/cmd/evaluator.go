package metathings_plugin_evaluator_cmd

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.uber.org/dig"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	cfg_helper "github.com/nayotta/metathings/pkg/common/config"
	constant_helper "github.com/nayotta/metathings/pkg/common/constant"
	service "github.com/nayotta/metathings/pkg/plugin/evaluator/service"
	dssdk "github.com/nayotta/metathings/sdk/data_storage"
)

type EvaluatorPluginOption struct {
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
	Evaluator                     struct {
		Endpoint string
	}
	DataStorage map[string]interface{}
}

func NewEvaluatorPluginOption() *EvaluatorPluginOption {
	return &EvaluatorPluginOption{
		ServiceBaseOption: cmd_contrib.CreateServiceBaseOption(),
	}
}

func init_data_storage(opt *EvaluatorPluginOption) {
	mds := map[string]interface{}{}
	vds := cmd_helper.GetFromStage().Sub("data_storage")
	for _, key := range vds.AllKeys() {
		mds[key] = vds.Get(key)
	}
	opt.DataStorage = mds
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

		init_data_storage(opt)

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
	opt := &service.EvaluatorPluginServiceOption{}

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
	c.Provide(NewEvaluatorPluginServiceOption)
	c.Provide(service.NewEvaluatorPluginService)
	if err := c.Invoke(func(evltr_plg_srv *service.EvaluatorPluginService) {
		srv = evltr_plg_srv
	}); err != nil {
		return nil, err
	}

	return srv, nil
}
