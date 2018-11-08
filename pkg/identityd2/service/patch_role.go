package metathings_identityd2_service

import (
	"context"
	"errors"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (self *MetathingsIdentitydService) PatchRole(ctx context.Context, req *pb.PatchRoleRequest) (*pb.PatchRoleResponse, error) {
	var rol = &storage.Role{}
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if req.GetId() == nil || req.GetId().GetValue() == "" {
		err = errors.New("role.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	idStr := req.GetId().GetValue()

	if req.GetAlias() != nil {
		aliasStr := req.GetAlias().GetValue()
		rol.Alias = &aliasStr
	}
	if req.GetDescription() != nil {
		descriptionStr := req.GetDescription().GetValue()
		rol.Description = &descriptionStr
	}
	if req.GetExtra() != nil {
		extraStr := must_parse_extra(req.GetExtra())
		rol.Extra = &extraStr
	}

	if rol, err = self.storage.PatchRole(idStr, rol); err != nil {
		self.logger.WithError(err).Errorf("failed to patch role in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.PatchRoleResponse{
		Role: copy_role(rol),
	}

	self.logger.WithField("id", idStr).Infof("patch role")

	return res, nil
}