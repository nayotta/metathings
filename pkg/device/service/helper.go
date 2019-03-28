package metathings_device_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/wrappers"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	session_helper "github.com/nayotta/metathings/pkg/common/session"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDeviceServiceImpl) generator_major_session() int64 {
	return int64(self.startup_session)<<32 | int64(session_helper.GenerateMajorSession())
}

func (self *MetathingsDeviceServiceImpl) generator_minor_session() int64 {
	return int64(self.startup_session)<<32 | int64(session_helper.GenerateMinorSession())
}

func (self *MetathingsDeviceServiceImpl) context() context.Context {
	return context_helper.WithToken(context.TODO(), self.tknr.GetToken())
}

func (self *MetathingsDeviceServiceImpl) pb_device() *deviced_pb.OpDevice {
	return &deviced_pb.OpDevice{Id: &wrappers.StringValue{Value: self.info.Id}}
}
