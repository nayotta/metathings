package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateListGroupsForSubject(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, subject_getter) {
				req := in.(*pb.ListGroupsForSubjectRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_subject_id},
	)
}

func (self *MetathingsIdentitydService) AuthorizeListGroupsForSubject(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.ListGroupsForSubjectRequest).GetSubject().GetId().GetValue(), "identityd2:list_groups_for_subject")
}

func (self *MetathingsIdentitydService) ListGroupsForSubject(ctx context.Context, req *pb.ListGroupsForSubjectRequest) (*pb.ListGroupsForSubjectResponse, error) {
	var err error

	sub_id := req.GetSubject().GetId().GetValue()
	grps, err := self.storage.ListGroupsForSubject(ctx, sub_id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListGroupsForSubjectResponse{
		Groups: copy_groups(grps),
	}

	self.logger.WithField("subject_id", sub_id).Debugf("list groups for subject")

	return res, nil
}
