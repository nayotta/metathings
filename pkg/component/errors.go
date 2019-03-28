package metathings_component

import "errors"

var (
	ErrInvalidArguments  = errors.New("invalid arguments")
	ErrComponentNotFound = errors.New("component not found")
)
