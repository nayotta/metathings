package metathings_identityd2_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateDeleteRole(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, role_getter) {
				req := in.(*pb.DeleteRoleRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_role_id},
	)
}

func (self *MetathingsIdentitydService) AuthorizeDeleteRole(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.DeleteRoleRequest).GetRole().GetId().GetValue(), "identityd2:delete_role")
}

func (self *MetathingsIdentitydService) DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest) (*empty.Empty, error) {
	var err error

	role := req.GetRole()
	role_id_str := role.GetId().GetValue()

	if err = self.storage.DeleteRole(role_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to delete role in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"role_id": role_id_str,
	}).Infof("delete role")

	return &empty.Empty{}, nil
}
