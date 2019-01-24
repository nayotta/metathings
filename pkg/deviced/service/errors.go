package metathings_deviced_service

import "errors"

var (
	ErrDuplicatedDevice  = errors.New("duplicated device")
	ErrUnconnectedDevice = errors.New("unconnected device")
)
