package metathings_servo_service

import (
	"github.com/nayotta/viper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"

	driver_helper "github.com/nayotta/metathings/pkg/common/driver"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	driver "github.com/nayotta/metathings/pkg/servo/driver"
)

type Servo struct {
	Name   string
	Driver driver.ServoDriver
}

type ServoManager struct {
	drv_fty *driver_helper.DriverFactory

	servos map[string]Servo
}

func (mgr *ServoManager) ListServos() []Servo {
	srvs := make([]Servo, 0, len(mgr.servos))
	for k := range mgr.servos {
		srvs = append(srvs, mgr.servos[k])
	}
	return srvs
}

func (mgr *ServoManager) GetServo(name string) (Servo, error) {
	srv, ok := mgr.servos[name]
	if !ok {
		return Servo{}, ErrServoNotFound
	}
	return srv, nil
}

func NewServoManager(opt opt_helper.Option) (*ServoManager, error) {
	drv_fty, err := driver_helper.NewDriverFactory(opt.GetString("driver.descriptor"))
	if err != nil {
		return nil, err
	}

	srvs_opt := opt.Get("servos").([]interface{})

	srvs := map[string]Servo{}
	for _, srv_opt := range srvs_opt {
		v := viper.NewWithData(cast.ToStringMap(srv_opt))
		var srv_opt_s struct {
			Name   string
			Driver struct {
				Name string
			}
		}
		v.Unmarshal(&srv_opt_s)

		new_srv_opt := opt_helper.Copy(opt)
		new_srv_opt.Set("name", srv_opt_s.Name)
		new_srv_opt.Set("driver", v.Sub("driver"))
		new_srv_opt.Set("logger", opt.Get("logger").(log.FieldLogger).WithFields(log.Fields{
			"#driver": srv_opt_s.Driver.Name,
			"#name":   srv_opt_s.Name,
		}))
		drv, err := drv_fty.New(srv_opt_s.Driver.Name, new_srv_opt)
		if err != nil {
			return nil, err
		}
		srv_drv, ok := drv.(driver.ServoDriver)
		if !ok {
			return nil, driver_helper.ErrUnmatchDriver
		}

		err = srv_drv.Init(new_srv_opt)
		if err != nil {
			return nil, err
		}

		srvs[srv_opt_s.Name] = Servo{
			Name:   srv_opt_s.Name,
			Driver: srv_drv,
		}
	}

	return &ServoManager{
		drv_fty: drv_fty,
		servos:  srvs,
	}, nil
}
