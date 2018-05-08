package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	helper "github.com/bigdatagz/metathings/pkg/common"
	cmd_helper "github.com/bigdatagz/metathings/pkg/common/cmd"
	opt_helper "github.com/bigdatagz/metathings/pkg/common/option"
	mt_plugin "github.com/bigdatagz/metathings/pkg/core/plugin"
)

type _coreRunOptions struct {
	_rootOptions                   `mapstructure:",squash"`
	mt_plugin.PluginCommandOptions `mapstructure:",squash"`
	ServiceDescriptorPath          string `mapstructure:"service_descriptor"`
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

			var opt _coreRunOptions
			cmd_helper.UnmarshalConfig(&opt)

			if opt.ServiceDescriptorPath == "" {
				opt.ServiceDescriptorPath = core_run_opts.ServiceDescriptorPath
			}

			if opt.Config == "" {
				opt.Config = root_opts.Config
			}

			core_run_opts = &opt
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
	path := helper.ExpendHomePath(core_run_opts.ServiceDescriptorPath)
	sd, err := mt_plugin.LoadServiceDescriptor(path)
	if err != nil {
		return err
	}

	plugin, err := sd.GetServicePlugin(core_run_opts.ServiceName)
	if err != nil {
		return err
	}

	// pass config to service plugin.
	args = append(args, "--config", root_opts.Config)
	err = plugin.Init(opt_helper.Option{"args": args})
	if err != nil {
		return err
	}

	return plugin.Run()
}

func init() {
	core_run_opts = &_coreRunOptions{}

	coreRunCmd.Flags().StringVarP(&core_run_opts.ServiceDescriptorPath, "service-descriptor-path", "p", "~/.metathings/service_descriptor.yaml", "Core Service Descriptor File Path")

	coreCmd.AddCommand(coreRunCmd)
}
