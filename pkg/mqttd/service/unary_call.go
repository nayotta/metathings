package metathingsmqttdservice

import (
	"context"

	pb "github.com/nayotta/metathings/pkg/proto/mqttd"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnaryCall UnaryCall
func (serv *MetathingsMqttdService) UnaryCall(ctx context.Context, req *pb.UnaryCallRequest) (*pb.UnaryCallResponse, error) {
	var val *pb.UnaryCallValue

	devIDStr := req.GetDevice().GetId().GetValue()
	_, err := serv.storage.GetDevice(devIDStr)
	if err != nil {
		serv.logger.WithError(err).Debugf("failed to get device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	//TODO(zh) mqtt UnaryCall

	res := &pb.UnaryCallResponse{
		Device: &pb.Device{Id: devIDStr},
		Value:  val,
	}

	serv.logger.WithField("id", devIDStr).Debugf("unary call")

	return res, nil
}
