package metathings_device_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/pkg/proto/device"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDeviceServiceImpl) PutObject(ctx context.Context, req *pb.PutObjectRequest) (*empty.Empty, error) {
	cli, cfn, err := self.cli_fty.NewDevicedServiceClient()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to connect to deviced service")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer cfn()

	obj := req.GetObject()
	obj.Device = self.pb_device()

	creq := &deviced_pb.PutObjectRequest{
		Object:  obj,
		Content: req.GetContent(),
	}
	_, err = cli.PutObject(self.context(), creq)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to put object from deviced service")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"object.prefix": obj.Prefix.Value,
		"object.name":   obj.Name.Value,
	}).Debugf("put object")

	return &empty.Empty{}, nil
}
