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

func (self *MetathingsIdentitydService) RemoveEntityFromGroup(ctx context.Context, req *pb.RemoveEntityFromGroupRequest) (*empty.Empty, error) {
	var g *storage.Group
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	grp := req.GetGroup()
	if grp == nil || grp.GetId() == nil || grp.GetId().GetValue() == "" {
		err = errors.New("group.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	grp_id_str := grp.GetId().GetValue()

	ent := req.GetEntity()
	if ent == nil || ent.GetId() == nil || ent.GetId().GetValue() == "" {
		err = errors.New("entity.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	ent_id_str := ent.GetId().GetValue()

	if g, err = self.storage.GetGroup(grp_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get group in storage")
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	if !entity_in_group(g, ent_id_str) {
		err = errors.New("entity not in group")
		self.logger.WithError(err).Warningf("failed to get entity in storage")
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	if err = self.storage.RemoveEntityFromGroup(ent_id_str, grp_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to remove entity from group in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"entity_id": ent_id_str,
		"group_id":  grp_id_str,
	}).Infof("remove entity from group")

	return &empty.Empty{}, nil
}
