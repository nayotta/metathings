package metathingsmqttdservice

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	identityd_policy "github.com/nayotta/metathings/pkg/identityd2/policy"
	storage "github.com/nayotta/metathings/pkg/mqttd/storage"
	identityd_pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	pb "github.com/nayotta/metathings/pkg/proto/mqttd"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Option MetathingsMqttdServiceOption struct
type Option struct {
}

// MetathingsMqttdService MetathingsMqttdService struct
type MetathingsMqttdService struct {
	grpc_helper.AuthorizationTokenParser

	tknr     token_helper.Tokener
	cliFty   *client_helper.ClientFactory
	opt      *Option
	logger   log.FieldLogger
	storage  storage.Storage
	enforcer identityd_policy.Enforcer
	vdr      token_helper.TokenValidator
}

func (sev *MetathingsMqttdService) getDeviceByContext(ctx context.Context) (*storage.Device, error) {
	var tkn *identityd_pb.Token
	var dev *storage.Device
	var err error

	tkn = context_helper.ExtractToken(ctx)

	if dev, err = sev.storage.GetDevice(tkn.Entity.Id); err != nil {
		return nil, err
	}

	return dev, nil
}

func (sev *MetathingsMqttdService) enforce(ctx context.Context, obj, act string) error {
	var err error

	tkn := context_helper.ExtractToken(ctx)

	var groups []string
	for _, g := range tkn.Groups {
		groups = append(groups, g.Id)
	}

	if err = sev.enforcer.Enforce(tkn.Domain.Id, groups, tkn.Entity.Id, obj, act); err != nil {
		if err == identityd_policy.ErrPermissionDenied {
			sev.logger.WithFields(log.Fields{
				"subject": tkn.Entity.Id,
				"domain":  tkn.Domain.Id,
				"groups":  groups,
				"object":  obj,
				"action":  act,
			}).Warningf("denied to do #action")
			return status.Errorf(codes.PermissionDenied, err.Error())
		} else {
			sev.logger.WithError(err).Errorf("failed to enforce")
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	return nil
}

func (sev *MetathingsMqttdService) validateChain(providers []interface{}, invokers []interface{}) error {
	defaultInvokers := []interface{}{policy_helper.ValidateValidator}
	invokers = append(defaultInvokers, invokers...)
	if err := policy_helper.ValidateChain(
		providers,
		invokers,
	); err != nil {
		sev.logger.WithError(err).Warningf("failed to validate request data")
		return status.Errorf(codes.InvalidArgument, err.Error())
	}

	return nil
}

func (sev *MetathingsMqttdService) isIgnoreMethod(md *grpc_helper.MethodDescription) bool {
	return false
}

// AuthFuncOverride AuthFuncOverride
func (sev *MetathingsMqttdService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	var tkn *identityd_pb.Token
	var tknTxt string
	var newCtx context.Context
	var err error
	var md *grpc_helper.MethodDescription

	if md, err = grpc_helper.ParseMethodDescription(fullMethodName); err != nil {
		sev.logger.WithError(err).Warningf("failed to parse method description")
		return ctx, err
	}

	if sev.isIgnoreMethod(md) {
		return ctx, nil
	}

	if tknTxt, err = sev.GetTokenFromContext(ctx); err != nil {
		sev.logger.WithError(err).Warningf("failed to get token from context")
		return ctx, err
	}

	if tkn, err = sev.vdr.Validate(tknTxt); err != nil {
		sev.logger.WithError(err).Warningf("failed to validate token in identity service")
		return ctx, err
	}

	newCtx = context.WithValue(ctx, "token", tkn)

	sev.logger.WithFields(log.Fields{
		"method":    md.Method,
		"entity_id": tkn.Entity.Id,
		"domain_id": tkn.Domain.Id,
	}).Debugf("authorize token")

	return newCtx, nil
}

func (sev *MetathingsMqttdService) ListDevices(context.Context, *pb.ListDevicesRequest) (*pb.ListDevicesResponse, error) {

}

func (sev *MetathingsMqttdService) ShowDevice(context.Context, *empty.Empty) (*pb.ShowDeviceResponse, error) {

}

func (sev *MetathingsMqttdService) UnaryCall(context.Context, *pb.UnaryCallRequest) (*pb.UnaryCallResponse, error) {

}

func (sev *MetathingsMqttdService) StreamCall(context.Context, *pb.StreamCallRequest) (*pb.StreamCallResponse, error) {

}

// NewMetathingsMqttdService NewMetathingsMqttdService
func NewMetathingsMqttdService(
	opt *Option,
	logger log.FieldLogger,
	storage storage.Storage,
	enforcer identityd_policy.Enforcer,
	vdr token_helper.TokenValidator,
	tknr token_helper.Tokener,
	cliFty *client_helper.ClientFactory,
) (pb.MqttdServiceServer, error) {
	return &MetathingsMqttdService{
		opt:      opt,
		logger:   logger,
		storage:  storage,
		enforcer: enforcer,
		vdr:      vdr,
		tknr:     tknr,
		cliFty:   cliFty,
	}, nil
}
