package metathings_device_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	pb "github.com/nayotta/metathings/pkg/proto/device"
)

func (self *MetathingsDeviceServiceImpl) PutObject(context.Context, *pb.PutObjectRequest) (*empty.Empty, error) {
	panic("unimplemented")
}
