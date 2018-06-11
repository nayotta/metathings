package driver_helper

import "errors"

var (
	ErrInitFail        = errors.New("initial failed")
	ErrDriverNotFound  = errors.New("driver not found")
	ErrUnmatchDriver   = errors.New("unmatch driver")
	ErrInvalidArgument = errors.New("invalid argument")
)
