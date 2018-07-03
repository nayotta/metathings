package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mtp "github.com/nayotta/metathings/pkg/cored/plugin"
)

var (
	rootCmd = &cobra.Command{
		Use: "servo",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {

		}),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runServod(); err != nil {
				log.WithError(err).Fatalf("failed to run servo(core) service")
			}
		},
	}
)

func runServod() error {
	return nil
}

type servoServicePlugin struct{}

func (p *servoServicePlugin) Run() error {
	return rootCmd.Execute()
}

func (p *servoServicePlugin) Init(opts opt_helper.Option) error {
	return nil
}

func NewServicePlugin() mtp.ServicePlugin {
	return &servoServicePlugin{}
}
