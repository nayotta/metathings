package cmd

import (
	"strings"

	"github.com/nayotta/viper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
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

type _rootOptions struct {
	cmd_helper.RootOptions  `mapstructure:",squash"`
	cmd_helper.TokenOptions `mapstructure:",squash"`
	ServiceConfig           _serviceConfigOptions `mapstructure:"service_config"`
}

var (
	root_opts *_rootOptions
)

var (
	_cli_fty *client_helper.ClientFactory
)

func getClientFactory() *client_helper.ClientFactory {
	if _cli_fty == nil {
		_cli_fty, _ = client_helper.NewClientFactory(
			client_helper.NewDefaultServiceConfigs(root_opts.ServiceConfig.Metathings.Address),
			client_helper.WithInsecureOptionFunc(),
		)
	}

	return _cli_fty
}

var (
	RootCmd = &cobra.Command{
		Use:   "metathings",
		Short: "MetaThings Command Line Toolkits",
	}
)

type PreRunHook func()

func preRunHooks(hooks ...PreRunHook) func(*cobra.Command, []string) {
	return func(*cobra.Command, []string) {
		for _, hook := range hooks {
			hook()
		}
	}
}

func defaultPreRunHooks(hook PreRunHook, defaults ...PreRunHook) func(*cobra.Command, []string) {
	if len(defaults) == 0 {
		defaults = []PreRunHook{initialize}
	}
	if hook != nil {
		defaults = append([]PreRunHook{hook}, defaults...)
	}
	return preRunHooks(defaults...)
}

func init() {
	root_opts = &_rootOptions{}

	cobra.OnInitialize(initConfig)
	viper.AutomaticEnv()
	viper.SetEnvPrefix(METATHINGS_PREFIX)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.BindEnv("stage")

	RootCmd.PersistentFlags().StringVarP(&root_opts.Config, "config", "c", "", "Config file")

	RootCmd.PersistentFlags().BoolVar(&root_opts.Verbose, "verbose", false, "Verbose mode")

	RootCmd.PersistentFlags().StringVar(&root_opts.Log.Level, "log-level", "info", "Logging Level[debug, info, warn, error]")

	RootCmd.PersistentFlags().StringVar(&root_opts.ServiceConfig.Metathings.Address, "addr", "mt-api.nayotta.com", "Metathings Service Address")

	RootCmd.PersistentFlags().StringVar(&root_opts.Token, "token", "", "MetaThings Token")
	viper.BindPFlag("token", RootCmd.PersistentFlags().Lookup("token"))

	RootCmd.PersistentFlags().StringVar(&root_opts.ApplicationCredential.Id, "application-credential-id", "", "MetaThings Application Credential ID")
	viper.BindPFlag("application-credential-id", RootCmd.PersistentFlags().Lookup("application-credential-id"))

	RootCmd.PersistentFlags().StringVar(&root_opts.ApplicationCredential.Secret, "application-credential-secret", "", "MetaThings Application Credential Secret")
	viper.BindPFlag("application-credential-secret", RootCmd.PersistentFlags().Lookup("application-credential-secret"))
}

func initialize() {
	token := viper.GetString("token")
	if token != "" {
		root_opts.Token = token
	}

	appCredId := viper.GetString("application-credential-id")
	if appCredId != "" {
		root_opts.ApplicationCredential.Id = appCredId
	}

	appCredSecret := viper.GetString("application-credential-secret")
	if appCredSecret != "" {
		root_opts.ApplicationCredential.Secret = appCredSecret
	}
}

var _GLOBAL_INITIALED = false

func initConfig() {
	if _GLOBAL_INITIALED {
		return
	}

	if root_opts.Config != "" {
		viper.SetConfigFile(root_opts.Config)
		if err := viper.ReadInConfig(); err != nil {
			log.WithError(err).Fatalf("failed to read config")
		}
	}
	_GLOBAL_INITIALED = true
}
