package metathings_device_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	pb "github.com/nayotta/metathings/pkg/proto/device"
)

func (self *MetathingsDeviceServiceImpl) RemoveObject(context.Context, *pb.RemoveObjectRequest) (*empty.Empty, error) {
	panic("unimplemented")
}
