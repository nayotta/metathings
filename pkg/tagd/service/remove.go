package metathings_tagd_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/proto/tagd"
)

func (ts *MetathingsTagdService) AuthorizeRemove(ctx context.Context, in interface{}) error {
	return ts.authorizer.Authorize(ctx, in.(*pb.RemoveRequest).GetId().GetValue(), TAGD_REMOVE)
}

func (ts *MetathingsTagdService) Remove(ctx context.Context, req *pb.RemoveRequest) (*empty.Empty, error) {
	logger := ts.GetLogger()

	ns := req.GetNamespace().GetValue()
	id := req.GetId().GetValue()

	err := ts.stor.Remove(ns, id)
	if err != nil {
		logger.WithError(err).Errorf("failed to remove id")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	logger.WithFields(log.Fields{
		"id":        id,
		"namespace": ns,
	}).Debugf("remove id")

	return &empty.Empty{}, nil
}
