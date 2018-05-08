package metathings_switcher_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	client_helper "github.com/bigdatagz/metathings/pkg/common/client"
	cs_helper "github.com/bigdatagz/metathings/pkg/common/core_service"
	log_helper "github.com/bigdatagz/metathings/pkg/common/log"
	mt_plugin "github.com/bigdatagz/metathings/pkg/core/plugin"
	pb "github.com/bigdatagz/metathings/pkg/proto/switcher"
	state_helper "github.com/bigdatagz/metathings/pkg/switcher/state"
)

type metathingsSwitcherService struct {
	mt_plugin.CoreService
	opts    cs_helper.Options
	logger  log.FieldLogger
	cli_fty *client_helper.ClientFactory

	switcher_st_psr state_helper.SwitcherStateParser
}

func (srv *metathingsSwitcherService) Get(ctx context.Context, _ *empty.Empty) (*pb.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")

}

func (srv *metathingsSwitcherService) Turn(ctx context.Context, req *pb.TurnRequest) (*pb.TurnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func NewSwitcherService(opts cs_helper.Options) (*metathingsSwitcherService, error) {
	opts.Set("service_name", "switcher")

	logger, err := log_helper.NewLogger("switcherd", opts.GetString("log.level"))
	if err != nil {
		return nil, err
	}

	cli_fty_cfgs := client_helper.NewDefaultServiceConfigs(opts.GetString("metathings.address"))
	cli_fty_cfgs[client_helper.AGENTD_CONFIG] = client_helper.ServiceConfig{opts.GetString("agent.address")}
	cli_fty, err := client_helper.NewClientFactory(
		cli_fty_cfgs,
		client_helper.WithInsecureOptionFunc(),
	)
	if err != nil {
		return nil, err
	}

	srv := &metathingsSwitcherService{
		opts:    opts,
		logger:  logger,
		cli_fty: cli_fty,

		switcher_st_psr: state_helper.NewSwitcherStateParser(),
	}
	srv.CoreService = mt_plugin.MakeCoreService(srv.opts, srv.logger, srv.cli_fty)

	return nil, nil
}
