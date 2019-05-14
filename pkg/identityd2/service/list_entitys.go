package metathings_identityd2_service

import (
	"context"

	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (self *MetathingsIdentitydService) ListEntities(ctx context.Context, req *pb.ListEntitiesRequest) (*pb.ListEntitiesResponse, error) {
	var err error

	if err = req.Validate(); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	ent_req := req.GetEntity()
	ent := &storage.Entity{}

	if ent_req.GetId() != nil && ent_req.GetId().GetValue() != "" {
		ent.Id = &ent_req.Id.Value
	}
	if ent_req.GetName() != nil && ent_req.GetName().GetValue() != "" {
		ent.Name = &ent_req.Name.Value
	}
	if ent_req.GetAlias() != nil && ent_req.GetAlias().GetValue() != "" {
		ent.Alias = &ent_req.Alias.Value
	}

	ents, err := self.storage.ListEntities(ent)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to list entitys in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListEntitiesResponse{}

	for _, ent = range ents {
		res.Entities = append(res.Entities, copy_entity(ent))
	}

	return res, nil
}
