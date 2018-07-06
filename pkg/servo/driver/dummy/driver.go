package main

import (
	"sync"

	log "github.com/sirupsen/logrus"

	driver_helper "github.com/nayotta/metathings/pkg/common/driver"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	driver "github.com/nayotta/metathings/pkg/servo/driver"
)

var (
	logger = log.WithFields(log.Fields{
		"#driver": "dummy",
		"#module": "motor",
	})
)

type dummyServoDriver struct {
	mutex *sync.Mutex
	name  string
	state driver.ServoState
	angle float32
}

func (drv *dummyServoDriver) toDriverServo() driver.Servo {
	return driver.Servo{
		Name:     drv.name,
		State:    drv.state,
		Angle:    drv.angle,
		MinAngle: 0,
		MaxAngle: 180,
	}
}

func (drv *dummyServoDriver) Init(opt opt_helper.Option) error {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	drv.state = driver.STATE_OFF
	drv.angle = 0

	logger.Infof("driver initialized")

	return nil
}

func (drv *dummyServoDriver) Close() error {
	logger.Infof("driver closed")

	return nil
}

func (drv *dummyServoDriver) Show() driver.Servo {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	return drv.toDriverServo()
}

func (drv *dummyServoDriver) Turn(x driver.ServoState) (driver.Servo, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	drv.state = x
	logger.WithField("state", x.ToString()).Infof("turn motor state")

	return drv.toDriverServo(), nil
}

func (drv *dummyServoDriver) SetAngle(ang float32) (driver.Servo, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	if !driver.IsValidAngle(ang) {
		return driver.Servo{}, driver_helper.ErrInvalidArgument
	}

	drv.angle = ang
	logger.WithField("angle", ang).Infof("set servo angle")

	return drv.toDriverServo(), nil
}

var NewDriver driver_helper.NewDriverMethod = func(opt opt_helper.Option) (driver_helper.Driver, error) {
	log.Infof("new srvo dummy driver")

	return &dummyServoDriver{
		mutex: &sync.Mutex{},
		name:  opt.GetString("name"),
		state: driver.STATE_UNKNOWN,
		angle: 0,
	}, nil
}
