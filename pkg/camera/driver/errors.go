package metathings_camera_driver

import "errors"

var (
	ErrUnstartable = errors.New("camera not startable")
	ErrUnstopable  = errors.New("camera not stopable")
)
