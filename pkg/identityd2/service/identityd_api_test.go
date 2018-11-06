package metathings_identityd2_service

import (
	"context"
	"fmt"
	"testing"

	gomock "github.com/golang/mock/gomock"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"
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

	if suite.s, err = NewMetathingsIdentitydService(newEnforcer, nil, log.New(), newStorage); err != nil {
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

	//test create role
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

	//test create role with no id
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
				Value: defaultDomainID,
			},
		},
	}
	_, err = suite.s.DeleteRole(suite.ctx, rolDeleteReq)
	suite.Nil(err)
}

func TestIdentifydTestSuite(t *testing.T) {
	suite.Run(t, new(identifydTestSuite))
}
