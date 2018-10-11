package metathings_identityd2_service

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) AddPolicyToRole(ctx context.Context, req *pb.AddPolicyToRoleRequest) (*empty.Empty, error) {
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	role := req.GetRole()
	if role == nil || role.GetId() == nil || role.GetId().GetValue() == "" {
		err = errors.New("role.id is empty")
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	role_id_str := role.GetId().GetValue()

	policy := req.GetPolicy()
	if policy == nil {
		err = errors.New("policy is null")
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	plc_rule := policy.GetRule()
	if plc_rule == nil || plc_rule.GetValue() == "" {
		err = errors.New("policy.rule is empty")
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	plc_rule_str := plc_rule.GetValue()

	plc_id_str := id_helper.NewId()
	plc_id := policy.GetId()
	if plc_id != nil && plc_id.GetValue() != "" {
		plc_id_str = plc_id.GetValue()
	}

	plc_desc_str := ""
	plc_desc := policy.GetDescription()
	if plc_desc != nil && plc_desc.GetValue() != "" {
		plc_desc_str = plc_desc.GetValue()
	}

	plc := &storage.Policy{
		Id:          &plc_id_str,
		Description: &plc_desc_str,
		RoleId:      &role_id_str,
		Rule:        &plc_rule_str,
	}

	if plc, err = self.storage.CreatePolicy(plc); err != nil {
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"role_id":     role_id_str,
		"policy_rule": plc_rule_str,
	}).Infof("add policy to role")

	return &empty.Empty{}, nil
}
