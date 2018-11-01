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

func (self *MetathingsIdentitydService) AddRoleToEntity(ctx context.Context, req *pb.AddRoleToEntityRequest) (*empty.Empty, error) {
	var r *storage.Role
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	ent := req.GetEntity()
	if ent == nil || ent.GetId() == nil || ent.GetId().GetValue() == "" {
		err = errors.New("entity.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	ent_id_str := ent.GetId().GetValue()

	role := req.GetRole()
	if role == nil || role.GetId() == nil || role.GetId().GetValue() == "" {
		err = errors.New("role.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	role_id_str := role.GetId().GetValue()

	if r, err = self.storage.GetRole(role_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get role in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = self.enforcer.AddSubjectToRole(ent_id_str, *r.Name); err != nil {
		self.logger.WithError(err).Errorf("failed to add subject to role in enforcer")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = self.storage.AddRoleToEntity(ent_id_str, role_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to add role to entity in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"entity_id": ent_id_str,
		"role_id":   role_id_str,
	}).Infof("add role to entity")

	return &empty.Empty{}, nil
}
