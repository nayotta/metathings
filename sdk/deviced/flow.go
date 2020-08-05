package metathings_deviced_sdk

import (
	"context"
	"sync"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type Flow interface {
	PushFrame(ctx context.Context, device, flow string, data interface{}) error
}

type FlowFactory func(...interface{}) (Flow, error)

var flow_factories_once sync.Once
var flow_factories map[string]FlowFactory

func register_flow_factory(name string, fty FlowFactory) {
	flow_factories_once.Do(func() {
		flow_factories = make(map[string]FlowFactory)
	})

	flow_factories[name] = fty
}

func NewFlow(name string, args ...interface{}) (Flow, error) {
	fty, ok := flow_factories[name]
	if !ok {
		return nil, ErrUnsupportedFlowFactory
	}

	return fty(args...)
}

func ToFlow(v *Flow) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		var ok bool

		if *v, ok = val.(Flow); !ok {
			return opt_helper.InvalidArgument(key)
		}

		return nil
	}
}
