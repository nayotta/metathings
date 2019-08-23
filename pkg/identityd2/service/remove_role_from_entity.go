package metathings_identityd2_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateRemoveRoleFromEntity(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, role_getter, entity_getter) {
				req := in.(*pb.RemoveRoleFromEntityRequest)
				return req, req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_role_id,
			ensure_get_entity_id,
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizeRemoveRoleFromEntity(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.RemoveRoleFromEntityRequest).GetEntity().GetId().GetValue(), "identityd2:remove_role_from_entity")
}

func (self *MetathingsIdentitydService) RemoveRoleFromEntity(ctx context.Context, req *pb.RemoveRoleFromEntityRequest) (*empty.Empty, error) {
	var err error

	ent := req.GetEntity()
	ent_id_str := ent.GetId().GetValue()

	role := req.GetRole()
	role_id_str := role.GetId().GetValue()

	ent_s, err := self.storage.GetEntity(ent_id_str)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get entity in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	rol_s, err := self.storage.GetRole(role_id_str)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get role in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	err = self.backend.RemoveRoleFromEntity(ent_s, rol_s)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to add role to entity in backend")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = self.storage.RemoveRoleFromEntity(ent_id_str, role_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to remove role from entity in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"entity": ent_id_str,
		"role":   role_id_str,
	}).Infof("remove role from entity")

	return &empty.Empty{}, nil
}
