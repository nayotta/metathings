package main

import (
	"errors"
	"sync"

	log "github.com/sirupsen/logrus"

	driver_helper "github.com/nayotta/metathings/pkg/common/driver"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	driver "github.com/nayotta/metathings/pkg/sensor/driver"
)

var (
	logger = log.WithFields(log.Fields{
		"#driver": "dummy",
		"#module": "sensor",
	})
)

type dummySensorDriver struct {
	mutex *sync.Mutex
	name  string
	state driver.SensorState
}

func (drv *dummySensorDriver) Init(opt opt_helper.Option) error {
	return errors.New("unimplemented")
}

func (drv *dummySensorDriver) Close() error {
	return errors.New("unimplemented")
}

func (drv *dummySensorDriver) Show() driver.Sensor {
	return driver.Sensor{}
}

func (drv *dummySensorDriver) Data() driver.Sensor {
	return driver.Sensor{}
}

func (drv *dummySensorDriver) Config(cfg driver.SensorConfig) driver.Sensor {
	return driver.Sensor{}
}

var NewDriver driver_helper.NewDriverMethod = func(opt opt_helper.Option) (driver_helper.Driver, error) {
	log.Infof("new sensor dummy driver")

	return &dummySensorDriver{
		mutex: &sync.Mutex{},
		name:  opt.GetString("name"),
		state: driver.STATE_UNKNOWN,
	}, nil
}
