package metathings_tagd_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	pb "github.com/nayotta/metathings/pkg/proto/tagd"
)

func (ts *MetathingsTagdService) Untag(context.Context, *pb.UntagRequest) (*empty.Empty, error) {
	panic("unimplemented")
}
