package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateDeleteConfig(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() config_getter {
				req := in.(*pb.DeleteConfigRequest)
				return req
			},
		},
		identityd_validator.Invokers{ensure_get_config_id},
	)
}

func (self *MetathingsDevicedService) AuthorizeDeleteConfig(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.DeleteConfigRequest).GetConfig().GetId().GetValue(), "deviced:delete_config")
}

func (self *MetathingsDevicedService) DeleteConfig(ctx context.Context, req *pb.DeleteConfigRequest) (*empty.Empty, error) {
	var err error

	cfg_id_str := req.GetConfig().GetId().GetValue()

	logger := self.get_logger().WithField("config", cfg_id_str)

	if err = self.storage.RemoveConfigFromDeviceByConfigId(ctx, cfg_id_str); err != nil {
		logger.WithError(err).Errorf("failed to remove config form device by config id in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if err = self.storage.DeleteConfig(ctx, cfg_id_str); err != nil {
		logger.WithError(err).Errorf("failed to delete config in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	logger.Infof("delete config")

	return &empty.Empty{}, nil
}
