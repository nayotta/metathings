package metathings_deviced_service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	connection "github.com/nayotta/metathings/pkg/deviced/connection"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_policy "github.com/nayotta/metathings/pkg/identityd2/policy"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
	identityd_pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

type MetathingsDevicedServiceOption struct {
}

type MetathingsDevicedService struct {
	grpc_helper.AuthorizationTokenParser

	tknr     token_helper.Tokener
	cli_fty  *client_helper.ClientFactory
	opt      *MetathingsDevicedServiceOption
	logger   log.FieldLogger
	storage  storage.Storage
	enforcer identityd_policy.Enforcer
	vdr      token_helper.TokenValidator
	cc       connection.ConnectionCenter
}

func (self *MetathingsDevicedService) get_device_by_context(ctx context.Context) (*storage.Device, error) {
	var tkn *identityd_pb.Token
	var dev_s *storage.Device
	var err error

	tkn = context_helper.ExtractToken(ctx)

	if dev_s, err = self.storage.GetDevice(tkn.Entity.Id); err != nil {
		return nil, err
	}

	return dev_s, nil
}

func (self *MetathingsDevicedService) enforce(ctx context.Context, obj, act string) error {
	var err error

	tkn := context_helper.ExtractToken(ctx)

	var groups []string
	for _, g := range tkn.Groups {
		groups = append(groups, g.Id)
	}

	if err = self.enforcer.Enforce(tkn.Domain.Id, groups, tkn.Entity.Id, obj, act); err != nil {
		if err == identityd_policy.ErrPermissionDenied {
			self.logger.WithFields(log.Fields{
				"subject": tkn.Entity.Id,
				"domain":  tkn.Domain.Id,
				"groups":  groups,
				"object":  obj,
				"action":  act,
			}).Warningf("denied to do #action")
			return status.Errorf(codes.PermissionDenied, err.Error())
		} else {
			self.logger.WithError(err).Errorf("failed to enforce")
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	return nil
}

func (self *MetathingsDevicedService) validate_chain(providers []interface{}, invokers []interface{}) error {
	default_invokers := []interface{}{policy_helper.ValidateValidator}
	invokers = append(default_invokers, invokers...)
	if err := policy_helper.ValidateChain(
		providers,
		invokers,
	); err != nil {
		self.logger.WithError(err).Warningf("failed to validate request data")
		return status.Errorf(codes.InvalidArgument, err.Error())
	}

	return nil
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

func NewMetathingsDevicedService(
	opt *MetathingsDevicedServiceOption,
	logger log.FieldLogger,
	storage storage.Storage,
	enforcer identityd_policy.Enforcer,
	vdr token_helper.TokenValidator,
	cc connection.ConnectionCenter,
	tknr token_helper.Tokener,
	cli_fty *client_helper.ClientFactory,
) (pb.DevicedServiceServer, error) {
	return &MetathingsDevicedService{
		opt:      opt,
		logger:   logger,
		storage:  storage,
		enforcer: enforcer,
		vdr:      vdr,
		cc:       cc,
		tknr:     tknr,
		cli_fty:  cli_fty,
	}, nil
}
