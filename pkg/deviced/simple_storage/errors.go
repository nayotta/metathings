package metathings_deviced_simple_storage

import "errors"

var (
	ErrInvalidSimpleStorageDriver = errors.New("invalid simple storage driver")
	ErrObjectNotFound             = errors.New("object not found")
)
