package metathings_streamd_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"

	app_cred_mgr "github.com/nayotta/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	pb "github.com/nayotta/metathings/pkg/proto/streamd"
	state_helper "github.com/nayotta/metathings/pkg/streamd/state"
	storage "github.com/nayotta/metathings/pkg/streamd/storage"
)

type options struct {
	logLevel                      string
	metathingsd_addr              string
	identityd_addr                string
	cored_addr                    string
	application_credential_id     string
	application_credential_secret string
	storage_driver                string
	storage_uri                   string
}

type ServiceOptions func(*options)

var defaultServiceOptions = options{
	logLevel: "info",
}

func SetLogLevel(lvl string) ServiceOptions {
	return func(o *options) {
		o.logLevel = lvl
	}
}

func SetMetathingsdAddr(addr string) ServiceOptions {
	return func(o *options) {
		o.metathingsd_addr = addr
	}
}

func SetIdentitydAddr(addr string) ServiceOptions {
	return func(o *options) {
		o.identityd_addr = addr
	}
}

func SetCoredAddr(addr string) ServiceOptions {
	return func(o *options) {
		o.cored_addr = addr
	}
}

func SetApplicationCredential(id, secret string) ServiceOptions {
	return func(o *options) {
		o.application_credential_id = id
		o.application_credential_secret = secret
	}
}

func SetStorage(driver, uri string) ServiceOptions {
	return func(o *options) {
		o.storage_driver = driver
		o.storage_uri = uri
	}
}

type metathingsStreamdService struct {
	grpc_helper.AuthorizationTokenParser

	cli_fty       *client_helper.ClientFactory
	stream_st_psr state_helper.StreamStateParser
	app_cred_mgr  app_cred_mgr.ApplicationCredentialManager
	logger        log.FieldLogger
	opts          options
	storage       storage.Storage
	tk_vdr        token_helper.TokenValidator
}

func (self *metathingsStreamdService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	token_str, err := self.GetTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}

	token, err := self.tk_vdr.Validate(token_str)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to validate token via identityd")
		return nil, err
	}

	ctx = context.WithValue(ctx, "token", token_str)
	ctx = context.WithValue(ctx, "credential", token)

	self.logger.WithFields(log.Fields{
		"method":   fullMethodName,
		"user_id":  token.User.Id,
		"username": token.User.Name,
	}).Debugf("validate token")

	return ctx, nil
}

func (self *metathingsStreamdService) Create(context.Context, *pb.CreateRequest) (*pb.CreateResponse, error) {
	panic("unimplemented")
}

func (self *metathingsStreamdService) Delete(context.Context, *pb.DeleteRequest) (*empty.Empty, error) {
	panic("unimplemented")
}

func (self *metathingsStreamdService) Start(context.Context, *pb.StartRequest) (*pb.StartResponse, error) {
	panic("unimplemented")
}

func (self *metathingsStreamdService) Stop(context.Context, *pb.StopRequest) (*pb.StopResponse, error) {
	panic("unimplemented")
}

func (self *metathingsStreamdService) Get(context.Context, *pb.GetRequest) (*pb.GetResponse, error) {
	panic("unimplemented")
}

func (self *metathingsStreamdService) List(context.Context, *pb.ListRequest) (*pb.ListResponse, error) {
	panic("unimplemented")
}

func (self *metathingsStreamdService) ListForUser(context.Context, *pb.ListForUserRequest) (*pb.ListForUserResponse, error) {
	panic("unimplemented")
}

func NewStreamdService(opt ...ServiceOptions) (*metathingsStreamdService, error) {
	opts := defaultServiceOptions
	for _, o := range opt {
		o(&opts)
	}

	logger, err := log_helper.NewLogger("streamd", opts.logLevel)
	if err != nil {
		return nil, err
	}

	cli_fty_cfgs := client_helper.NewDefaultServiceConfigs(opts.metathingsd_addr)
	cli_fty_cfgs[client_helper.CORED_CONFIG] = client_helper.ServiceConfig{Address: opts.cored_addr}
	cli_fty_cfgs[client_helper.IDENTITYD_CONFIG] = client_helper.ServiceConfig{Address: opts.identityd_addr}
	cli_fty, err := client_helper.NewClientFactory(
		cli_fty_cfgs,
		client_helper.WithInsecureOptionFunc(),
	)
	if err != nil {
		log.WithError(err).Errorf("failed to new client factory")
		return nil, err
	}

	storage, err := storage.NewStorage(opts.storage_driver, opts.storage_uri, logger)
	if err != nil {
		logger.WithError(err).Errorf("failed to connect storage")
		return nil, err
	}

	app_cred_mgr, err := app_cred_mgr.NewApplicationCredentialManager(
		cli_fty,
		opts.application_credential_id,
		opts.application_credential_secret,
	)
	if err != nil {
		logger.WithError(err).Errorf("failed to new application credential manager")
		return nil, err
	}

	tk_vdr := token_helper.NewTokenValidator(app_cred_mgr, cli_fty, logger)

	srv := &metathingsStreamdService{
		cli_fty:       cli_fty,
		stream_st_psr: state_helper.NewStreamStateParser(),
		app_cred_mgr:  app_cred_mgr,
		opts:          opts,
		logger:        logger,
		storage:       storage,
		tk_vdr:        tk_vdr,
	}

	logger.Debugf("new streamd service")

	return srv, nil
}
