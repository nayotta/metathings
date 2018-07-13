package metathings_core_plugin

import opt_helper "github.com/nayotta/metathings/pkg/common/option"

func DefaultOptions() opt_helper.Option {
	return opt_helper.NewOption(
		"heartbeat.interval", 5,
	)

}
