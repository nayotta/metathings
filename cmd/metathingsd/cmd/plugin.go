package cmd

import (
	"github.com/spf13/cobra"
)

var (
	pluginCmd = &cobra.Command{
		Use:   "plugin",
		Short: "Metathings Service Plugin",
	}
)

func init() {
	RootCmd.AddCommand(pluginCmd)
}
