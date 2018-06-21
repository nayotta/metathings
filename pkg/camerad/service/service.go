package metathings_camerad_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	storage "github.com/nayotta/metathings/pkg/camerad/storage"
	app_cred_mgr "github.com/nayotta/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	pb "github.com/nayotta/metathings/pkg/proto/camerad"
)

type options struct {
	logLevel                      string
	identityd_addr                string
	cored_addr                    string
	application_credential_id     string
	application_credential_secret string
	storage_driver                string
	storage_uri                   string
	rtmp_addr                     string
}

var defaultServiceOptions = options{
	logLevel: "info",
}

type ServiceOptions func(*options)

func SetLogLevel(lvl string) ServiceOptions {
	return func(o *options) {
		o.logLevel = lvl
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

func SetRtmpAddr(addr string) ServiceOptions {
	return func(o *options) {
		o.rtmp_addr = addr
	}
}

type metathingsCameradService struct {
	grpc_helper.AuthorizationTokenParser

	cli_fty      *client_helper.ClientFactory
	app_cred_mgr app_cred_mgr.ApplicationCredentialManager
	logger       log.FieldLogger
	opts         options
	storage      storage.Storage
	tk_vdr       token_helper.TokenValidator
}

func (srv *metathingsCameradService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	token_str, err := srv.GetTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}

	token, err := srv.tk_vdr.Validate(token_str)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate token via identityd")
		return nil, err
	}

	ctx = context.WithValue(ctx, "token", token_str)
	ctx = context.WithValue(ctx, "credential", token)

	srv.logger.WithFields(log.Fields{
		"method":   fullMethodName,
		"user_id":  token.User.Id,
		"username": token.User.Name,
	}).Debugf("validator token")

	return ctx, nil
}

func (srv *metathingsCameradService) Create(context.Context, *pb.CreateRequest) (*pb.CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsCameradService) Delete(context.Context, *pb.DeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsCameradService) Patch(context.Context, *pb.PatchRequest) (*pb.PatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsCameradService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsCameradService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsCameradService) ListForUser(ctx context.Context, req *pb.ListForUserRequest) (*pb.ListForUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsCameradService) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsCameradService) Stop(ctx context.Context, req *pb.StopRequest) (*pb.StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsCameradService) Callback(ctx context.Context, req *pb.CallbackRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func NewCameradService(opt ...ServiceOptions) (*metathingsCameradService, error) {
	opts := defaultServiceOptions
	for _, o := range opt {
		o(&opts)
	}

	logger, err := log_helper.NewLogger("camerad", opts.logLevel)
	if err != nil {
		log.WithError(err).Errorf("failed to new logger")
		return nil, err
	}

	cli_fty_cfgs := client_helper.NewDefaultServiceConfigs(opts.identityd_addr)
	cli_fty_cfgs[client_helper.CORED_CONFIG] = client_helper.ServiceConfig{Address: opts.cored_addr}
	cli_fty_cfgs[client_helper.IDENTITYD_CONFIG] = client_helper.ServiceConfig{Address: opts.identityd_addr}
	cli_fty, err := client_helper.NewClientFactory(
		cli_fty_cfgs,
		client_helper.WithInsecureOptionFunc(),
	)

	storage, err := storage.NewStorage(opts.storage_driver, opts.storage_uri, logger)
	if err != nil {
		log.WithError(err).Errorf("failed to connect storage")
		return nil, err
	}

	app_cred_mgr, err := app_cred_mgr.NewApplicationCredentialManager(
		cli_fty,
		opts.application_credential_id,
		opts.application_credential_secret,
	)
	if err != nil {
		log.WithError(err).Errorf("failed to new application credential manager")
		return nil, err
	}

	tk_vdr := token_helper.NewTokenValidator(app_cred_mgr, cli_fty, logger)

	srv := &metathingsCameradService{
		cli_fty:      cli_fty,
		app_cred_mgr: app_cred_mgr,
		opts:         opts,
		logger:       logger,
		storage:      storage,
		tk_vdr:       tk_vdr,
	}
	return srv, nil
}
