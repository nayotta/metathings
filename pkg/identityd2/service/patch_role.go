package metathings_identityd2_service

import (
	"context"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (self *MetathingsIdentitydService) ValidatePatchRole(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, role_getter) {
				req := in.(*pb.PatchRoleRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_role_id,
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizePatchRole(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.PatchRoleRequest).GetRole().GetId().GetValue(), "patch_role")
}

func (self *MetathingsIdentitydService) PatchRole(ctx context.Context, req *pb.PatchRoleRequest) (*pb.PatchRoleResponse, error) {
	var err error

	rol_req := req.GetRole()
	rol := &storage.Role{}

	idStr := rol_req.GetId().GetValue()

	if rol_req.GetAlias() != nil {
		rol.Alias = &rol_req.Alias.Value
	}
	if rol_req.GetDescription() != nil {
		rol.Description = &rol_req.Description.Value
	}
	if rol_req.GetExtra() != nil {
		extraStr := must_parse_extra(rol_req.GetExtra())
		rol.Extra = &extraStr
	}

	if rol, err = self.storage.PatchRole(idStr, rol); err != nil {
		self.logger.WithError(err).Errorf("failed to patch role in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.PatchRoleResponse{
		Role: copy_role(rol),
	}

	self.logger.WithField("id", idStr).Infof("patch role")

	return res, nil
}
