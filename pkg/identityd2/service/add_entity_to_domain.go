package metathings_identityd2_service

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) AddEntityToDomain(ctx context.Context, req *pb.AddEntityToDomainRequest) (*empty.Empty, error) {
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

	if err = self.storage.AddEntityToDomain(dom_id_str, ent_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to add entity to domain in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"entity_id": ent_id_str,
		"domain_id": dom_id_str,
	}).Infof("add entity to domain")

	return &empty.Empty{}, nil
}
