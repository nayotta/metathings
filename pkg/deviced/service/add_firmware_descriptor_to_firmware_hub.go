package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) ValidateAddFirmwareDescriptorToFirmwareHub(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, firmware_hub_getter) {
				req := in.(*pb.AddFirmwareDescriptorToFirmwareHubRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{ensure_get_firmware_hub_id},
	)
}

func (self *MetathingsDevicedService) AuthorizeAddFirmwareDescriptorToFirmwareHub(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.AddFirmwareDescriptorToFirmwareHubRequest).GetFirmwareHub().GetId().GetValue(), "deviced:add_firmware_descriptor_to_firmware_hub")
}

/*
 * Descriptor Format:
 *
 * device.sha256.next:
 * device.uri.next:
 * device.version.next:
 * modules.<name>.sha256.next:
 * modules.<name>.uri.next:
 * modules.<name>.version.next:
 *
 * Example:
 *
 * {
 *   "device": {
 *     "sha256": { "next": "8484848484848484848484848484848484848484848484848484848484848484" },
 *     "uri": { "next": "http://example.com/firmwares/device/device/v1_0_0/device_v1_0_0.bin" },
 *     "version": { "next": "v1.0.0" }
 *   },
 *   "modules": {
 *     "answer42": {
 *       "sha256": { "next": "4242424242424242424242424242424242424242424242424242424242424242" },
 *       "uri": { "next": "http://example.com/firmwares/modules/answer42/v1_0_0/answeer_v1_0_0.bin" },
 *       "version": { "next": "v1.0.0" }
 *     }
 *   }
 * }
 */
func (self *MetathingsDevicedService) AddFirmwareDescriptorToFirmwareHub(ctx context.Context, req *pb.AddFirmwareDescriptorToFirmwareHubRequest) (*empty.Empty, error) {
	var err error

	fh := req.GetFirmwareHub()
	fh_id_str := fh.GetId().GetValue()

	desc := req.GetFirmwareDescriptor()

	desc_id_str := id_helper.NewId()
	desc_name_str := desc.GetName().GetValue()

	logger := self.get_logger().WithFields(log.Fields{
		"firmware_hub":        fh_id_str,
		"firmware_descriptor": desc_name_str,
	})

	desc_desc := desc.GetDescriptor_()
	desc_desc_str, err := grpc_helper.JSONPBMarshaler.MarshalToString(desc_desc)
	if err != nil {
		logger.WithError(err).Errorf("failed to marshal descriptor to json string")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	desc_s := &storage.FirmwareDescriptor{
		Id:            &desc_id_str,
		Name:          &desc_name_str,
		FirmwareHubId: &fh_id_str,
		Descriptor:    &desc_desc_str,
	}

	if err = self.storage.CreateFirmwareDescriptor(ctx, desc_s); err != nil {
		logger.WithError(err).Errorf("failed to create firmware descriptor in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	logger.Infof("add firmware descriptor to firmware hub")

	return &empty.Empty{}, nil
}
