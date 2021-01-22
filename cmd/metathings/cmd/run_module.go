package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	fx_helper "github.com/nayotta/metathings/pkg/common/fx"
	component "github.com/nayotta/metathings/pkg/component"
)

type RunModuleOption struct {
	cmd_contrib.ClientBaseOption `mapstructure:",squash"`
	BinarySynchronizer           map[string]interface{}
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
		Short: "Run Module Service Daemon",
		Run:   cmd_helper.Run("run module", run_module),
	}
)

func NewMetathingsSodaModule() (*component.Module, error) {
	return component.NewModule(
		component.SetVersion(version_str),
		component.SetArgs(os.Args[3:]),
	)
}

func run_module() error {
	var err error
	var app *fx.App

	app = fx.New(
		fx.NopLogger,
		fx.Provide(
			fx_helper.NewFxAppGetter(&app),
			NewMetathingsSodaModule,
		),
		fx.Invoke(
			func(lc fx.Lifecycle, mdl *component.Module) error {
				lc.Append(fx.Hook{
					OnStart: func(context.Context) error {
						if err = mdl.Init(); err != nil {
							return err
						}

						go func() {
							mdl.Serve()
						}()

						return nil
					},
					OnStop: func(context.Context) error {
						mdl.Stop()

						return nil
					},
				})

				return nil
			},
		),
	)

	if err = app.Start(context.Background()); err != nil {
		return err
	}

	<-app.Done()
	if err = app.Err(); err != nil {
		return err
	}

	return nil
}

func init() {
	moduleCmd.AddCommand(runModuleCmd)
}
