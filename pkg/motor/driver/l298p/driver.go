package main

import (
	"strconv"
	"sync"

	"github.com/nayotta/viper"
	log "github.com/sirupsen/logrus"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"

	driver_helper "github.com/nayotta/metathings/pkg/common/driver"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	driver "github.com/nayotta/metathings/pkg/motor/driver"
)

type driverOption struct {
	Pins struct {
		Enable int
		Input1 int
		Input2 int
	}
}

type l298pMotorDriver struct {
	mutex     *sync.Mutex
	name      string
	state     driver.MotorState
	direction driver.MotorDirection
	speed     float32

	adaptor gobot.Connection
	pin_en  *gpio.DirectPinDriver
	pin_in1 *gpio.DirectPinDriver
	pin_in2 *gpio.DirectPinDriver

	logger log.FieldLogger
	opt    driverOption
}

func (drv *l298pMotorDriver) toDriverMotor() driver.Motor {
	return driver.Motor{
		Name:      drv.name,
		State:     drv.state,
		Direction: drv.direction,
		Speed:     drv.speed,
	}
}

func (drv *l298pMotorDriver) Init(opt opt_helper.Option) error {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	drv.state = driver.STATE_OFF
	drv.direction = driver.DIRECTION_FORWARD
	drv.speed = 0.0

	v, ok := opt.Get("driver").(*viper.Viper)
	if !ok {
		drv.logger.Debugf("failed to get driver")
		return driver.ErrInitFail
	}

	logger, ok := opt.Get("logger").(log.FieldLogger)
	if !ok {
		drv.logger.Debugf("failed to get logger")
		return driver.ErrInitFail
	}
	drv.logger = logger

	err := v.Unmarshal(&drv.opt)
	if err != nil {
		drv.logger.Debugf("failed to unmarshal data")
		return err
	}

	drv.initPins()

	drv.logger.Debugf("driver initialized")

	return nil
}

func (drv *l298pMotorDriver) Close() error {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	drv.logger.Debugf("driver closed")

	return nil
}

func (drv *l298pMotorDriver) Show() driver.Motor {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	return drv.toDriverMotor()
}

func (drv *l298pMotorDriver) Turn(x driver.MotorState) (driver.Motor, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	err := drv.turn(x)
	if err != nil {
		return driver.Motor{}, err
	}
	drv.state = x
	drv.logger.WithField("state", x.ToString()).Debugf("turn motor state")

	return drv.toDriverMotor(), nil
}

func (drv *l298pMotorDriver) turn(x driver.MotorState) error {
	if x == driver.STATE_OFF {
		if err := drv.setSpeed(0); err != nil {
			return err
		}
		if err := drv.setDirection(driver.DIRECTION_UNKNOWN); err != nil {
			return err
		}
	} else if x == driver.STATE_ON {
		if err := drv.setSpeed(drv.speed); err != nil {
			return err
		}
		if err := drv.setDirection(drv.direction); err != nil {
			return err
		}
	}

	return nil
}

func (drv *l298pMotorDriver) SetDirection(d driver.MotorDirection) (driver.Motor, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	if err := drv.setDirection(d); err != nil {
		return driver.Motor{}, err
	}
	drv.direction = d
	drv.logger.WithField("direction", d).Debugf("set motor direction")

	return drv.toDriverMotor(), nil
}

func (drv *l298pMotorDriver) setDirection(d driver.MotorDirection) error {
	if d == driver.DIRECTION_FORWARD {
		drv.pin_in1.On()
		drv.pin_in2.Off()
	} else if d == driver.DIRECTION_BACKWARD {
		drv.pin_in1.Off()
		drv.pin_in2.On()
	} else if d == driver.DIRECTION_UNKNOWN {
		drv.pin_in1.Off()
		drv.pin_in2.Off()
	}

	return nil
}

func (drv *l298pMotorDriver) SetSpeed(spd float32) (driver.Motor, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	if drv.state == driver.STATE_ON {
		drv.setSpeed(spd)
	}

	drv.speed = spd
	drv.logger.WithField("speed", spd).Debugf("set motor speed")

	return drv.toDriverMotor(), nil
}

func (drv *l298pMotorDriver) initPins() error {
	if drv.opt.Pins.Enable == 0 || drv.opt.Pins.Input1 == 0 || drv.opt.Pins.Input2 == 0 {
		return driver.ErrInitFail
	}

	drv.adaptor = raspi.NewAdaptor()
	drv.pin_en = gpio.NewDirectPinDriver(drv.adaptor, strconv.Itoa(drv.opt.Pins.Enable))
	drv.pin_in1 = gpio.NewDirectPinDriver(drv.adaptor, strconv.Itoa(drv.opt.Pins.Input1))
	drv.pin_in2 = gpio.NewDirectPinDriver(drv.adaptor, strconv.Itoa(drv.opt.Pins.Input2))

	drv.turn(drv.state)

	return nil
}

func (drv *l298pMotorDriver) setSpeed(spd float32) error {
	if spd < 0 {
		spd = 0
	} else if spd > 1 {
		spd = 1
	}

	drv.pin_en.PwmWrite(byte(spd * 254))
	return nil
}

var NewDriver driver_helper.NewDriverMethod = func(opt opt_helper.Option) (driver_helper.Driver, error) {
	log.Debugf("new motor dummy driver")

	return &l298pMotorDriver{
		mutex:     &sync.Mutex{},
		name:      opt.GetString("name"),
		state:     driver.STATE_UNKNOWN,
		direction: driver.DIRECTION_UNKNOWN,
		speed:     0.0,
	}, nil
}
