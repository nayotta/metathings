package cmd

import (
	"github.com/bigdatagz/metathings/pkg/common/cmd/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	METATHINGS_PREFIX = "MT_"
)

var (
	V func(string) string
)

var (
	root_opts struct {
		verbose bool
		addr    string
		token   string
	}
)

var (
	RootCmd = &cobra.Command{
		Use:   "metathings",
		Short: "MetaThings Command Line Toolkits",
	}
)

func init() {
	V = helper.NewArgumentHelper(METATHINGS_PREFIX).GetString
	viper.AutomaticEnv()

	RootCmd.PersistentFlags().BoolVar(&root_opts.verbose, "verbose", false, "Verbose mode")
	RootCmd.PersistentFlags().StringVar(&root_opts.token, "mt-token", "", "MetaThings Token")
	viper.BindPFlag("MT_TOKEN", RootCmd.PersistentFlags().Lookup("mt-token"))
	RootCmd.PersistentFlags().StringVar(&root_opts.addr, "addr", "127.0.0.1:5000", "Metathings Service Address")

	RootCmd.AddCommand(tokenCmd)
}
