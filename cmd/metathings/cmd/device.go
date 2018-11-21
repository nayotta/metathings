package cmd

import "github.com/spf13/cobra"

var (
	deviceCmd = &cobra.Command{
		Use:   "device",
		Short: "Device Toolkits",
	}
)

func init() {
	RootCmd.AddCommand(deviceCmd)
}
