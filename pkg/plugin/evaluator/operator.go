package metathings_plugin_evaluator

import (
	"context"
	"sync"

	esdk "github.com/nayotta/metathings/sdk/evaluatord"
)

type Operator interface {
	Run(gctx context.Context, ctx, dat esdk.Data) (esdk.Data, error)
	Close() error
}

type OperatorFactory func(args ...interface{}) (Operator, error)

var operator_factories map[string]OperatorFactory
var operator_factories_once sync.Once

func registry_operator_factory(driver string, fty OperatorFactory) {
	operator_factories_once.Do(func() {
		operator_factories = make(map[string]OperatorFactory)
	})

	operator_factories[driver] = fty
}

func IsValidOperatorName(driver string) bool {
	for key, _ := range operator_factories {
		if key == driver {
			return true
		}
	}
	return false
}

func NewOperator(driver string, args ...interface{}) (Operator, error) {
	fty, ok := operator_factories[driver]
	if !ok {
		return nil, ErrUnsupportedOperatorName
	}

	return fty(args...)
}
