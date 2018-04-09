package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "metathings",
		Short: "MetaThings Command Line Tools",
	}
)

func main() {
	rootCmd.AddCommand(identitydCmd)

	rootCmd.Execute()
}
