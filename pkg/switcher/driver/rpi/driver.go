package main

import (
	"sync"

	"github.com/nayotta/viper"
	log "github.com/sirupsen/logrus"
	rpio "github.com/stianeikeland/go-rpio"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	driver "github.com/nayotta/metathings/pkg/switcher/driver"
)

type driverOption struct {
	Model string
	Pin   int
}

type rpiSwitcherDriver struct {
	mutex *sync.Mutex
	state driver.SwitcherState

	logger log.FieldLogger
	opt    driverOption

	pin rpio.Pin
}

func (drv *rpiSwitcherDriver) Init(opt opt_helper.Option) error {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	drv.state = driver.STATE_OFF
	v, ok := opt.Get("driver").(*viper.Viper)
	if !ok {
		return driver.ErrInitFail
	}

	logger, ok := opt.Get("logger").(log.FieldLogger)
	if !ok {
		return driver.ErrInitFail
	}
	drv.logger = logger

	err := v.Unmarshal(&drv.opt)
	if err != nil {
		return err
	}

	err = rpio.Open()
	if err != nil {
		return err
	}

	if drv.opt.Model == "" {
		drv.opt.Model = "modern"
	}

	if drv.opt.Pin == 0 {
		drv.opt.Pin = 18
	}

	drv.pin, err = Pin(drv.opt.Model, drv.opt.Pin)
	if err != nil {
		return err
	}
	drv.pin.Output()
	drv.pin.Low()

	drv.logger.WithFields(log.Fields{
		"model": drv.opt.Model,
		"pin":   drv.opt.Pin,
	}).Debugf("pin initialized")

	return nil
}

func (drv *rpiSwitcherDriver) Close() error {
	rpio.Close()

	drv.logger.Debugf("close")

	return nil
}

func (drv *rpiSwitcherDriver) Get() (driver.Switcher, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	drv.logger.WithField("state", drv.state.ToString()).Debugf("get switcher state")
	return driver.Switcher{drv.state}, nil
}

func (drv *rpiSwitcherDriver) Turn(x driver.SwitcherState) (driver.Switcher, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	drv.state = x
	switch x {
	case driver.STATE_ON:
		drv.pin.High()
	case driver.STATE_OFF:
		drv.pin.Low()
	}

	drv.logger.WithField("state", drv.state.ToString()).Debugf("turn switcher state")
	return driver.Switcher{drv.state}, nil
}

var NewDriver driver.NewDriverMethod = func(opt opt_helper.Option) (driver.SwitcherDriver, error) {
	return &rpiSwitcherDriver{
		mutex: &sync.Mutex{},
		state: driver.STATE_UNKNOWN,
	}, nil
}
