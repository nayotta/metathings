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

func (self *MetathingsIdentitydService) ValidateRemoveSubjectFromGroup(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (subject_getter, group_getter) {
				req := in.(*pb.RemoveSubjectFromGroupRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_subject_id,
			ensure_get_group_id,
			ensure_group_exists_s(ctx, self.storage),
			ensure_subject_exists_in_group_s(ctx, self.storage),
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizeRemoveSubjectFromGroup(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.RemoveSubjectFromGroupRequest).GetGroup().GetId().GetValue(), "identityd2:remove_subject_from_group")
}

func (self *MetathingsIdentitydService) RemoveSubjectFromGroup(ctx context.Context, req *pb.RemoveSubjectFromGroupRequest) (*empty.Empty, error) {
	var err error

	grp_id_str := req.GetGroup().GetId().GetValue()
	sub_id_str := req.GetSubject().GetId().GetValue()

	grp_s, err := self.storage.GetGroup(ctx, grp_id_str)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get group in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	sub_s, err := self.storage.GetEntity(ctx, sub_id_str)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get entity in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	err = self.backend.RemoveSubjectFromGroup(ctx, grp_s, sub_s)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to remove subject from group in backend")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = self.storage.RemoveSubjectFromGroup(ctx, grp_id_str, sub_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to remove subject from group in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"subject_id": sub_id_str,
		"group_id":   grp_id_str,
	}).Infof("remove subject from group")

	return &empty.Empty{}, nil
}
