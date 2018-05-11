package metathings_switcher_driver

import "errors"

var (
	ErrInitFail       = errors.New("initial fail")
	ErrDriverNotFound = errors.New("driver not found")
)
