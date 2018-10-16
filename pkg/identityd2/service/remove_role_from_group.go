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

func (self *MetathingsIdentitydService) RemoveRoleFromGroup(ctx context.Context, req *pb.RemoveRoleFromGroupRequest) (*empty.Empty, error) {
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

	role := req.GetRole()
	if role == nil || role.GetId() == nil || role.GetId().GetValue() == "" {
		err = errors.New("role.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	role_id_str := role.GetId().GetValue()

	if g, err = self.storage.GetGroup(grp_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get group in storage")
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	if !role_in_group(g, role_id_str) {
		err = errors.New("role not in group")
		self.logger.WithError(err).Warningf("failed to get role in storage")
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	if err = self.storage.RemoveRoleFromGroup(grp_id_str, role_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to remove role from group in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"group_id": grp_id_str,
		"role_id":  role_id_str,
	}).Infof("remove role from group")

	return &empty.Empty{}, nil
}
