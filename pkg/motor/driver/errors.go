package metathings_motor_driver

import "errors"

var (
	ErrInitFail        = errors.New("initial fail")
	ErrDriverNotFound  = errors.New("driver not found")
	ErrInvalidArgument = errors.New("invalid argument")
)
