package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) RemoveObject(context.Context, *pb.RemoveObjectRequest) (*empty.Empty, error) {
	panic("unimplemented")
}
