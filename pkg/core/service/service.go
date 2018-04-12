package metathings_core_service

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	pb "github.com/bigdatagz/metathings/pkg/proto/core"
)

type options struct {
	logLevel string
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

type metathingsCoreService struct {
	logger *log.Logger
	opts   options
}

func (srv *metathingsCoreService) CreateCore(context.Context, *pb.CreateCoreRequest) (*pb.CreateCoreResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

func (srv *metathingsCoreService) DeleteCore(context.Context, *pb.DeleteCoreRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

func (srv *metathingsCoreService) GetCore(context.Context, *pb.GetCoreRequest) (*pb.GetCoreResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

func (srv *metathingsCoreService) ListCores(context.Context, *pb.ListCoresRequest) (*pb.ListCoresResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

func (srv *metathingsCoreService) Heartbeat(context.Context, *pb.HeartbeatRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}
func (srv *metathingsCoreService) Pipeline(pb.CoreService_PipelineServer) error {
	return grpc.Errorf(codes.Unimplemented, "unimplement")
}

func (srv *metathingsCoreService) ListCoresForUser(context.Context, *pb.ListCoresForUserRequest) (*pb.ListCoresForUserResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

func (srv *metathingsCoreService) SendUnaryCall(context.Context, *pb.SendUnaryCallRequest) (*pb.SendUnaryCallResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}
