package metathings_plugin_evaluator

import (
	"context"

	log "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	cfg_helper "github.com/nayotta/metathings/pkg/common/config"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	dssdk "github.com/nayotta/metathings/sdk/data_storage"
	dsdk "github.com/nayotta/metathings/sdk/deviced"
	esdk "github.com/nayotta/metathings/sdk/evaluatord"
	smssdk "github.com/nayotta/metathings/sdk/sms"
)

type EvaluatorImplOption struct {
	Operator map[string]interface{}
}

type EvaluatorImpl struct {
	opt        *EvaluatorImplOption
	dat_stor   dssdk.DataStorage
	smpl_stor  dsdk.SimpleStorage
	flow       dsdk.Flow
	info       esdk.Data
	ctx        esdk.Data
	logger     log.FieldLogger
	caller     dsdk.Caller
	sms_sender smssdk.SmsSender
}

func (e *EvaluatorImpl) get_eval_context() esdk.Data {
	return e.ctx
}

func (e *EvaluatorImpl) get_logger() log.FieldLogger {
	return e.logger
}

func (e *EvaluatorImpl) Id() string {
	return e.info.Get("id").(string)
}

func (e *EvaluatorImpl) Eval(ctx context.Context, dat esdk.Data) (esdk.Data, error) {
	logger := e.get_logger().WithField("evaluator", e.Id())

	op_drv, args, err := cfg_helper.ParseConfigOption(
		"driver", e.opt.Operator,
		"logger", e.get_logger(),
		"data_storage", e.dat_stor,
		"simple_storage", e.smpl_stor,
		"flow", e.flow,
		"caller", e.caller,
		"sms_sender", e.sms_sender,
	)
	if err != nil {
		logger.WithError(err).Debugf("failed to parse operator config option")
		return nil, err
	}

	op, err := NewOperator(op_drv, args...)
	if err != nil {
		logger.WithError(err).Debugf("failed to new operator")
		return nil, err
	}
	defer op.Close()

	ret, err := op.Run(ctx, e.get_eval_context(), dat)
	if err != nil {
		logger.WithError(err).Debugf("failed to run operator")
		return nil, err
	}

	return ret, nil
}

func NewEvaluatorImpl(args ...interface{}) (*EvaluatorImpl, error) {
	var logger log.FieldLogger
	var context map[string]interface{}
	var info map[string]interface{}
	var ds dssdk.DataStorage
	var ss dsdk.SimpleStorage
	var flw dsdk.Flow
	var caller dsdk.Caller
	var sms_sender smssdk.SmsSender
	var cli_fty *client_helper.ClientFactory
	opt := &EvaluatorImplOption{}

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger":         opt_helper.ToLogger(&logger),
		"info":           opt_helper.ToStringMap(&info),
		"operator":       opt_helper.ToStringMap(&opt.Operator),
		"caller":         dsdk.ToCaller(&caller),
		"sms_sender":     smssdk.ToSmsSender(&sms_sender),
		"context":        opt_helper.ToStringMap(&context),
		"data_storage":   dssdk.ToDataStorage(&ds),
		"simple_storage": dsdk.ToSimpleStorage(&ss),
		"flow":           dsdk.ToFlow(&flw),
		"client_factory": client_helper.ToClientFactory(&cli_fty),
	}, opt_helper.SetSkip(true))(args...); err != nil {
		return nil, err
	}

	ctx, err := esdk.DataFromMap(context)
	if err != nil {
		return nil, err
	}

	inf, err := esdk.DataFromMap(info)
	if err != nil {
		return nil, err
	}

	evltr := &EvaluatorImpl{
		opt:        opt,
		info:       inf,
		ctx:        ctx,
		dat_stor:   ds,
		smpl_stor:  ss,
		flow:       flw,
		logger:     logger,
		caller:     caller,
		sms_sender: sms_sender,
	}

	return evltr, nil
}
