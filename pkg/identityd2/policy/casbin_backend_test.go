package metathings_identityd2_policy

import (
	"context"
	"testing"

	casbin_pb "github.com/casbin/casbin-server/proto"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	cmd_contrib "github.com/nayotta/metathings/cmd/contrib"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	test_helper "github.com/nayotta/metathings/pkg/common/test"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
)

var (
	test_domain_id          = "test-domain-id"
	test_group_id           = "test-group-id"
	test_subject_id         = "test-subject-id"
	test_subject2_id        = "test-subject2-id"
	test_object_id          = "test-object-id"
	test_object2_id         = "test-object2-id"
	test_role_id            = "test-role-id"
	test_role_name          = "test-role-name"
	test_role2_id           = "test-role2-id"
	test_role2_name         = "test-role2-name"
	test_action_id          = "test-action-id"
	test_action_name        = "test-action-name"
	test_action2_id         = "test-action2-id"
	test_action2_name       = "test-action2-name"
	test_entity_id          = "test-entity-id"
	test_sysadmin_role_id   = "test-sysadmin-role-id"
	test_sysadmin_role_name = "sysadmin"
)

var (
	test_action        *storage.Action
	test_action2       *storage.Action
	test_role          *storage.Role
	test_role2         *storage.Role
	test_group         *storage.Group
	test_subject       *storage.Entity
	test_subject2      *storage.Entity
	test_object        *storage.Entity
	test_object2       *storage.Entity
	test_entity        *storage.Entity
	test_sysadmin_role *storage.Role
)

type casbinBackendTestSuite struct {
	suite.Suite
	b                *CasbinBackend
	enforcer_handler int32
}

func (s *casbinBackendTestSuite) SetupTest() {
	var err error
	var b Backend
	var logger log.FieldLogger
	var cli_fty *client_helper.ClientFactory
	opt := new_casbin_backend_option()
	mdl_txt := `[request_definition]
r = sub, grp, obj, act

[policy_definition]
p = sub, grp, obj, act

[role_definition]
g = _, _
g2 = _, _, _
g3 = _, _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = (g2(r.sub, r.grp, p.sub) && g3(r.obj, r.grp, p.obj) && r.grp == p.grp && r.act == p.act) || (g2(r.sub, r.grp, p.sub) && p.sub == "rol.sysadmin")
`

	test_action = &storage.Action{
		Id:   &test_action_id,
		Name: &test_action_name,
	}
	test_action2 = &storage.Action{
		Id:   &test_action2_id,
		Name: &test_action2_name,
	}
	test_role = &storage.Role{
		Id:   &test_role_id,
		Name: &test_role_name,
		Actions: []*storage.Action{
			test_action,
		},
	}
	test_role2 = &storage.Role{
		Id:   &test_role2_id,
		Name: &test_role2_name,
		Actions: []*storage.Action{
			test_action2,
		},
	}
	test_group = &storage.Group{
		Id:       &test_domain_id,
		DomainId: &test_domain_id,
		Domain:   &storage.Domain{Id: &test_domain_id},
		Roles: []*storage.Role{
			test_role,
		},
	}
	test_subject = &storage.Entity{
		Id: &test_subject_id,
		Groups: []*storage.Group{
			test_group,
		},
	}
	test_subject2 = &storage.Entity{
		Id: &test_subject2_id,
		Groups: []*storage.Group{
			test_group,
		},
	}
	test_object = &storage.Entity{Id: &test_object_id}
	test_object2 = &storage.Entity{Id: &test_object2_id}
	test_entity = &storage.Entity{Id: &test_entity_id}
	test_sysadmin_role = &storage.Role{
		Id:   &test_sysadmin_role_id,
		Name: &test_sysadmin_role_name,
	}

	test_group.Subjects = []*storage.Entity{test_subject}
	test_group.Objects = []*storage.Entity{test_object}

	logger, err = log_helper.NewLogger("test", "debug")
	s.Nil(err)

	srv_opt := cmd_contrib.CreateServiceEndpointsOption()
	srv_opt.ServiceEndpoint[client_helper.DEFAULT_CONFIG.String()].Address = test_helper.GetTestPolicydAddress()
	cli_fty, err = cmd_contrib.NewClientFactory(&srv_opt, nil, logger)
	s.Nil(err)

	cli, cfn, err := cli_fty.NewPolicydServiceClient()
	s.Nil(err)
	defer cfn()

	drv_name := test_helper.GetTestPolicydDriverName()
	conn_str := test_helper.GetTestPolicydConnectString()
	new_adapter_res, err := cli.NewAdapter(context.TODO(), &casbin_pb.NewAdapterRequest{
		DriverName:    drv_name,
		ConnectString: conn_str,
		DbSpecified:   true,
	})
	s.Nil(err)

	new_enforcer_res, err := cli.NewEnforcer(context.TODO(), &casbin_pb.NewEnforcerRequest{
		ModelText:     mdl_txt,
		AdapterHandle: new_adapter_res.Handler,
	})
	s.Nil(err)

	opt.EnforcerHandler = new_enforcer_res.Handler
	s.enforcer_handler = new_adapter_res.Handler

	_, err = cli.AddPolicy(context.TODO(), &casbin_pb.PolicyRequest{
		EnforcerHandler: new_adapter_res.Handler,
		Params:          []string{"rol.sysadmin", CASBIN_BACKEND_UNGROUPING, "any", "any"},
	})
	s.Nil(err)

	b, err = casbin_backend_factory("logger", logger, "client_factory", cli_fty, "casbin_enforcer_handler", s.enforcer_handler)
	s.Nil(err)

	s.b = b.(*CasbinBackend)

	s.Nil(s.b.AddRoleToGroup(test_group, test_role))
	s.Nil(s.b.AddSubjectToGroup(test_group, test_subject))
	s.Nil(s.b.AddObjectToGroup(test_group, test_object))
	s.Nil(s.b.AddRoleToEntity(test_entity, test_sysadmin_role))
}

