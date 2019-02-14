package metathings_tagd_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/pkg/proto/tagd"
)

func (ts *MetathingsTagdService) AuthorizeUntag(ctx context.Context, in interface{}) error {
	return ts.authorizer.Authorize(ctx, in.(*pb.UntagRequest).GetId().GetValue(), TAGD_UNTAG)
}

func (ts *MetathingsTagdService) Untag(ctx context.Context, req *pb.UntagRequest) (*empty.Empty, error) {
	logger := ts.GetLogger()

	id := req.GetId().GetValue()
	var tags []string

	for _, tag := range req.GetTags() {
		tags = append(tags, tag.GetValue())
	}

	err := ts.stor.Untag(id, tags)
	if err != nil {
		logger.WithError(err).Errorf("failed to untag")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	logger.WithFields(log.Fields{
		"id":   id,
		"tags": tags,
	}).Debugf("untag")

	return &empty.Empty{}, nil
}
