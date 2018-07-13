package metathings_sensor_service

import (
	"github.com/nayotta/viper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"

	driver_helper "github.com/nayotta/metathings/pkg/common/driver"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	driver "github.com/nayotta/metathings/pkg/sensor/driver"
)

type Sensor struct {
	Name   string
	Driver driver.SensorDriver
}

type SensorManager struct {
	drv_fty *driver_helper.DriverFactory

	sensors map[string]Sensor
}

func (mgr *SensorManager) ListSensors() []Sensor {
	snrs := make([]Sensor, 0, len(mgr.sensors))
	for k := range mgr.sensors {
		snrs = append(snrs, mgr.sensors[k])
	}
	return snrs
}

func (mgr *SensorManager) GetSensor(name string) (Sensor, error) {
	snr, ok := mgr.sensors[name]
	if !ok {
		return Sensor{}, ErrSensorNotFound
	}
	return snr, nil
}

func NewSensorManager(opt opt_helper.Option) (*SensorManager, error) {
	drv_fty, err := driver_helper.NewDriverFactory(opt.GetString("driver.descriptor"))
	if err != nil {
		return nil, err
	}

	snrs_opt := opt.Get("sensors").([]interface{})

	snrs := map[string]Sensor{}
	for _, snr_opt := range snrs_opt {
		v := viper.NewWithData(cast.ToStringMap(snr_opt))
		var snr_opt_s struct {
			Name   string
			Driver struct {
				Name string
			}
		}
		v.Unmarshal(&snr_opt_s)

		new_snr_opt := opt_helper.Copy(opt)
		new_snr_opt.Set("name", snr_opt_s.Name)
		new_snr_opt.Set("driver", v.Sub("driver"))
		new_snr_opt.Set("logger", opt.Get("logger").(log.FieldLogger).WithFields(log.Fields{
			"#driver": snr_opt_s.Driver.Name,
			"#name":   snr_opt_s.Name,
		}))
		drv, err := drv_fty.New(snr_opt_s.Driver.Name, new_snr_opt)
		if err != nil {
			return nil, err
		}
		snr_drv, ok := drv.(driver.SensorDriver)
		if !ok {
			return nil, driver_helper.ErrUnmatchDriver
		}
		err = snr_drv.Init(new_snr_opt)
		if err != nil {
			return nil, err
		}
		snrs[snr_opt_s.Name] = Sensor{
			Name:   snr_opt_s.Name,
			Driver: snr_drv,
		}
	}

	return &SensorManager{
		drv_fty: drv_fty,
		sensors: snrs,
	}, nil
}
