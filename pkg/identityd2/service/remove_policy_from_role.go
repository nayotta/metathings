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

func (self *MetathingsIdentitydService) RemovePolicyFromRole(ctx context.Context, req *pb.RemovePolicyFromRoleRequest) (*empty.Empty, error) {
	var plc *storage.Policy
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	policy := req.GetPolicy()
	if policy == nil || policy.GetId() == nil || policy.GetId().GetValue() == "" {
		err = errors.New("policy.id is empty")
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	plc_id_str := policy.GetId().GetValue()

	role := req.GetRole()
	if role == nil || role.GetId() == nil || role.GetId().GetValue() == "" {
		err = errors.New("role.id is empty")
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	role_id_str := role.GetId().GetValue()

	if plc, err = self.storage.GetPolicy(plc_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get policy in storage")
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	if *plc.RoleId != role_id_str {
		err = errors.New("failed to found policy")
		self.logger.WithError(err).Errorf("failed to get policy in storage")
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	if err = self.storage.DeletePolicy(plc_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to delete policy in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"policy_id": plc_id_str,
		"role_id":   role_id_str,
	}).Infof("remove policy from role")

	return &empty.Empty{}, nil
}
