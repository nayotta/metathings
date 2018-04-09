package cmd

import (
	"github.com/spf13/cobra"
)

var (
	root_opts struct {
		verbose bool
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

	RootCmd.AddCommand(identitydCmd)
}
