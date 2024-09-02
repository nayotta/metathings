package metathings_device_cloud_profile_storage_core

import "github.com/PeerXu/errors-go"

var (
	ErrProfileNotFound, ErrProfileNotFoundFn = errors.NewErrorAndErrorFunc[string]("profile not found")
	ErrDeviceNotFound, ErrDeviceNotFoundFn   = errors.NewErrorAndErrorFunc[string]("device not found")
)
