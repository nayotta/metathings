package metathings_tagd_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	pb "github.com/nayotta/metathings/pkg/proto/tagd"
)

func (ts *MetathingsTagdService) Remove(context.Context, *pb.RemoveRequest) (*empty.Empty, error) {
	panic("unimplemented")
}
