package metathings_motor_service

import (
	"github.com/nayotta/viper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"

	driver_helper "github.com/nayotta/metathings/pkg/common/driver"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	driver "github.com/nayotta/metathings/pkg/motor/driver"
)

type Motor struct {
	Name   string
	Driver driver.MotorDriver
}

type MotorManager struct {
	drv_fty *driver_helper.DriverFactory

	motors map[string]Motor
}

func (mgr *MotorManager) ListMotors() []Motor {
	mtrs := make([]Motor, 0, len(mgr.motors))
	for k := range mgr.motors {
		mtrs = append(mtrs, mgr.motors[k])
	}
	return mtrs
}

func (mgr *MotorManager) GetMotor(name string) (Motor, error) {
	mtr, ok := mgr.motors[name]
	if !ok {
		return Motor{}, ErrMotorNotFound
	}
	return mtr, nil
}

func NewMotorManager(opt opt_helper.Option) (*MotorManager, error) {
	drv_fty, err := driver_helper.NewDriverFactory(opt.GetString("driver.descriptor"))
	if err != nil {
		return nil, err
	}

	mtrs_opt := opt.Get("motors").([]interface{})

	mtrs := map[string]Motor{}
	for _, mtr_opt := range mtrs_opt {
		v := viper.NewWithData(cast.ToStringMap(mtr_opt))
		var mtr_opt_s struct {
			Name   string
			Driver struct {
				Name string
			}
		}
		v.Unmarshal(&mtr_opt_s)

		new_mtr_opt := opt_helper.Copy(opt)
		new_mtr_opt.Set("name", mtr_opt_s.Name)
		new_mtr_opt.Set("driver", v.Sub("driver"))
		new_mtr_opt.Set("logger", opt.Get("logger").(log.FieldLogger).WithFields(log.Fields{
			"#driver": mtr_opt_s.Driver.Name,
			"#name":   mtr_opt_s.Name,
		}))
		drv, err := drv_fty.New(mtr_opt_s.Driver.Name, new_mtr_opt)
		if err != nil {
			return nil, err
		}
		mtr_drv, ok := drv.(driver.MotorDriver)
		if !ok {
			return nil, driver_helper.ErrUnmatchDriver
		}

		err = mtr_drv.Init(new_mtr_opt)
		if err != nil {
			return nil, err
		}

		mtrs[mtr_opt_s.Name] = Motor{
			Name:   mtr_opt_s.Name,
			Driver: mtr_drv,
		}
	}

	return &MotorManager{
		drv_fty: drv_fty,
		motors:  mtrs,
	}, nil
}
