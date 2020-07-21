package metathings_component

import "github.com/stretchr/objx"

type NewModuleOption func(objx.Map)

func SetVersion(v string) NewModuleOption {
	return func(opt objx.Map) {
		if v != "" {
			opt.Set("version", v)
		}
	}
}
