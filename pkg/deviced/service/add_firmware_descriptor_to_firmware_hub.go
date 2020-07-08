package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
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

func (self *MetathingsDevicedService) AddFirmwareDescriptorToFirmwareHub(ctx context.Context, req *pb.AddFirmwareDescriptorToFirmwareHubRequest) (*empty.Empty, error) {
	var err error

	fh := req.GetFirmwareHub()
	fh_id_str := fh.GetId().GetValue()

	desc := req.GetFirmwareDescriptor()

	desc_id_str := id_helper.NewId()
	desc_name_str := desc.GetName().GetValue()

	logger := self.logger.WithFields(log.Fields{
		"firmware_hub":        fh_id_str,
		"firmware_descriptor": desc_name_str,
	})

	desc_desc := desc.GetDescriptor_()
	desc_desc_str, err := new(jsonpb.Marshaler).MarshalToString(desc_desc)
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
