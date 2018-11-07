package metathings_deviced_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"

	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
	identityd_pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

type MetathingsDevicedServiceOption struct {
}

type MetathingsDevicedService struct {
	grpc_helper.AuthorizationTokenParser

	storage storage.Storage
	opt     *MetathingsDevicedServiceOption
	logger  log.FieldLogger
	vdr     token_helper.TokenValidator
}

func (self *MetathingsDevicedService) get_device_by_context(ctx context.Context) (*storage.Device, error) {
	var tkn *identityd_pb.Token
	var dev_s *storage.Device
	var err error

	tkn = context_helper.ExtractToken(ctx)

	if dev_s, err = self.storage.GetDeviceByEntityId(tkn.Entity.Id); err != nil {
		return nil, err
	}

	return dev_s, nil
}

func (self *MetathingsDevicedService) is_ignore_method(md *grpc_helper.MethodDescription) bool {
	return false
}

func (self *MetathingsDevicedService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	var tkn *identityd_pb.Token
	var tkn_txt string
	var new_ctx context.Context
	var err error
	var md *grpc_helper.MethodDescription

	if md, err = grpc_helper.ParseMethodDescription(fullMethodName); err != nil {
		self.logger.WithError(err).Warningf("failed to parse method description")
		return ctx, err
	}

	if self.is_ignore_method(md) {
		return ctx, nil
	}

	if tkn_txt, err = self.GetTokenFromContext(ctx); err != nil {
		self.logger.WithError(err).Warningf("failed to get token from context")
		return ctx, err
	}

	if tkn, err = self.vdr.Validate(tkn_txt); err != nil {
		self.logger.WithError(err).Warningf("failed to validate token in identity service")
		return ctx, err
	}

	new_ctx = context.WithValue(ctx, "token", tkn)

	self.logger.WithFields(log.Fields{
		"method":    md.Method,
		"entity_id": tkn.Entity.Id,
		"domain_id": tkn.Domain.Id,
	}).Debugf("authorize token")

	return new_ctx, nil
}

func (self *MetathingsDevicedService) DeleteDevice(context.Context, *pb.DeleteDeviceRequest) (*empty.Empty, error) {
	panic("unimplemented")
}

func (self *MetathingsDevicedService) PatchDevice(context.Context, *pb.PatchDeviceRequest) (*pb.PatchDeviceResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsDevicedService) ListDevices(context.Context, *pb.ListDevicesRequest) (*pb.ListDevicesResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsDevicedService) Connect(pb.DevicedService_ConnectServer) error {
	panic("unimplemented")
}

func (self *MetathingsDevicedService) UnaryCall(context.Context, *pb.UnaryCallRequest) (*pb.UnaryCallResponse, error) {
	panic("unimplemented")
}

func (self *MetathingsDevicedService) StreamCall(pb.DevicedService_StreamCallServer) error {
	panic("unimplemented")
}

func NewMetathingsDevicedService(
	opt *MetathingsDevicedServiceOption,
	logger log.FieldLogger,
) (pb.DevicedServiceServer, error) {
	return &MetathingsDevicedService{
		opt:    opt,
		logger: logger,
	}, nil
}
