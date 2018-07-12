package metathings_sensor_driver

import (
	driver_helper "github.com/nayotta/metathings/pkg/common/driver"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type SensorConfig opt_helper.Option
type SensorData opt_helper.Option

type SensorState int32

const (
	STATE_UNKNOWN SensorState = iota
	STATE_ON
	STATE_OFF
	STATE_OVERFLOW
)

type Sensor struct {
	Name   string
	State  SensorState
	Config SensorConfig
}

type SensorDriver interface {
	driver_helper.Driver
	Show() Sensor
	Data() SensorData
	Config(SensorConfig) Sensor
}

type NewDriverMethod func(opt_helper.Option) (SensorDriver, error)
