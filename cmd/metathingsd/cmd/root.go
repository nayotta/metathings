package cmd

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cmd_helper "github.com/bigdatagz/metathings/pkg/common/cmd"
)

const (
	METATHINGSD_PREFIX = "mtd"
)

type _rootOptions struct {
	Config                string
	Stage                 string
	Service               string
	Verbose               bool
	Log                   _logOptions
	ApplicationCredential _applicationCredentialOptions `mapstructure:"application_credential"`
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

func initialize() {
	lvl, err := log.ParseLevel(root_opts.Log.Level)
	if err != nil {
		log.Fatalf("bad log level %v: %v", root_opts.Log.Level, err)
	}
	log.SetLevel(lvl)
	log.WithField("log.level", root_opts.Log.Level).Debugf("set log level")
}

func initConfig() {
	if root_opts.Config != "" {
		viper.SetConfigFile(root_opts.Config)
		if err := viper.ReadInConfig(); err != nil {
			log.WithError(err).Fatalf("failed to read config")
		}
	}
}

func init() {
	root_opts = &_rootOptions{}
	cmd_helper.SetDefaultHooks(initialize)

	cobra.OnInitialize(initConfig)
	viper.AutomaticEnv()
	viper.SetEnvPrefix(METATHINGSD_PREFIX)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.BindEnv("stage")

	RootCmd.PersistentFlags().StringVarP(&root_opts.Config, "config", "c", "", "Config file")

	RootCmd.PersistentFlags().BoolVar(&root_opts.Verbose, "verbose", false, "Verbose mode")

	RootCmd.PersistentFlags().StringVar(&root_opts.Log.Level, "log-level", "info", "Logging Level[debug, info, warn, error]")

	RootCmd.PersistentFlags().StringVar(&root_opts.ApplicationCredential.Id, "application-credential-id", "", "MetaThings Application Credential ID")

	RootCmd.PersistentFlags().StringVar(&root_opts.ApplicationCredential.Secret, "application-credential-secret", "", "MetaThings Application Credential Secret")
}
