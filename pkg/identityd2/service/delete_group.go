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

func (self *MetathingsIdentitydService) DeleteGroup(ctx context.Context, req *pb.DeleteGroupRequest) (*empty.Empty, error) {
	var grp_s *storage.Group
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	grp := req.GetGroup()
	if grp.GetId() == nil {
		err = errors.New("group.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	grp_id_str := grp.GetId().GetValue()

	if grp_s, err = self.storage.GetGroup(grp_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get group in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = self.enforcer.RemoveGroup(*grp_s.DomainId, grp_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to remove group in enforcer")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = self.storage.DeleteGroup(grp_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to delete group in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"group_id": grp_id_str,
	}).Infof("delete group")

	return &empty.Empty{}, nil
}
