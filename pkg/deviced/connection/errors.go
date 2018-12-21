package metathings_deviced_connection

import "errors"

var (
	ErrInvalidArgument    = errors.New("invalid argument")
	ErrUnexpectedResponse = errors.New("unexpected response")
	ErrBridgeClosed       = errors.New("bridge closed")
)
