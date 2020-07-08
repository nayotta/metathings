package metathings_deviced_storage

import "errors"

var (
	InvalidArgument = errors.New("invalid argument")
	RecordNotFound  = errors.New("record not found")
)
