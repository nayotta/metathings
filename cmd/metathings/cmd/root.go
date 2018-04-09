package cmd

import (
	"github.com/spf13/cobra"
)

var (
	root_opts struct {
		verbose bool
		addr    string
	}
)

var (
	RootCmd = &cobra.Command{
		Use:   "metathings",
		Short: "MetaThings Command Line Toolkits",
	}
)

func init() {
	RootCmd.PersistentFlags().BoolVar(&root_opts.verbose, "verbose", false, "Verbose mode")
	RootCmd.PersistentFlags().StringVar(&root_opts.addr, "addr", "127.0.0.1:5000", "Metathings Service Address")

	RootCmd.AddCommand(tokenCmd)
}
