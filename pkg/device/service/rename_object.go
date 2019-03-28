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

func (self *MetathingsDeviceServiceImpl) RenameObject(ctx context.Context, req *pb.RenameObjectRequest) (*empty.Empty, error) {
	cli, cfn, err := self.cli_fty.NewDevicedServiceClient()
	if err != nil {
		self.logger.WithError(err).Warningf("failed to connect to deviced service")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer cfn()

	src := req.GetSource()
	dst := req.GetDestination()

	src.Device = self.pb_device()

	creq := &deviced_pb.RenameObjectRequest{
		Source:      src,
		Destination: dst,
	}
	_, err = cli.RenameObject(self.context(), creq)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to rename object from deviced service")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	self.logger.WithFields(log.Fields{
		"source.prefix":      src.Prefix.Value,
		"source.name":        src.Name.Value,
		"destination.prefix": dst.Prefix.Value,
		"destination.name":   dst.Name.Value,
	}).Debugf("rename object")

	return &empty.Empty{}, nil
}
