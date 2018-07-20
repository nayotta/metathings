package metathings_sensord_storage

import (
	"fmt"
	"strings"
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

func (self *storageImpl) PatchSensor(snr_id string, snr Sensor) (Sensor, error) {
	values := []string{}
	arguments := []interface{}{}
	i := 1
	s := Sensor{}

	if snr.Name != nil {
		values = append(values, fmt.Sprintf("name=$%v", i))
		arguments = append(arguments, *snr.Name)
		i++
	}

	if snr.State != nil {
		values = append(values, fmt.Sprintf("state=$%v", i))
		arguments = append(arguments, *snr.State)
		i++
	}

	if len(values) > 0 {
		values = append(values, fmt.Sprintf("updated_at=$%v", i))
		arguments = append(arguments, time.Now())
		i++

		val := strings.Join(values, ", ")
		arguments = append(arguments, snr_id)

		sql_str := "UPDATE sensor SET " + val + fmt.Sprintf(" WHERE id$%v", i)
		self.logger.WithFields(log.Fields{
			"sql":  sql_str,
			"args": arguments,
		}).Debugf("execute sql")
		_, err := self.db.Exec(sql_str, arguments...)
		if err != nil {
			self.logger.WithError(err).WithField("snr_id", snr_id).Errorf("failed to patch sensor")
			return s, err
		}
		s, _ = self.get_sensor(snr_id)
		self.logger.WithField("snr_id", snr_id).Debugf("update sensor")
		return s, nil
	}

	self.logger.WithField("snr_id", snr_id).Debugf("nothing changed when update sensor")
	return s, ErrNothingChanged
}

func (self *storageImpl) GetSensor(snr_id string) (Sensor, error) {
	var err error
	var s Sensor
	s, err = self.get_sensor(snr_id)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get sensor")
		return s, err
	}
	self.logger.WithField("snr_id", snr_id).Debugf("get sensor")
	return s, nil
}

func (self *storageImpl) get_sensor(snr_id string) (Sensor, error) {
	var s Sensor
	err := self.db.Get(&s, "SELECT * FROM sensor WHERE id=$1", snr_id)
	if err != nil {
		return s, err
	}

	tags, err := self.get_sensor_tags(snr_id)
	if err != nil {
		self.logger.WithError(err).WithField("snr_id", snr_id).Warningf("failed to get sensor tags")
	}
	s.Tags = tags

	return s, nil
}

func (self *storageImpl) ListSensors(snr Sensor) ([]Sensor, error) {
	ss, err := self.list_sensors(snr)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to list sensors")
		return nil, err
	}
	self.logger.Debugf("list sensors")
	return ss, nil
}

func (self *storageImpl) ListSensorsForUser(owner_id string, snr Sensor) ([]Sensor, error) {
	snr.OwnerId = &owner_id
	ss, err := self.list_sensors(snr)
	if err != nil {
		self.logger.WithField("owner_id", owner_id).WithError(err).Errorf("failed to list sensors for user")
		return nil, err
	}
	return ss, nil
}

func (self *storageImpl) list_sensors(sensor Sensor) ([]Sensor, error) {
	var err error
	values := []string{}
	arguments := []interface{}{}
	i := 0
	sensors := []Sensor{}

	if sensor.Name != nil {
		values = append(values, fmt.Sprintf("name=$%v", i))
		arguments = append(arguments, *sensor.Name)
		i++
	}

	if sensor.CoreId != nil {
		values = append(values, fmt.Sprintf("core_id=$%v", i))
		arguments = append(arguments, *sensor.CoreId)
		i++
	}

	if sensor.EntityName != nil {
		values = append(values, fmt.Sprintf("entity_name=$%v", i))
		arguments = append(arguments, *sensor.EntityName)
		i++
	}

	if sensor.OwnerId != nil {
		values = append(values, fmt.Sprintf("owner_id=$%v", i))
		arguments = append(arguments, *sensor.OwnerId)
		i++
	}

	if sensor.ApplicationCredentialId != nil {
		values = append(values, fmt.Sprintf("application_credential_id=$%v", i))
		arguments = append(arguments, *sensor.ApplicationCredentialId)
		i++
	}

	if sensor.State != nil {
		values = append(values, fmt.Sprintf("state=$%v", i))
		arguments = append(arguments, *sensor.State)
		i++
	}

	if len(values) == 0 {
		err = self.db.Select(&sensors, "SELECT * FROM sensor")
	} else {
		val := strings.Join(values, " and ")
		sql_str := fmt.Sprintf("SELECT * FROM sensor WHERE %v", val)
		self.logger.WithFields(log.Fields{
			"sql":  sql_str,
			"args": arguments,
		}).Debugf("execute sql")
		err = self.db.Select(&sensors, sql_str, arguments...)
	}
	if err != nil {
		return nil, err
	}
	return sensors, nil
}

func (self *storageImpl) get_sensor_tags(snr_id string) ([]SensorTag, error) {
	tags := []SensorTag{}

	err := self.db.Select(&tags, "SELECT * FROM sensor_tag WHERE sensor_id=$1", snr_id)
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (self *storageImpl) GetSensorTags(snr_id string) ([]SensorTag, error) {
	tags, err := self.get_sensor_tags(snr_id)
	if err != nil {
		self.logger.WithField("snr_id", snr_id).Errorf("failed to get sensor tags")
		return nil, err
	}
	return tags, nil
}

func (s *storageImpl) AddSensorTag(SensorTag) (SensorTag, error) {
	panic("unimplemented")
}

func (s *storageImpl) RemoveSensorTag(snr_tag_id string) error {
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
