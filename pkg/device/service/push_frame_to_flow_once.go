package metathings_device_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/proto/device"
	deviced_pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDeviceServiceImpl) PushFrameToFlowOnce(ctx context.Context, req *pb.PushFrameToFlowOnceRequest) (*empty.Empty, error) {
	cli, cfn, err := self.cli_fty.NewDevicedServiceClient()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to connect to deviced service")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer cfn()

	creq := &deviced_pb.PushFrameToFlowOnceRequest{
		Id: req.GetId(),
		Device: &deviced_pb.OpDevice{
			Id: self.pb_device().GetId(),
			Flows: []*deviced_pb.OpFlow{
				req.GetFlow(),
			},
		},
		Frame: req.GetFrame(),
	}

	_, err = cli.PushFrameToFlowOnce(self.context(), creq)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to push frame to flow once to deviced service")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}
