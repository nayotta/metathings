package metathings_tagd_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/pkg/proto/tagd"
)

func (ts *MetathingsTagdService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	logger := ts.GetLogger()

	id := req.GetId().GetValue()

	tags, err := ts.tagtk.Get(id)
	if err != nil {
		logger.WithError(err).Errorf("failed to get tags")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	logger.WithField("id", id).Debugf("get tags")

	res := &pb.GetResponse{
		Id:   id,
		Tags: tags,
	}

	return res, nil
}
