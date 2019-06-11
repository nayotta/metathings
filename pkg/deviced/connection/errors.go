package metathings_deviced_connection

import "errors"

var (
	ErrInvalidArgument          = errors.New("invalid argument")
	ErrChannelClosed            = errors.New("channel closed")
	ErrUnexpectedResponse       = errors.New("unexpected response")
	ErrBridgeClosed             = errors.New("bridge closed")
	ErrTimeout                  = errors.New("timeout")
	ErrDuplicatedDeviceInstance = errors.New("duplicated device instance")
	ErrDeviceOffline            = errors.New("device offline")
)
