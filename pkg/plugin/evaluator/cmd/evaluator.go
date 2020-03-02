package metathings_plugin_evaluator_cmd

import (
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/dig"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	constant_helper "github.com/nayotta/metathings/pkg/common/constant"
	service "github.com/nayotta/metathings/pkg/plugin/evaluator/service"
)

type EvaluatorPluginOption struct {
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
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
	return &service.EvaluatorPluginServiceOption{}, nil
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
	c.Provide(NewEvaluatorPluginServiceOption)
	c.Provide(service.NewEvaluatorPluginService)
	if err := c.Invoke(func(evltr_plg_srv *service.EvaluatorPluginService) {
		srv = evltr_plg_srv
	}); err != nil {
		return nil, err
	}

	return srv, nil
}
