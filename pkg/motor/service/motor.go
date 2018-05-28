package metathings_motor_service

import (
	"github.com/spf13/cast"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	driver "github.com/nayotta/metathings/pkg/motor/driver"
	"github.com/nayotta/viper"
)

type Motor struct {
	Name   string
	Driver driver.MotorDriver
}

type MotorManager struct {
	drv_fty *driver.DriverFactory

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
	drv_fty, err := driver.NewDriverFactory(opt.GetString("driver.descriptor"))
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

		new_mtr_opt := opt_helper.Option{
			"driver": v.Sub("driver"),
		}

		drv, err := drv_fty.New(mtr_opt_s.Driver.Name, new_mtr_opt)
		if err != nil {
			return nil, err
		}

		mtrs[mtr_opt_s.Name] = Motor{
			Name:   mtr_opt_s.Name,
			Driver: drv,
		}
	}

	return &MotorManager{
		drv_fty: drv_fty,
		motors:  mtrs,
	}, nil
}
