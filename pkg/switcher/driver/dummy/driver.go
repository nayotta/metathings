package main

import (
	"sync"

	log "github.com/sirupsen/logrus"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	driver "github.com/nayotta/metathings/pkg/switcher/driver"
)

type dummySwitcherDriver struct {
	mutex *sync.Mutex
	state driver.SwitcherState
}

func (drv *dummySwitcherDriver) Init(opt opt_helper.Option) error {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	drv.state = driver.OFF
	return nil
}

func (drv *dummySwitcherDriver) Close() error {
	return nil
}

func (drv *dummySwitcherDriver) Get() (driver.Switcher, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	return driver.Switcher{drv.state}, nil
}

func (drv *dummySwitcherDriver) Turn(x driver.SwitcherState) (driver.Switcher, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	drv.state = x
	log.WithFields(log.Fields{
		"state":   x.ToString(),
		"#driver": "dummy",
		"#module": "switcher",
	}).Infof("turn siwtcher state")
	return driver.Switcher{drv.state}, nil
}

var NewDriver driver.NewDriverMethod = func(opt opt_helper.Option) (driver.SwitcherDriver, error) {
	return &dummySwitcherDriver{
		mutex: &sync.Mutex{},
		state: driver.UNKNOWN,
	}, nil
}
