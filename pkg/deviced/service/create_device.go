package metathings_deviced_service

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	context_helper "github.com/nayotta/metathings/pkg/common/context"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	deviced_helper "github.com/nayotta/metathings/pkg/deviced/helper"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb_kind "github.com/nayotta/metathings/pkg/proto/constant/kind"
	pb_state "github.com/nayotta/metathings/pkg/proto/constant/state"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
	identityd_pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

func (self *MetathingsDevicedService) ValidateCreateDevice(ctx context.Context, in interface{}) error {
	return self.validator.Validate(
		identityd_validator.Providers{
			func() (policy_helper.Validator, device_getter) {
				req := in.(*pb.CreateDeviceRequest)
				return req, req
			},
		},
		identityd_validator.Invokers{
			func(x device_getter) error {
				dev := x.GetDevice()

				if dev.GetKind() == pb_kind.DeviceKind_DEVICE_KIND_UNKNOWN {
					return errors.New("device.kind is invalid value")
				}

				if dev.GetName() == nil {
					return errors.New("device.name is empty")
				}

				mdls := dev.GetModules()
				if len(mdls) == 0 {
					return errors.New("device.modules too short")
				}

				for _, mdl := range mdls {
					if mdl.GetEndpoint() == nil {
						return errors.New("model.endpoint is empty")
					}

					if mdl.GetComponent() == nil {
						return errors.New("model.component is empty")
					}

					if mdl.GetName() == nil {
						return errors.New("model.name is empty")
					}
				}

				return nil
			},
		},
	)
}

func (self *MetathingsDevicedService) create_entity(ctx context.Context, ent_id, ent_name string) error {
	cli, cfn, err := self.cli_fty.NewIdentityd2ServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	new_ctx := context_helper.WithToken(ctx, self.tknr.GetToken())
	req := &identityd_pb.CreateEntityRequest{
		Entity: &identityd_pb.OpEntity{
			Id:   &wrappers.StringValue{Value: ent_id},
			Name: &wrappers.StringValue{Value: ent_name},
		},
	}
	if _, err = cli.CreateEntity(new_ctx, req); err != nil {
		return err
	}

	return nil
}

func (self *MetathingsDevicedService) create_device_entity(ctx context.Context, dev *storage.Device) error {
	return self.create_entity(ctx, *dev.Id, "/deviced/device/"+*dev.Name)
}

func (self *MetathingsDevicedService) create_module_entity(ctx context.Context, mdl *storage.Module) error {
	return self.create_entity(ctx, *mdl.Id, "/deviced/module/"+*mdl.Name)
}

func (self *MetathingsDevicedService) create_flow_entity(ctx context.Context, flw *storage.Flow) error {
	return self.create_entity(ctx, *flw.Id, "/deviced/flow/"+*flw.Name)
}

func (self *MetathingsDevicedService) CreateDevice(ctx context.Context, req *pb.CreateDeviceRequest) (*pb.CreateDeviceResponse, error) {
	var err error

	dev := req.GetDevice()

	dev_id_str := id_helper.NewId()
	if dev.GetId() != nil {
		dev_id_str = dev.GetId().GetValue()
	}
	dev_kind_str := deviced_helper.DEVICE_KIND_ENUMER.ToString(dev.GetKind())
	dev_state_str := deviced_helper.DEVICE_STATE_ENUMER.ToString(pb_state.DeviceState_DEVICE_STATE_OFFLINE)
	dev_name_str := dev.GetName().GetValue()
	dev_alias_str := dev_name_str
	if dev.GetAlias() != nil {
		dev_alias_str = dev.GetAlias().GetValue()
	}

	logger := self.get_logger().WithField("device", dev_id_str)

	dev_s := &storage.Device{
		Id:    &dev_id_str,
		Kind:  &dev_kind_str,
		State: &dev_state_str,
		Name:  &dev_name_str,
		Alias: &dev_alias_str,
	}

	if extra := dev.GetExtra(); extra != nil {
		dev_s.ExtraHelper = pb_helper.ExtractStringMapToString(extra)
	}

	if err = self.create_device_entity(ctx, dev_s); err != nil {
		logger.WithError(err).Errorf("failed to create entity for device")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for _, mdl := range dev.GetModules() {
		mdl_id_str := id_helper.NewId()
		if mdl.GetId() != nil {
			mdl_id_str = mdl.GetId().GetValue()
		}
		mdl_state_str := deviced_helper.MODULE_STATE_ENUMER.ToString(pb_state.ModuleState_MODULE_STATE_OFFLINE)
		mdl_component_str := mdl.GetComponent().GetValue()
		mdl_name_str := mdl.GetName().GetValue()
		mdl_alias_str := mdl.GetAlias().GetValue()
		mdl_endpoint_str := mdl.GetEndpoint().GetValue()

		mdl_s := &storage.Module{
			DeviceId:  &dev_id_str,
			Id:        &mdl_id_str,
			State:     &mdl_state_str,
			Component: &mdl_component_str,
			Name:      &mdl_name_str,
			Alias:     &mdl_alias_str,
			Endpoint:  &mdl_endpoint_str,
		}

		if err = self.create_module_entity(ctx, mdl_s); err != nil {
			logger.WithError(err).Errorf("failed to create entity for module")
			return nil, status.Errorf(codes.Internal, err.Error())
		}

		if _, err = self.storage.CreateModule(ctx, mdl_s); err != nil {
			logger.WithError(err).Errorf("failed to create module in storage")
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	for _, flw := range dev.GetFlows() {
		flw_id_str := id_helper.NewId()
		if flw.GetId() != nil {
			flw_id_str = flw.GetId().GetValue()
		}

		flw_name_str := flw.GetName().GetValue()
		flw_alias_str := flw.GetAlias().GetValue()

		flw_s := &storage.Flow{
			DeviceId: &dev_id_str,
			Id:       &flw_id_str,
			Name:     &flw_name_str,
			Alias:    &flw_alias_str,
		}

		if err = self.create_flow_entity(ctx, flw_s); err != nil {
			logger.WithError(err).Errorf("failed to create entity to flow")
			return nil, status.Errorf(codes.Internal, err.Error())
		}

		if _, err = self.storage.CreateFlow(ctx, flw_s); err != nil {
			logger.WithError(err).Errorf("failed to create flow in storage")
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	if dev_s, err = self.storage.CreateDevice(ctx, dev_s); err != nil {
		logger.WithError(err).Errorf("failed to create device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CreateDeviceResponse{
		Device: copy_device(dev_s),
	}

	logger.Infof("create device")

	return res, nil
}
