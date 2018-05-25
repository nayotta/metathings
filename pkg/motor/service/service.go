package metathings_motor_service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mt_plugin "github.com/nayotta/metathings/pkg/core/plugin"
	pb "github.com/nayotta/metathings/pkg/proto/motor"
)

type metathingsMotorService struct {
	mt_plugin.CoreService
	opt     opt_helper.Option
	logger  log.FieldLogger
	cli_fty *client_helper.ClientFactory
}

func (srv *metathingsMotorService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsMotorService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsMotorService) StreamNormalMotor(stream pb.MotorService_StreamNormalMotorServer) error {
	return status.Errorf(codes.Unimplemented, "unimplemented")
}

func NewMotorService(opt opt_helper.Option) (*metathingsMotorService, error) {
	opt.Set("service_name", "motor")

	logger, err := log_helper.NewLogger("motord", opt.GetString("log.level"))
	if err != nil {
		return nil, err
	}

	cli_fty_cfgs := client_helper.NewDefaultServiceConfigs(opt.GetString("metathings.address"))
	cli_fty_cfgs[client_helper.AGENTD_CONFIG] = client_helper.ServiceConfig{opt.GetString("agent.address")}
	cli_fty, err := client_helper.NewClientFactory(
		cli_fty_cfgs,
		client_helper.WithInsecureOptionFunc(),
	)
	if err != nil {
		return nil, err
	}

	srv := &metathingsMotorService{
		opt:     opt,
		logger:  logger,
		cli_fty: cli_fty,
	}
	srv.CoreService = mt_plugin.MakeCoreService(srv.opt, srv.logger, srv.cli_fty)

	return srv, nil
}
