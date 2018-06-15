package metathings_core_plugin

import opt_helper "github.com/nayotta/metathings/pkg/common/option"

func DefaultOptions() opt_helper.Option {
	return opt_helper.Option{
		"heartbeat.interval": 5,
	}
}
