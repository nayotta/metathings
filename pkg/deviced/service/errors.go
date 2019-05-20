package metathings_deviced_service

import "errors"

var (
	ErrUnexpectedMessage = errors.New("unexpected message")
	ErrDuplicatedDevice  = errors.New("duplicated device")
	ErrUnconnectedDevice = errors.New("unconnected device")
	ErrFlowNotFound      = errors.New("flow not found")
	ErrDeviceOffline     = errors.New("device offline")
)
