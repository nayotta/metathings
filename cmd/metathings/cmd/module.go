package cmd

import "github.com/spf13/cobra"

var (
	moduleCmd = &cobra.Command{
		Use:   "module",
		Short: "Module Toolkits",
	}
)

func init() {
	RootCmd.AddCommand(moduleCmd)
}
