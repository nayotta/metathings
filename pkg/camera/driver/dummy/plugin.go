package main

import (
	"sync"

	log "github.com/sirupsen/logrus"

	driver "github.com/nayotta/metathings/pkg/camera/driver"
	driver_helper "github.com/nayotta/metathings/pkg/common/driver"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

var (
	logger = log.WithFields(log.Fields{})
)

type dummyCameraDriver struct {
	mutex  *sync.Mutex
	state  driver.CameraState
	config driver.CameraConfig
}

func (drv *dummyCameraDriver) show() driver.Camera {
	return driver.Camera{
		State:  drv.state,
		Config: drv.config,
	}
}

func (drv *dummyCameraDriver) Init(opt opt_helper.Option) error {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	drv.state = driver.STATE_STOP

	logger.Debugf("driver initialzed")

	return nil
}

func (drv *dummyCameraDriver) Close() error {
	logger.Debugf("driver closed")

	return nil
}

func (drv *dummyCameraDriver) Show() (driver.Camera, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	return drv.show(), nil
}

func (drv *dummyCameraDriver) Start(cfg driver.CameraConfig) (driver.Camera, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	if drv.state != driver.STATE_STOP {
		return drv.show(), driver.ErrUnstartable
	}

	drv.config = cfg
	drv.state = driver.STATE_STARTING
	go func() {
		drv.mutex.Lock()
		defer drv.mutex.Unlock()

		logger.WithFields(log.Fields{
			"device":    drv.config.Device,
			"url":       drv.config.Url,
			"width":     drv.config.Width,
			"height":    drv.config.Height,
			"framerate": drv.config.Framerate,
			"bitrate":   drv.config.Bitrate,
		}).Debugf("start camera")
		drv.state = driver.STATE_RUNNING
		logger.Debugf("camera is running")
	}()
	logger.Debugf("camera is starting")

	return drv.show(), nil
}

func (drv *dummyCameraDriver) Stop() (driver.Camera, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	if drv.state != driver.STATE_RUNNING {
		return drv.show(), driver.ErrUnstopable
	}

	drv.state = driver.STATE_TERMINATING
	go func() {
		drv.mutex.Lock()
		defer drv.mutex.Unlock()

		logger.Debugf("stop camera")
		drv.state = driver.STATE_STOP
		logger.Debugf("camera stopped")
	}()
	logger.Debugf("camera is terminating")

	return drv.show(), nil
}

var NewDriver driver_helper.NewDriverMethod = func(opt opt_helper.Option) (driver_helper.Driver, error) {
	logger.Debugf("new camera dummy driver")

	return &dummyCameraDriver{
		mutex: &sync.Mutex{},
		state: driver.STATE_UNKNOWN,
	}, nil
}
