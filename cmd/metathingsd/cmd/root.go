package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/bigdatagz/metathings/pkg/common/cmd/helper"
)

const (
	METATHINGSD_PREFIX = "MTD_"
)

var (
	V helper.GetArgument
	A helper.WithPrefix
)

var (
	root_opts struct {
		verbose   bool
		log_level string
	}
)

var (
	RootCmd = &cobra.Command{
		Use:   "metathings",
		Short: "MetaThings Command Line Toolkits",
	}
)

func init() {
	h := helper.NewArgumentHelper(METATHINGSD_PREFIX)
	V = h.GetString
	A = h.PrefixWith
	viper.AutomaticEnv()

	RootCmd.PersistentFlags().BoolVar(&root_opts.verbose, "verbose", false, "Verbose mode")
	RootCmd.PersistentFlags().StringVar(&root_opts.log_level, "log-level", "info", "Logging Level[debug, info, warn, error]")
	viper.BindPFlag(A("LOG_LEVEL"), RootCmd.PersistentFlags().Lookup("log-level"))
}

func initialize() {
	lvl, err := log.ParseLevel(V("log_level"))
	if err != nil {
		log.Fatalf("bad log level %v: %v", V("log_level"), err)
	}
	log.SetLevel(lvl)
}
