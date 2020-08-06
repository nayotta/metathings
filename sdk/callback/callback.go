package metathings_callback_sdk

import (
	"sync"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type Callback interface {
	Emit(data interface{}) error
}

type CallbackFactory func(args ...interface{}) (Callback, error)

var callback_factories_once sync.Once
var callback_factories map[string]CallbackFactory

func register_callback_factory(name string, fty CallbackFactory) {
	callback_factories_once.Do(func() {
		callback_factories = make(map[string]CallbackFactory)
	})

	callback_factories[name] = fty
}

func NewCallback(name string, args ...interface{}) (Callback, error) {
	fty, ok := callback_factories[name]
	if !ok {
		return nil, ErrUnsupportedCallbackDriver
	}

	return fty(args...)
}

func ToCallback(v *Callback) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		var ok bool

		if *v, ok = val.(Callback); !ok {
			return opt_helper.InvalidArgument(key)
		}

		return nil
	}
}
