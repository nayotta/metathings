package metathings_switcher_driver

import (
	opt_helper "github.com/bigdatagz/metathings/pkg/common/option"
	state_helper "github.com/bigdatagz/metathings/pkg/switcher/state"
)

var _switcher_st_psr = state_helper.NewSwitcherStateParser()

type SwitcherState int32

const (
	UNKNOWN SwitcherState = iota
	ON
	OFF
)

type Switcher struct {
	State SwitcherState
}

type SwitcherDriver interface {
	Init(opt_helper.Option) error
	Get() (Switcher, error)
	Turn(SwitcherState) (Switcher, error)
}
