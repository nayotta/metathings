package main

import (
	"sync"

	driver_helper "github.com/nayotta/metathings/pkg/common/driver"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	driver "github.com/nayotta/metathings/pkg/servo/driver"
)

type driverOption struct {
	I2c struct {
		Bus  int
		Addr int
		Idx  int
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
}

func (drv *i2cXiaoRServoDriver) Init(opt opt_helper.Option) error {
	return nil
}

func (drv *i2cXiaoRServoDriver) Close() error {
	return nil
}

func (drv *i2cXiaoRServoDriver) Show() driver.Servo {
	return driver.Servo{}
}

func (drv *i2cXiaoRServoDriver) Turn(x driver.ServoState) (driver.Servo, error) {
	return driver.Servo{}, nil
}

func (drv *i2cXiaoRServoDriver) SetAngle(ang float32) (driver.Servo, error) {
	return driver.Servo{}, nil
}

var NewDriver driver_helper.NewDriverMethod = func(opt opt_helper.Option) (driver_helper.Driver, error) {
	return &i2cXiaoRServoDriver{}, nil
}
