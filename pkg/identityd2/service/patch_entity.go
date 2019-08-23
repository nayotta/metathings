package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidatePatchEntity(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, entity_getter) {
				req := in.(*pb.PatchEntityRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_entity_id,
		},
	)
}

func (self *MetathingsIdentitydService) AuthorizePatchEntity(ctx context.Context, in interface{}) error {
	return self.authorize(ctx, in.(*pb.PatchEntityRequest).GetEntity().GetId().GetValue(), "identityd2:patch_entity")
}

func (self *MetathingsIdentitydService) PatchEntity(ctx context.Context, req *pb.PatchEntityRequest) (*pb.PatchEntityResponse, error) {
	var err error

	ent_req := req.GetEntity()
	ent := &storage.Entity{}

	idStr := ent_req.GetId().GetValue()

	if ent_req.GetAlias() != nil {
		ent.Alias = &ent_req.Alias.Value
	}
	if ent_req.GetPassword() != nil {
		passwordStr := passwd_helper.MustParsePassword(ent_req.GetPassword().GetValue())
		ent.Password = &passwordStr
	}
	if ent_req.GetExtra() != nil {
		extraStr := must_parse_extra(ent_req.GetExtra())
		ent.Extra = &extraStr
	}

	if ent, err = self.storage.PatchEntity(idStr, ent); err != nil {
		self.logger.WithError(err).Errorf("failed to patch entity in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.PatchEntityResponse{
		Entity: copy_entity(ent),
	}

	self.logger.WithField("id", idStr).Infof("patch entity")

	return res, nil
}
