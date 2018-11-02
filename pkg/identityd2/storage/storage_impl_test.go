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
	testGroupDomainID    = "test-group-domainid"
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

	testTokenID           = "test-token-id"
	testTokenCredentialID = "test-token-credentialid"
	testTokenText         = "test-token-text"
)

type storageImplTestSuite struct {
	suite.Suite
	s *StorageImpl
}

func (suite *storageImplTestSuite) SetupTest() {
	var err error
	suite.s, err = NewStorageImpl("sqlite3", ":memory:", "logger", logrus.New())
	if err != nil {
		fmt.Println("SetupTest newstorage error:", err.Error())
	}

	dom := Domain{
		Id:    &testDomainID,
		Name:  &testDomainName,
		Alias: &testDomainAlias,
		Extra: &testDomainExtra,
	}

	_, err = suite.s.CreateDomain(&dom)
	if err != nil {
		fmt.Println("SetupTest create Domain error:", err.Error())
	}

	grp := Group{
		Id:          &testGroupID,
		DomainId:    &testGroupDomainID,
		Name:        &testGroupName,
		Alias:       &testGroupAlias,
		Description: &testGroupDescription,
		Extra:       &testGroupExtra,
	}

	_, err = suite.s.CreateGroup(&grp)
	if err != nil {
		fmt.Println("SetupTest create Group error:", err.Error())
	}

	ent := Entity{
		Id:       &testEntityID,
		Name:     &testEntityName,
		Alias:    &testEntityAlias,
		Password: &testEntityPassword,
		Extra:    &testEntityExtra,
	}

	_, err = suite.s.CreateEntity(&ent)
	if err != nil {
		fmt.Println("SetupTest create Entity error:", err.Error())
	}

	err = suite.s.AddEntityToGroup(testEntityID, testGroupID)
	if err != nil {
		fmt.Println("SetupTest add entity to group error:", err.Error())
	}

	err = suite.s.AddEntityToDomain(testDomainID, testEntityID)
	if err != nil {
		fmt.Println("SetupTest add entity to domain error:", err.Error())
	}

	rol := Role{
		Id:          &testRoleID,
		DomainId:    &testRoleDomainID,
		Name:        &testRoleName,
		Alias:       &testRoleAlias,
		Description: &testRoleDescription,
		Extra:       &testRoleExtra,
	}

	_, err = suite.s.CreateRole(&rol)
	if err != nil {
		fmt.Println("SetupTest create role error:", err.Error())
	}

	err = suite.s.AddRoleToEntity(testEntityID, testRoleID)
	if err != nil {
		fmt.Println("SetupTest add role to entity error:", err.Error())
	}

	err = suite.s.AddRoleToGroup(testGroupID, testRoleID)
	if err != nil {
		fmt.Println("SetupTest add role to group error:", err.Error())
	}

	cred := Credential{
		Id:          &testCredentialID,
		DomainId:    &testDomainID,
		EntityId:    &testEntityID,
		Name:        &testCredentialName,
		Alias:       &testCredentialAlias,
		Secret:      &testCredentialSecret,
		Description: &testCredentialDecription,
	}

	_, err = suite.s.CreateCredential(&cred)
	if err != nil {
		fmt.Println("SetupTest create credentia error:", err.Error())
	}

	tkn := Token{
		Id:           &testTokenID,
		DomainId:     &testDomainID,
		EntityId:     &testEntityID,
		CredentialId: &testCredentialID,
		Text:         &testTokenText,
	}

	_, err = suite.s.CreateToken(&tkn)
	if err != nil {
		fmt.Println("SetupTest create token error:", err.Error())
	}
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

func (suite *storageImplTestSuite) TestPatchDomainAlias() {
	testStr := "test"
	dom := &Domain{
		Alias: &testStr,
	}
	dom, err := suite.s.PatchDomain(testDomainID, dom)
	suite.Nil(err)
	suite.Equal(testStr, *dom.Alias)
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
	var err error
	testStr := "test"
	rol := &Role{}

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
		Id:       &testStr,
		Name:     &testStr,
		Alias:    &testStr,
		Password: &testStr,
		Extra:    &testStr,
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
	var err error
	testStr := "test"
	ent := &Entity{}

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

func (suite *storageImplTestSuite) TestGetEntity() {
	ent, err := suite.s.GetEntity(testEntityID)
	suite.Nil(err)
	suite.Equal(testEntityName, *ent.Name)
}

func (suite *storageImplTestSuite) TestGetEntityByName() {
	ent, err := suite.s.GetEntityByName(testEntityName)
	suite.Nil(err)
	suite.Equal(testEntityID, *ent.Id)
}

func (suite *storageImplTestSuite) TestListEntities() {
	//list by Name
	ent := &Entity{
		Name: &testEntityName,
	}
	ents, err := suite.s.ListEntities(ent)
	suite.Nil(err)
	suite.Len(ents, 1)

	//list by Alias
	ent = &Entity{
		Alias: &testEntityAlias,
	}
	ents, err = suite.s.ListEntities(ent)
	suite.Nil(err)
	suite.Len(ents, 1)

	//list by Extra
	ent = &Entity{
		Extra: &testEntityExtra,
	}
	ents, err = suite.s.ListEntities(ent)
	suite.Nil(err)
	suite.Len(ents, 1)
}

func (suite *storageImplTestSuite) TestListEntitiesByDomainId() {
	ents, err := suite.s.ListEntitiesByDomainId(testDomainID)
	suite.Nil(err)
	suite.Len(ents, 1)
}

func (suite *storageImplTestSuite) TestAddRoleToEntity() {
	err := suite.s.AddRoleToEntity(testEntityID, testRoleID)
	suite.Nil(err)
}

func (suite *storageImplTestSuite) TestRemoveRoleFromEntity() {
	err := suite.s.RemoveRoleFromEntity(testEntityID, testRoleID)
	suite.Nil(err)
}

func (suite *storageImplTestSuite) TestCreateGroup() {
	testStr := "test"
	grp := &Group{
		Id:          &testStr,
		DomainId:    &testStr,
		Name:        &testStr,
		Alias:       &testStr,
		Description: &testStr,
		Extra:       &testStr,
	}

	grpRet, err := suite.s.CreateGroup(grp)
	suite.Nil(err)
	suite.Equal(testStr, *grpRet.Id)
	suite.Equal(testStr, *grpRet.DomainId)
	suite.Equal(testStr, *grpRet.Name)
	suite.Equal(testStr, *grpRet.Alias)
	suite.Equal(testStr, *grpRet.Description)
	suite.Equal(testStr, *grpRet.Extra)
}

func (suite *storageImplTestSuite) TestDeleteGroup() {
	err := suite.s.DeleteGroup(testGroupID)
	suite.Nil(err)
	grp, err := suite.s.GetGroup(testGroupID)
	suite.NotNil(err)
	suite.Nil(grp)
}

func (suite *storageImplTestSuite) TestPatchGroup() {
	var err error
	testStr := "test"
	grp := &Group{}

	//Alias
	grp = &Group{
		Alias: &testStr,
	}
	grp, err = suite.s.PatchGroup(testGroupID, grp)
	suite.Nil(err)
	suite.Equal(testStr, *grp.Alias)

	//Description
	grp = &Group{
		Description: &testStr,
	}
	grp, err = suite.s.PatchGroup(testGroupID, grp)
	suite.Nil(err)
	suite.Equal(testStr, *grp.Description)

	//Extra
	grp = &Group{
		Extra: &testStr,
	}
	grp, err = suite.s.PatchGroup(testGroupID, grp)
	suite.Nil(err)
	suite.Equal(testStr, *grp.Extra)
}

func (suite *storageImplTestSuite) TestGetGroup() {
	grp, err := suite.s.GetGroup(testGroupID)
	suite.Nil(err)
	suite.Equal(testGroupName, *grp.Name)
}

func (suite *storageImplTestSuite) TestListGroups() {
	//list by DomainId
	grp := &Group{
		DomainId: &testGroupDomainID,
	}
	grps, err := suite.s.ListGroups(grp)
	suite.Nil(err)
	suite.Len(grps, 1)

	//list by Name
	grp = &Group{
		Name: &testGroupName,
	}
	grps, err = suite.s.ListGroups(grp)
	suite.Nil(err)
	suite.Len(grps, 1)

	//list by Alias
	grp = &Group{
		Alias: &testGroupAlias,
	}
	grps, err = suite.s.ListGroups(grp)
	suite.Nil(err)
	suite.Len(grps, 1)

	//list by Description
	grp = &Group{
		Description: &testGroupDescription,
	}
	grps, err = suite.s.ListGroups(grp)
	suite.Nil(err)
	suite.Len(grps, 1)

	//list by Extra
	grp = &Group{
		Extra: &testGroupExtra,
	}
	grps, err = suite.s.ListGroups(grp)
	suite.Nil(err)
	suite.Len(grps, 1)
}

func (suite *storageImplTestSuite) TestAddRoleToGroup() {
	err := suite.s.AddRoleToGroup(testGroupID, testRoleID)
	suite.Nil(err)
}

func (suite *storageImplTestSuite) TestRemoveRoleFromGroup() {
	err := suite.s.AddRoleToGroup(testGroupID, testRoleID)
	suite.Nil(err)

	err = suite.s.RemoveRoleFromGroup(testGroupID, testRoleID)
	suite.Nil(err)
}

func (suite *storageImplTestSuite) TestAddEntityToGroup() {
	err := suite.s.AddEntityToGroup(testEntityID, testRoleID)
	suite.Nil(err)

	err = suite.s.RemoveEntityFromGroup(testGroupID, testRoleID)
	suite.Nil(err)
}

func (suite *storageImplTestSuite) TestRemoveEntityFromGroup() {
	err := suite.s.AddEntityToGroup(testEntityID, testRoleID)
	suite.Nil(err)
}

func (suite *storageImplTestSuite) TestCreateCredential() {
	testStr := "test"
	cred := &Credential{
		Id:          &testStr,
		DomainId:    &testStr,
		EntityId:    &testStr,
		Name:        &testStr,
		Alias:       &testStr,
		Secret:      &testStr,
		Description: &testStr,
	}

	credRet, err := suite.s.CreateCredential(cred)
	suite.Nil(err)
	suite.Equal(testStr, *credRet.Id)
	suite.Equal(testStr, *credRet.DomainId)
	suite.Equal(testStr, *credRet.EntityId)
	suite.Equal(testStr, *credRet.Name)
	suite.Equal(testStr, *credRet.Alias)
	suite.Equal(testStr, *credRet.Secret)
	suite.Equal(testStr, *credRet.Description)
}

func (suite *storageImplTestSuite) TestDeleteCredential() {
	err := suite.s.DeleteCredential(testCredentialID)
	suite.Nil(err)
	cred, err := suite.s.GetCredential(testCredentialID)
	suite.NotNil(err)
	suite.Nil(cred)
}

func (suite *storageImplTestSuite) TestPatchCredential() {
	var err error
	testStr := "test"
	cred := &Credential{}

	//Alias
	cred = &Credential{
		Alias: &testStr,
	}
	cred, err = suite.s.PatchCredential(testCredentialID, cred)
	suite.Nil(err)
	suite.Equal(testStr, *cred.Alias)

	//Decription
	cred = &Credential{
		Description: &testStr,
	}
	cred, err = suite.s.PatchCredential(testCredentialID, cred)
	suite.Nil(err)
	suite.Equal(testStr, *cred.Description)
}

func (suite *storageImplTestSuite) TestGetCredential() {
	cred, err := suite.s.GetCredential(testCredentialID)
	suite.Nil(err)
	suite.Equal(testCredentialName, *cred.Name)
}

func (suite *storageImplTestSuite) TestListCredentials() {
	//list by DomainId
	cred := &Credential{
		DomainId: &testDomainID,
	}
	creds, err := suite.s.ListCredentials(cred)
	suite.Nil(err)
	suite.Len(creds, 1)

	//list by EntityId
	cred = &Credential{
		EntityId: &testEntityID,
	}
	creds, err = suite.s.ListCredentials(cred)
	suite.Nil(err)
	suite.Len(creds, 1)

	//list by Name
	cred = &Credential{
		Name: &testCredentialName,
	}
	creds, err = suite.s.ListCredentials(cred)
	suite.Nil(err)
	suite.Len(creds, 1)

	//list by Alias
	cred = &Credential{
		Alias: &testCredentialAlias,
	}
	creds, err = suite.s.ListCredentials(cred)
	suite.Nil(err)
	suite.Len(creds, 1)

	//list by Sectret
	cred = &Credential{
		Secret: &testCredentialSecret,
	}
	creds, err = suite.s.ListCredentials(cred)
	suite.Nil(err)
	suite.Len(creds, 1)

	//list by Description
	cred = &Credential{
		Description: &testCredentialDecription,
	}
	creds, err = suite.s.ListCredentials(cred)
	suite.Nil(err)
	suite.Len(creds, 1)

	//list by EntityId
	cred = &Credential{
		EntityId: &testEntityID,
	}
	creds, err = suite.s.ListCredentials(cred)
	suite.Nil(err)
	suite.Len(creds, 1)
}

func (suite *storageImplTestSuite) TestCreateToken() {
	testStr := "test"
	tkn := &Token{
		Id:           &testStr,
		DomainId:     &testDomainID,
		EntityId:     &testEntityID,
		CredentialId: &testCredentialID,
		Text:         &testStr,
	}

	tknRet, err := suite.s.CreateToken(tkn)
	suite.Nil(err)
	suite.Equal(testStr, *tknRet.Id)
	suite.Equal(testDomainID, *tknRet.DomainId)
	suite.Equal(testEntityID, *tknRet.EntityId)
	suite.Equal(testCredentialID, *tknRet.CredentialId)
	suite.Equal(testStr, *tknRet.Text)
}

func (suite *storageImplTestSuite) TestDeleteToken() {
	err := suite.s.DeleteToken(testTokenID)
	suite.Nil(err)
	tkn, err := suite.s.GetToken(testTokenID)
	suite.NotNil(err)
	suite.Nil(tkn)
}

func (suite *storageImplTestSuite) TestGetToken() {
	tkn, err := suite.s.GetToken(testTokenID)
	suite.Nil(err)
	suite.Equal(testTokenText, *tkn.Text)
}

func (suite *storageImplTestSuite) TestGetTokenByText() {
	tkn, err := suite.s.GetTokenByText(testTokenText)
	suite.Nil(err)
	suite.Equal(testTokenID, *tkn.Id)
}

func (suite *storageImplTestSuite) TestListTokens() {
	//list by DomainId
	tkn := &Token{
		DomainId: &testDomainID,
	}
	tkns, err := suite.s.ListTokens(tkn)
	suite.Nil(err)
	suite.Len(tkns, 1)

	//list by EntityId
	tkn = &Token{
		EntityId: &testEntityID,
	}
	tkns, err = suite.s.ListTokens(tkn)
	suite.Nil(err)
	suite.Len(tkns, 1)

	//list by CredentialId
	tkn = &Token{
		CredentialId: &testCredentialID,
	}
	tkns, err = suite.s.ListTokens(tkn)
	suite.Nil(err)
	suite.Len(tkns, 1)

	//list by Text
	tkn = &Token{
		Text: &testTokenText,
	}
	tkns, err = suite.s.ListTokens(tkn)
	suite.Nil(err)
	suite.Len(tkns, 1)
}

func TestStorageImplTestSuite(t *testing.T) {
	suite.Run(t, new(storageImplTestSuite))
}
