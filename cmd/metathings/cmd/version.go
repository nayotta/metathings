package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	version_helper "github.com/nayotta/metathings/pkg/common/version"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Metathings Toolkit Version",
		Run:   cmd_helper.Run("version", version),
	}
)

func version() error {
	fmt.Println(version_helper.NewVersioner(version_str)().GetVersion())
	return nil
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
