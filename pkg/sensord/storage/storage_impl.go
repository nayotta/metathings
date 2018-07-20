package metathings_sensord_storage

import (
	"time"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var schemas = `
CREATE TABLE IF NOT EXISTS sensor (
    id VARCHAR(255),
    name VARCHAR(255),
    core_id VARCHAR(255),
    entity_name VARCHAR(255),
    owner_id VARCHAR(255),
    application_credential_id VARCHAR(255),
    state VARCHAR(255),

    created_at DATETIME,
    updated_at DATETIME
);

CREATE TABLE IF NOT EXISTS sensor_tag (
    id VARCHAR(255),
    sensor_id VARCHAR(255),
    tag VARCHAR(255),

    created_at DATETIME
);
`

type storageImpl struct {
	logger log.FieldLogger
	db     *sqlx.DB
}

func (self *storageImpl) CreateSensor(snr Sensor) (Sensor, error) {
	var s Sensor

	now := time.Now()
	snr.CreatedAt = now
	snr.UpdatedAt = now
	_, err := self.db.NamedExec(`
INSERT INTO sensor(id, name, core_id, entity_name, owner_id, application_credential_id, state)
values (:id, :name, :core_id, :entity_name, :owner_id, :application_credential_id, :state)`, &snr)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to create sensor")
		return s, err
	}

	self.db.Get(&s, "SELECT * FROM sensor WHERE id=$1", *snr.Id)

	self.logger.WithField("snr_id", *snr.Id).Debugf("create sensor")

	return s, nil
}

func (self *storageImpl) DeleteSensor(snr_id string) error {
	tx, err := self.db.Begin()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to begin tx")
		return err
	}

	tx.Exec("DELETE FROM sensor_tag WHERE sensor_id=$1", snr_id)
	tx.Exec("DELETE FROM sensor WHERE id=$1", snr_id)
	err = tx.Commit()
	if err != nil {
		self.logger.WithError(err).WithField("snr_id", snr_id).Errorf("failed to delete sensor")
		return err
	}

	self.logger.WithField("snr_id", snr_id).Debugf("delete sensor")
	return nil
}

func (s *storageImpl) PatchSensor(snr_id string, snr Sensor) (Sensor, error) {
	panic("unimplemented")
}

func (self *storageImpl) GetSensor(snr_id string) (Sensor, error) {
	var s Sensor
	err := self.db.Get(&s, "SELECT * FROM sensor WHERE id=$1", snr_id)
	if err != nil {
		self.logger.WithError(err).WithField("snr_id", snr_id).Errorf("failed to get sensor")
		return s, err
	}

	err = self.db.Select(&s.Tags, "SELECT * FROM sensor_tag WHERE sensor_id=$1", snr_id)
	if err != nil {
		self.logger.WithError(err).WithField("snr_id", snr_id).Warningf("failed to get sensor tags")
	}

	return s, nil
}

func (s *storageImpl) ListSensors(Sensor) ([]Sensor, error) {
	panic("unimplemented")
}

func (s *storageImpl) ListSensorsForUser(owner_id string, snr Sensor) ([]Sensor, error) {
	panic("unimplemented")
}

func (s *storageImpl) GetSensorTags(snr_id string) ([]SensorTag, error) {
	panic("unimplemented")
}

func (s *storageImpl) AddSensorTag(snr_id, tag string) (SensorTag, error) {
	panic("unimplemented")
}

func (s *storageImpl) RemoveSensorTag(snr_id, tag string) error {
	panic("unimplemented")
}

func newStorageImpl(driver, uri string, logger log.FieldLogger) (*storageImpl, error) {
	if driver != "sqlite3" {
		logger.WithField("driver", driver).Errorf("not supported driver")
		return nil, ErrUnknownStorageDriver
	}
	db, err := sqlx.Connect(driver, uri)
	if err != nil {
		logger.WithFields(log.Fields{
			"driver": driver,
			"uri":    uri,
		}).WithError(err).Errorf("failed to connect database")
	}
	db.MustExec(schemas)
	return &storageImpl{
		logger: logger.WithField("#module", "storage"),
		db:     db,
	}, nil
}
