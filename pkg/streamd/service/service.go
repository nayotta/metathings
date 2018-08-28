package metathings_streamd_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lovoo/goka/storage"
	log "github.com/sirupsen/logrus"

	app_cred_mgr "github.com/nayotta/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	pb "github.com/nayotta/metathings/pkg/proto/streamd"
	state_helper "github.com/nayotta/metathings/pkg/streamd/state"
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
	panic("unimplemented")
}
