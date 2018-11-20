package cmd

import (
	"strings"

	"github.com/nayotta/viper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
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

var (
	root_opts *_rootOptions
	base_opt  *cmd_contrib.BaseOption
)

var (
	RootCmd = &cobra.Command{
		Use:   "metathingsd",
		Short: "MetaThingsd Command Line Toolkits",
	}
)

func initConfig() {
	cfg := base_opt.GetConfig()
	if cfg != "" {
		viper.SetConfigFile(cfg)
		if err := viper.ReadInConfig(); err != nil {
			log.WithError(err).Fatalf("failed to read config")
		}
	}
}

func init() {
	base_opt = &cmd_contrib.BaseOption{}

	cobra.OnInitialize(initConfig)
	viper.AutomaticEnv()
	viper.SetEnvPrefix(METATHINGSD_PREFIX)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.BindEnv("stage")

	flags := RootCmd.PersistentFlags()

	flags.StringVarP(base_opt.GetConfigP(), "config", "c", "", "Config file")
	flags.BoolVar(base_opt.GetVerboseP(), "verbose", false, "Verbose mode")
	flags.StringVar(base_opt.GetLevelP(), "log-level", "info", "Logging Level[debug, info, warn, error]")
	flags.StringVar(base_opt.GetCredentialIdP(), "application-credential-id", "", "MetaThings Application Credential ID")
	flags.StringVar(base_opt.GetCredentialSecretP(), "application-credential-secret", "", "MetaThings Application Credential Secret")
}
