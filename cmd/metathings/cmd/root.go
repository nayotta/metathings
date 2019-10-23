package cmd

import (
	"strings"

	"github.com/nayotta/viper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	constant_helper "github.com/nayotta/metathings/pkg/common/constant"
)

const (
	METATHINGS_PREFIX = "mt"
)

type _metathingsServiceOptions struct {
	Address string
}

type _serviceConfigOptions struct {
	Metathings _metathingsServiceOptions
}

// DEPRECATED(Peer): rename to _rootOption
type _rootOptions struct {
	cmd_helper.RootOptions `mapstructure:",squash"`
	ServiceConfig          _serviceConfigOptions `mapstructure:"service_config"`
}

var (
	root_opts *_rootOptions
	base_opt  *cmd_contrib.BaseOption
)

var (
	_cli_fty *client_helper.ClientFactory
)

var (
	RootCmd = &cobra.Command{
		Use:   "metathings",
		Short: "MetaThings Command Line Toolkits",
	}
)

var _GLOBAL_INITIALED = false

func initConfig() {
	if _GLOBAL_INITIALED {
		return
	}

	cfg := base_opt.GetConfig()
	if cfg != "" {
		viper.SetConfigFile(cfg)
		if err := viper.ReadInConfig(); err != nil {
			log.WithError(err).Fatalf("failed to read config")
		}
	}

	_GLOBAL_INITIALED = true
}

func init() {
	opt := cmd_contrib.CreateBaseOption()
	base_opt = &opt

	cobra.OnInitialize(initConfig)
	viper.AutomaticEnv()
	viper.SetEnvPrefix(METATHINGS_PREFIX)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.BindEnv("stage")

	flags := RootCmd.PersistentFlags()

	flags.StringVarP(base_opt.GetConfigP(), "config", "c", "", "Config file")
	flags.BoolVar(base_opt.GetVerboseP(), "verbose", false, "Verbose mode")
	flags.StringVar(base_opt.GetLevelP(), "log-level", "info", "Logging level[debug, info, warn, error]")
	flags.StringVar(base_opt.GetKeyFileP(), "key", "", "Transport Credential Key")
	flags.StringVar(base_opt.GetCertFileP(), "cert", "", "Transport Credential Cert")
	flags.BoolVar(base_opt.GetInsecureP(), "insecure", false, "Do not verify transport credential")
	flags.BoolVar(base_opt.GetPlainTextP(), "plaintext", false, "Transport data without tls")

	flags.StringVar(base_opt.GetServiceEndpoint(client_helper.DEFAULT_CONFIG).GetAddressP(), "addr", constant_helper.CONSTANT_METATHINGSD_DEFAULT_HOST, "MetaThings Service Address")
}
