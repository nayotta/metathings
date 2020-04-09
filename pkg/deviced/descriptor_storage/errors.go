package metathings_deviced_descriptor_storage

import "errors"

var (
	ErrUnknownDescriptorStorageDriver = errors.New("unknown descriptor storage driver")
	ErrDescriptorNotFound             = errors.New("descriptor not found")
)
