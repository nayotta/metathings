package metathings_tagd_service

import (
	"context"

	pb "github.com/nayotta/metathings/pkg/proto/tagd"
)

func (ts *MetathingsTagdService) Query(context.Context, *pb.QueryRequest) (*pb.QueryResponse, error) {
	panic("unimplemented")
}
