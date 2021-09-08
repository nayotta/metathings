package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	context_helper "github.com/nayotta/metathings/pkg/common/context"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/proto/deviced"
	identityd_pb "github.com/nayotta/metathings/proto/identityd2"
)

func (self *MetathingsDevicedService) ValidateDeleteDevice(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() device_getter {
				req := in.(*pb.DeleteDeviceRequest)
				return req
			},
		},
		identityd_validator.Invokers{ensure_get_device_id},
	)
}

func (self *MetathingsDevicedService) AuthorizeDeleteDevice(ctx context.Context, in interface{}) error {
	return self.authorizer.Authorize(ctx, in.(*pb.DeleteDeviceRequest).GetDevice().GetId().GetValue(), "deviced:delete_device")
}

func (self *MetathingsDevicedService) delete_entity(cli identityd_pb.IdentitydServiceClient, ctx context.Context, ent_id string) error {
	var err error

	new_ctx := context_helper.WithToken(ctx, self.tknr.GetToken())
	req := &identityd_pb.DeleteEntityRequest{
		Entity: &identityd_pb.OpEntity{
			Id: &wrappers.StringValue{Value: ent_id},
		},
	}
	if _, err = cli.DeleteEntity(new_ctx, req); err != nil {
		return err
	}

	return nil
}

func (self *MetathingsDevicedService) DeleteDevice(ctx context.Context, req *pb.DeleteDeviceRequest) (*empty.Empty, error) {
	var dev *storage.Device
	var err error

	dev_id_str := req.GetDevice().GetId().GetValue()
	logger := self.get_logger().WithField("device", dev_id_str)

	if dev, err = self.storage.GetDevice(ctx, dev_id_str); err != nil {
		logger.WithError(err).Errorf("failed to get device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	identityd_cli, identityd_cfn, err := self.cli_fty.NewIdentityd2ServiceClient()
	if err != nil {
		logger.WithError(err).Errorf("failed to new identityd2 service client")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer identityd_cfn()

	for _, m := range dev.Modules {
		mdl_id_str := *m.Id

		if err = self.delete_entity(identityd_cli, ctx, mdl_id_str); err != nil {
			logger.WithError(err).WithField("id", mdl_id_str).Warningf("failed to delete module in identityd2")
		}

		if err = self.storage.DeleteModule(ctx, mdl_id_str); err != nil {
			logger.WithError(err).WithField("id", mdl_id_str).Warningf("failed to delete module in storage")
		}
	}

	for _, f := range dev.Flows {
		flw_id_str := *f.Id
		if err = self.storage.DeleteFlow(ctx, flw_id_str); err != nil {
			logger.WithError(err).WithField("id", flw_id_str).Warningf("failed to delete flow in storage")
		}
	}

	if err = self.delete_entity(identityd_cli, ctx, dev_id_str); err != nil {
		logger.WithError(err).Warningf("failed to delete device in identityd2")
	}

	if err = self.storage.DeleteDevice(ctx, dev_id_str); err != nil {
		logger.WithError(err).Debugf("failed to delete device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	logger.WithField("id", dev_id_str).Infof("delete device")

	return &empty.Empty{}, nil
}
