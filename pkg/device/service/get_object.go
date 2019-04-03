package metathings_device_service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/pkg/proto/device"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDeviceServiceImpl) GetObject(ctx context.Context, req *pb.GetObjectRequest) (*pb.GetObjectResponse, error) {
	cli, cfn, err := self.cli_fty.NewDevicedServiceClient()
	if err != nil {
		self.logger.WithError(err).Warningf("failed to connect to deviced service")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer cfn()

	obj := req.GetObject()
	obj.Device = self.pb_device()

	creq := &deviced_pb.GetObjectRequest{Object: obj}
	cres, err := cli.GetObject(self.context(), creq)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get object from deviced service")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.GetObjectResponse{Object: cres.GetObject()}

	self.logger.WithFields(log.Fields{
		"object.prefix": obj.Prefix.Value,
		"object.name":   obj.Name.Value,
	}).Debugf("get object")

	return res, nil
}
