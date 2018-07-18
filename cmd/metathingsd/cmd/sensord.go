package cmd

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
)

type _sensordOptions struct{}

var (
	sensordCmd = &cobra.Command{
		Use:   "sensord",
		Short: "Sensor Service Daemon",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
		}),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runSensord(); err != nil {
				log.WithError(err).Fatalf("failed to run sensord")
			}
		},
	}
)

func runSensord() error {
	return errors.New("unimplemented")
}

func init() {
	RootCmd.AddCommand(sensordCmd)
}
