package metathings_identityd2_service

import (
	"context"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (self *MetathingsIdentitydService) ValidateListGroupsForObject(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, object_getter) {
				req := in.(*pb.ListGroupsForObjectRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_object_id},
	)
}

func (self *MetathingsIdentitydService) AuthorizeListGroupsForObject(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.ListGroupsForObjectRequest).GetObject().GetId().GetValue(), "list_groups_for_object")
}

func (self *MetathingsIdentitydService) ListGroupsForObject(ctx context.Context, req *pb.ListGroupsForObjectRequest) (*pb.ListGroupsForObjectResponse, error) {
	var err error

	obj_id := req.GetObject().GetId().GetValue()
	grps, err := self.storage.ListGroupsForObject(obj_id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListGroupsForObjectResponse{
		Groups: copy_groups(grps),
	}

	self.logger.WithField("object_id", obj_id).Debugf("list groups for object")

	return res, nil
}