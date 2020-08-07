package metathings_deviced_sdk

import (
	"context"
	"sync"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type Caller interface {
	UnaryCall(ctx context.Context, device, module, method string, arguments map[string]interface{}) (map[string]interface{}, error)
	// TODO(Peer): StreamCall
}

type CallerFactory func(...interface{}) (Caller, error)

var caller_factories_once sync.Once
var caller_factories map[string]CallerFactory

func register_caller_factory(name string, fty CallerFactory) {
	caller_factories_once.Do(func() {
		caller_factories = make(map[string]CallerFactory)
	})

	caller_factories[name] = fty
}

func NewCaller(name string, args ...interface{}) (Caller, error) {
	fty, ok := caller_factories[name]
	if !ok {
		return nil, ErrUnsupportedCallerFactory
	}

	return fty(args...)
}

func ToCaller(v *Caller) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		var ok bool

		if *v, ok = val.(Caller); !ok {
			return opt_helper.InvalidArgument(key)
		}

		return nil
	}
}
