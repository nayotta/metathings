package metathingsmqttdservice

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/mqttd/storage"
	pb "github.com/nayotta/metathings/pkg/proto/mqttd"
)

// ShowDevice ShowDevice
func (serv *MetathingsMqttdService) ShowDevice(ctx context.Context, _ *empty.Empty) (*pb.ShowDeviceResponse, error) {
	var devS *storage.Device
	var err error

	if devS, err = serv.getDeviceByContext(ctx); err != nil {
		serv.logger.WithError(err).Errorf("failed to get device by context in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ShowDeviceResponse{
		Device: copyDevice(devS),
	}

	serv.logger.WithField("id", *devS.ID).Debugf("show device")

	return res, nil
}
