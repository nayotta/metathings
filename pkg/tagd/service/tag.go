package metathings_tagd_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/proto/tagd"
)

func (ts *MetathingsTagdService) AuthorizeTag(ctx context.Context, in interface{}) error {
	return ts.authorizer.Authorize(ctx, in.(*pb.TagRequest).GetId().GetValue(), TAGD_TAG)
}

func (ts *MetathingsTagdService) Tag(ctx context.Context, req *pb.TagRequest) (*empty.Empty, error) {
	logger := ts.GetLogger()

	id := req.GetId().GetValue()
	ns := req.GetNamespace().GetValue()

	var tags []string
	for _, tag := range req.GetTags() {
		tags = append(tags, tag.GetValue())
	}

	err := ts.stor.Tag(ns, id, tags)
	if err != nil {
		logger.WithError(err).Errorf("failed to tag")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err != nil {
		defer ts.stor.Remove(ns, id)
		logger.WithError(err).Errorf("failed to add tag in enforcer")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	logger.WithFields(log.Fields{
		"id":        id,
		"namespace": ns,
		"tags":      tags,
	}).Debugf("tag")

	return &empty.Empty{}, nil
}
