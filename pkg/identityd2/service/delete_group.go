package metathings_identityd2_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateDeleteGroup(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, group_getter) {
				req := in.(*pb.DeleteGroupRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_group_id},
	)
}

func (self *MetathingsIdentitydService) AuthorizeDeleteGroup(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.DeleteGroupRequest).GetGroup().GetId().GetValue(), "identityd2:delete_group")
}

func (self *MetathingsIdentitydService) DeleteGroup(ctx context.Context, req *pb.DeleteGroupRequest) (*empty.Empty, error) {
	var grp_s *storage.Group
	var err error

	grp_id_str := req.GetGroup().GetId().GetValue()
	if grp_s, err = self.storage.GetGroup(ctx, grp_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get group in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	rols := []*storage.Role{}
	for _, rol := range grp_s.Roles {
		r, err := self.storage.GetRoleWithFullActions(ctx, *rol.Id)
		if err != nil {
			return nil, err
		}
		rols = append(rols, r)
	}
	grp_s.Roles = rols

	if err = self.backend.DeleteGroup(ctx, grp_s); err != nil {
		self.logger.WithError(err).Errorf("failed to delet group in backend")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = self.storage.DeleteGroup(ctx, grp_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to delete group in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"group_id": grp_id_str,
	}).Infof("delete group")

	return &empty.Empty{}, nil
}
