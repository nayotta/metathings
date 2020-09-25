package metathings_device_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/proto/device"
	deviced_pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDeviceServiceImpl) ListObjects(ctx context.Context, req *pb.ListObjectsRequest) (*pb.ListObjectsResponse, error) {
	cli, cfn, err := self.cli_fty.NewDevicedServiceClient()
	if err != nil {
		self.logger.WithError(err).Warningf("failed to connect to deviced service")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer cfn()

	obj := req.GetObject()
	obj.Device = self.pb_device()

	creq := &deviced_pb.ListObjectsRequest{Object: obj}
	cres, err := cli.ListObjects(self.context(), creq)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to list objects from deviced service")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListObjectsResponse{
		Objects: cres.GetObjects(),
	}

	return res, nil
}
