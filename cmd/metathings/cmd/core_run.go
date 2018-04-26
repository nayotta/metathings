package cmd

import (
	"plugin"

	cmd_helper "github.com/bigdatagz/metathings/pkg/common/cmd"
	mt_plugin "github.com/bigdatagz/metathings/pkg/core/plugin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type _pluginOptions struct {
	Path string
}

type _coreRunOptions struct {
	_rootOptions `mapstructure:",squash"`
	Plugin       _pluginOptions
}

var (
	core_run_opts *_coreRunOptions
)

var (
	coreRunCmd = &cobra.Command{
		Use:   "run",
		Short: "Run a service in core runtime",
		PreRun: defaultPreRunHooks(func() {
			if root_opts.Config == "" {
				return
			}

			cmd_helper.UnmarshalConfig(core_run_opts)
			root_opts = &core_run_opts._rootOptions
			core_run_opts.Stage = cmd_helper.GetStageFromEnv()
		}),
		Run: func(cmd *cobra.Command, args []string) {
			if err := runCore(args); err != nil {
				log.WithError(err).Fatalf("failed to run service in core runtime")
			}
		},
	}
)

func runCore(args []string) error {
	lib, err := plugin.Open(core_run_opts.Plugin.Path)
	if err != nil {
		log.Fatalf("failed to open core service plugin %v: %v", core_run_opts.Plugin.Path, err)
	}

	NewPlugin, err := lib.Lookup("NewPlugin")
	if err != nil {
		log.Fatalf("failed to lookup NewPlugin method: %v", err)
	}

	opt := mt_plugin.Option{Args: args}
	p := NewPlugin.(func() mt_plugin.CorePlugin)()
	if err = p.Init(opt); err != nil {
		return err
	}

	return p.Run()
}

func init() {
	core_run_opts = &_coreRunOptions{}

	coreRunCmd.Flags().StringVarP(&core_run_opts.Plugin.Path, "plugin", "p", "", "Core plugin path")

	coreCmd.AddCommand(coreRunCmd)
}
