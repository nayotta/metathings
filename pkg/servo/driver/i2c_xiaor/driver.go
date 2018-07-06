package main

import (
	"errors"
	"sync"

	"github.com/go-daq/smbus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	driver_helper "github.com/nayotta/metathings/pkg/common/driver"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	driver "github.com/nayotta/metathings/pkg/servo/driver"
)

var (
	ErrAngleOutOfRange = errors.New("angle out of range")
)

type driverOption struct {
	I2c struct {
		Bus  int
		Addr uint8
		Idx  uint8
	}
	Angle struct {
		Min  float32
		Max  float32
		Init float32
	}
}

type i2cXiaoRServoDriver struct {
	mutex *sync.Mutex
	name  string
	state driver.ServoState
	angle float32

	logger log.FieldLogger
	opt    driverOption

	i2c_conn *smbus.Conn
}

func (drv *i2cXiaoRServoDriver) toDriverServo() driver.Servo {
	return driver.Servo{
		Name:     drv.name,
		State:    drv.state,
		Angle:    drv.angle,
		MinAngle: drv.opt.Angle.Min,
		MaxAngle: drv.opt.Angle.Max,
	}
}

func (drv *i2cXiaoRServoDriver) init() error {
	var err error

	if drv.opt.Angle.Min < 0 ||
		drv.opt.Angle.Min > 180 ||
		drv.opt.Angle.Max < 0 ||
		drv.opt.Angle.Max > 180 ||
		drv.opt.Angle.Min > drv.opt.Angle.Max {
		return ErrAngleOutOfRange
	}

	if drv.opt.Angle.Init < drv.opt.Angle.Min || drv.opt.Angle.Init > drv.opt.Angle.Max {
		drv.opt.Angle.Init = drv.opt.Angle.Min
		drv.logger.Warningf("init value out of range, change init value to min value")
	}

	drv.i2c_conn, err = smbus.Open(drv.opt.I2c.Bus, 0)
	if err != nil {
		return err
	}

	drv.setAngle(drv.opt.Angle.Init)
	drv.logger.WithField("angle", drv.opt.Angle.Init).Debugf("set angle to init value")

	return nil
}

func (drv *i2cXiaoRServoDriver) setAngle(ang float32) error {
	err := drv.i2c_conn.WriteReg(drv.opt.I2c.Addr, 0xff, drv.opt.I2c.Idx)
	if err != nil {
		return err
	}

	err = drv.i2c_conn.WriteReg(drv.opt.I2c.Addr, uint8(ang), 0xff)
	if err != nil {
		return err
	}

	drv.angle = ang

	return nil
}

func (drv *i2cXiaoRServoDriver) Init(opt opt_helper.Option) error {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	v, ok := opt.Get("driver").(*viper.Viper)
	if !ok {
		drv.logger.Debugf("failed to get driver")
		return driver_helper.ErrInitFail
	}

	logger, ok := opt.Get("logger").(log.FieldLogger)
	if !ok {
		drv.logger.Debugf("failed to get logger")
		return driver_helper.ErrInitFail
	}
	drv.logger = logger

	err := v.Unmarshal(&drv.opt)
	if err != nil {
		drv.logger.Debugf("failed to unmarshal data")
		return err
	}

	drv.init()

	drv.logger.WithFields(log.Fields{
		"i2c_bus":    drv.opt.I2c.Bus,
		"i2c_addr":   drv.opt.I2c.Addr,
		"i2c_idx":    drv.opt.I2c.Idx,
		"min_angle":  drv.opt.Angle.Min,
		"max_angle":  drv.opt.Angle.Max,
		"init_angle": drv.opt.Angle.Init,
	}).Debugf("driver initialized")
	return nil
}

func (drv *i2cXiaoRServoDriver) Close() error {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	drv.setAngle(drv.opt.Angle.Init)
	drv.i2c_conn.Close()

	drv.logger.Debugf("driver closed")

	return nil
}

func (drv *i2cXiaoRServoDriver) Show() driver.Servo {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	return drv.toDriverServo()
}

func (drv *i2cXiaoRServoDriver) Turn(x driver.ServoState) (driver.Servo, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	if x == driver.STATE_ON || x == driver.STATE_OFF {
		drv.state = x
	}

	drv.logger.WithField("state", x.ToString()).Debugf("turn servo state")

	return drv.toDriverServo(), nil
}

func (drv *i2cXiaoRServoDriver) SetAngle(ang float32) (driver.Servo, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	if ang < drv.opt.Angle.Min || ang > drv.opt.Angle.Max {
		return driver.Servo{}, ErrAngleOutOfRange
	}

	err := drv.setAngle(ang)
	if err != nil {
		return driver.Servo{}, err
	}
	drv.logger.WithField("angle", ang).Debugf("set angle")

	return drv.toDriverServo(), nil
}

var NewDriver driver_helper.NewDriverMethod = func(opt opt_helper.Option) (driver_helper.Driver, error) {
	log.Debugf("new servo i2c_xiaor driver")

	return &i2cXiaoRServoDriver{
		mutex: &sync.Mutex{},
		name:  opt.GetString("name"),
		state: driver.STATE_UNKNOWN,
		angle: 0.0,
	}, nil
}
