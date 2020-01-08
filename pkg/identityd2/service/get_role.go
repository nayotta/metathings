package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateGetRole(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, role_getter) {
				req := in.(*pb.GetRoleRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_role_id},
	)
}

func (self *MetathingsIdentitydService) AuthorizeGetRole(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.GetRoleRequest).GetRole().GetId().GetValue(), "identityd2:get_role")
}

func (self *MetathingsIdentitydService) GetRole(ctx context.Context, req *pb.GetRoleRequest) (*pb.GetRoleResponse, error) {
	var role_s *storage.Role
	var err error

	id_str := req.GetRole().GetId().GetValue()
	if role_s, err = self.storage.GetRole(ctx, id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get role in storage")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	res := &pb.GetRoleResponse{
		Role: copy_role(role_s),
	}

	self.logger.WithField("id", id_str).Debugf("get role")

	return res, nil
}
