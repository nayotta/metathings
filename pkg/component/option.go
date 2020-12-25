package metathings_component

import (
	"github.com/stretchr/objx"
)

type NewModuleOption func(objx.Map)

func SetVersion(v string) NewModuleOption {
	return func(opt objx.Map) {
		if v != "" {
			opt.Set("version", v)
		}
	}
}

func SetArgs(vs []string) NewModuleOption {
	return func(opt objx.Map) {
		opt.Set("args", vs)
	}
}

func SetTarget(v interface{}) NewModuleOption {
	return func(opt objx.Map) {
		opt.Set("target", v)
	}
}
