package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateGetEntity(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, entity_getter) {
				req := in.(*pb.GetEntityRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_entity_id,
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizeGetEntity(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.GetEntityRequest).GetEntity().GetId().GetValue(), "get_entity")
}

func (self *MetathingsIdentitydService) GetEntity(ctx context.Context, req *pb.GetEntityRequest) (*pb.GetEntityResponse, error) {
	var ent_s *storage.Entity
	var err error

	id_str := req.GetEntity().GetId().GetValue()
	if ent_s, err = self.storage.GetEntity(id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get entity in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.GetEntityResponse{
		Entity: copy_entity(ent_s),
	}

	self.logger.WithField("id", id_str).Debugf("get entity")

	return res, nil
}
