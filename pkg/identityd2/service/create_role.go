package metathings_identityd2_service

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateCreateRole(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, role_getter) {
				req := in.(*pb.CreateRoleRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			func(x role_getter) error {
				role := x.GetRole()

				if role.GetName() == nil {
					return errors.New("role.name is empty")
				}

				return nil
			},
		},
	)
}

func (self *MetathingsIdentitydService) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
	var role_s *storage.Role
	var err error

	role := req.GetRole()

	id_str := id_helper.NewId()
	if role.GetId() != nil {
		id_str = role.GetId().GetValue()
	}

	desc_str := ""
	if role.GetDescription() != nil {
		desc_str = role.GetDescription().GetValue()
	}
	extra_str := must_parse_extra(role.GetExtra())
	name_str := role.GetName().GetValue()
	alias_str := name_str
	if role.GetAlias() != nil {
		alias_str = role.GetAlias().GetValue()
	}

	role_s = &storage.Role{
		Id:          &id_str,
		Name:        &name_str,
		Alias:       &alias_str,
		Description: &desc_str,
		Extra:       &extra_str,
	}

	role_s, err = self.storage.CreateRole(role_s)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to create role in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CreateRoleResponse{
		Role: copy_role(role_s),
	}

	self.logger.WithField("id", id_str).Infof("create role")

	return res, nil
}
