package metathings_device_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	context_helper "github.com/nayotta/metathings/pkg/common/context"
	pb "github.com/nayotta/metathings/pkg/proto/device"
)

func (self *MetathingsDeviceServiceImpl) ShowModuleFirmwareDescriptor(ctx context.Context, _ *empty.Empty) (*pb.ShowModuleFirmwareDescriptorResponse, error) {
	tkn := context_helper.ExtractToken(ctx)
	mdl_id := tkn.GetEntity().GetId()
	logger := self.logger.WithField("module", mdl_id)

	mdl, err := self.get_module_info(mdl_id)
	if err != nil {
		logger.WithError(err).Errorf("failed to get module info")
		return nil, self.ParseError(err)
	}
	logger = logger.WithField("module.name", mdl.GetName())

	cli, cfn, err := self.cli_fty.NewDevicedServiceClient()
	if err != nil {
		logger.WithError(err).Errorf("failed to new deviced service client")
		return nil, err
	}
	defer cfn()

	desc, err := self.get_device_firmware_descriptor(cli, self.context())
	if err != nil {
		logger.WithError(err).Errorf("failed to get device firmware descriptor")
		return nil, err
	}

	res := &pb.ShowModuleFirmwareDescriptorResponse{
		FirmwareDescriptor: desc,
	}

	logger.Infof("show module firmware descriptor")

	return res, nil
}
