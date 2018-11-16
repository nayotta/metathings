package metathingsmqttdservice

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/mqttd/storage"
	pb "github.com/nayotta/metathings/pkg/proto/mqttd"
)

// ValidateGetDevice ValidateGetDevice
func (serv *MetathingsMqttdService) ValidateGetDevice(ctx context.Context, in interface{}) error {
	return serv.validateChain(
		[]interface{}{
			func() (policy_helper.Validator, getDevicer) {
				req := in.(*pb.GetDeviceRequest)
				return req, req
			},
		},
		[]interface{}{ensureGetDeviceID},
	)
}

// AuthorizeGetDevice AuthorizeGetDevice
func (serv *MetathingsMqttdService) AuthorizeGetDevice(ctx context.Context, in interface{}) error {
	return serv.enforce(ctx, in.(*pb.GetDeviceRequest).GetDevice().GetId().GetValue(), "get_device")
}

// GetDevice GetDevice
func (serv *MetathingsMqttdService) GetDevice(ctx context.Context, req *pb.GetDeviceRequest) (*pb.GetDeviceResponse, error) {
	var devS *storage.Device
	var err error

	devIDStr := req.GetDevice().GetId().GetValue()
	if devS, err = serv.storage.GetDevice(devIDStr); err != nil {
		serv.logger.WithError(err).Errorf("failed to get device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.GetDeviceResponse{
		Device: copyDevice(devS),
	}

	serv.logger.WithField("id", devIDStr).Debugf("get device")

	return res, nil
}
