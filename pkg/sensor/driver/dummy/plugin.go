package main

import (
	"math/rand"
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
	mutex  *sync.Mutex
	name   string
	state  driver.SensorState
	config driver.SensorConfig
}

func (drv *dummySensorDriver) toDriverSensor() driver.Sensor {
	return driver.Sensor{
		Name:   drv.name,
		State:  drv.state,
		Config: drv.config,
	}
}

func (drv *dummySensorDriver) data() driver.SensorData {
	random := rand.Int31n(100)
	return driver.NewSensorData("random", random)
}

func (drv *dummySensorDriver) Init(opt opt_helper.Option) error {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	drv.state = driver.STATE_ON

	logger.Infof("driver initialized")

	return nil
}

func (drv *dummySensorDriver) Close() error {
	logger.Infof("driver closed")

	return nil
}

func (drv *dummySensorDriver) Show() driver.Sensor {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	return drv.toDriverSensor()
}

func (drv *dummySensorDriver) Data() driver.SensorData {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	data := drv.data()
	logger.WithField("random", data.GetInt("random")).Infof("data")

	return data
}

func (drv *dummySensorDriver) Config(cfg driver.SensorConfig) driver.Sensor {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	drv.config.Update(cfg)
	logger.WithField("keys", drv.config.Keys()).Infof("config")

	return drv.toDriverSensor()
}

var NewDriver driver_helper.NewDriverMethod = func(opt opt_helper.Option) (driver_helper.Driver, error) {
	log.Infof("new sensor dummy driver")

	return &dummySensorDriver{
		mutex: &sync.Mutex{},
		name:  opt.GetString("name"),
		state: driver.STATE_UNKNOWN,
	}, nil
}
