package cmd

import "github.com/spf13/cobra"

var (
	credentialCmd = &cobra.Command{
		Use:   "credential",
		Short: "Credential Toolkits",
	}
)

func init() {
	RootCmd.AddCommand(credentialCmd)
}
