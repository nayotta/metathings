package metathings_identityd2_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateRemoveRoleFromGroup(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (group_getter, role_getter) {
				req := in.(*pb.RemoveRoleFromGroupRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_group_id,
			ensure_get_role_id,
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizeRemoveRoleFromGroup(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.RemoveRoleFromGroupRequest).GetGroup().GetId().GetValue(), "identityd2:remove_role_from_group")
}

func (self *MetathingsIdentitydService) RemoveRoleFromGroup(ctx context.Context, req *pb.RemoveRoleFromGroupRequest) (*empty.Empty, error) {
	var err error

	grp := req.GetGroup()
	rol := req.GetRole()

	grp_id_str := grp.GetId().GetValue()
	rol_id_str := rol.GetId().GetValue()

	grp_s, err := self.storage.GetGroup(ctx, grp_id_str)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get group in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	rol_s, err := self.storage.GetRoleWithFullActions(ctx, rol_id_str)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get role in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	err = self.backend.RemoveRoleFromGroup(ctx, grp_s, rol_s)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to remove role from group")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = self.storage.RemoveRoleFromGroup(ctx, grp_id_str, rol_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to remove role from group in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"group": grp_id_str,
		"role":  rol_id_str,
	}).Infof("remove role from group")

	return &empty.Empty{}, nil
}
