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
	testRoleDomainID    = "test-role-domainid"
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
		DomainId:    &testRoleDomainID,
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
	suite.NotNil(err)
	suite.Nil(dom)
}

func (suite *storageImplTestSuite) TestPatchDomainName() {
	testStr := "test"
	dom := &Domain{
		Name: &testStr,
	}
	dom, err := suite.s.PatchDomain(testDomainID, dom)
	suite.Nil(err)
	suite.Equal(testStr, *dom.Name)
}

func (suite *storageImplTestSuite) TestPatchDomainAlias() {
	testStr := "test"
	dom := &Domain{
		Alias: &testStr,
	}
	dom, err := suite.s.PatchDomain(testDomainID, dom)
	suite.Nil(err)
	suite.Equal(testStr, *dom.Alias)
}

func (suite *storageImplTestSuite) TestPatchDomainParentId() {
	testStr := "test"
	dom := &Domain{
		ParentId: &testStr,
	}
	dom, err := suite.s.PatchDomain(testDomainID, dom)
	suite.Nil(err)
	suite.Equal(testStr, *dom.ParentId)
}

func (suite *storageImplTestSuite) TestPatchDomainExtra() {
	testStr := "test"
	dom := &Domain{
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
	//list by Name
	dom := &Domain{
		Name: &testDomainName,
	}
	doms, err := suite.s.ListDomains(dom)
	suite.Nil(err)
	suite.Len(doms, 1)

	//list by Alias
	dom = &Domain{
		Alias: &testDomainAlias,
	}
	doms, err = suite.s.ListDomains(dom)
	suite.Nil(err)
	suite.Len(doms, 1)

	//list by Extra
	dom = &Domain{
		Extra: &testDomainExtra,
	}
	doms, err = suite.s.ListDomains(dom)
	suite.Nil(err)
	suite.Len(doms, 1)
}

func (suite *storageImplTestSuite) TestAddEntityToDomain() {
	err := suite.s.AddEntityToDomain(testDomainID, testEntityID)
	suite.Nil(err)
}

func (suite *storageImplTestSuite) TestRemoveEntityFormDomain() {
	err := suite.s.AddEntityToDomain(testDomainID, testEntityID)
	suite.Nil(err)

	err = suite.s.RemoveEntityFromDomain(testDomainID, testEntityID)
	suite.Nil(err)
}

func (suite *storageImplTestSuite) TestCreateRole() {
	testStr := "test"
	rol := Role{
		Id:          &testStr,
		DomainId:    &testStr,
		Name:        &testStr,
		Alias:       &testStr,
		Description: &testStr,
	}

	rolRet, err := suite.s.CreateRole(&rol)
	suite.Nil(err)
	suite.Equal(testStr, *rolRet.Id)
	suite.Equal(testStr, *rolRet.DomainId)
	suite.Equal(testStr, *rolRet.Name)
	suite.Equal(testStr, *rolRet.Alias)
	suite.Equal(testStr, *rolRet.Description)
}

func (suite *storageImplTestSuite) TestDeleteRole() {
	err := suite.s.DeleteRole(testRoleID)
	suite.Nil(err)
	rol, err := suite.s.GetRole(testRoleID)
	suite.NotNil(err)
	suite.Nil(rol)
}

func (suite *storageImplTestSuite) TestPatchRole() {
	testStr := "test"

	//DomainId
	rol := &Role{
		DomainId: &testStr,
	}
	rol, err := suite.s.PatchRole(testRoleID, rol)
	suite.Nil(err)
	suite.Equal(testStr, *rol.DomainId)

	//Name
	rol = &Role{
		Name: &testStr,
	}
	rol, err = suite.s.PatchRole(testRoleID, rol)
	suite.Nil(err)
	suite.Equal(testStr, *rol.Name)

	//Alias
	rol = &Role{
		Alias: &testStr,
	}
	rol, err = suite.s.PatchRole(testRoleID, rol)
	suite.Nil(err)
	suite.Equal(testStr, *rol.Alias)

	//Description
	rol = &Role{
		Description: &testStr,
	}
	rol, err = suite.s.PatchRole(testRoleID, rol)
	suite.Nil(err)
	suite.Equal(testStr, *rol.Description)

	//Extra
	rol = &Role{
		Extra: &testStr,
	}
	rol, err = suite.s.PatchRole(testRoleID, rol)
	suite.Nil(err)
	suite.Equal(testStr, *rol.Extra)
}

func (suite *storageImplTestSuite) TestGetRole() {
	rol, err := suite.s.GetRole(testRoleID)
	suite.Nil(err)
	suite.Equal(testRoleName, *rol.Name)
}

func (suite *storageImplTestSuite) TestlistRoles() {
	//list by DomainId
	rol := &Role{
		DomainId: &testRoleDomainID,
	}
	rols, err := suite.s.ListRoles(rol)
	suite.Nil(err)
	suite.Len(rols, 1)

	//list by Name
	rol = &Role{
		Name: &testRoleName,
	}
	rols, err = suite.s.ListRoles(rol)
	suite.Nil(err)
	suite.Len(rols, 1)

	//list by Alias
	rol = &Role{
		Alias: &testRoleAlias,
	}
	rols, err = suite.s.ListRoles(rol)
	suite.Nil(err)
	suite.Len(rols, 1)

	//list by Description
	rol = &Role{
		Description: &testRoleDescription,
	}
	rols, err = suite.s.ListRoles(rol)
	suite.Nil(err)
	suite.Len(rols, 1)

	//list by Extra
	rol = &Role{
		Extra: &testRoleExtra,
	}
	rols, err = suite.s.ListRoles(rol)
	suite.Nil(err)
	suite.Len(rols, 1)
}

func (suite *storageImplTestSuite) TestCreateEntity() {
	testStr := "test"
	ent := &Entity{
		Id:			&testStr,
		Name:		&testStr,
		Alias:  	&testStr,
		Password: 	&testStr,
		Extra:		&testStr,
	}
	entRet, err := suite.s.CreateEntity(ent)
	suite.Nil(err)
	suite.Equal(testStr, *entRet.Id)
	suite.Equal(testStr, *entRet.Name)
	suite.Equal(testStr, *entRet.Alias)
	suite.Equal(testStr, *entRet.Password)
	suite.Equal(testStr, *entRet.Extra)
}

func (suite *storageImplTestSuite) TestDeleteEntity() {
	err := suite.s.DeleteEntity(testEntityID)
	suite.Nil(err)
	ent, err := suite.s.GetEntity(testEntityID)
	suite.NotNil(err)
	suite.Nil(ent)
}

func (suite *storageImplTestSuite) TestPatchEntity() {
	testStr := "test"

	//Name
	ent := &Entity{
		Name: &testStr,
	}
	ent, err := suite.s.PatchEntity(testEntityID, ent)
	suite.Nil(err)
	suite.Equal(testStr, *ent.Name)

	//Alias
	ent = &Entity{
		Alias: &testStr,
	}
	ent, err = suite.s.PatchEntity(testEntityID, ent)
	suite.Nil(err)
	suite.Equal(testStr, *ent.Alias)

	//Password
	ent = &Entity{
		Password: &testStr,
	}
	ent, err = suite.s.PatchEntity(testEntityID, ent)
	suite.Nil(err)
	suite.Equal(testStr, *ent.Password)

	//Extra
	ent = &Entity{
		Extra: &testStr,
	}
	ent, err = suite.s.PatchEntity(testEntityID, ent)
	suite.Nil(err)
	suite.Equal(testStr, *ent.Extra)
}

func TestStorageImplTestSuite(t *testing.T) {
	suite.Run(t, new(storageImplTestSuite))
}
