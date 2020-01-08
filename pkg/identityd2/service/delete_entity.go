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

func (self *MetathingsIdentitydService) ValidateDeleteEntity(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, entity_getter) {
				req := in.(*pb.DeleteEntityRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_entity_id},
	)
}

func (self *MetathingsIdentitydService) AuthorizeDeleteEntity(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.DeleteEntityRequest).GetEntity().GetId().GetValue(), "identityd2:delete_entity")
}

func (self *MetathingsIdentitydService) DeleteEntity(ctx context.Context, req *pb.DeleteEntityRequest) (*empty.Empty, error) {
	var err error

	ent := req.GetEntity()
	ent_id_str := ent.GetId().GetValue()

	ent_s, err := self.storage.GetEntity(ctx, ent_id_str)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get entity in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	grps_s, err := self.storage.ListGroupsForSubject(ctx, ent_id_str)
	if err != nil {
		self.logger.WithError(err).Warningf("failed to list groups for subject")
	} else {
		for _, grp_s := range grps_s {
			err = self.backend.RemoveSubjectFromGroup(ctx, grp_s, ent_s)
			if err != nil {
				self.logger.WithError(err).WithFields(log.Fields{
					"subject": ent_id_str,
					"group":   *grp_s.Id,
				}).Warningf("failed to remove subject from group")
			}
		}
	}

	grps_s, err = self.storage.ListGroupsForObject(ctx, ent_id_str)
	if err != nil {
		self.logger.WithError(err).Warningf("failed to list groups for object")
	} else {
		for _, grp_s := range grps_s {
			err = self.backend.RemoveObjectFromGroup(ctx, grp_s, ent_s)
			if err != nil {
				self.logger.WithError(err).WithFields(log.Fields{
					"object": ent_id_str,
					"group":  *grp_s.Id,
				}).Warningf("failed to remove object from group")
			}
		}
	}

	if err = self.storage.DeleteEntity(ctx, ent_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to delete entity in storage")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"entity_id": ent_id_str,
	}).Infof("delete entity")

	return &empty.Empty{}, nil
}
