package cmd

import (
	"net/url"
	"plugin"

	"github.com/spf13/cobra"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	component "github.com/nayotta/metathings/pkg/component"
)

type RunModuleOption struct {
	cmd_contrib.ClientBaseOption `mapstructure:",squash"`

	Libraries []string `mapstructure:"libraries"`
	Component string   `mapstructure:"component"`
}

func NewRunModuleOption() *RunModuleOption {
	return &RunModuleOption{
		ClientBaseOption: cmd_contrib.CreateClientBaseOption(),
	}
}

var (
	run_module_opt *RunModuleOption
)

var (
	runModuleCmd = &cobra.Command{
		Use:   "run",
		Short: "Run Module Daemon With Configure",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if base_opt.Config == "" {
				run_module_opt.BaseOption = *base_opt
				return
			}

			cmd_helper.UnmarshalConfig(run_module_opt)
			base_opt = &run_module_opt.BaseOption

			run_module_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: cmd_helper.RunWithArgs("run module", run_module),
	}
)

func load_component_from_url(uri string) (component.Component, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "file":
		return load_component_from_file(u.Host + u.Path)
	}

	return nil, component.ErrInvalidArguments
}

func load_component_from_file(path string) (component.Component, error) {
	p, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	sym, err := p.Lookup("NewComponent")
	if err != nil {
		return nil, err
	}

	cmp, err := (*(sym.(*component.NewComponent)))()
	if err != nil {
		return nil, err
	}

	return cmp, nil
}

func load_components(libraries []string) (map[string]component.Component, error) {
	components := map[string]component.Component{}

	for _, uri := range libraries {
		cmp, err := load_component_from_url(uri)
		if err != nil {
			return nil, err
		}
		components[cmp.Name()] = cmp
	}

	return components, nil
}

func run_module(args []string) error {
	components, err := load_components(run_module_opt.Libraries)
	if err != nil {
		return err
	}

	cmp, ok := components[run_module_opt.Component]
	if !ok {
		return component.ErrComponentNotFound
	}

	mdl, err := cmp.NewModule(args)

	if err = mdl.Start(); err != nil {
		return err
	}
	defer mdl.Stop()

	<-cmd_helper.Done()

	if err = mdl.Err(); err != nil {
		return err
	}

	return nil
}

func init() {
	run_module_opt = NewRunModuleOption()

	flags := runModuleCmd.Flags()

	flags.StringSliceVar(&run_module_opt.Libraries, "libraries", nil, "Component Libraries")
	flags.StringVar(&run_module_opt.Component, "component", "", "Run module by component")

	moduleCmd.AddCommand(runModuleCmd)
}
