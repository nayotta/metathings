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

func (self *MetathingsIdentitydService) ValidateRemoveObjectFromGroup(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, object_getter, group_getter) {
				req := in.(*pb.RemoveObjectFromGroupRequest)
				return req, req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_object_id,
			ensure_get_group_id,
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizeRemoveObjectFromGroup(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.RemoveObjectFromGroupRequest).GetObject().GetId().GetValue(), "remove_object_from_group")
}

func (self *MetathingsIdentitydService) RemoveObjectFromGroup(ctx context.Context, req *pb.RemoveObjectFromGroupRequest) (*empty.Empty, error) {
	var err error

	grp_id_str := req.GetGroup().GetId().GetValue()
	obj_id_str := req.GetObject().GetId().GetValue()

	grp_s, err := self.storage.GetGroup(grp_id_str)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get group in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	obj_s, err := self.storage.GetEntity(obj_id_str)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get entity in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	err = self.backend.RemoveObjectFromGroup(grp_s, obj_s)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to remove object from group in backend")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = self.storage.RemoveObjectFromGroup(grp_id_str, obj_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to remove object from group in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"object_id": obj_id_str,
		"group_id":  grp_id_str,
	}).Infof("remove object from group")

	return &empty.Empty{}, nil
}
