package metathings_deviced_session_storage

import "errors"

var (
	ErrInvalidArgument             = errors.New("invalid argument")
	ErrUnknownSessionStorageDriver = errors.New("unknown session storage driver")
)
