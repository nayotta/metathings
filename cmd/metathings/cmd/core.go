package cmd

import (
	"github.com/spf13/cobra"
)

var (
	coreCmd = &cobra.Command{
		Use:   "core",
		Short: "Core Toolkits (deprecated)",
	}
)

func init() {
	RootCmd.AddCommand(coreCmd)
}
