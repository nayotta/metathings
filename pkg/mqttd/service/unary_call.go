package metathingsmqttdservice

import (
	"context"

	conn "github.com/nayotta/metathings/pkg/mqttd/connection"
	pb "github.com/nayotta/metathings/pkg/proto/mqttd"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnaryCall UnaryCall
func (serv *MetathingsMqttdService) UnaryCall(ctx context.Context, req *pb.UnaryCallRequest) (*pb.UnaryCallResponse, error) {
	var err error

	devIDStr := req.GetDeviceId().GetValue()
	msg := req.GetPayload().GetValue()

	path := conn.EncodeDownPath(devIDStr)
	err = serv.cc.Pub(msg, path)
	if err != nil {
		serv.logger.WithError(err).Errorf("failed to pub msg")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	//TODO(zh) mqtt UnaryCall

	res := &pb.UnaryCallResponse{
		Payload: req.GetPayload(),
	}

	serv.logger.WithField("id", devIDStr).Debugf("unary call")

	return res, nil
}
