package metathings_identityd2_service

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) PatchGroup(ctx context.Context, req *pb.PatchGroupRequest) (*pb.PatchGroupResponse, error) {
	var grp = &storage.Group{}
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if req.GetId() == nil || req.GetId().GetValue() == "" {
		err = errors.New("group.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	idStr := req.GetId().GetValue()

	if req.GetAlias() != nil {
		aliasStr := req.GetAlias().GetValue()
		grp.Alias = &aliasStr
	}
	if req.GetDescription() != nil {
		descriptionStr := req.GetDescription().GetValue()
		grp.Description = &descriptionStr
	}
	if req.GetExtra() != nil {
		extraStr := must_parse_extra(req.GetExtra())
		grp.Extra = &extraStr
	}

	if grp, err = self.storage.PatchGroup(idStr, grp); err != nil {
		self.logger.WithError(err).Errorf("failed to patch group in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.PatchGroupResponse{
		Group: copy_group(grp),
	}

	self.logger.WithField("id", idStr).Infof("patch group")

	return res, nil
}
