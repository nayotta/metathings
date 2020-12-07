package metathings_tagd_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/proto/tagd"
	log "github.com/sirupsen/logrus"
)

func (ts *MetathingsTagdService) AuthorizeGet(ctx context.Context, in interface{}) error {
	return ts.authorizer.Authorize(ctx, in.(*pb.GetRequest).GetId().GetValue(), TAGD_GET)
}

func (ts *MetathingsTagdService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	logger := ts.GetLogger()

	ns := req.GetNamespace().GetValue()
	id := req.GetId().GetValue()

	tags, err := ts.stor.Get(ns, id)
	if err != nil {
		logger.WithError(err).Errorf("failed to get tags")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	logger.WithFields(log.Fields{
		"id":        id,
		"namespace": ns,
	}).Debugf("get tags")

	res := &pb.GetResponse{
		Id:   id,
		Tags: tags,
	}

	return res, nil
}
