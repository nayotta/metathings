package driver_helper

import "errors"

var (
	ErrDriverNotFound = errors.New("driver not found")
	ErrUnmatchDriver  = errors.New("unmatch driver")
)
