package metathings_identityd2_service

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) RemoveRoleFromEntity(ctx context.Context, req *pb.RemoveRoleFromEntityRequest) (*empty.Empty, error) {
	var e *storage.Entity
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	ent := req.GetEntity()
	if ent == nil || ent.GetId() == nil || ent.GetId().GetValue() == "" {
		err = errors.New("entity.id is empty")
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	ent_id_str := ent.GetId().GetValue()

	role := req.GetRole()
	if role == nil || role.GetId() == nil || role.GetId().GetValue() == "" {
		err = errors.New("role.id is empty")
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	role_id_str := role.GetId().GetValue()

	if e, err = self.storage.GetEntity(ent_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get entity in storage")
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	if !role_in_entity(e, role_id_str) {
		err = errors.New("role not in entity")
		self.logger.WithError(err).Errorf("failed to get role in storage")
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	if err = self.storage.RemoveRoleFromEntity(ent_id_str, role_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to remove role from entity in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"entity_id": ent_id_str,
		"role_id":   role_id_str,
	}).Infof("remove role from entity")

	return &empty.Empty{}, nil
}
