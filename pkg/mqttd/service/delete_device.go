package metathingsmqttdservice

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/mqttd/storage"
	pb "github.com/nayotta/metathings/pkg/proto/mqttd"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ValidateDeleteDevice ValidateDeleteDevice
func (serv *MetathingsMqttdService) ValidateDeleteDevice(ctx context.Context, in interface{}) error {
	return serv.validateChain(
		[]interface{}{
			func() (policy_helper.Validator, getDevicer) {
				req := in.(*pb.DeleteDeviceRequest)
				return req, req
			},
		},
		[]interface{}{ensureGetDeviceID},
	)
}

// AuthorizeDeleteDevice AuthorizeDeleteDevice
func (serv *MetathingsMqttdService) AuthorizeDeleteDevice(ctx context.Context, in interface{}) error {
	return serv.enforce(ctx, in.(*pb.DeleteDeviceRequest).GetDevice().GetId().GetValue(), "delete_device")
}

// DeleteDevice DeleteDevice
func (serv *MetathingsMqttdService) DeleteDevice(ctx context.Context, req *pb.DeleteDeviceRequest) (*empty.Empty, error) {
	var dev *storage.Device
	var err error

	devIDStr := req.GetDevice().GetId().GetValue()
	if dev, err = serv.storage.GetDevice(devIDStr); err != nil {
		serv.logger.WithError(err).Errorf("failed to get device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = serv.enforcer.RemoveObjectFromKind(devIDStr, KINDDEVICE); err != nil {
		serv.logger.WithError(err).Warningf("failed to remove device from kind in enforcer")
	}
	if err = serv.storage.DeleteDevice(devIDStr); err != nil {
		serv.logger.WithError(err).Debugf("failed to delete device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	serv.logger.WithField("id", devIDStr).Infof("delete device")

	return &empty.Empty{}, nil
}
