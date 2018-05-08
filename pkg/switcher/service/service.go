package metathings_switcher_service

import (
	"context"
	"sync"

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

	mtx             *sync.Mutex
	state           string
	switcher_st_psr state_helper.SwitcherStateParser
}

func (srv *metathingsSwitcherService) Info(ctx context.Context, _ *empty.Empty) (*pb.InfoResponse, error) {
	srv.mtx.Lock()
	defer srv.mtx.Unlock()

	srv.logger.WithField("state", srv.state).Debugf("switcher info")
	return &pb.InfoResponse{
		Switcher: &pb.Switcher{
			State: srv.switcher_st_psr.ToValue(srv.state),
		},
	}, nil
}
func (srv *metathingsSwitcherService) Toggle(ctx context.Context, _ *empty.Empty) (*pb.ToggleResponse, error) {
	srv.mtx.Lock()
	defer srv.mtx.Unlock()

	st := srv.switcher_st_psr.ToValue(srv.state)
	if st == pb.SwitcherState_SWITCHER_STATE_UNKNOWN {
		return nil, status.Errorf(codes.Internal, "switcher state unknown")
	}

	res := &pb.ToggleResponse{
		Switcher: &pb.Switcher{},
	}
	if st == pb.SwitcherState_SWITCHER_STATE_ON {
		srv.state = "off"
		res.Switcher.State = pb.SwitcherState_SWITCHER_STATE_OFF
	} else if st == pb.SwitcherState_SWITCHER_STATE_OFF {
		srv.state = "on"
		res.Switcher.State = pb.SwitcherState_SWITCHER_STATE_ON
	}

	srv.logger.WithField("state", srv.state).Infof("switcher toggle")
	return res, nil
}
func (srv *metathingsSwitcherService) TurnOn(ctx context.Context, _ *empty.Empty) (*pb.TurnOnResponse, error) {
	srv.mtx.Lock()
	defer srv.mtx.Unlock()

	srv.state = "on"

	res := &pb.TurnOnResponse{
		Switcher: &pb.Switcher{State: pb.SwitcherState_SWITCHER_STATE_ON},
	}
	srv.logger.Infof("switcher turn on")

	return res, nil
}
func (srv *metathingsSwitcherService) TurnOff(ctx context.Context, _ *empty.Empty) (*pb.TurnOffResponse, error) {
	srv.mtx.Lock()
	defer srv.mtx.Unlock()

	srv.state = "off"

	res := &pb.TurnOffResponse{
		Switcher: &pb.Switcher{State: pb.SwitcherState_SWITCHER_STATE_OFF},
	}

	return res, nil
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
	}
	srv.CoreService = mt_plugin.MakeCoreService(srv.opts, srv.logger, srv.cli_fty)

	return nil, nil
}
