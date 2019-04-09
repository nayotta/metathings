package metathings_component

import "errors"

var (
	ErrBadScheme              = errors.New("bad scheme")
	ErrBadServiceEndpoint     = errors.New("bad service endpoint")
	ErrDefaultAddressRequired = errors.New("default address required")
	ErrDeviceAddressRequired  = errors.New("device address required")
	ErrInvalidArguments       = errors.New("invalid arguments")
)
