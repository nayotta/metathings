package main

import (
	log "github.com/sirupsen/logrus"

	opt_helper "github.com/bigdatagz/metathings/pkg/common/option"
	driver "github.com/bigdatagz/metathings/pkg/switcher/driver"
)

type dummySwitcherDriver struct {
	state driver.SwitcherState
}

func (drv *dummySwitcherDriver) Init(opt opt_helper.Option) error {
	drv.state = driver.OFF
	return nil
}

func (drv *dummySwitcherDriver) Get() (driver.Switcher, error) {
	return driver.Switcher{drv.state}, nil
}

func (drv *dummySwitcherDriver) Turn(x driver.SwitcherState) (driver.Switcher, error) {
	drv.state = x
	log.WithFields(log.Fields{
		"state":   x.ToString(),
		"#driver": "dummy",
		"#module": "switcher",
	}).Infof("turn siwtcher state")
	return driver.Switcher{drv.state}, nil
}

func NewDriver(opt opt_helper.Option) (driver.SwitcherDriver, error) {
	return &dummySwitcherDriver{
		state: driver.UNKNOWN,
	}, nil
}
