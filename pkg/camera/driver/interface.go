package metathings_camera_driver

import (
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type CameraState int32

const (
	STATE_UNKNOWN CameraState = iota
	STATE_STOP
	STATE_STARTING
	STATE_TERMINATING
	STATE_RUNNING
	STATE_OVERFLOW
)

type CameraConfig struct {
	Url       string
	Device    string
	Width     uint32
	Height    uint32
	Bitrate   uint32
	Framerate uint32
}

type Camera struct {
	State  CameraState
	Name   string
	Config CameraConfig
}

type CameraDriver interface {
	Init(opt_helper.Option) error
	Close() error
	Show() (Camera, error)
	Start(cfg CameraConfig) (Camera, error)
	Stop() (Camera, error)
}

type NotificationCenter interface {
	GetStateNotificationChannel() chan CameraState
	CloseStateNotificationChannel(chan CameraState)
}
