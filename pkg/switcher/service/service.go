package metathings_switcher_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/bigdatagz/metathings/pkg/proto/switcher"
)

type metathingsSwitcherService struct{}

func (srv *metathingsSwitcherService) Info(ctx context.Context, _ *empty.Empty) (*pb.InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}
func (srv *metathingsSwitcherService) Toggle(ctx context.Context, _ *empty.Empty) (*pb.ToggleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}
func (srv *metathingsSwitcherService) TurnOn(ctx context.Context, _ *empty.Empty) (*pb.TurnOnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}
func (srv *metathingsSwitcherService) TurnOff(ctx context.Context, _ *empty.Empty) (*pb.TurnOffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}
