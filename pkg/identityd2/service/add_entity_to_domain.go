package metathings_identityd2_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateAddEntityToDomain(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (entity_getter, domain_getter) {
				req := in.(*pb.AddEntityToDomainRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_entity_id,
			ensure_get_domain_id,
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizeAddEntityToDomain(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.AddEntityToDomainRequest).GetDomain().GetId().GetValue(), "identityd2:add_entity_to_domain")
}

func (self *MetathingsIdentitydService) AddEntityToDomain(ctx context.Context, req *pb.AddEntityToDomainRequest) (*empty.Empty, error) {
	var err error

	ent_id_str := req.GetEntity().GetId().GetValue()
	dom_id_str := req.GetDomain().GetId().GetValue()

	if err = self.storage.AddEntityToDomain(ctx, dom_id_str, ent_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to add entity to domain in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"entity_id": ent_id_str,
		"domain_id": dom_id_str,
	}).Infof("add entity to domain")

	return &empty.Empty{}, nil
}
