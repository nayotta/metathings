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
func (serv *MetathingsMqttdService) GpioDigitalWrite(ctx context.Context, req *pb.MqttRequest) (*pb.MqttResponse, error) {
	var err error

	reqType := req.GetType()
	if reqType != pb.MessageReqType_REQ_GPIO_DIGITAL {
		serv.logger.WithError(ErrUnsupportRequestType).Errorf("failed to get gpio digital req")
		return nil, status.Errorf(codes.Internal, ErrUnsupportRequestType.Error())
	}

	gpio, ok := req.GetPayload().(*pb.MqttRequest_GpioDigital)
	if !ok {
		serv.logger.WithError(err).Errorf("failed to get gpio digital payload")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	devID := req.GetDeviceId().GetValue()

	dat := &pb.MqttDeviceRequest{
		Type: reqType,
		Payload: &pb.MqttDeviceRequest_GpioDigital{
			GpioDigital: &pb.GpioDigitalPayload{
				Pin:   gpio.GpioDigital.GetPin().GetValue(),
				Value: gpio.GpioDigital.GetValue().GetValue(),
			},
		},
	}

	msg, err := proto.Marshal(dat)
	if err != nil {
		serv.logger.WithError(err).Errorf("failed to marshal gpio msg")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	path := conn.EncodeDownPath(devID)

	err = serv.cc.Pub(msg, path)
	if err != nil {
		serv.logger.WithError(err).Errorf("failed to pub gpio msg")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	//TODO(zh) ACK need
	res := &pb.MqttResponse{
		Type: pb.MessageResType_RES_GPIO_DIGITAL,
		Payload: &pb.MqttResponse_GpioDigital{
			GpioDigital: &pb.GpioDigitalPayload{
				Pin:   gpio.GpioDigital.GetPin().GetValue(),
				Value: gpio.GpioDigital.GetValue().GetValue(),
			},
		},
	}

	return res, nil
}
