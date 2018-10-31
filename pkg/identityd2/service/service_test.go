package metathings_identityd2_service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
)

type mock_enforcer struct {
	mock.Mock
}

func (self *mock_enforcer) Enforce(domain, group, subject, object, action interface{}) error {
	self.Called(domain, group, subject, object, action)
	if action.(string) == "pass" {
		return nil
	} else {
		return ErrPermissionDenied
	}
}

func (self *mock_enforcer) AddGroup(domain, group string) error {
	panic("unimplemented")
}

func (self *mock_enforcer) RemoveGroup(domain, group string) error {
	panic("unimplemented")
}

func (self *mock_enforcer) AddSubjectToRole(domain, group, subject, role string) error {
	panic("unimplemented")
}

func (self *mock_enforcer) RemoveSubjectFromRole(domain, group, subject, role string) error {
	panic("unimplemented")
}

func (self *mock_enforcer) AddObjectToKind(domain, group, object, kind string) error {
	panic("unimplemented")
}

func (self *mock_enforcer) RemoveObjectFromKind(domain, group, object, kind string) error {
	panic("unimplemented")
}

type metathingsIdentitydService_enforceTestSuite struct {
	suite.Suite

	subject       string
	domain        string
	group         string
	groups        []string
	object        string
	action_pass   string
	action_nopass string
	enforcer      *mock_enforcer
	service       *MetathingsIdentitydService
	token         *storage.Token
	ctx           context.Context
}

func (self *metathingsIdentitydService_enforceTestSuite) SetupTest() {

	self.subject = "subject"
	self.domain = "domain"
	self.group = "group"
	self.groups = []string{self.group}
	self.object = "object"
	self.action_pass = "pass"
	self.action_nopass = "nopass"
	self.enforcer = &mock_enforcer{}
	self.enforcer.On("Enforce", self.domain, self.groups, self.subject, self.object, self.action_pass).Return(nil)
	self.enforcer.On("Enforce", self.domain, self.groups, self.subject, self.object, self.action_nopass).Return(ErrPermissionDenied)

	logger, _ := log_helper.NewLogger("test", "debug")
	self.service = &MetathingsIdentitydService{
		enforcer: self.enforcer,
		logger:   logger,
	}
	self.token = &storage.Token{
		EntityId: &self.subject,
		DomainId: &self.domain,
		Groups: []*storage.Group{
			&storage.Group{
				Id: &self.group,
			},
		},
	}
	self.ctx = context.WithValue(context.Background(), "token", self.token)
}

func (self *metathingsIdentitydService_enforceTestSuite) Test_enforce_pass() {
	err := self.service.enforce(self.ctx, self.object, self.action_pass)
	self.Nil(err)
	self.enforcer.AssertCalled(self.T(), "Enforce", self.domain, self.groups, self.subject, self.object, self.action_pass)
}

func (self *metathingsIdentitydService_enforceTestSuite) Test_enforce_nopass() {
	err := self.service.enforce(self.ctx, self.object, self.action_nopass)
	self.NotNil(err)
	self.enforcer.AssertCalled(self.T(), "Enforce", self.domain, self.groups, self.subject, self.object, self.action_nopass)
}

func TestMetathingsIdentitydService_enforceTestSuite(t *testing.T) {
	suite.Run(t, new(metathingsIdentitydService_enforceTestSuite))
}
