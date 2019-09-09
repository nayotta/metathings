package metathings_deviced_service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateGetDeviceByModule(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, module_getter) {
				req := in.(*pb.GetDeviceByModuleRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_module_id},
	)
}

func (self *MetathingsDevicedService) AuthorizeGetDeviceByModule(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.GetDeviceByModuleRequest).GetModule().GetId().GetValue(), "deviced:get_device_by_module")
}

func (self *MetathingsDevicedService) GetDeviceByModule(ctx context.Context, req *pb.GetDeviceByModuleRequest) (*pb.GetDeviceByModuleResponse, error) {
	var dev_s *storage.Device
	var err error

	mdl_id_str := req.GetModule().GetId().GetValue()
	if dev_s, err = self.storage.GetDeviceByModuleId(mdl_id_str); err != nil {
		self.logger.WithError(err).Errorf("failed to get device by module id in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.GetDeviceByModuleResponse{
		Device: copy_device(dev_s),
	}

	self.logger.WithFields(log.Fields{
		"module_id": mdl_id_str,
	}).Debugf("get device by module")

	return res, nil
}
