package metathings_cored_storage

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

var (
	test_core_id                = "test-core-id"
	test_core_name              = "test-core-name"
	test_core_state             = "test-core-state"
	test_project_id             = "test-project-id"
	test_owner_id               = "test-owner-id"
	test_application_credential = "test-application-credential"
	test_entity_id              = "test-entity-id"
	test_entity_name            = "test-entity-name"
	test_entity_service_name    = "test-entity-service-name"
	test_entity_endpoint        = "test-entity-endpoint"
	test_entity_state           = "test-entity-state"
)

type storageImplTestSuite struct {
	suite.Suite
	s Storage
}

func (suite *storageImplTestSuite) SetupTest() {
	suite.s, _ = NewStorage("sqlite3", ":memory:", logrus.New())

	c := Core{
		Id:        &test_core_id,
		Name:      &test_core_name,
		ProjectId: &test_project_id,
		OwnerId:   &test_owner_id,
		State:     &test_core_state,
	}
	suite.s.CreateCore(c)
	suite.s.AssignCoreToApplicationCredential(test_core_id, test_application_credential)

	e := Entity{
		Id:          &test_entity_id,
		CoreId:      &test_core_id,
		Name:        &test_entity_name,
		ServiceName: &test_entity_service_name,
		Endpoint:    &test_entity_endpoint,
	}
	suite.s.CreateEntity(e)
}

func (suite *storageImplTestSuite) TestCreateCoreAndAssignCoreToApplicationCredential() {
	id := "test-id"
	app_cred_id := "test-app-cred-id"
	name := "test"
	project_id := "project-id"
	owner_id := "owner-id"
	state := "unknown"
	c := Core{
		Id:        &id,
		Name:      &name,
		ProjectId: &project_id,
		OwnerId:   &owner_id,
		State:     &state,
	}

	nc, err := suite.s.CreateCore(c)
	suite.Empty(err)
	suite.Equal(id, *nc.Id)
	suite.Equal(name, *nc.Name)
	suite.Equal(project_id, *nc.ProjectId)
	suite.Equal(owner_id, *nc.OwnerId)
	suite.Equal(state, *nc.State)
	suite.NotEqual("0001-01-01 00:00:00 +0000 UTC", nc.HeartbeatAt.String())
	suite.NotEqual("0001-01-01 00:00:00 +0000 UTC", nc.CreatedAt.String())
	suite.NotEqual("0001-01-01 00:00:00 +0000 UTC", nc.UpdatedAt.String())

	err = suite.s.AssignCoreToApplicationCredential(id, app_cred_id)
	suite.Empty(err)
}

func (suite *storageImplTestSuite) TestDeleteCore() {
	err := suite.s.DeleteCore(test_core_id)
	suite.Empty(err)

	cs, err := suite.s.ListCores(Core{})
	suite.Empty(err)
	suite.Len(cs, 0)
}

func (suite *storageImplTestSuite) TestGetCore() {
	c, err := suite.s.GetCore(test_core_id)
	suite.Empty(err)
	suite.Equal(test_core_name, *c.Name)
}

func (suite *storageImplTestSuite) TestGetNotExistedCore() {
	_, err := suite.s.GetCore("not-existed-id")
	suite.NotEmpty(err)
}

func (suite *storageImplTestSuite) TestListCores() {
	cs, err := suite.s.ListCores(Core{})
	suite.Empty(err)
	suite.Len(cs, 1)
}

func (suite *storageImplTestSuite) TestListCoresForUser() {
	cs, err := suite.s.ListCoresForUser(test_owner_id, Core{})
	suite.Empty(err)
	suite.Len(cs, 1)
}

func (suite *storageImplTestSuite) TestGetAssignedCoreFromApplicationCredential() {
	c, err := suite.s.GetAssignedCoreFromApplicationCredential(test_application_credential)
	suite.Empty(err)
	suite.Equal(test_core_name, *c.Name)
}

func (suite *storageImplTestSuite) TestCreateEntity() {
	id := "test-id"
	name := "test-name"
	service_name := "test-service-name"
	endpoint := "test-endpoint"
	state := "test-state"
	e := Entity{
		Id:          &id,
		Name:        &name,
		ServiceName: &service_name,
		Endpoint:    &endpoint,
		State:       &state,
	}
	ne, err := suite.s.CreateEntity(e)
	suite.Empty(err)
	suite.Equal(id, *ne.Id)
	suite.Equal(name, *ne.Name)
	suite.Equal(service_name, *ne.ServiceName)
	suite.Equal(endpoint, *ne.Endpoint)
	suite.Equal(state, *ne.State)
	suite.NotEqual("0001-01-01 00:00:00 +0000 UTC", ne.HeartbeatAt.String())
	suite.NotEqual("0001-01-01 00:00:00 +0000 UTC", ne.CreatedAt.String())
	suite.NotEqual("0001-01-01 00:00:00 +0000 UTC", ne.UpdatedAt.String())
}

func (suite *storageImplTestSuite) TestDeleteEntity() {
	err := suite.s.DeleteEntity(test_entity_id)
	suite.Empty(err)

	es, err := suite.s.ListEntities(Entity{})
	suite.Empty(err)
	suite.Len(es, 0)
}

func (suite *storageImplTestSuite) TestGetEntity() {
	e, err := suite.s.GetEntity(test_entity_id)
	suite.Empty(err)
	suite.Equal(test_entity_name, *e.Name)
}

func (suite *storageImplTestSuite) TestListEntities() {
	es, err := suite.s.ListEntities(Entity{})
	suite.Empty(err)
	suite.Len(es, 1)
}

func (suite *storageImplTestSuite) TestListEntitiesForCore() {
	es, err := suite.s.ListEntitiesForCore(test_core_id, Entity{})
	suite.Empty(err)
	suite.Len(es, 1)
}

func TestStorageImplTestSuite(t *testing.T) {
	suite.Run(t, new(storageImplTestSuite))
}
