package metathings_tagd_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/proto/tagd"
)

func (ts *MetathingsTagdService) Query(ctx context.Context, req *pb.QueryRequest) (*pb.QueryResponse, error) {
	var tags []string
	logger := ts.GetLogger()

	ns := req.GetNamespace().GetValue()
	for _, tag := range req.GetTags() {
		tags = append(tags, tag.GetValue())
	}

	ids, err := ts.stor.Query(ns, tags)
	if err != nil {
		logger.WithError(err).Errorf("failed to query tags")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	logger.WithField("namespace", ns).Debugf("query tags")

	res := &pb.QueryResponse{
		Tags: tags,
		Ids:  ids,
	}

	return res, nil
}
