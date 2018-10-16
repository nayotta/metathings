package cmd

import (
	"strings"

	"github.com/nayotta/viper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
)

const (
	METATHINGSD_PREFIX = "mtd"
)

// DEPRECATED(Peer): rename to _rootOption
type _rootOptions struct {
	cmd_helper.RootOptions `mapstructure:",squash"`
	Service                string
}

type _rootOption struct {
	cmd_helper.RootOptions `mapstructure:",squash"`
	Service                string
}

var (
	// DEPRECATED(Peer): rename to root_opt
	root_opts *_rootOptions

	root_opt *_rootOption
)

var (
	RootCmd = &cobra.Command{
		Use:   "metathings",
		Short: "MetaThings Command Line Toolkits",
	}
)

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
