package metathings_sensord_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nayotta/metathings/pkg/proto/sensord"
)

type options struct{}

type ServiceOptions func(*options)

type metathingsSensordService struct{}

func (srv *metathingsSensordService) Create(context.Context, *pb.CreateRequest) (*pb.CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsSensordService) Delete(context.Context, *pb.DeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsSensordService) Patch(context.Context, *pb.PatchRequest) (*pb.PatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsSensordService) Get(context.Context, *pb.GetRequest) (*pb.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsSensordService) List(context.Context, *pb.ListRequest) (*pb.ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsSensordService) ListForUser(context.Context, *pb.ListForUserRequest) (*pb.ListForUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsSensordService) Subscribe(pb.SensordService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsSensordService) Publish(pb.SensordService_PublishServer) error {
	return status.Errorf(codes.Unimplemented, "unimplemented")
}

func NewSensordService(opt ...ServiceOptions) (*metathingsSensordService, error) {
	srv := &metathingsSensordService{}
	return srv, nil
}
