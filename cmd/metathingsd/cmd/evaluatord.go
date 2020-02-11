package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	service "github.com/nayotta/metathings/pkg/evaluatord/service"
	authorizer "github.com/nayotta/metathings/pkg/identityd2/authorizer"
	pb "github.com/nayotta/metathings/pkg/proto/evaluatord"
)

type EvaluatordOption struct {
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
}

func NewEvaluatordOption() *EvaluatordOption {
	return &EvaluatordOption{
		ServiceBaseOption: cmd_contrib.CreateServiceBaseOption(),
	}
}

var (
	evaluatord_opt *EvaluatordOption
)

var (
	evaluatordCmd = &cobra.Command{
		Use:   "evaluatord",
		Short: "Evaluator Service Daemon",
		PreRun: cmd_helper.DefaultPreRunHooks(func() {
			if base_opt.Config == "" {
				return
			}

			opt_t := NewEvaluatordOption()
			cmd_helper.UnmarshalConfig(opt_t)
			base_opt = &opt_t.BaseOption

			evaluatord_opt = opt_t
			evaluatord_opt.SetServiceName("evaluatord")
			evaluatord_opt.SetStage(cmd_helper.GetStageFromEnv())
		}),
		Run: cmd_helper.Run("evaluatord", runEvaluatord),
	}
)

func GetEvaluatordOptions() (
	*EvaluatordOption,
	cmd_contrib.ServiceOptioner,
	cmd_contrib.ListenOptioner,
	cmd_contrib.TransportCredentialOptioner,
	cmd_contrib.LoggerOptioner,
	cmd_contrib.ServiceEndpointsOptioner,
	cmd_contrib.CredentialOptioner,
	cmd_contrib.OpentracingOptioner,
) {
	return evaluatord_opt,
		evaluatord_opt,
		evaluatord_opt,
		evaluatord_opt,
		evaluatord_opt,
		evaluatord_opt,
		evaluatord_opt,
		evaluatord_opt
}

func NewMetathingsEvaulatordServiceOption(opt *EvaluatordOption) *service.MetathingsEvaluatordServiceOption {
	o := &service.MetathingsEvaluatordServiceOption{}

	return o
}

func runEvaluatord() error {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			GetEvaluatordOptions,
			cmd_contrib.NewServerTransportCredentials,
			cmd_contrib.NewLogger("evaluatord"),
			cmd_contrib.NewListener,
			cmd_contrib.NewOpentracing,
			cmd_contrib.NewGrpcServer,
			cmd_contrib.NewClientFactory,
			cmd_contrib.NewNoExpireTokener,
			token_helper.NewTokenValidator,
			NewMetathingsEvaulatordServiceOption,
			authorizer.NewAuthorizer,
			service.NewMetathingsEvaludatorService,
		),
		fx.Invoke(
			pb.RegisterEvaluatordServiceServer,
		),
	)

	if err := app.Start(context.Background()); err != nil {
		return err
	}
	defer app.Stop(context.Background())

	<-app.Done()
	if err := app.Err(); err != nil {
		return err
	}

	return nil
}

func init() {
	evaluatord_opt = NewEvaluatordOption()

	RootCmd.AddCommand(evaluatordCmd)
}
