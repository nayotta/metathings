package metathings_device_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/proto/device"
	deviced_pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDeviceServiceImpl) GetObjectContent(ctx context.Context, req *pb.GetObjectContentRequest) (*pb.GetObjectContentResponse, error) {
	logger := self.get_logger().WithField("method", "GetObjectContent")

	cli, cfn, err := self.cli_fty.NewDevicedServiceClient()
	if err != nil {
		logger.WithError(err).Warningf("failed to connect to deviced service")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer cfn()

	obj := req.GetObject()
	obj.Device = self.pb_device()

	creq := &deviced_pb.GetObjectContentRequest{
		Object: obj,
	}
	cres, err := cli.GetObjectContent(self.context(), creq)
	if err != nil {
		logger.WithError(err).Errorf("failed to get object content from deviced service")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.GetObjectContentResponse{
		Content: cres.GetContent(),
	}

	return res, nil
}
