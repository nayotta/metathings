package cmd

import (
	"github.com/spf13/cobra"
)

var (
	tokenCmd = &cobra.Command{
		Use:   "token",
		Short: "Token Toolkits",
	}
)

func init() {
	RootCmd.AddCommand(tokenCmd)
}
