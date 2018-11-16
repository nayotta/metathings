package metathingsmqttdservice

import (
	"context"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/mqttd/storage"
	pb "github.com/nayotta/metathings/pkg/proto/mqttd"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ValidatePatchDevice ValidatePatchDevice
func (serv *MetathingsMqttdService) ValidatePatchDevice(ctx context.Context, in interface{}) error {
	return serv.validateChain(
		[]interface{}{
			func() (policy_helper.Validator, getDevicer) {
				req := in.(*pb.PatchDeviceRequest)
				return req, req
			},
		},
		[]interface{}{ensureGetDeviceID},
	)
}

// AuthorizePatchDevice AuthorizePatchDevice
func (serv *MetathingsMqttdService) AuthorizePatchDevice(ctx context.Context, in interface{}) error {
	return serv.enforce(ctx, in.(*pb.PatchDeviceRequest).GetDevice().GetId().GetValue(), "patch_device")
}

// PatchDevice PatchDevice
func (serv *MetathingsMqttdService) PatchDevice(ctx context.Context, req *pb.PatchDeviceRequest) (*pb.PatchDeviceResponse, error) {
	devS := &storage.Device{}
	var err error

	dev := req.GetDevice()
	devIDStr := dev.GetId().GetValue()

	alias := dev.GetAlias()
	if alias != nil {
		devS.Alias = &alias.Value
	}

	if devS, err = serv.storage.PatchDevice(devIDStr, devS); err != nil {
		serv.logger.WithError(err).Errorf("failed to patch device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.PatchDeviceResponse{
		Device: copyDevice(devS),
	}

	serv.logger.WithField("id", devIDStr).Infof("patch device")

	return res, nil
}
