package cmd

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

type _applicationCredentialOptions struct {
	Id     string
	Secret string
}

type _logOptions struct {
	Level string
}

type _rootOptions struct {
	Config                string
	Stage                 string
	Verbose               bool
	Token                 string
	Log                   _logOptions
	ApplicationCredential _applicationCredentialOptions `mapstructure:"application_credential"`
	ServiceConfig         _serviceConfigOptions         `mapstructure:"service_config"`
}

var (
	root_opts *_rootOptions
)

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

	RootCmd.PersistentFlags().BoolVar(&root_opts.Verbose, "verbose", false, "Verbose mode")

	RootCmd.PersistentFlags().StringVar(&root_opts.Log.Level, "log-level", "info", "Logging Level[debug, info, warn, error]")

	RootCmd.PersistentFlags().StringVar(&root_opts.ServiceConfig.Metathings.Address, "addr", "", "Metathings Service Address")

	RootCmd.PersistentFlags().StringVar(&root_opts.Token, "token", "", "MetaThings Token")
	viper.BindPFlag("token", RootCmd.PersistentFlags().Lookup("token"))

	RootCmd.PersistentFlags().StringVar(&root_opts.ApplicationCredential.Id, "application-credential-id", "", "MetaThings Application Credential ID")
	viper.BindPFlag("application-credential-id", RootCmd.PersistentFlags().Lookup("application-credential-id"))

	RootCmd.PersistentFlags().StringVar(&root_opts.ApplicationCredential.Secret, "application-credential-secret", "", "MetaThings Application Credential Secret")
	viper.BindPFlag("application-credential-secret", RootCmd.PersistentFlags().Lookup("application-credential-secret"))
}

func initialize() {
	lvl, err := log.ParseLevel(root_opts.Log.Level)
	if err != nil {
		log.WithField("log.level", root_opts.Log.Level).Fatalf("bad log level")
	}
	log.SetLevel(lvl)

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

func initConfig() {
	if root_opts.Config != "" {
		viper.SetConfigFile(root_opts.Config)
		if err := viper.ReadInConfig(); err != nil {
			log.WithError(err).Fatalf("failed to read config")
		}
	}
}
