package metathings_deviced_service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateListConfigsByDevice(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, device_getter) {
				req := in.(*pb.ListConfigsByDeviceRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			ensure_get_device_id,
		},
	)
}

func (self *MetathingsDevicedService) AuthorizeListConfigsByDevice(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.ListConfigsByDeviceRequest).GetDevice().GetId().GetValue(), "deviced:list_configs_by_device")
}

func (self *MetathingsDevicedService) ListConfigsByDevice(ctx context.Context, req *pb.ListConfigsByDeviceRequest) (*pb.ListConfigsByDeviceResponse, error) {
	var cfgs_s []*storage.Config
	var err error

	dev := req.GetDevice()
	dev_id_str := dev.GetId().GetValue()
	logger := self.logger.WithField("device", dev_id_str)

	if cfgs_s, err = self.storage.ListConfigsByDeviceId(ctx, dev_id_str); err != nil {
		logger.WithError(err).Errorf("failed to list configs by device")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListConfigsByDeviceResponse{
		Configs: copy_configs(cfgs_s),
	}

	logger.Debugf("list configs by device")

	return res, nil
}
