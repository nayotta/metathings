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

func (self *MetathingsIdentitydService) ValidateDeleteDomain(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, domain_getter) {
				req := in.(*pb.DeleteDomainRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_domain_id},
	)
}

func (self *MetathingsIdentitydService) AuthorizeDeleteDomain(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.DeleteDomainRequest).GetDomain().GetId().GetValue(), "delete_domain")
}

func (self *MetathingsIdentitydService) DeleteDomain(ctx context.Context, req *pb.DeleteDomainRequest) (*empty.Empty, error) {
	var dom_s *storage.Domain
	var ents_s []*storage.Entity
	var grp_s *storage.Group
	var grps_s []*storage.Group
	var err error

	dom_id_str := req.GetDomain().GetId().GetValue()

	if dom_s, err = self.storage.GetDomain(dom_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get domain in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if len(dom_s.Children) > 0 {
		err = errors.New("more than 0 children in domain")
		self.logger.WithError(err).Warningf("any children in domain")
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	if ents_s, err = self.storage.ListEntitiesByDomainId(dom_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to list entities by domain_id in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if len(ents_s) > 0 {
		err = errors.New("more than 0 entities in domain")
		self.logger.WithError(err).Warningf("any entities in domain")
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	grp_s = &storage.Group{
		DomainId: &dom_id_str,
	}
	if grps_s, err = self.storage.ListGroups(grp_s); err != nil {
		self.logger.WithError(err).Errorf("failed to list groups by domain_id in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if len(grps_s) > 0 {
		err = errors.New("more than 0 groups in domain")
		self.logger.WithError(err).Warningf("any groups in domain")
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	if err = self.enforcer.RemoveObjectFromKind(dom_id_str, KIND_DOMAIN); err != nil {
		self.logger.WithError(err).Warningf("failed to remove domain from kind in enforcer")
	}

	if err = self.storage.DeleteDomain(dom_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to delete domain in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"domain_id": dom_id_str,
	}).Infof("delete domain")

	return &empty.Empty{}, nil
}
