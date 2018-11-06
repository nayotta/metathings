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

func (self *MetathingsIdentitydService) RemoveEntityFromDomain(ctx context.Context, req *pb.RemoveEntityFromDomainRequest) (*empty.Empty, error) {
	var e *storage.Entity
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	dom := req.GetDomain()
	if dom == nil || dom.GetId() == nil || dom.GetId().GetValue() == "" {
		err = errors.New("domain.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	dom_id_str := dom.GetId().GetValue()

	ent := req.GetEntity()
	if ent == nil || ent.GetId() == nil || ent.GetId().GetValue() == "" {
		err = errors.New("entity.id is empty")
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	ent_id_str := ent.GetId().GetValue()

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
