package main

import (
	"sync"

	log "github.com/sirupsen/logrus"

	driver_helper "github.com/nayotta/metathings/pkg/common/driver"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	driver "github.com/nayotta/metathings/pkg/motor/driver"
)

var (
	logger = log.WithFields(log.Fields{
		"#driver": "dummy",
		"#module": "motor",
	})
)

type dummyMotorDriver struct {
	mutex     *sync.Mutex
	name      string
	state     driver.MotorState
	direction driver.MotorDirection
	speed     float32
}

func (drv *dummyMotorDriver) toDriverMotor() driver.Motor {
	return driver.Motor{
		Name:      drv.name,
		State:     drv.state,
		Direction: drv.direction,
		Speed:     drv.speed,
	}
}

func (drv *dummyMotorDriver) Init(opt opt_helper.Option) error {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	drv.state = driver.STATE_OFF
	drv.direction = driver.DIRECTION_FORWARD
	drv.speed = 0.0

	logger.Infof("driver initialized")

	return nil
}

func (drv *dummyMotorDriver) Close() error {
	logger.Infof("driver closed")

	return nil
}

func (drv *dummyMotorDriver) Show() driver.Motor {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	return drv.toDriverMotor()
}

func (drv *dummyMotorDriver) Turn(x driver.MotorState) (driver.Motor, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	drv.state = x
	logger.WithField("state", x.ToString()).Infof("turn motor state")

	return drv.toDriverMotor(), nil
}

func (drv *dummyMotorDriver) SetDirection(d driver.MotorDirection) (driver.Motor, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	drv.direction = d
	logger.WithField("direction", d).Infof("set motor direction")

	return drv.toDriverMotor(), nil
}

func (drv *dummyMotorDriver) SetSpeed(spd float32) (driver.Motor, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	if !driver.IsValidSpeed(spd) {
		return driver.Motor{}, driver_helper.ErrInvalidArgument
	}

	drv.speed = spd
	logger.WithField("speed", spd).Infof("set motor speed")

	return drv.toDriverMotor(), nil
}

var NewDriver driver_helper.NewDriverMethod = func(opt opt_helper.Option) (driver_helper.Driver, error) {
	log.Infof("new motor dummy driver")

	return &dummyMotorDriver{
		mutex:     &sync.Mutex{},
		name:      opt.GetString("name"),
		direction: driver.DIRECTION_UNKNOWN,
		state:     driver.STATE_UNKNOWN,
		speed:     0.0,
	}, nil
}
