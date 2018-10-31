package metathings_identityd2_storage

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

var (
	testDomainID    = "test-domain-id"
	testDomainName  = "test-domain-name"
	testDomainAlias = "test-domain-alias"
	testDomainExtra = "test-domain-extra"

	testGroupID          = "test-group-id"
	testGroupName        = "test-group-name"
	testGroupAlias       = "test-group-alias"
	testGroupDescription = "test-group-decription"
	testGroupExtra       = "test-group-extra"

	testEntityID       = "test-entity-id"
	testEntityName     = "test-entity-name"
	testEntityAlias    = "test-entity-alias"
	testEntityPassword = "test-entity-password"
	testEntityExtra    = "test-entity-extra"

	testRoleID          = "test-role-id"
	testRoleName        = "test-role-name"
	testRoleAlias       = "test-role-alias"
	testRoleDescription = "test-role-description"
	testRoleExtra       = "test-role-extra"

	testCredentialID         = "test-credential-id"
	testCredentialName       = "test-credential-name"
	testCredentialAlias      = "test-credential-alias"
	testCredentialSecret     = "test-credential-secret"
	testCredentialDecription = "test-credential-decription"
)

type storageImplTestSuite struct {
	suite.Suite
	s *StorageImpl
}

func (suite *storageImplTestSuite) SetupTest() {
	var err error
	suite.s, err = NewStorageImpl("sqlite3", ":memory:", "logger", logrus.New())
	if err != nil {
		fmt.Println("err", err.Error())
	}

	dom := Domain{
		Id:    &testDomainID,
		Name:  &testDomainName,
		Alias: &testDomainAlias,
		Extra: &testDomainExtra,
	}

	_, err = suite.s.CreateDomain(&dom)
	if err != nil {
		fmt.Println("err create Domain")
	}

	grp := Group{
		Id:          &testGroupID,
		DomainId:    &testDomainID,
		Name:        &testGroupName,
		Alias:       &testGroupAlias,
		Description: &testGroupDescription,
		Extra:       &testGroupExtra,
	}

	suite.s.CreateGroup(&grp)

	ent := Entity{
		Id:       &testEntityID,
		Name:     &testEntityName,
		Alias:    &testEntityAlias,
		Password: &testEntityPassword,
		Extra:    &testEntityExtra,
	}

	suite.s.CreateEntity(&ent)
	suite.s.AddEntityToGroup(testEntityID, testGroupID)

	rol := Role{
		Id:          &testRoleID,
		Name:        &testRoleName,
		Alias:       &testRoleAlias,
		Description: &testRoleDescription,
		Extra:       &testRoleExtra,
	}

	suite.s.CreateRole(&rol)
	suite.s.AddRoleToEntity(testEntityID, testRoleID)
	suite.s.AddRoleToGroup(testGroupID, testRoleID)

	cred := Credential{
		Id:          &testCredentialID,
		DomainId:    &testDomainID,
		EntityId:    &testEntityID,
		Name:        &testCredentialName,
		Alias:       &testCredentialAlias,
		Secret:      &testCredentialSecret,
		Description: &testCredentialDecription,
	}

	suite.s.CreateCredential(&cred)
}

func (suite *storageImplTestSuite) TestCreteDomain() {
	testStr := "test"
	dom := Domain{
		Id:    &testStr,
		Name:  &testStr,
		Alias: &testStr,
		Extra: &testStr,
	}
	domRet, err := suite.s.CreateDomain(&dom)
	suite.Nil(err)
	suite.Equal(testStr, *domRet.Id)
	suite.Equal(testStr, *domRet.Name)
	suite.Equal(testStr, *domRet.Alias)
	suite.Equal(testStr, *domRet.Extra)
}

func (suite *storageImplTestSuite) TestDeleteDomain() {
	err := suite.s.DeleteDomain(testDomainID)
	suite.Nil(err)
	dom, err := suite.s.GetDomain(testDomainID)
	suite.Nil(dom)
}

func (suite *storageImplTestSuite) TestPatchDomainName() {
	testStr := "test"
	dom := &Domain {
		Name: &testStr,
	}
	dom, err := suite.s.PatchDomain(testDomainID, dom)
	suite.Nil(err)
	suite.Equal(testStr, *dom.Name)
}

func (suite *storageImplTestSuite) TestPatchDomainAlias() {
	testStr := "test"
	dom := &Domain {
		Alias: &testStr,
	}
	dom, err := suite.s.PatchDomain(testDomainID, dom)
	suite.Nil(err)
	suite.Equal(testStr, *dom.Alias)
}

func (suite *storageImplTestSuite) TestPatchDomainParentId() {
	testStr := "test"
	dom := &Domain {
		ParentId: &testStr,
	}
	dom, err := suite.s.PatchDomain(testDomainID, dom)
	suite.Nil(err)
	suite.Equal(testStr, *dom.ParentId)
}

func (suite *storageImplTestSuite) TestPatchDomainExtra() {
	testStr := "test"
	dom := &Domain {
		Extra: &testStr,
	}
	dom, err := suite.s.PatchDomain(testDomainID, dom)
	suite.Nil(err)
	suite.Equal(testStr, *dom.Extra)
}

func (suite *storageImplTestSuite) TestGetDomain() {
	dom, err := suite.s.GetDomain(testDomainID)
	suite.Nil(err)
	suite.Equal(testDomainName, *dom.Name)
}

func (suite *storageImplTestSuite) TestListDomains() {
	dom := &Domain {
		Name: 	&testDomainName,
	}
	doms, err := suite.s.ListDomains(dom)
	suite.Nil(err)
	suite.Len(doms, 1)
}

func TestStorageImplTestSuite(t *testing.T) {
	suite.Run(t, new(storageImplTestSuite))
}
