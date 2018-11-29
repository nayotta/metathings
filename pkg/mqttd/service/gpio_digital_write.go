package metathingsmqttdservice

import (
	"context"

	"github.com/golang/protobuf/proto"
	conn "github.com/nayotta/metathings/pkg/mqttd/connection"
	pb "github.com/nayotta/metathings/pkg/proto/mqttd"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GpioDigitalWrite GpioDigitalWrite
func (serv *MetathingsMqttdService) GpioDigitalWrite(ctx context.Context, req *pb.GpioDigitalWriteRequest) (*pb.GpioDigitalWriteResponse, error) {
	var err error

	gpio := req.GetGpio()
	devID := gpio.GetDeviceId().GetValue()

	dat := &pb.GpioValue{
		Pin:   gpio.GetPin().GetValue(),
		Value: gpio.GetValue().GetValue(),
	}

	msg, err := proto.Marshal(dat)
	if err != nil {
		serv.logger.WithError(err).Errorf("failed to marshal gpio msg")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	path := conn.EncodeDownPath(devID, conn.GpioType)

	err = serv.cc.Pub(msg, path)
	if err != nil {
		serv.logger.WithError(err).Errorf("failed to pub gpio msg")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	//TODO(zh) ACK need

	return nil, nil
}
