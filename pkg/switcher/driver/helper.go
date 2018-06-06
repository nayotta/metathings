package metathings_switcher_driver

import (
	pb "github.com/nayotta/metathings/pkg/proto/switcher"
	state_helper "github.com/nayotta/metathings/pkg/switcher/state"
)

var _switcher_st_psr = state_helper.NewSwitcherStateParser()

func (s SwitcherState) ToString() string {
	return _switcher_st_psr.ToString(pb.SwitcherState(s))
}

func FromValue(x int32) SwitcherState {
	if x >= int32(STATE_OVERFLOW) || x < int32(STATE_UNKNOWN) {
		return STATE_UNKNOWN
	}
	return SwitcherState(x)
}
