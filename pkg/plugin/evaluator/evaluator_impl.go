package metathings_plugin_evaluator

import (
	"context"

	log "github.com/sirupsen/logrus"

	cfg_helper "github.com/nayotta/metathings/pkg/common/config"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	dssdk "github.com/nayotta/metathings/sdk/data_storage"
	esdk "github.com/nayotta/metathings/sdk/evaluatord"
)

type EvaluatorImplOption struct {
	Operator map[string]interface{}
}

type EvaluatorImpl struct {
	opt      *EvaluatorImplOption
	dat_stor dssdk.DataStorage
	info     esdk.Data
	cfg      esdk.Data
	logger   log.FieldLogger
}

func (e *EvaluatorImpl) get_config() esdk.Data {
	return e.cfg
}

func (e *EvaluatorImpl) get_eval_context() esdk.Data {
	ctx, _ := esdk.DataFromMap(map[string]interface{}{
		"config": e.cfg.Iter(),
	})
	return ctx
}

func (e *EvaluatorImpl) get_logger() log.FieldLogger {
	return e.logger
}

func (e *EvaluatorImpl) Id() string {
	return e.info.Get("id").(string)
}

func (e *EvaluatorImpl) Eval(ctx context.Context, dat esdk.Data) error {
	logger := e.get_logger().WithField("evaluator", e.Id())

	op_drv, args, err := cfg_helper.ParseConfigOption(
		"driver", e.opt.Operator,
		"logger", e.get_logger(),
		"data_storage", e.dat_stor,
	)
	if err != nil {
		logger.WithError(err).Debugf("failed to parse operator config option")
		return err
	}

	op, err := NewOperator(op_drv, args...)
	if err != nil {
		logger.WithError(err).Debugf("failed to new operator")
		return err
	}
	defer op.Close()

	// TODO(Peer): handle operator result
	_, err = op.Run(e.get_eval_context(), dat)
	if err != nil {
		logger.WithError(err).Debugf("failed to run operator")
		return err
	}

	return nil
}

func NewEvaluatorImpl(args ...interface{}) (*EvaluatorImpl, error) {
	var logger log.FieldLogger
	var config map[string]interface{}
	var info map[string]interface{}
	var ds dssdk.DataStorage
	opt := &EvaluatorImplOption{}

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger":       opt_helper.ToLogger(&logger),
		"info":         opt_helper.ToStringMap(&info),
		"operator":     opt_helper.ToStringMap(&opt.Operator),
		"config":       opt_helper.ToStringMap(&config),
		"data_storage": dssdk.ToDataStorage(&ds),
	}, opt_helper.SetSkip(true))(args...); err != nil {
		return nil, err
	}

	cfg, err := esdk.DataFromMap(config)
	if err != nil {
		return nil, err
	}

	inf, err := esdk.DataFromMap(info)
	if err != nil {
		return nil, err
	}

	evltr := &EvaluatorImpl{
		opt:    opt,
		info:   inf,
		cfg:    cfg,
		logger: logger,
	}

	return evltr, nil
}
