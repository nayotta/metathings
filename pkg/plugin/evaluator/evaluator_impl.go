package metathings_plugin_evaluator

import (
	"context"

	cfg_helper "github.com/nayotta/metathings/pkg/common/config"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	esdk "github.com/nayotta/metathings/sdk/evaluatord"
	log "github.com/sirupsen/logrus"
)

type EvaluatorImplOption struct {
	Operator map[string]interface{}
}

type EvaluatorImpl struct {
	opt    *EvaluatorImplOption
	cfg    esdk.Data
	logger log.FieldLogger
}

func (e *EvaluatorImpl) get_config() esdk.Data {
	return e.cfg
}

func (e *EvaluatorImpl) get_logger() log.FieldLogger {
	return e.logger
}

func (e *EvaluatorImpl) Id() string {
	return e.cfg.Get("id").(string)
}

func (e *EvaluatorImpl) Eval(ctx context.Context, dat esdk.Data) error {
	logger := e.get_logger().WithField("evaluator", e.Id())

	op_drv, args, err := cfg_helper.ParseConfigOption("driver", e.opt.Operator, "logger", e.get_logger())
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
	_, err = op.Run(e.get_config(), dat)
	if err != nil {
		logger.WithError(err).Debugf("failed to run operator")
		return err
	}

	logger.Debugf("eval")

	return nil
}

func NewEvaluatorImpl(args ...interface{}) (*EvaluatorImpl, error) {
	var logger log.FieldLogger
	var config map[string]interface{}
	opt := &EvaluatorImplOption{}

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger":   opt_helper.ToLogger(&logger),
		"operator": opt_helper.ToStringMap(&opt.Operator),
		"config":   opt_helper.ToStringMap(&config),
	})(args...); err != nil {
		return nil, err
	}

	cfg, err := esdk.DataFromMap(config)
	if err != nil {
		return nil, err
	}

	evltr := &EvaluatorImpl{
		opt:    opt,
		cfg:    cfg,
		logger: logger,
	}

	return evltr, nil
}