func (s *casbinBackendTestSuite) TestEnforce() {
	s.Nil(s.b.Enforce(test_subject, test_object, test_action))
	s.NotNil(s.b.Enforce(test_subject2, test_object, test_action))
	s.NotNil(s.b.Enforce(test_subject, test_object2, test_action))
	s.NotNil(s.b.Enforce(test_subject, test_object, test_action2))
	s.Nil(s.b.Enforce(test_entity, test_object, test_action))
}

func (s *casbinBackendTestSuite) TestDeleteGroup() {
	cli, cfn, err := s.b.cli_fty.NewPolicydServiceClient()
	s.Nil(err)
	defer cfn()

	s.Nil(s.b.DeleteGroup(test_group))

	rs, err := s.b._list_grouping_policies(cli, CASBIN_BACKEND_SUBJECT_PTYPE, ConvertSubject(test_subject), ConvertGroup(test_group))
	s.Nil(err)
	s.Len(rs, 0)

	rs, err = s.b._list_grouping_policies(cli, CASBIN_BACKEND_OBJECT_PTYPE, ConvertObject(test_object), ConvertGroup(test_group))
	s.Nil(err)
	s.Len(rs, 0)

	s.NotNil(s.b.Enforce(test_subject, test_object, test_action))
}

func (s *casbinBackendTestSuite) TestAddSubjectToGroup() {
	cli, cfn, err := s.b.cli_fty.NewPolicydServiceClient()
	s.Nil(err)
	defer cfn()

	s.Nil(s.b.AddSubjectToGroup(test_group, test_subject2))

	rs, err := s.b._list_grouping_policies(cli, CASBIN_BACKEND_SUBJECT_PTYPE, ConvertSubject(test_subject2), ConvertGroup(test_group))
	s.Nil(err)
	s.Len(rs, 1)

	s.Nil(s.b.Enforce(test_subject2, test_object, test_action))
}

func (s *casbinBackendTestSuite) TestRemoveSubjectFromGroup() {
	cli, cfn, err := s.b.cli_fty.NewPolicydServiceClient()
	s.Nil(err)
	defer cfn()

	s.Nil(s.b.RemoveSubjectFromGroup(test_group, test_subject))

	rs, err := s.b._list_grouping_policies(cli, CASBIN_BACKEND_SUBJECT_PTYPE, ConvertSubject(test_subject), ConvertGroup(test_group))
	s.Nil(err)
	s.Len(rs, 0)

	s.NotNil(s.b.Enforce(test_subject, test_object, test_action))
}

func (s *casbinBackendTestSuite) TestAddObjectToGroup() {
	cli, cfn, err := s.b.cli_fty.NewPolicydServiceClient()
	s.Nil(err)
	defer cfn()

	s.Nil(s.b.AddObjectToGroup(test_group, test_object2))

	rs, err := s.b._list_grouping_policies(cli, CASBIN_BACKEND_OBJECT_PTYPE, ConvertObject(test_object2), ConvertGroup(test_group))
	s.Nil(err)
	s.Len(rs, 1)

	s.Nil(s.b.Enforce(test_subject, test_object2, test_action))
}

func (s *casbinBackendTestSuite) TestRemoveObjectFromGroup() {
	cli, cfn, err := s.b.cli_fty.NewPolicydServiceClient()
	s.Nil(err)
	defer cfn()

	s.Nil(s.b.RemoveObjectFromGroup(test_group, test_object))

	rs, err := s.b._list_grouping_policies(cli, CASBIN_BACKEND_OBJECT_PTYPE, ConvertObject(test_object), ConvertGroup(test_group))
	s.Nil(err)
	s.Len(rs, 0)

	s.NotNil(s.b.Enforce(test_subject, test_object, test_action))
}

func (s *casbinBackendTestSuite) TestAddRoleToGroup() {
	s.Nil(s.b.AddRoleToGroup(test_group, test_role2))
	s.Nil(s.b.Enforce(test_subject, test_object, test_action2))
}

func (s *casbinBackendTestSuite) TestRemoveRoleFromGroup() {
	s.Nil(s.b.RemoveRoleFromGroup(test_group, test_role))
	s.NotNil(s.b.Enforce(test_subject, test_object, test_action))
}

func (s *casbinBackendTestSuite) TestAddRoleToEntity() {
	s.Nil(s.b.AddRoleToEntity(test_subject, test_sysadmin_role))
	s.Nil(s.b.Enforce(test_subject, test_object, test_action2))
}

func (s *casbinBackendTestSuite) TestRemoveRoleFromEntity() {
	s.Nil(s.b.RemoveRoleFromEntity(test_entity, test_sysadmin_role))
	s.NotNil(s.b.Enforce(test_entity, test_object, test_action))
}

func TestCasbinBackendTestSuite(t *testing.T) {
	suite.Run(t, new(casbinBackendTestSuite))
}
