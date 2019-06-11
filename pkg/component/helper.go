package metathings_component

import (
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

func ToModule(v **Module) func(string, interface{}) error {
	return func(key string, val interface{}) error {
		var ok bool

		if *v, ok = val.(*Module); !ok {
			return opt_helper.InvalidArgument(key)
		}

		return nil
	}
}
