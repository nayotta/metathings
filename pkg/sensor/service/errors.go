package metathings_sensor_service

import "errors"

var (
	ErrInitialFailed      = errors.New("failed to initialed")
	ErrUnsupportedTrigger = errors.New("unsupported trigger")
	ErrSensorNotFound     = errors.New("sensor not found")
)
