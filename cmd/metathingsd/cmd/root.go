package cmd

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	const_helper "github.com/nayotta/metathings/pkg/common/constant"
)

var (
	base_opt *cmd_contrib.BaseOption
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
	viper.SetEnvPrefix(const_helper.PREFIX_METATHINGSD)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.BindEnv("stage")

	flags := RootCmd.PersistentFlags()

	flags.StringVarP(base_opt.GetConfigP(), "config", "c", "", "Config file")
	flags.BoolVar(base_opt.GetVerboseP(), "verbose", false, "Verbose mode")
	flags.StringVar(base_opt.GetLevelP(), "log-level", "info", "Logging Level[debug, info, warn, error]")
	flags.StringVar(base_opt.GetCredentialIdP(), "application-credential-id", "", "MetaThings Application Credential ID")
	flags.StringVar(base_opt.GetCredentialSecretP(), "application-credential-secret", "", "MetaThings Application Credential Secret")
	flags.StringVar(base_opt.GetKeyFileP(), "key", "", "Transport Credential Key")
	flags.StringVar(base_opt.GetCertFileP(), "cert", "", "Transport Credential Cert")
	flags.BoolVar(base_opt.GetInsecureP(), "insecure", false, "Do not verify transport credential")
	flags.BoolVar(base_opt.GetPlainTextP(), "plaintext", false, "Transport data without tls")
}
