package metathings_motor_driver

import (
	driver_helper "github.com/nayotta/metathings/pkg/common/driver"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

const (
	SPEED_MAX float32 = 1.0
	SPEED_MIN float32 = 0.0
)

type MotorState int32

const (
	STATE_UNKNOWN MotorState = iota
	STATE_ON
	STATE_OFF
	STATE_OVERFLOW
)

type MotorDirection int32

const (
	DIRECTION_UNKNOWN MotorDirection = iota
	DIRECTION_FORWARD
	DIRECTION_BACKWARD
	DIRECTION_OVERFLOW
)

type Motor struct {
	Name      string
	State     MotorState
	Direction MotorDirection
	Speed     float32
}

type MotorDriver interface {
	driver_helper.Driver
	Show() Motor
	Turn(MotorState) (Motor, error)
	SetDirection(MotorDirection) (Motor, error)
	SetSpeed(float32) (Motor, error)
}

type NewDriverMethod func(opt_helper.Option) (MotorDriver, error)
