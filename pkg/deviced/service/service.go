package metathings_deviced_service

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	log "github.com/sirupsen/logrus"

	afo_helper "github.com/nayotta/metathings/pkg/common/auth_func_overrider"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	connection "github.com/nayotta/metathings/pkg/deviced/connection"
	flow "github.com/nayotta/metathings/pkg/deviced/flow"
	session_storage "github.com/nayotta/metathings/pkg/deviced/session_storage"
	simple_storage "github.com/nayotta/metathings/pkg/deviced/simple_storage"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	identityd_authorizer "github.com/nayotta/metathings/pkg/identityd2/authorizer"
	identityd_validator "github.com/nayotta/metathings/pkg/identityd2/validator"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
	identityd_pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

type MetathingsDevicedServiceOption struct{}

type MetathingsDevicedService struct {
	grpc_auth.ServiceAuthFuncOverride
	tknr            token_helper.Tokener
	cli_fty         *client_helper.ClientFactory
	opt             *MetathingsDevicedServiceOption
	logger          log.FieldLogger
	storage         storage.Storage
	session_storage session_storage.SessionStorage
	simple_storage  simple_storage.SimpleStorage
	authorizer      identityd_authorizer.Authorizer
	validator       identityd_validator.Validator
	tkvdr           token_helper.TokenValidator
	cc              connection.ConnectionCenter
	flw_fty         flow.FlowFactory
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

func (self *MetathingsDevicedService) new_flow(dev_id, flw_id string) (flow.Flow, error) {
	return self.flw_fty.New(&flow.FlowOption{
		FlowId:   flw_id,
		DeviceId: dev_id,
	})
}

func (self *MetathingsDevicedService) IsIgnoreMethod(md *grpc_helper.MethodDescription) bool {
	return false
}

func NewMetathingsDevicedService(
	opt *MetathingsDevicedServiceOption,
	logger log.FieldLogger,
	storage storage.Storage,
	session_storage session_storage.SessionStorage,
	simple_storage simple_storage.SimpleStorage,
	authorizer identityd_authorizer.Authorizer,
	validator identityd_validator.Validator,
	tkvdr token_helper.TokenValidator,
	cc connection.ConnectionCenter,
	tknr token_helper.Tokener,
	cli_fty *client_helper.ClientFactory,
	flw_fty flow.FlowFactory,
) (pb.DevicedServiceServer, error) {
	srv := &MetathingsDevicedService{
		opt:             opt,
		logger:          logger,
		storage:         storage,
		session_storage: session_storage,
		simple_storage:  simple_storage,
		authorizer:      authorizer,
		validator:       validator,
		tkvdr:           tkvdr,
		cc:              cc,
		tknr:            tknr,
		cli_fty:         cli_fty,
		flw_fty:         flw_fty,
	}
	srv.ServiceAuthFuncOverride = afo_helper.NewAuthFuncOverrider(tkvdr, srv, logger)

	return srv, nil
}
