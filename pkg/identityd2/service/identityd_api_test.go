package metathings_identityd2_service

import (
	"context"
	"fmt"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	mock_enf "github.com/nayotta/metathings/pkg/identityd2/policy/mock"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type identifydTestSuite struct {
	suite.Suite
	ctx context.Context
	enf mock_enf.MockEnforcer
	s   *MetathingsIdentitydService
}

type mockEnf struct {
	mock.Mock
}

var (
	defaultDomainID       = "default"
	defaultDomainName     = "default"
	defaultDomainAlias    = "default"
	defaultDomainParentID = ""

	defaultSysadminID    = id_helper.NewId()
	defaultSysadminName  = "sysadmin"
	defaultSysadminAlias = "sysadmin"

	defaultAdminID       = id_helper.NewId()
	defaultAdminName     = "admin"
	defaultAdminAlias    = "admin"
	defaultAdminPassword = passwd_helper.MustParsePassword("admin")

	testDomainID    = "test-domain-id"
	testDomainName  = "test-domain-name"
	testDomainAlias = "test-domain-alias"
	testDomainExtra = map[string]*wrappers.StringValue{
		"test": &wrappers.StringValue{
			Value: "test",
		},
	}

	testRoleID          = "test-role-id"
	testRoleName        = "test-role-name"
	testRoleAlias       = "test-role-alias"
	testRoleDescription = "test-role-description"
	testRoleExtra       = map[string]*wrappers.StringValue{
		"test": &wrappers.StringValue{
			Value: "test",
		},
	}

	testEntityID       = "test-entity-id"
	testEntityName     = "test-entity-name"
	testEntityAlias    = "test-entity-alias"
	testEntityPassword = "test-entity-password"
	testEntityExtra    = map[string]*wrappers.StringValue{
		"test": &wrappers.StringValue{
			Value: "test",
		},
	}

	testGroupID         = "test-group-id"
	testGroupName       = "test-group-name"
	testGroupAlias      = "test-group-alias"
	testGroupDescrition = "test-group-description"
	testGroupExtra      = map[string]*wrappers.StringValue{
		"test": &wrappers.StringValue{
			Value: "test",
		},
	}

	testCredentialID          = "test-credential-id"
	testCredentialName        = "test-credential-name"
	testCredentialAlias       = "test-credential-alias"
	testCredentialDescription = "test-credential-description"
	day, _                    = time.ParseDuration("24h")
	testCredentialExpires     = time.Now().Add(day)
)

func (suite *identifydTestSuite) SetupTest() {
	var err error
	suite.ctx = context.Background()

	newStorage, err := storage.NewStorage("sqlite3", ":memory:", "logger", log.New())
	if err != nil {
		fmt.Println("identity SetupTest NewStorage error:", err.Error())
	}

	ctrl := gomock.NewController(suite.T())
	newEnforcer := mock_enf.NewMockEnforcer(ctrl)
	newEnforcer.EXPECT().AddObjectToKind(gomock.Any(), gomock.Any()).Return(nil).Times(9999)
	newEnforcer.EXPECT().RemoveObjectFromKind(gomock.Any(), gomock.Any()).Return(nil).Times(9999)
	newEnforcer.EXPECT().AddGroup(gomock.Any(), gomock.Any()).Return(nil).Times(9999)
	newEnforcer.EXPECT().AddSubjectToRole(gomock.Any(), gomock.Any()).Return(nil).Times(9999)
	newEnforcer.EXPECT().RemoveGroup(gomock.Any(), gomock.Any()).Return(nil).Times(9999)
	newEnforcer.EXPECT().RemoveSubjectFromRole(gomock.Any(), gomock.Any()).Return(nil).Times(9999)

	day, _ = time.ParseDuration("24h")
	newOpt := &MetathingsIdentitydServiceOption{
		TokenExpire: day,
	}

	if suite.s, err = NewMetathingsIdentitydService(newEnforcer, newOpt, log.New(), newStorage); err != nil {
		fmt.Println("identity SetupTest NewService error:", err.Error())
	}

	defaultDomain := &storage.Domain{
		Id:       &defaultDomainID,
		Name:     &defaultDomainName,
		Alias:    &defaultDomainAlias,
		ParentId: &defaultDomainParentID,
	}

	if _, err := suite.s.storage.CreateDomain(defaultDomain); err != nil {
		fmt.Println("identity SetupTest newDefaultDomain error:", err.Error())
	}

	//TODO(zh) enforce add domain

	defaultAdminRole := &storage.Role{
		Id:    &defaultSysadminID,
		Name:  &defaultSysadminName,
		Alias: &defaultSysadminAlias,
	}

	if _, err := suite.s.storage.CreateRole(defaultAdminRole); err != nil {
		fmt.Println("identity SetupTest newDefaultAdminRole error:", err.Error())
	}

	//TODO(zh) enforce add role

	defaultAdminEntity := &storage.Entity{
		Id:       &defaultAdminID,
		Name:     &defaultAdminName,
		Alias:    &defaultAdminAlias,
		Password: &defaultAdminPassword,
	}

	if _, err := suite.s.storage.CreateEntity(defaultAdminEntity); err != nil {
		fmt.Println("identity SetupTest newDefaultAdminEntity error:", err.Error())
	}

	if err = suite.s.storage.AddRoleToEntity(defaultAdminID, defaultSysadminID); err != nil {
		fmt.Println("identity SetupTest AddRoleToEntity error:", err.Error())
	}

	if err = suite.s.storage.AddEntityToDomain(defaultDomainID, defaultAdminID); err != nil {
		fmt.Println("identity SetupTest AddEntityToDomain error:", err.Error())
	}

	//TODO(zh) enforce add entity and sysadmin
}

func (suite *identifydTestSuite) TestDomain() {
	testStr := "test"

	domCreateReq := &pb.CreateDomainRequest{
		Id: &wrappers.StringValue{
			Value: testDomainID,
		},
		Name: &wrappers.StringValue{
			Value: testDomainName,
		},
		Alias: &wrappers.StringValue{
			Value: testDomainAlias,
		},
		Parent: &pb.OpDomain{
			Id: &wrappers.StringValue{
				Value: defaultDomainID,
			},
		},
		Extra: map[string]*wrappers.StringValue{
			"test": &wrappers.StringValue{
				Value: "",
			},
		},
	}
	_, err := suite.s.CreateDomain(suite.ctx, domCreateReq)
	suite.Nil(err)

	//test create domain with no id
	domCreateReq.Id = nil
	_, err = suite.s.CreateDomain(suite.ctx, domCreateReq)
	suite.Nil(err)

	//test get domain
	domGetReq := &pb.GetDomainRequest{
		Domain: &pb.OpDomain{
			Id: &wrappers.StringValue{
				Value: testDomainID,
			},
		},
	}
	domGetRet, err := suite.s.GetDomain(suite.ctx, domGetReq)
	suite.Nil(err)
	suite.Equal(domGetRet.GetDomain().GetName(), testDomainName)

	//test patch domain alias
	domPatchReq := &pb.PatchDomainRequest{
		Id: &wrappers.StringValue{
			Value: testDomainID,
		},
		Alias: &wrappers.StringValue{
			Value: testStr,
		},
	}
	domPatchRet, err := suite.s.PatchDomain(suite.ctx, domPatchReq)
	suite.Nil(err)
	suite.Equal(testStr, domPatchRet.GetDomain().GetAlias())

	//test patch domain extra
	domPatchReq.Extra = testDomainExtra
	domPatchRet, err = suite.s.PatchDomain(suite.ctx, domPatchReq)
	extraMap := domPatchRet.GetDomain().GetExtra()
	suite.Nil(err)
	suite.Equal(testDomainExtra["test"].GetValue(), extraMap["test"])

	//test list domains by name (create 2 above)
	domListReq := &pb.ListDomainsRequest{
		Name: &wrappers.StringValue{
			Value: testDomainName,
		},
	}
	domsRet, err := suite.s.ListDomains(suite.ctx, domListReq)
	suite.Nil(err)
	suite.Len(domsRet.GetDomains(), 2)

	//test list domains by Alias (create 2 above, change 1, left 1)
	domListReq = &pb.ListDomainsRequest{
		Alias: &wrappers.StringValue{
			Value: testDomainAlias,
		},
	}
	domsRet, err = suite.s.ListDomains(suite.ctx, domListReq)
	suite.Nil(err)
	suite.Len(domsRet.GetDomains(), 1)

	//test delete domain default (must be err: more than 0 children in domain)
	domDeleteReq := &pb.DeleteDomainRequest{
		Domain: &pb.OpDomain{
			Id: &wrappers.StringValue{
				Value: defaultDomainID,
			},
		},
	}
	_, err = suite.s.DeleteDomain(suite.ctx, domDeleteReq)
	suite.NotNil(err)

	//test delete domain not default and no children
	domDeleteReq = &pb.DeleteDomainRequest{
		Domain: &pb.OpDomain{
			Id: &wrappers.StringValue{
				Value: testDomainID,
			},
		},
	}
	_, err = suite.s.DeleteDomain(suite.ctx, domDeleteReq)
	suite.Nil(err)
}

func (suite *identifydTestSuite) TestRole() {
	testStr := "test"

	//test create role(create 1st)
	rolCreateReq := &pb.CreateRoleRequest{
		Id: &wrappers.StringValue{
			Value: testRoleID,
		},
		Domain: &pb.OpDomain{
			Id: &wrappers.StringValue{
				Value: defaultDomainID,
			},
		},
		Name: &wrappers.StringValue{
			Value: testRoleName,
		},
		Alias: &wrappers.StringValue{
			Value: testRoleAlias,
		},
		Description: &wrappers.StringValue{
			Value: testRoleDescription,
		},
		Extra: map[string]*wrappers.StringValue{
			"test": &wrappers.StringValue{
				Value: "",
			},
		},
	}
	_, err := suite.s.CreateRole(suite.ctx, rolCreateReq)
	suite.Nil(err)

	//test create role with no id(create 2st)
	rolCreateReq.Id = nil
	_, err = suite.s.CreateRole(suite.ctx, rolCreateReq)
	suite.Nil(err)

	//test get role
	roleGetReq := &pb.GetRoleRequest{
		Role: &pb.OpRole{
			Id: &wrappers.StringValue{
				Value: testRoleID,
			},
		},
	}
	rolGetRet, err := suite.s.GetRole(suite.ctx, roleGetReq)
	suite.Nil(err)
	suite.Equal(rolGetRet.GetRole().GetName(), testRoleName)

	//test patch role alias
	rolPatchReq := &pb.PatchRoleRequest{
		Id: &wrappers.StringValue{
			Value: testRoleID,
		},
		Alias: &wrappers.StringValue{
			Value: testStr,
		},
	}
	rolPatchRet, err := suite.s.PatchRole(suite.ctx, rolPatchReq)
	suite.Nil(err)
	suite.Equal(testStr, rolPatchRet.GetRole().GetAlias())

	//test patch role Description
	rolPatchReq.Alias = nil
	rolPatchReq.Description = &wrappers.StringValue{
		Value: testStr,
	}
	rolPatchRet, err = suite.s.PatchRole(suite.ctx, rolPatchReq)
	suite.Nil(err)
	suite.Equal(testStr, rolPatchRet.GetRole().GetDescription())

	//test patch role Extra
	rolPatchReq.Description = nil
	rolPatchReq.Extra = testRoleExtra
	rolPatchRet, err = suite.s.PatchRole(suite.ctx, rolPatchReq)
	extraMap := rolPatchRet.GetRole().GetExtra()
	suite.Nil(err)
	suite.Equal(testRoleExtra["test"].GetValue(), extraMap["test"])

	//test list roles by name (create 2 above)
	rolListReq := &pb.ListRolesRequest{
		Name: &wrappers.StringValue{
			Value: testRoleName,
		},
	}
	rolsRet, err := suite.s.ListRoles(suite.ctx, rolListReq)
	suite.Nil(err)
	suite.Len(rolsRet.GetRoles(), 2)

	//test list roles by Alias (create 2 above, change 1, left 1)
	rolListReq = &pb.ListRolesRequest{
		Alias: &wrappers.StringValue{
			Value: testRoleAlias,
		},
	}
	rolsRet, err = suite.s.ListRoles(suite.ctx, rolListReq)
	suite.Nil(err)
	suite.Len(rolsRet.GetRoles(), 1)

	//test list roles by domain (create 2 above)
	rolListReq = &pb.ListRolesRequest{
		Domain: &pb.OpDomain{
			Id: &wrappers.StringValue{
				Value: defaultDomainID,
			},
		},
	}
	rolsRet, err = suite.s.ListRoles(suite.ctx, rolListReq)
	suite.Nil(err)
	suite.Len(rolsRet.GetRoles(), 2)

	//test delete role
	rolDeleteReq := &pb.DeleteRoleRequest{
		Role: &pb.OpRole{
			Id: &wrappers.StringValue{
				Value: testRoleID,
			},
		},
	}
	_, err = suite.s.DeleteRole(suite.ctx, rolDeleteReq)
	suite.Nil(err)
}

func (suite *identifydTestSuite) TestEntity() {
	testStr := "test"

	//test create Entity(create 1st)
	entCreateReq := &pb.CreateEntityRequest{
		Id: &wrappers.StringValue{
			Value: testEntityID,
		},
		Name: &wrappers.StringValue{
			Value: testEntityName,
		},
		Alias: &wrappers.StringValue{
			Value: testEntityAlias,
		},
		Password: &wrappers.StringValue{
			Value: testEntityPassword,
		},
		Extra: map[string]*wrappers.StringValue{
			"test": &wrappers.StringValue{
				Value: "",
			},
		},
	}
	_, err := suite.s.CreateEntity(suite.ctx, entCreateReq)
	suite.Nil(err)

	//test create entity with no id(create 2st)
	entCreateReq.Id = nil
	_, err = suite.s.CreateEntity(suite.ctx, entCreateReq)
	suite.Nil(err)

	//test get entity
	entGetReq := &pb.GetEntityRequest{
		Entity: &pb.OpEntity{
			Id: &wrappers.StringValue{
				Value: testEntityID,
			},
		},
	}
	entGetRet, err := suite.s.GetEntity(suite.ctx, entGetReq)
	suite.Nil(err)
	suite.Equal(entGetRet.GetEntity().GetName(), testEntityName)

	//test patch entity alias
	entPatchReq := &pb.PatchEntityRequest{
		Id: &wrappers.StringValue{
			Value: testEntityID,
		},
		Alias: &wrappers.StringValue{
			Value: testStr,
		},
	}
	entPatchRet, err := suite.s.PatchEntity(suite.ctx, entPatchReq)
	suite.Nil(err)
	suite.Equal(testStr, entPatchRet.GetEntity().GetAlias())

	//test patch entity Password(password no return)
	entPatchReq.Alias = nil
	entPatchReq.Password = &wrappers.StringValue{
		Value: testEntityPassword,
	}
	entPatchRet, err = suite.s.PatchEntity(suite.ctx, entPatchReq)
	suite.Nil(err)

	//test patch entity Extra
	entPatchReq.Password = nil
	entPatchReq.Extra = testEntityExtra
	entPatchRet, err = suite.s.PatchEntity(suite.ctx, entPatchReq)
	extraMap := entPatchRet.GetEntity().GetExtra()
	suite.Nil(err)
	suite.Equal(testEntityExtra["test"].GetValue(), extraMap["test"])

	//test list entities by name (create 2 above)
	entListReq := &pb.ListEntitiesRequest{
		Name: &wrappers.StringValue{
			Value: testEntityName,
		},
	}
	entsRet, err := suite.s.ListEntities(suite.ctx, entListReq)
	suite.Nil(err)
	suite.Len(entsRet.GetEntities(), 2)

	//test list entities by alias (create 2 above, change 1, left 1)
	entListReq = &pb.ListEntitiesRequest{
		Alias: &wrappers.StringValue{
			Value: testEntityAlias,
		},
	}
	entsRet, err = suite.s.ListEntities(suite.ctx, entListReq)
	suite.Nil(err)
	suite.Len(entsRet.GetEntities(), 1)

	//test add role to entity
	entAddRoleToEntityReq := &pb.AddRoleToEntityRequest{
		Entity: &pb.OpEntity{
			Id: &wrappers.StringValue{
				Value: testEntityID,
			},
		},
		Role: &pb.OpRole{
			Id: &wrappers.StringValue{
				Value: defaultSysadminID,
			},
		},
	}
	_, err = suite.s.AddRoleToEntity(suite.ctx, entAddRoleToEntityReq)
	suite.Nil(err)

	//test remove role from entity
	entRemoveRoleFromEntityReq := &pb.RemoveRoleFromEntityRequest{
		Entity: &pb.OpEntity{
			Id: &wrappers.StringValue{
				Value: testEntityID,
			},
		},
		Role: &pb.OpRole{
			Id: &wrappers.StringValue{
				Value: defaultSysadminID,
			},
		},
	}
	_, err = suite.s.RemoveRoleFromEntity(suite.ctx, entRemoveRoleFromEntityReq)
	suite.Nil(err)

	//test add Entity to domain
	entAddEntityToDomainReq := &pb.AddEntityToDomainRequest{
		Entity: &pb.OpEntity{
			Id: &wrappers.StringValue{
				Value: testEntityID,
			},
		},
		Domain: &pb.OpDomain{
			Id: &wrappers.StringValue{
				Value: defaultDomainID,
			},
		},
	}
	_, err = suite.s.AddEntityToDomain(suite.ctx, entAddEntityToDomainReq)
	suite.Nil(err)

	//test remove entity from domain
	entRemoveEntityFromDomainReq := &pb.RemoveEntityFromDomainRequest{
		Entity: &pb.OpEntity{
			Id: &wrappers.StringValue{
				Value: testEntityID,
			},
		},
		Domain: &pb.OpDomain{
			Id: &wrappers.StringValue{
				Value: defaultDomainID,
			},
		},
	}
	_, err = suite.s.RemoveEntityFromDomain(suite.ctx, entRemoveEntityFromDomainReq)
	suite.Nil(err)

	//test delete entity
	entDeleteReq := &pb.DeleteEntityRequest{
		Entity: &pb.OpEntity{
			Id: &wrappers.StringValue{
				Value: testEntityID,
			},
		},
	}
	_, err = suite.s.DeleteEntity(suite.ctx, entDeleteReq)
	suite.Nil(err)
}

func (suite *identifydTestSuite) TestGroup() {
	testStr := "test"

	//test create Group(create 1st)
	grpCreateReq := &pb.CreateGroupRequest{
		Id: &wrappers.StringValue{
			Value: testGroupID,
		},
		Domain: &pb.OpDomain{
			Id: &wrappers.StringValue{
				Value: defaultDomainID,
			},
		},
		Name: &wrappers.StringValue{
			Value: testGroupName,
		},
		Alias: &wrappers.StringValue{
			Value: testGroupAlias,
		},
		Description: &wrappers.StringValue{
			Value: testGroupDescrition,
		},
		Extra: map[string]*wrappers.StringValue{
			"test": &wrappers.StringValue{
				Value: "",
			},
		},
	}
	_, err := suite.s.CreateGroup(suite.ctx, grpCreateReq)
	suite.Nil(err)

	//test create group with no id(create 2st)
	grpCreateReq.Id = nil
	_, err = suite.s.CreateGroup(suite.ctx, grpCreateReq)
	suite.Nil(err)

	//test get group
	grpGetReq := &pb.GetGroupRequest{
		Group: &pb.OpGroup{
			Id: &wrappers.StringValue{
				Value: testGroupID,
			},
		},
	}
	grpGetRet, err := suite.s.GetGroup(suite.ctx, grpGetReq)
	suite.Nil(err)
	suite.Equal(grpGetRet.GetGroup().GetName(), testGroupName)

	//test patch group alias
	grpPatchReq := &pb.PatchGroupRequest{
		Id: &wrappers.StringValue{
			Value: testGroupID,
		},
		Alias: &wrappers.StringValue{
			Value: testStr,
		},
	}
	grpPatchRet, err := suite.s.PatchGroup(suite.ctx, grpPatchReq)
	suite.Nil(err)
	suite.Equal(testStr, grpPatchRet.GetGroup().GetAlias())

	//test patch group description
	grpPatchReq.Alias = nil
	grpPatchReq.Description = &wrappers.StringValue{
		Value: testGroupDescrition,
	}
	grpPatchRet, err = suite.s.PatchGroup(suite.ctx, grpPatchReq)
	suite.Nil(err)

	//test patch group Extra
	grpPatchReq.Description = nil
	grpPatchReq.Extra = testGroupExtra
	grpPatchRet, err = suite.s.PatchGroup(suite.ctx, grpPatchReq)
	extraMap := grpPatchRet.GetGroup().GetExtra()
	suite.Nil(err)
	suite.Equal(testGroupExtra["test"].GetValue(), extraMap["test"])

	//test list groups by name (create 2 above)
	grpListReq := &pb.ListGroupsRequest{
		Name: &wrappers.StringValue{
			Value: testGroupName,
		},
	}
	grpsRet, err := suite.s.ListGroups(suite.ctx, grpListReq)
	suite.Nil(err)
	suite.Len(grpsRet.GetGroups(), 2)

	//test list groups by alias (create 2 above, change 1, left 1)
	grpListReq = &pb.ListGroupsRequest{
		Alias: &wrappers.StringValue{
			Value: testGroupAlias,
		},
	}
	grpsRet, err = suite.s.ListGroups(suite.ctx, grpListReq)
	suite.Nil(err)
	suite.Len(grpsRet.GetGroups(), 1)

	//test add role to group
	grpAddRoleToGroupReq := &pb.AddRoleToGroupRequest{
		Group: &pb.OpGroup{
			Id: &wrappers.StringValue{
				Value: testGroupID,
			},
		},
		Role: &pb.OpRole{
			Id: &wrappers.StringValue{
				Value: defaultSysadminID,
			},
		},
	}
	_, err = suite.s.AddRoleToGroup(suite.ctx, grpAddRoleToGroupReq)
	suite.Nil(err)

	//test remove role from group
	grpRemoveRoleFromGroupReq := &pb.RemoveRoleFromGroupRequest{
		Group: &pb.OpGroup{
			Id: &wrappers.StringValue{
				Value: testGroupID,
			},
		},
		Role: &pb.OpRole{
			Id: &wrappers.StringValue{
				Value: defaultSysadminID,
			},
		},
	}
	_, err = suite.s.RemoveRoleFromGroup(suite.ctx, grpRemoveRoleFromGroupReq)
	suite.Nil(err)

	//test add Entity to group
	///create Entity first
	entCreateReq := &pb.CreateEntityRequest{
		Id: &wrappers.StringValue{
			Value: testEntityID,
		},
		Name: &wrappers.StringValue{
			Value: testEntityName,
		},
		Alias: &wrappers.StringValue{
			Value: testEntityAlias,
		},
		Password: &wrappers.StringValue{
			Value: testEntityPassword,
		},
		Extra: map[string]*wrappers.StringValue{
			"test": &wrappers.StringValue{
				Value: "",
			},
		},
	}
	_, err = suite.s.CreateEntity(suite.ctx, entCreateReq)
	suite.Nil(err)
	grpAddEntityToGroupReq := &pb.AddEntityToGroupRequest{
		Entity: &pb.OpEntity{
			Id: &wrappers.StringValue{
				Value: testEntityID,
			},
		},
		Group: &pb.OpGroup{
			Id: &wrappers.StringValue{
				Value: testGroupID,
			},
		},
	}
	_, err = suite.s.AddEntityToGroup(suite.ctx, grpAddEntityToGroupReq)
	suite.Nil(err)

	//test remove entity from group
	grpRemoveEntityFromGroupReq := &pb.RemoveEntityFromGroupRequest{
		Entity: &pb.OpEntity{
			Id: &wrappers.StringValue{
				Value: testEntityID,
			},
		},
		Group: &pb.OpGroup{
			Id: &wrappers.StringValue{
				Value: testGroupID,
			},
		},
	}
	_, err = suite.s.RemoveEntityFromGroup(suite.ctx, grpRemoveEntityFromGroupReq)
	suite.Nil(err)

	//test delete group
	grpDeleteReq := &pb.DeleteGroupRequest{
		Group: &pb.OpGroup{
			Id: &wrappers.StringValue{
				Value: testGroupID,
			},
		},
	}
	_, err = suite.s.DeleteGroup(suite.ctx, grpDeleteReq)
	suite.Nil(err)
}

func (suite *identifydTestSuite) TestCredential() {
	testStr := "test"

	//test create Credential(create 1st)
	credCreateReq := &pb.CreateCredentialRequest{
		Id: &wrappers.StringValue{
			Value: testCredentialID,
		},
		Domain: &pb.OpDomain{
			Id: &wrappers.StringValue{
				Value: testDomainID,
			},
		},
		Entity: &pb.OpEntity{
			Id: &wrappers.StringValue{
				Value: testEntityID,
			},
		},
		Name: &wrappers.StringValue{
			Value: testCredentialName,
		},
		Alias: &wrappers.StringValue{
			Value: testCredentialAlias,
		},
		Description: &wrappers.StringValue{
			Value: testCredentialDescription,
		},
		ExpiresAt: &timestamp.Timestamp{
			Seconds: pb_helper.FromTime(testCredentialExpires).Seconds,
			Nanos:   pb_helper.FromTime(testCredentialExpires).Nanos,
		},
	}
	_, err := suite.s.CreateCredential(suite.ctx, credCreateReq)
	suite.Nil(err)

	//test create cred with no id(create 2st)
	credCreateReq.Id = nil
	_, err = suite.s.CreateCredential(suite.ctx, credCreateReq)
	suite.Nil(err)

	//test get credential
	credGetReq := &pb.GetCredentialRequest{
		Credential: &pb.OpCredential{
			Id: &wrappers.StringValue{
				Value: testCredentialID,
			},
		},
	}
	credGetRet, err := suite.s.GetCredential(suite.ctx, credGetReq)
	suite.Nil(err)
	suite.Equal(credGetRet.GetCredential().GetName(), testCredentialName)

	//test patch credential alias
	credPatchReq := &pb.PatchCredentialRequest{
		Id: &wrappers.StringValue{
			Value: testCredentialID,
		},
		Alias: &wrappers.StringValue{
			Value: testStr,
		},
	}
	credPatchRet, err := suite.s.PatchCredential(suite.ctx, credPatchReq)
	suite.Nil(err)
	suite.Equal(testStr, credPatchRet.GetCredential().GetAlias())

	//test patch credential description
	credPatchReq.Alias = nil
	credPatchReq.Description = &wrappers.StringValue{
		Value: testCredentialDescription,
	}
	credPatchRet, err = suite.s.PatchCredential(suite.ctx, credPatchReq)
	suite.Nil(err)

	//test list credentials by name (create 2 above)
	credListReq := &pb.ListCredentialsRequest{
		Name: &wrappers.StringValue{
			Value: testCredentialName,
		},
	}
	credsRet, err := suite.s.ListCredentials(suite.ctx, credListReq)
	suite.Nil(err)
	suite.Len(credsRet.GetCredentials(), 2)

	//test list credential by alias (create 2 above, change 1, left 1)
	credListReq = &pb.ListCredentialsRequest{
		Alias: &wrappers.StringValue{
			Value: testCredentialAlias,
		},
	}
	credsRet, err = suite.s.ListCredentials(suite.ctx, credListReq)
	suite.Nil(err)
	suite.Len(credsRet.GetCredentials(), 1)

	//test delete credential
	credDeleteReq := &pb.DeleteCredentialRequest{
		Credential: &pb.OpCredential{
			Id: &wrappers.StringValue{
				Value: testCredentialID,
			},
		},
	}
	_, err = suite.s.DeleteCredential(suite.ctx, credDeleteReq)
	suite.Nil(err)
}

func (suite *identifydTestSuite) TestIssueTokenByCredential() {
	//create Credential
	credCreateReq := &pb.CreateCredentialRequest{
		Id: &wrappers.StringValue{
			Value: testCredentialID,
		},
		Domain: &pb.OpDomain{
			Id: &wrappers.StringValue{
				Value: testDomainID,
			},
		},
		Entity: &pb.OpEntity{
			Id: &wrappers.StringValue{
				Value: testEntityID,
			},
		},
		Name: &wrappers.StringValue{
			Value: testCredentialName,
		},
		Alias: &wrappers.StringValue{
			Value: testCredentialAlias,
		},
		Description: &wrappers.StringValue{
			Value: testCredentialDescription,
		},
		ExpiresAt: &timestamp.Timestamp{
			Seconds: pb_helper.FromTime(testCredentialExpires).Seconds,
			Nanos:   pb_helper.FromTime(testCredentialExpires).Nanos,
		},
	}
	credCreateRet, err := suite.s.CreateCredential(suite.ctx, credCreateReq)
	suite.Nil(err)
	credSecretStr := credCreateRet.GetCredential().GetSecret()

	tknIssueTokenByCredentialReq := &pb.IssueTokenByCredentialRequest{
		Credential: &pb.OpCredential{
			Id: &wrappers.StringValue{
				Value: testCredentialID,
			},
			Domain: &pb.OpDomain{
				Id: &wrappers.StringValue{
					Value: testDomainID,
				},
			},
			Secret: &wrappers.StringValue{
				Value: credSecretStr,
			},
		},
	}
	tknIssueTokenByCredentialRet, err := suite.s.IssueTokenByCredential(suite.ctx, tknIssueTokenByCredentialReq)
	suite.Nil(err)
	suite.NotNil(tknIssueTokenByCredentialRet)
}

func (suite *identifydTestSuite) TestIssueToken() {
	//1st create Entity
	entCreateReq := &pb.CreateEntityRequest{
		Id: &wrappers.StringValue{
			Value: testEntityID,
		},
		Name: &wrappers.StringValue{
			Value: testEntityName,
		},
		Alias: &wrappers.StringValue{
			Value: testEntityAlias,
		},
		Password: &wrappers.StringValue{
			Value: testEntityPassword,
		},
		Extra: map[string]*wrappers.StringValue{
			"test": &wrappers.StringValue{
				Value: "",
			},
		},
	}
	_, err := suite.s.CreateEntity(suite.ctx, entCreateReq)
	suite.Nil(err)

	//2st add entity to default domain
	entAddEntityToDomainReq := &pb.AddEntityToDomainRequest{
		Entity: &pb.OpEntity{
			Id: &wrappers.StringValue{
				Value: testEntityID,
			},
		},
		Domain: &pb.OpDomain{
			Id: &wrappers.StringValue{
				Value: defaultDomainID,
			},
		},
	}
	_, err = suite.s.AddEntityToDomain(suite.ctx, entAddEntityToDomainReq)
	suite.Nil(err)

	//test IssueTokenByPassword use entity id
	tknIssueTokenByPasswordReq := &pb.IssueTokenByPasswordRequest{
		Entity: &pb.OpEntity{
			Id: &wrappers.StringValue{
				Value: testEntityID,
			},
			Password: &wrappers.StringValue{
				Value: testEntityPassword,
			},
		},
	}
	tknIssueTokenByPasswordReq.Entity.Domains = append(tknIssueTokenByPasswordReq.Entity.Domains, &pb.OpDomain{
		Id: &wrappers.StringValue{
			Value: defaultDomainID,
		},
	})

	tknIssueTokenByPasswordRet, err := suite.s.IssueTokenByPassword(suite.ctx, tknIssueTokenByPasswordReq)
	suite.Nil(err)
	suite.NotNil(tknIssueTokenByPasswordRet)

	//test IssueTokenByPassword use entity name
	tknIssueTokenByPasswordReq.Entity.Id = nil
	tknIssueTokenByPasswordReq.Entity.Name = &wrappers.StringValue{
		Value: testEntityName,
	}

	tknIssueTokenByPasswordRet, err = suite.s.IssueTokenByPassword(suite.ctx, tknIssueTokenByPasswordReq)
	suite.Nil(err)
	suite.NotNil(tknIssueTokenByPasswordRet)

	//test IssueTokenByToken
	tkn := tknIssueTokenByPasswordRet.Token.Text
	tknIssueTokenByTokenReq := &pb.IssueTokenByTokenRequest{
		Token: &pb.OpToken{
			Domain: &pb.OpDomain{
				Id: &wrappers.StringValue{
					Value: defaultDomainID,
				},
			},
			Text: &wrappers.StringValue{
				Value: tkn,
			},
		},
	}
	tknIssueTokenByTokenRet, err := suite.s.IssueTokenByToken(suite.ctx, tknIssueTokenByTokenReq)
	suite.Nil(err)
	suite.NotNil(tknIssueTokenByTokenRet)

	//test validateToken
	tkn = tknIssueTokenByTokenRet.Token.Text
	tknValidateTokenReq := &pb.ValidateTokenRequest{
		Token: &pb.OpToken{
			Text: &wrappers.StringValue{
				Value: tkn,
			},
		},
	}
	tknValidateTokenRet, err := suite.s.ValidateToken(suite.ctx, tknValidateTokenReq)
	suite.Nil(err)
	suite.NotNil(tknValidateTokenRet)

	//test checktoken
	tkn = tknIssueTokenByTokenRet.Token.Text
	tknCheckTokenReq := &pb.CheckTokenRequest{
		Token: &pb.OpToken{
			Text: &wrappers.StringValue{
				Value: tkn,
			},
			Domain: &pb.OpDomain{
				Id: &wrappers.StringValue{
					Value: defaultDomainID,
				},
			},
		},
	}
	_, err = suite.s.CheckToken(suite.ctx, tknCheckTokenReq)
	suite.Nil(err)

	//test revoke token
	tkn = tknIssueTokenByTokenRet.Token.Text
	tknRevokeTokenReq := &pb.RevokeTokenRequest{
		Token: &pb.OpToken{
			Text: &wrappers.StringValue{
				Value: tkn,
			},
		},
	}
	_, err = suite.s.RevokeToken(suite.ctx, tknRevokeTokenReq)
	suite.Nil(err)
}

func TestIdentifydTestSuite(t *testing.T) {
	suite.Run(t, new(identifydTestSuite))
}