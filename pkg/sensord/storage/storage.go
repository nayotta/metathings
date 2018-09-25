package metathings_sensord_storage

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type Sensor struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	Name                    *string `gorm:"column:name"`
	CoreId                  *string `gorm:"column:core_id"`
	EntityName              *string `gorm:"column:entity_name"`
	OwnerId                 *string `gorm:"column:owner_id"`
	ApplicationCredentialId *string `gorm:"column:application_credential_id"`
	State                   *string `gorm:"column:state"`

	Tags []SensorTag `gorm:"-"`
}

type SensorTag struct {
	Id        *string
	CreatedAt time.Time
	UpdatedAt time.Time

	SensorId *string `gorm:"column:sensor_id"`
	Tag      *string `gorm:"column:tag"`
}

type Storage interface {
	CreateSensor(Sensor) (Sensor, error)
	DeleteSensor(snr_id string) error
	PatchSensor(snr_id string, snr Sensor) (Sensor, error)
	GetSensor(snr_id string) (Sensor, error)
	ListSensors(Sensor) ([]Sensor, error)
	ListSensorsForUser(owner_id string, snr Sensor) ([]Sensor, error)
	GetSensorTags(snr_id string) ([]SensorTag, error)
	AddSensorTag(SensorTag) (SensorTag, error)
	RemoveSensorTag(snr_tag_id string) error
}

func NewStorage(driver, uri string, logger log.FieldLogger) (Storage, error) {
	return newStorageImpl(driver, uri, logger)
}
