package metathings_sensor_service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mt_plugin "github.com/nayotta/metathings/pkg/cored/plugin"
	pb "github.com/nayotta/metathings/pkg/proto/sensor"
)

type metathingsSensorService struct {
	mt_plugin.CoreService
	opts    opt_helper.Option
	logger  log.FieldLogger
	cli_fty *client_helper.ClientFactory
}

func (srv *metathingsSensorService) Get(context.Context, *pb.GetRequest) (*pb.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsSensorService) List(context.Context, *pb.ListRequest) (*pb.ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsSensorService) Patch(context.Context, *pb.PatchRequest) (*pb.PatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsSensorService) GetData(context.Context, *pb.GetDataRequest) (*pb.GetDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsSensorService) ListData(context.Context, *pb.ListDataRequest) (*pb.ListDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func NewSensorService(opts opt_helper.Option) (*metathingsSensorService, error) {
	opts.Set("service_name", "sensor")

	logger, err := log_helper.NewLogger("sensor", opts.GetString("log.level"))
	if err != nil {
		return nil, err
	}

	cli_fty_cfgs := client_helper.NewDefaultServiceConfigs(opts.GetString("metathings.address"))
	cli_fty_cfgs[client_helper.AGENT_CONFIG] = client_helper.ServiceConfig{opts.GetString("agent.address")}
	cli_fty, err := client_helper.NewClientFactory(
		cli_fty_cfgs,
		client_helper.WithInsecureOptionFunc(),
	)
	if err != nil {
		return nil, err
	}

	srv := &metathingsSensorService{
		opts:    opts,
		logger:  logger,
		cli_fty: cli_fty,
	}

	srv.CoreService = mt_plugin.MakeCoreService(srv.opts, srv.logger, srv.cli_fty)

	return srv, nil
}
