package metathings_plugin_evaluator

import (
	"context"

	esdk "github.com/nayotta/metathings/sdk/evaluatord"
)

type Evaluator interface {
	Id() string
	Eval(context.Context, esdk.Data) error
}

func NewEvaluator(args ...interface{}) (Evaluator, error) {
	return NewEvaluatorImpl(args...)
}
