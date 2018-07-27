package metathings_sensor_service

import (
	"time"

	"github.com/nayotta/viper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"

	driver_helper "github.com/nayotta/metathings/pkg/common/driver"
	"github.com/nayotta/metathings/pkg/common/emitter"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	driver "github.com/nayotta/metathings/pkg/sensor/driver"
)

type Sensor struct {
	Name   string
	Driver driver.SensorDriver
}

type SensorManager struct {
	logger      log.FieldLogger
	drv_fty     *driver_helper.DriverFactory
	emitter     emitter.Emitter
	publishable bool

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

type DataEvent struct {
	Name string
	Data driver.SensorData
}

func (mgr *SensorManager) DataEvent() chan DataEvent {
	ch := make(chan DataEvent)
	mgr.emitter.OnEvent(func(evt emitter.Event) error {
		data := evt.Data().(map[string]interface{})
		ch <- DataEvent{
			Name: data["name"].(string),
			Data: data["data"].(driver.SensorData),
		}
		return nil
	})
	return ch
}

func (mgr *SensorManager) Publishable() bool {
	return mgr.publishable
}

func (mgr *SensorManager) initTriggers(drv driver.SensorDriver, snr_v *viper.Viper) error {
	tgrs_i := snr_v.Get("triggers")
	if tgrs_i == nil {
		return nil
	}

	tgrs_opt, ok := tgrs_i.([]interface{})
	if !ok {
		return ErrInitialFailed
	}

	for _, tgr_opt := range tgrs_opt {
		tgr_v := viper.NewWithData(cast.ToStringMap(tgr_opt))
		err := mgr.initTrigger(drv, snr_v, tgr_v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (mgr *SensorManager) initTrigger(drv driver.SensorDriver, snr_v, tgr_v *viper.Viper) error {
	var err error
	name := tgr_v.GetString("name")

	switch name {
	case "period":
		err = mgr.initPeriodTrigger(drv, snr_v, tgr_v)
	case "change":
		err = mgr.initChangeTrigger(drv, snr_v, tgr_v)
	default:
		err = ErrUnsupportedTrigger
	}

	if err != nil {
		mgr.logger.WithField("trigger", name).WithError(err).Debugf("failed to init trigger")
		return err
	}

	mgr.publishable = true
	return nil
}

func (mgr *SensorManager) initPeriodTrigger(drv driver.SensorDriver, snr_v, tgr_v *viper.Viper) error {
	tvl := tgr_v.GetFloat64("interval")
	if tvl == 0 {
		tvl = 30
	}

	go func() {
		snr_name := snr_v.GetString("name")
		for {
			go func() {
				data := drv.Data()
				evt := emitter.NewEvent("period", map[string]interface{}{
					"name": snr_name,
					"data": data,
				})
				mgr.emitter.Trigger(evt)
			}()
			time.Sleep(time.Duration(tvl*1000) * time.Millisecond)
		}
	}()

	return nil
}

func (mgr *SensorManager) initChangeTrigger(drv driver.SensorDriver, snr_v, tgr_v *viper.Viper) error {
	chg, ok := drv.(driver.Changer)
	if !ok {
		return ErrUnsupportedTrigger
	}

	snr_name := snr_v.GetString("name")
	chg.OnChange(func(data driver.SensorData) error {
		evt := emitter.NewEvent("change", map[string]interface{}{
			"name": snr_name,
			"data": data,
		})
		mgr.emitter.Trigger(evt)
		return nil
	})

	return nil
}

func NewSensorManager(opt opt_helper.Option) (*SensorManager, error) {
	snr_mgr := &SensorManager{
		publishable: false,
	}

	logger := opt.Get("logger").(log.FieldLogger)
	snr_mgr.logger = logger.WithField("#module", "sensor_manager")

	drv_fty, err := driver_helper.NewDriverFactory(opt.GetString("driver.descriptor"))
	if err != nil {
		return nil, err
	}
	snr_mgr.drv_fty = drv_fty

	emt := emitter.NewEmitter()
	snr_mgr.emitter = emt

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
		err := v.Unmarshal(&snr_opt_s)
		if err != nil {
			return nil, err
		}

		new_snr_opt := opt_helper.Copy(opt)
		new_snr_opt.Set("name", snr_opt_s.Name)
		new_snr_opt.Set("driver", v.Sub("driver"))
		new_snr_opt.Set("logger", logger.WithFields(log.Fields{
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

		err = snr_mgr.initTriggers(snr_drv, v)
		if err != nil {
			return nil, err
		}

		snrs[snr_opt_s.Name] = Sensor{
			Name:   snr_opt_s.Name,
			Driver: snr_drv,
		}
	}
	snr_mgr.sensors = snrs

	return snr_mgr, nil
}
