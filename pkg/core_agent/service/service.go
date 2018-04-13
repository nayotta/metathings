package meatathings_core_agent_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	log_helper "github.com/bigdatagz/metathings/pkg/common/log"
	pb "github.com/bigdatagz/metathings/pkg/proto/core_agent"
)

type options struct {
	token    string
	logLevel string

	cored_addr string
	core_id    string
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

func SetToken(tkn string) ServiceOptions {
	return func(o *options) {
		o.token = tkn
	}
}

func SetCoreId(id string) ServiceOptions {
	return func(o *options) {
		o.core_id = id
	}
}

type coreAgentService struct {
	logger log.FieldLogger
	opts   options
}

func (srv *coreAgentService) CreateEntity(context.Context, *pb.CreateEntityRequest) (*pb.CreateEntityResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *coreAgentService) DeleteEntity(context.Context, *pb.DeleteEntityRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *coreAgentService) PatchEntity(context.Context, *pb.PatchEntityRequest) (*pb.PatchEntityResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *coreAgentService) GetEntity(context.Context, *pb.GetEntityRequest) (*pb.GetEntityResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *coreAgentService) ListEntities(context.Context, *pb.ListEntitiesRequest) (*pb.ListEntitiesResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplemented")
}

func NewCoreAgentService(opt ...ServiceOptions) *coreAgentService {
	opts := defaultServiceOptions
	for _, o := range opt {
		o(&opts)
	}

	logger, err := log_helper.NewLogger("core-agent", opts.logLevel)
	if err != nil {
		log.Fatalf("failed to new logger: %v", err)
	}

	srv := &coreAgentService{
		logger: logger,
		opts:   opts,
	}
	return srv
}
