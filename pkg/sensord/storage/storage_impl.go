package metathings_sensord_storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
)

var (
	empty_sensor     = Sensor{}
	empty_sensor_tag = SensorTag{}
)

type storageImpl struct {
	db     *gorm.DB
	logger log.FieldLogger
}

func (self *storageImpl) get_sensor(snr_id string) (Sensor, error) {
	var snr Sensor
	err := self.db.Where("id = ?", snr_id).First(&snr).Error
	if err != nil {
		return empty_sensor, err
	}

	snr.Tags, err = self.get_sensor_tags_by_sensor_id(snr_id)
	if err != nil {
		return empty_sensor, err
	}

	return snr, nil
}

func (self *storageImpl) get_sensor_tags_by_sensor_id(snr_id string) ([]SensorTag, error) {
	var snr_tags []SensorTag
	err := self.db.Where("sensor_id = ?", snr_id).Find(&snr_tags).Error
	if err != nil {
		return nil, err
	}

	return snr_tags, nil
}

func (self *storageImpl) CreateSensor(snr Sensor) (Sensor, error) {
	err := self.db.Create(&snr).Error
	if err != nil {
		return empty_sensor, err
	}

	stm, err := self.get_sensor(*snr.Id)
	if err != nil {
		return empty_sensor, err
	}

	self.logger.WithField("id", *stm.Id).Debugf("create sensor")
	return stm, nil
}

func (self *storageImpl) DeleteSensor(snr_id string) error {
	tx := self.db.Begin()
	tx.Delete(&Sensor{}, "id = ?", snr_id)
	tx.Delete(&SensorTag{}, "sensor_id = ?", snr_id)
	err := tx.Commit().Error
	if err != nil {
		return err
	}

	self.logger.WithField("id", snr_id).Debugf("delete sensor")
	return nil
}

func (self *storageImpl) PatchSensor(snr_id string, snr Sensor) (Sensor, error) {
	var s Sensor

	if snr.Name != nil {
		s.Name = snr.Name
	}

	if snr.State != nil {
		s.State = snr.State
	}

	err := self.db.Model(&Sensor{}).Where("id = ?", snr_id).Updates(s).Error
	if err != nil {
		return empty_sensor, err
	}

	snr, err = self.get_sensor(snr_id)
	if err != nil {
		return empty_sensor, err
	}

	self.logger.WithField("id", snr_id).Debugf("patch sensor")
	return snr, nil
}

func (self *storageImpl) GetSensor(snr_id string) (Sensor, error) {
	snr, err := self.get_sensor(snr_id)
	if err != nil {
		return empty_sensor, err
	}

	self.logger.WithField("id", snr_id).Debugf("get sensor")
	return snr, nil
}

// TODO(Peer): poor performance
func (self *storageImpl) list_sensors(snr Sensor) ([]Sensor, error) {
	var snrs_t []Sensor
	err := self.db.Select("id").Find(&snrs_t, snr).Error
	if err != nil {
		return nil, err
	}

	var sensors []Sensor
	for _, s := range snrs_t {
		snr, err := self.get_sensor(*s.Id)
		if err != nil {
			return nil, err
		}
		sensors = append(sensors, snr)
	}

	return sensors, nil
}

func (self *storageImpl) ListSensors(snr Sensor) ([]Sensor, error) {
	sensors, err := self.list_sensors(snr)
	if err != nil {
		return nil, err
	}

	self.logger.Debugf("list sensors")
	return sensors, nil
}

func (self *storageImpl) ListSensorsForUser(owner_id string, snr Sensor) ([]Sensor, error) {
	snr.OwnerId = &owner_id
	sensors, err := self.list_sensors(snr)
	if err != nil {
		return nil, err
	}

	self.logger.Debugf("list sensors for user")
	return sensors, nil
}

func (self *storageImpl) GetSensorTags(snr_id string) ([]SensorTag, error) {
	tags, err := self.get_sensor_tags_by_sensor_id(snr_id)
	if err != nil {
		return nil, err
	}

	self.logger.WithField("sensor_id", snr_id).Debugf("list sensor tags")
	return tags, nil
}

func (self *storageImpl) AddSensorTag(tag SensorTag) (SensorTag, error) {
	err := self.db.Create(&tag).Error
	if err != nil {
		return empty_sensor_tag, err
	}

	self.logger.WithField("sensor_id", *tag.SensorId).Debugf("create sensor tag")
	return tag, nil
}

func (self *storageImpl) RemoveSensorTag(id string) error {
	err := self.db.Delete(&SensorTag{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	self.logger.WithField("id", id).Debugf("delete sensor tag")
	return nil
}

func newStorageImpl(driver, uri string, logger log.FieldLogger) (*storageImpl, error) {
	db, err := gorm.Open(driver, uri)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Sensor{})
	db.AutoMigrate(&SensorTag{})

	return &storageImpl{
		logger: logger.WithField("#module", "storage"),
		db:     db,
	}, nil
}
