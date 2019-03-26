package client_helper

import "errors"

var (
	ErrInvalidArgument      = errors.New("invalid argument")
	ErrMissingDefaultConfig = errors.New("require default config")
)
