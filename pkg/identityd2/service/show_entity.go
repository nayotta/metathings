package metathings_identityd2_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsIdentitydService) ShowEntity(ctx context.Context, _ *empty.Empty) (*pb.ShowEntityResponse, error) {
	tkn := ctx.Value("token").(*pb.Token)
	ent_id := tkn.GetEntity().GetId()
	ent, err := self.storage.GetEntity(ctx, ent_id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	self.logger.WithFields(log.Fields{
		"entity_id": ent_id,
	}).Infof("show entity")

	return &pb.ShowEntityResponse{
		Entity: copy_entity(ent),
	}, nil
}
