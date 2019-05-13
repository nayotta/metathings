package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateGetGroup(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, group_getter) {
				req := in.(*pb.GetGroupRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_group_id,
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizeGetGroup(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.GetGroupRequest).GetGroup().GetId().GetValue(), "get_group")
}

func (self *MetathingsIdentitydService) GetGroup(ctx context.Context, req *pb.GetGroupRequest) (*pb.GetGroupResponse, error) {
	var grp_s *storage.Group
	var err error

	id_str := req.GetGroup().GetId().GetValue()
	if grp_s, err = self.storage.GetGroup(id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get group in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.GetGroupResponse{
		Group: copy_group(grp_s),
	}

	self.logger.WithField("id", id_str).Debugf("get group")

	return res, nil
}
