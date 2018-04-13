package cmd

import (
	"plugin"

	mt_plugin "github.com/bigdatagz/metathings/pkg/core/plugin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	core_run_opts struct {
		config string
		plugin string
	}
)

var (
	coreRunCmd = &cobra.Command{
		Use:    "run",
		Short:  "Run a service in core runtime",
		PreRun: globalPreRunHook,
		Run: func(cmd *cobra.Command, args []string) {
			if err := runCore(); err != nil {
				log.Fatalf("failed to run service in core runtime: %v", err)
			}
		},
	}
)

func runCore() error {
	lib, err := plugin.Open(core_run_opts.plugin)
	if err != nil {
		log.Fatalf("failed to open core service plugin %v: %v", core_run_opts.plugin, err)
	}

	NewPlugin, err := lib.Lookup("NewPlugin")
	if err != nil {
		log.Fatalf("failed to lookup NewPlugin method: %v", err)
	}

	opt := mt_plugin.Option{Config: core_run_opts.config}
	p := NewPlugin.(func() mt_plugin.CorePlugin)()
	if err = p.Init(opt); err != nil {
		return err
	}

	return p.Run()
}

func init() {
	coreRunCmd.Flags().StringVarP(&core_run_opts.plugin, "plugin", "p", "", "Core plugin path")
	coreRunCmd.MarkFlagRequired("plugin")
	coreRunCmd.Flags().StringVarP(&core_run_opts.config, "config", "c", "", "Core plugin config path")
	coreRunCmd.MarkFlagRequired("config")

	coreCmd.AddCommand(coreRunCmd)
}
