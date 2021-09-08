package metathings_identityd2_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidatePatchEntity(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() entity_getter {
				req := in.(*pb.PatchEntityRequest)
				return req
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

	if alias := ent_req.GetAlias(); alias != nil {
		alias_str := alias.GetValue()
		ent.Alias = &alias_str
	}

	if passwd := ent_req.GetPassword(); passwd != nil {
		passwd_str := passwd_helper.MustParsePassword(passwd.GetValue())
		ent.Password = &passwd_str
	}

	if extra := ent_req.GetExtra(); extra != nil {
		ent.ExtraHelper = pb_helper.ExtractStringMapToString(extra)
	}

	if ent, err = self.storage.PatchEntity(ctx, idStr, ent); err != nil {
		self.logger.WithError(err).Errorf("failed to patch entity in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.PatchEntityResponse{
		Entity: copy_entity(ent),
	}

	self.logger.WithField("id", idStr).Infof("patch entity")

	return res, nil
}
