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

func (self *MetathingsIdentitydService) DeleteDomain(ctx context.Context, req *pb.DeleteDomainRequest) (*empty.Empty, error) {
	var dom_s *storage.Domain
	var role_s *storage.Role
	var roles_s []*storage.Role
	var ents_s []*storage.Entity
	var grp_s *storage.Group
	var grps_s []*storage.Group
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	dom := req.GetDomain()
	if dom.GetId() == nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	dom_id_str := dom.GetId().GetValue()

	if dom_s, err = self.storage.GetDomain(dom_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get domain in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if len(dom_s.Children) > 0 {
		err = errors.New("more than 0 children in domain")
		self.logger.WithError(err).Warningf("any children in domain")
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	role_s = &storage.Role{
		DomainId: &dom_id_str,
	}
	if roles_s, err = self.storage.ListRoles(role_s); err != nil {
		self.logger.WithError(err).Errorf("failed to list roles by domain_id in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if len(roles_s) > 0 {
		err = errors.New("more than 0 roles in domain")
		self.logger.WithError(err).Warningf("any roles in domain")
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

	if err = self.storage.DeleteDomain(dom_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to delete domain in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"domain_id": dom_id_str,
	}).Infof("delete domain")

	return &empty.Empty{}, nil
}
