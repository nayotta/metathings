package main

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
)

const (
	METATHINGS_PLUGIN_PREFIX = "mtp"
)

type EvaluatorPluginOption struct {
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
}

func NewEvaluatorPluginOption() *EvaluatorPluginOption {
	return &EvaluatorPluginOption{
		ServiceBaseOption: cmd_contrib.CreateServiceBaseOption(),
	}
}

func LoadEvaluatorPluginOption() (*EvaluatorPluginOption, error) {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(METATHINGS_PLUGIN_PREFIX)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.BindEnv("stage")

	cfg := os.Getenv("MTP_CONFIG")
	viper.SetConfigFile(cfg)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	opt := NewEvaluatorPluginOption()

	cmd_helper.UnmarshalConfig(&opt)

	return opt, nil
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

func NewEvaluatorPluginServiceOption(o *EvaluatorPluginOption) (*EvaluatorPluginServiceOption, error) {
	return &EvaluatorPluginServiceOption{}, nil
}

func NewEvaluatorPluginService(
	opt *EvaluatorPluginServiceOption,
	logger log.FieldLogger,
	tknr token_helper.Tokener,
	cli_fty *client_helper.ClientFactory,
) (*EvaluatorPluginService, error) {
	return &EvaluatorPluginService{
		opt:     opt,
		logger:  logger,
		tknr:    tknr,
		cli_fty: cli_fty,
	}, nil
}
