package main

import (
	"sync"

	log "github.com/sirupsen/logrus"

	driver_helper "github.com/nayotta/metathings/pkg/common/driver"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	driver "github.com/nayotta/metathings/pkg/switcher/driver"
)

var (
	logger = log.WithFields(log.Fields{
		"#driver": "dummy",
		"#module": "switcher",
	})
)

type dummySwitcherDriver struct {
	mutex *sync.Mutex
	state driver.SwitcherState
}

func (drv *dummySwitcherDriver) Init(opt opt_helper.Option) error {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	drv.state = driver.STATE_OFF

	logger.Debugf("driver initialized")

	return nil
}

func (drv *dummySwitcherDriver) Close() error {
	logger.Debugf("driver closed")

	return nil
}

func (drv *dummySwitcherDriver) Show() (driver.Switcher, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	return driver.Switcher{drv.state}, nil
}

func (drv *dummySwitcherDriver) Turn(x driver.SwitcherState) (driver.Switcher, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	drv.state = x
	logger.WithField("state", x.ToString()).Debugf("turn siwtcher state")
	return driver.Switcher{drv.state}, nil
}

var NewDriver driver_helper.NewDriverMethod = func(opt opt_helper.Option) (driver_helper.Driver, error) {
	logger.Infof("new switcher dummy driver")

	return &dummySwitcherDriver{
		mutex: &sync.Mutex{},
		state: driver.STATE_UNKNOWN,
	}, nil
}
