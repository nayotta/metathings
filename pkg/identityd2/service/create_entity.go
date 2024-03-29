package metathings_identityd2_service

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsIdentitydService) ValidateCreateEntity(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() entity_getter {
				req := in.(*pb.CreateEntityRequest)
				return req
			},
		},
		identityd_validator.Invokers{
			func(x entity_getter) error {
				ent := x.GetEntity()

				if ent.GetId() == nil {
					return errors.New("entity.id is empty")
				}

				if ent.GetName() == nil {
					return errors.New("entity.name is empty")
				}

				if ent.GetPassword() != nil {
					passwd := ent.GetPassword().GetValue()
					if len(passwd) < 8 || len(passwd) > 128 {
						return errors.New("entity.password size range from 8 to 128 bytes")
					}
				}

				return nil
			},
		},
	)
}

func (self *MetathingsIdentitydService) CreateEntity(ctx context.Context, req *pb.CreateEntityRequest) (*pb.CreateEntityResponse, error) {
	var err error

	ent := req.GetEntity()

	ent_id_str := ent.GetId().GetValue()
	pwd_str := passwd_helper.MustParsePassword("")
	if ent.GetPassword() != nil {
		pwd_str = passwd_helper.MustParsePassword(ent.GetPassword().GetValue())
	}
	name_str := ent.GetName().GetValue()
	alias_str := name_str
	if ent.GetAlias() != nil {
		alias_str = ent.GetAlias().GetValue()
	}

	ent_s := &storage.Entity{
		Id:       &ent_id_str,
		Name:     &name_str,
		Alias:    &alias_str,
		Password: &pwd_str,
	}

	if extra := ent.GetExtra(); extra != nil {
		ent_s.ExtraHelper = pb_helper.ExtractStringMapToString(extra)
	}

	if ent_s, err = self.storage.CreateEntity(ctx, ent_s); err != nil {
		self.logger.WithError(err).Errorf("failed to create entity in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CreateEntityResponse{
		Entity: copy_entity(ent_s),
	}

	self.logger.WithField("entity", ent_id_str).Infof("create entity")

	return res, nil
}
