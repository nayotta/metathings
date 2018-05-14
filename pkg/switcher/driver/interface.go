package metathings_switcher_driver

import (
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type SwitcherState int32

const (
	UNKNOWN SwitcherState = iota
	ON
	OFF
	OVERFLOW
)

type Switcher struct {
	State SwitcherState
}

type SwitcherDriver interface {
	Init(opt_helper.Option) error
	Get() (Switcher, error)
	Turn(SwitcherState) (Switcher, error)
	Close() error
}

type NewDriverMethod func(opt_helper.Option) (SwitcherDriver, error)
