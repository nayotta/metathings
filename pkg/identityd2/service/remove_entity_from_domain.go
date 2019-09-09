package metathings_identityd2_service

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateRemoveEntityFromDomain(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, entity_getter, domain_getter) {
				req := in.(*pb.RemoveEntityFromDomainRequest)
				return req, req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_entity_id,
			ensure_get_domain_id,
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizeRemoveEntityFromDomain(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.RemoveEntityFromDomainRequest).GetDomain().GetId().GetValue(), "identityd2:remove_entity_from_domain")
}

func (self *MetathingsIdentitydService) RemoveEntityFromDomain(ctx context.Context, req *pb.RemoveEntityFromDomainRequest) (*empty.Empty, error) {
	var e *storage.Entity
	var err error

	dom_id_str := req.GetDomain().GetId().GetValue()
	ent_id_str := req.GetEntity().GetId().GetValue()

	if e, err = self.storage.GetEntity(ent_id_str); err != nil {
		self.logger.WithError(err).Error("failed to get entity in storage")
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	if !domain_in_entity(e, dom_id_str) {
		err = errors.New("domain not in entity")
		self.logger.WithError(err).Warningf("failed to get domain in storage")
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	if err = self.storage.RemoveEntityFromDomain(dom_id_str, ent_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to remove entity from domain in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"entity_id": ent_id_str,
		"domain_id": dom_id_str,
	}).Infof("remove entity from domain")

	return &empty.Empty{}, nil
}
