package metathings_device_cloud_service

import "errors"

var (
	ErrModuleNotFound                    = errors.New("module not found")
	ErrBadModuleEndpoint                 = errors.New("bad module endpoint")
	ErrUnsupportedModuleProxyDriver      = errors.New("unsupported module proxy driver")
	ErrUnsupportedFlowChannelDriver      = errors.New("unsupported flow channel driver")
	ErrUnsupportedDeviceConnectionDriver = errors.New("unsupported device connection driver")
	ErrUnmatchedRequestId                = errors.New("unmatched request id")
)
