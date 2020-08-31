package cmd

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	cmd_helper "github.com/nayotta/metathings/pkg/common/cmd"
	cfg_helper "github.com/nayotta/metathings/pkg/common/config"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	service "github.com/nayotta/metathings/pkg/evaluatord/service"
	storage "github.com/nayotta/metathings/pkg/evaluatord/storage"
	timer_backend "github.com/nayotta/metathings/pkg/evaluatord/timer"
	authorizer "github.com/nayotta/metathings/pkg/identityd2/authorizer"
	pb "github.com/nayotta/metathings/pkg/proto/evaluatord"
	dssdk "github.com/nayotta/metathings/sdk/data_storage"
)

type EvaluatordOption struct {
	cmd_contrib.ServiceBaseOption `mapstructure:",squash"`
	TaskStorage                   map[string]interface{}
	DataStorage                   map[string]interface{}
	TimerStorage                  map[string]interface{}
	TimerBackend                  map[string]interface{}
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

			cmd_helper.InitManyStringMapFromConfigWithStage([]cmd_helper.InitManyOption{
				{Dst: &opt_t.TaskStorage, Key: "task_storage"},
				{Dst: &opt_t.DataStorage, Key: "data_storage"},
				{Dst: &opt_t.TimerStorage, Key: "timer_storage"},
				{Dst: &opt_t.TimerBackend, Key: "timer_backend"},
			})

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
	cmd_contrib.StorageOptioner,
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
		evaluatord_opt,
		evaluatord_opt
}

func NewMetathingsEvaulatordServiceOption(opt *EvaluatordOption) *service.MetathingsEvaluatordServiceOption {
	var o service.MetathingsEvaluatordServiceOption

	o.Methods.QueryStorageByDevice.DefaultRangeFromDuration = -1 * time.Hour
	o.Methods.QueryStorageByDevice.DefaultPageSize = 50

	return &o
}

type NewEvaluatordStorageParams struct {
	fx.In

	Option cmd_contrib.StorageOptioner
	Logger logrus.FieldLogger
	Tracer opentracing.Tracer `name:"opentracing_tracer" optional:"true"`
}

func NewEvaluatordStorage(p NewEvaluatordStorageParams) (storage.Storage, error) {
	return storage.NewStorage(p.Option.GetDriver(), p.Option.GetUri(), "logger", p.Logger, "tracer", p.Tracer)
}

type NewEvaluatordTaskStorageParams struct {
	fx.In

	Option *EvaluatordOption
	Logger logrus.FieldLogger
	Tracer opentracing.Tracer `name:"opentracing_tracer" optional:"true"`
}

func NewEvaluatordTaskStorage(p NewEvaluatordTaskStorageParams) (storage.TaskStorage, error) {
	var drv string
	var args []interface{}
	var err error

	if drv, args, err = cfg_helper.ParseConfigOption("driver", p.Option.TaskStorage, "logger", p.Logger, "tracer", p.Tracer); err != nil {
		return nil, err
	}

	return storage.NewTaskStorage(drv, args...)
}

type NewEvaluatordDataStorageParams struct {
	fx.In

	Option *EvaluatordOption
	Logger logrus.FieldLogger
}

func NewEvaluatordDataStorage(p NewEvaluatordDataStorageParams) (dssdk.DataStorage, error) {
	name, args, err := cfg_helper.ParseConfigOption("name", p.Option.DataStorage, "logger", p.Logger)
	if err != nil {
		return nil, err
	}

	return dssdk.NewDataStorage(name, args...)
}

type NewEvaluatordTimerStorageParams struct {
	fx.In

	Option *EvaluatordOption
	Logger logrus.FieldLogger
	Tracer opentracing.Tracer `name:"opentracing_tracer" optional:"true"`
}

func NewEvaluatordTimerStorage(p NewEvaluatordTimerStorageParams) (storage.TimerStorage, error) {
	var drv string
	var args []interface{}
	var err error

	if drv, args, err = cfg_helper.ParseConfigOption("driver", p.Option.TimerStorage, "logger", p.Logger, "tracer", p.Tracer); err != nil {
		return nil, err
	}

	return storage.NewTimerStorage(drv, args...)
}

type NewEvaluatordTimerBackendParams struct {
	fx.In

	Option *EvaluatordOption
	Logger logrus.FieldLogger
}

func NewEvaluatordTimerBackend(p NewEvaluatordTimerBackendParams) (timer_backend.TimerBackend, error) {
	var drv string
	var args []interface{}
	var err error

	if drv, args, err = cfg_helper.ParseConfigOption("driver", p.Option.TimerBackend, "logger", p.Logger); err != nil {
		return nil, err
	}

	return timer_backend.NewTimerBackend(drv, args...)
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
			NewEvaluatordStorage,
			NewEvaluatordTaskStorage,
			NewEvaluatordDataStorage,
			NewEvaluatordTimerStorage,
			NewEvaluatordTimerBackend,
			authorizer.NewAuthorizer,
			cmd_contrib.NewValidator,
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
