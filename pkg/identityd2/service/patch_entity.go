package metathings_identityd2_service

import (
	"context"
	"errors"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (self *MetathingsIdentitydService) PatchEntity(ctx context.Context, req *pb.PatchEntityRequest) (*pb.PatchEntityResponse, error) {
	var ent *storage.Entity
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if req.GetId() == nil || req.GetId().GetValue() == "" {
		err = errors.New("entity.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	idStr := req.GetId().GetValue()

	if req.GetAlias() != nil {
		*ent.Alias = req.GetAlias().GetValue()
	}
	if req.GetPassword() != nil {
		*ent.Password = req.GetPassword().GetValue()
	}

	if ent, err = self.storage.PatchEntity(idStr, ent); err != nil {
		self.logger.WithError(err).Errorf("failed to patch entity in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.PatchEntityResponse{
		Entity: copy_entity(ent),
	}

	self.logger.WithField("id", idStr).Debugf("patch entity")

	return res, nil
}
