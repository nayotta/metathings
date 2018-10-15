package metathings_identityd2_service

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) DeleteEntity(ctx context.Context, req *pb.DeleteEntityRequest) (*empty.Empty, error) {
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	ent := req.GetEntity()
	if ent.GetId() == nil {
		err = errors.New("entity.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	ent_id_str := ent.GetId().GetValue()

	if err = self.storage.DeleteEntity(ent_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to delete entity in storage")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"entity_id": ent_id_str,
	}).Infof("delete entity")

	return &empty.Empty{}, nil
}
