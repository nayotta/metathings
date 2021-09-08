package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidatePatchGroup(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() group_getter {
				req := in.(*pb.PatchGroupRequest)
				return req
			},
		},
		identityd_validator.Invokers{
			ensure_get_group_id,
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizePatchGroup(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.PatchGroupRequest).GetGroup().GetId().GetValue(), "identityd2:patch_group")
}

func (self *MetathingsIdentitydService) PatchGroup(ctx context.Context, req *pb.PatchGroupRequest) (*pb.PatchGroupResponse, error) {
	var err error

	grp_req := req.GetGroup()
	grp := &storage.Group{}

	idStr := grp_req.GetId().GetValue()

	if grp_req.GetAlias() != nil {
		grp.Alias = &grp_req.Alias.Value
	}
	if grp_req.GetDescription() != nil {
		grp.Description = &grp_req.Description.Value
	}
	if extra := grp_req.GetExtra(); extra != nil {
		grp.ExtraHelper = pb_helper.ExtractStringMapToString(extra)
	}

	if grp, err = self.storage.PatchGroup(ctx, idStr, grp); err != nil {
		self.logger.WithError(err).Errorf("failed to patch group in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.PatchGroupResponse{
		Group: copy_group(grp),
	}

	self.logger.WithField("id", idStr).Infof("patch group")

	return res, nil
}
