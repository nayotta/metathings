package cmd

import "github.com/spf13/cobra"

var (
	descriptorCmd = &cobra.Command{
		Use:   "descriptor",
		Short: "Descriptor Toolkits",
	}
)

func init() {
	RootCmd.AddCommand(descriptorCmd)
}
