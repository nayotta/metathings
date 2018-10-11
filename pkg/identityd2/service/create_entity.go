package metathings_identityd2_service

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsIdentitydService) CreateEntity(ctx context.Context, req *pb.CreateEntityRequest) (*pb.CreateEntityResponse, error) {
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	ent_id_str := id_helper.NewId()
	if req.GetId() != nil && req.GetId().GetValue() != "" {
		ent_id_str = req.GetId().GetValue()
	}

	dom_id := req.GetDomain().GetId()
	if dom_id == nil || dom_id.GetValue() == "" {
		err = errors.New("domain.id is empty")
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	dom_id_str := dom_id.Value

	extra_str := must_parse_extra(req.Extra)
	pwd_str := must_parse_password(req.GetPassword().GetValue())

	ent := &storage.Entity{
		Id:       &ent_id_str,
		DomainId: &dom_id_str,
		Name:     &req.Name.Value,
		Alias:    &req.Alias.Value,
		Password: &pwd_str,
		Extra:    &extra_str,
	}

	if ent, err = self.storage.CreateEntity(ent); err != nil {
		self.logger.WithError(err).Errorf("failed to create entity in storage")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	res := &pb.CreateEntityResponse{
		Entity: copy_entity(ent),
	}

	self.logger.WithField("id", ent_id_str).Infof("create entity")

	return res, nil
}
