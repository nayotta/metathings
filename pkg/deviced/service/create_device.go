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
					return errors.New("domain.kind is invalid value")
				}

				if dev.GetName() == nil {
					return errors.New("domain.name is empty")
				}

				mdls := dev.GetModules()
				if len(mdls) == 0 {
					return errors.New("domain.modules too short")
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

func (self *MetathingsDevicedService) create_entity(ent_id, ent_name string) error {
	cli, cfn, err := self.cli_fty.NewIdentityd2ServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	ctx := context_helper.WithToken(context.Background(), self.tknr.GetToken())
	req := &identityd_pb.CreateEntityRequest{
		Entity: &identityd_pb.OpEntity{
			Id:   &wrappers.StringValue{Value: ent_id},
			Name: &wrappers.StringValue{Value: ent_name},
		},
	}
	if _, err = cli.CreateEntity(ctx, req); err != nil {
		return err
	}

	return nil
}

func (self *MetathingsDevicedService) create_device_entity(dev *storage.Device) error {
	return self.create_entity(*dev.Id, "/deviced/device/"+*dev.Name)
}

func (self *MetathingsDevicedService) create_module_entity(mdl *storage.Module) error {
	return self.create_entity(*mdl.Id, "/deviced/module/"+*mdl.Name)
}

func (self *MetathingsDevicedService) create_flow_entity(flw *storage.Flow) error {
	return self.create_entity(*flw.Id, "/deviced/flow/"+*flw.Name)
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

	dev_s := &storage.Device{
		Id:    &dev_id_str,
		Kind:  &dev_kind_str,
		State: &dev_state_str,
		Name:  &dev_name_str,
		Alias: &dev_alias_str,
	}

	if err = self.create_device_entity(dev_s); err != nil {
		self.logger.WithError(err).Errorf("failed to create entity for device")
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

		if err = self.create_module_entity(mdl_s); err != nil {
			self.logger.WithError(err).Errorf("failed to create entity for module")
			return nil, status.Errorf(codes.Internal, err.Error())
		}

		if _, err = self.storage.CreateModule(mdl_s); err != nil {
			self.logger.WithError(err).Errorf("failed to create module in storage")
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

		if err = self.create_flow_entity(flw_s); err != nil {
			self.logger.WithError(err).Errorf("failed to create entity to flow")
			return nil, status.Errorf(codes.Internal, err.Error())
		}

		if _, err = self.storage.CreateFlow(flw_s); err != nil {
			self.logger.WithError(err).Errorf("failed to create flow in storage")
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	if dev_s, err = self.storage.CreateDevice(dev_s); err != nil {
		self.logger.WithError(err).Errorf("failed to create device in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CreateDeviceResponse{
		Device: copy_device(dev_s),
	}

	self.logger.WithField("id", dev_id_str).Infof("create device")

	return res, nil
}
