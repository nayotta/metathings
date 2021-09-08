package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateListGroupsForObject(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() object_getter {
				req := in.(*pb.ListGroupsForObjectRequest)
				return req
			},
		},
		identityd_validator.Invokers{ensure_get_object_id},
	)
}

func (self *MetathingsIdentitydService) AuthorizeListGroupsForObject(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.ListGroupsForObjectRequest).GetObject().GetId().GetValue(), "identityd2:list_groups_for_object")
}

func (self *MetathingsIdentitydService) ListGroupsForObject(ctx context.Context, req *pb.ListGroupsForObjectRequest) (*pb.ListGroupsForObjectResponse, error) {
	var err error

	obj_id := req.GetObject().GetId().GetValue()
	grps, err := self.storage.ListGroupsForObject(ctx, obj_id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListGroupsForObjectResponse{
		Groups: copy_groups(grps),
	}

	self.logger.WithField("object_id", obj_id).Debugf("list groups for object")

	return res, nil
}
