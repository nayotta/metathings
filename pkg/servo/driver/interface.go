package metathings_servo_driver

import (
	driver_helper "github.com/nayotta/metathings/pkg/common/driver"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

const (
	ANGLE_MIN float32 = 0
	ANGLE_MAX float32 = 1
)

type ServoState int32

const (
	STATE_UNKNOWN ServoState = iota
	STATE_ON
	STATE_OFF
	STATE_OVERFLOW
)

type Servo struct {
	Name     string
	State    ServoState
	Angle    float32
	MaxAngle float32
	MinAngle float32
}

type ServoDriver interface {
	driver_helper.Driver
	Show() Servo
	Turn(ServoState) (Servo, error)
	SetAngle(float32) (Servo, error)
}

type NewDriverMethod func(opt_helper.Option) (ServoDriver, error)
