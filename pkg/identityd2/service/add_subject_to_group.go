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

func (self *MetathingsIdentitydService) ValidateAddSubjectToGroup(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, subject_getter, group_getter) {
				req := in.(*pb.AddSubjectToGroupRequest)
				return req, req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_subject_id,
			ensure_get_group_id,
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizeAddSubjectToGroup(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.AddSubjectToGroupRequest).GetSubject().GetId().GetValue(), "add_subject_to_group")
}

func (self *MetathingsIdentitydService) AddSubjectToGroup(ctx context.Context, req *pb.AddSubjectToGroupRequest) (*empty.Empty, error) {
	var err error

	grp_id_str := req.GetGroup().GetId().GetValue()
	sub_id_str := req.GetSubject().GetId().GetValue()

	if err = self.storage.AddSubjectToGroup(grp_id_str, sub_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to add subject to group in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"subject_id": sub_id_str,
		"group_id":   grp_id_str,
	}).Infof("add subject to group")

	return &empty.Empty{}, nil
}