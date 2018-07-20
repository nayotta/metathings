package metathings_sensord_storage

import (
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

var (
	test_sensor_id                 = "test-sensor-id"
	test_name                      = "test-name"
	test_core_id                   = "test-core-id"
	test_entity_name               = "test-entity-name"
	test_owner_id                  = "test-owner-id"
	test_application_credential_id = "test-application-credential-id"
	test_state                     = "test-state"
	test_tag0_id                   = "test-tag0-id"
	test_tag1_id                   = "test-tag1-id"
)

type storageImplTestSuite struct {
	suite.Suite
	s *storageImpl
}

func (suite *storageImplTestSuite) SetupTest() {
	suite.s, _ = newStorageImpl("sqlite3", ":memory:", logrus.New())

	s := Sensor{
		Id:                      &test_sensor_id,
		Name:                    &test_name,
		CoreId:                  &test_core_id,
		EntityName:              &test_entity_name,
		OwnerId:                 &test_owner_id,
		ApplicationCredentialId: &test_application_credential_id,
		State: &test_state,
	}
	suite.s.CreateSensor(s)

	tag0 := SensorTag{
		Id:       &test_tag0_id,
		SensorId: &test_sensor_id,
		Tag:      &test_tag0_id,
	}
	suite.s.AddSensorTag(tag0)

	tag1 := SensorTag{
		Id:       &test_tag1_id,
		SensorId: &test_sensor_id,
		Tag:      &test_tag1_id,
	}
	suite.s.AddSensorTag(tag1)
}

func (suite *storageImplTestSuite) TestGetSensor() {
	snr, err := suite.s.GetSensor(test_sensor_id)
	suite.Nil(err)
	suite.Equal(test_sensor_id, *snr.Id)
	suite.Equal(test_name, *snr.Name)
	suite.Equal(test_core_id, *snr.CoreId)
	suite.Equal(test_entity_name, *snr.EntityName)
	suite.Equal(test_owner_id, *snr.OwnerId)
	suite.Equal(test_application_credential_id, *snr.ApplicationCredentialId)
	suite.Equal(test_state, *snr.State)
	suite.NotEqual("0001-01-01 00:00:00 +0000 UTC", snr.CreatedAt)
	suite.NotEqual("0001-01-01 00:00:00 +0000 UTC", snr.UpdatedAt)
}

func (suite *storageImplTestSuite) TestListSensors() {
	snrs, err := suite.s.ListSensors(Sensor{})
	suite.Nil(err)
	suite.Len(snrs, 1)
}

func (suite *storageImplTestSuite) TestListSensorsForUser() {
	snrs, err := suite.s.ListSensorsForUser(test_owner_id, Sensor{})
	suite.Nil(err)
	suite.Len(snrs, 1)

	snrs, err = suite.s.ListSensorsForUser("not-existed-id", Sensor{})
	suite.Nil(err)
	suite.Len(snrs, 0)
}

func (suite *storageImplTestSuite) TestCreateSensor() {
	test_str := "test"
	snr := Sensor{
		Id:                      &test_str,
		Name:                    &test_str,
		CoreId:                  &test_str,
		EntityName:              &test_str,
		OwnerId:                 &test_str,
		ApplicationCredentialId: &test_str,
		State: &test_str,
	}
	snr, err := suite.s.CreateSensor(snr)
	suite.Nil(err)
	suite.Equal(test_str, *snr.Id)
	suite.Equal(test_str, *snr.Name)
	suite.Equal(test_str, *snr.CoreId)
	suite.Equal(test_str, *snr.EntityName)
	suite.Equal(test_str, *snr.OwnerId)
	suite.Equal(test_str, *snr.ApplicationCredentialId)
	suite.Equal(test_str, *snr.State)
}

func (suite *storageImplTestSuite) TestDeleteSensor() {
	err := suite.s.DeleteSensor(test_sensor_id)
	suite.Nil(err)
	ss, err := suite.s.ListSensors(Sensor{})
	suite.Nil(err)
	suite.Len(ss, 0)
}

func (suite *storageImplTestSuite) TestGetSensorTags() {
	tags, err := suite.s.GetSensorTags(test_sensor_id)
	suite.Nil(err)
	suite.Len(tags, 2)
}

func (suite *storageImplTestSuite) TestAddSensorTag() {
	test_str := "test"
	_, err := suite.s.AddSensorTag(SensorTag{
		CreatedAt: time.Now(),
		Id:        &test_str,
		SensorId:  &test_sensor_id,
		Tag:       &test_str,
	})
	suite.Nil(err)

	tags, err := suite.s.GetSensorTags(test_sensor_id)
	suite.Nil(err)
	suite.Len(tags, 3)
}

func (suite *storageImplTestSuite) TestRemoveSensorTag() {
	err := suite.s.RemoveSensorTag(test_tag0_id)
	suite.Nil(err)

	tags, err := suite.s.GetSensorTags(test_sensor_id)
	suite.Nil(err)
	suite.Len(tags, 1)
}

func TestStorageImplTestSuite(t *testing.T) {
	suite.Run(t, new(storageImplTestSuite))
}
