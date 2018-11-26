package metathingsmqttdservice

import (
	"errors"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	deviced_helper "github.com/nayotta/metathings/pkg/mqttd/helper"
	storage "github.com/nayotta/metathings/pkg/mqttd/storage"
	pb "github.com/nayotta/metathings/pkg/proto/mqttd"
)

func (srv *MetathingsMqttdService) handleGRPCError(err error, format string, args ...interface{}) error {
	return grpc_helper.HandleGRPCError(srv.logger, err, format, args...)
}

func copyDevice(x *storage.Device) *pb.Device {
	y := &pb.Device{
		Id:    *x.ID,
		State: deviced_helper.DEVICESTATEENUMER.ToValue(*x.State),
		Name:  *x.Name,
		Alias: *x.Alias,
	}

	return y
}

func copyDevices(xs []*storage.Device) []*pb.Device {
	var ys []*pb.Device

	for _, x := range xs {
		ys = append(ys, copyDevice(x))
	}

	return ys
}

type getDevicer interface {
	GetDevice() *pb.OpDevice
}

func ensureGetDeviceID(x getDevicer) error {
	if x.GetDevice().GetId() == nil {
		return errors.New("device.id is empty")
	}
	return nil
}
