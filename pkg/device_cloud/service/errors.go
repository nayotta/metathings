package metathings_device_cloud_service

import "errors"

var (
	ErrModuleNotFound               = errors.New("module not found")
	ErrBadModuleEndpoint            = errors.New("bad module endpoint")
	ErrUnsupportedModuleProxyDriver = errors.New("unsupported module proxy driver")
)