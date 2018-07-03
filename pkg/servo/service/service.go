package metathings_servo_service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mt_plugin "github.com/nayotta/metathings/pkg/cored/plugin"
	pb "github.com/nayotta/metathings/pkg/proto/servo"
)

type metathingsServoService struct {
	mt_plugin.CoreService
	opts    opt_helper.Option
	logger  log.FieldLogger
	cli_fty *client_helper.ClientFactory
}

func (srv *metathingsServoService) List(context.Context, *pb.ListRequest) (*pb.ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}
func (srv *metathingsServoService) Get(context.Context, *pb.GetRequest) (*pb.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}
func (srv *metathingsServoService) Stream(pb.ServoService_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "unimplemented")
}

func NewServoService(opts opt_helper.Option) (*metathingsServoService, error) {
	srv := &metathingsServoService{}
	return srv, nil
}
