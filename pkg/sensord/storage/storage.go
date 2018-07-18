package metathings_sensord_service

import (
	"errors"
	"time"

	log "github.com/sirupsen/logrus"
)

type Sensor struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	Id                      *string
	Name                    *string
	CoreId                  *string `db:"core_id"`
	EntityName              *string `db:"entity_name"`
	OwnerId                 *string `db:"owner_id"`
	ApplicationCredentialId *string `db:"application_credential_id"`
	State                   *string

	Tags []SensorTag `db:"-"`
}

type SensorTag struct {
	Id       *string
	SensorId *string `db:"sensor_id"`
	Tag      *string
}

type Storage interface {
	CreateSensor(Sensor) (Sensor, error)
	DeleteSensor(snr_id string) error
	PatchSensor(snr_id string, snr Sensor) (Sensor, error)
	GetSensor(snr_id string) (Sensor, error)
	ListSensor(Sensor) ([]Sensor, error)
	ListSensorForUser(owner_id string, snr Sensor) ([]Sensor, error)
	GetSensorTags(snr_id string) ([]SensorTag, error)
	AddSensorTag(snr_id, tag string) (SensorTag, error)
	RemoveSensorTag(snr_id, tag string) error
}

func NewStorage(driver, uri string, logger log.FieldLogger) (Storage, error) {
	return nil, errors.New("unimplemented")
}
