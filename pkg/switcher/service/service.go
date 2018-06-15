package metathings_switcher_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	driver_helper "github.com/nayotta/metathings/pkg/common/driver"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mt_plugin "github.com/nayotta/metathings/pkg/cored/plugin"
	pb "github.com/nayotta/metathings/pkg/proto/switcher"
	driver "github.com/nayotta/metathings/pkg/switcher/driver"
	state_helper "github.com/nayotta/metathings/pkg/switcher/state"
)

type metathingsSwitcherService struct {
	mt_plugin.CoreService
	opt     opt_helper.Option
	logger  log.FieldLogger
	cli_fty *client_helper.ClientFactory
	drv     driver.SwitcherDriver

	switcher_st_psr state_helper.SwitcherStateParser
}

func (srv *metathingsSwitcherService) copySwitcher(sw driver.Switcher) *pb.Switcher {
	return &pb.Switcher{
		State: srv.switcher_st_psr.ToValue(sw.State.ToString()),
	}
}

func (srv *metathingsSwitcherService) Show(ctx context.Context, _ *empty.Empty) (*pb.ShowResponse, error) {
	sw, err := srv.drv.Show()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	srv.logger.WithField("switcher", sw).Debugf("show switcher")

	return &pb.ShowResponse{Switcher: srv.copySwitcher(sw)}, nil

}

func (srv *metathingsSwitcherService) Turn(ctx context.Context, req *pb.TurnRequest) (*pb.TurnResponse, error) {
	st := driver.FromValue(int32(req.State))
	if st == driver.STATE_UNKNOWN {
		return nil, status.Errorf(codes.InvalidArgument, "unsupported switcher state")
	}

	sw, err := srv.drv.Turn(st)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	srv.logger.WithField("switcher", sw).Infof("switcher state turning")

	return &pb.TurnResponse{Switcher: srv.copySwitcher(sw)}, nil
}

func (srv *metathingsSwitcherService) Close() {
	err := srv.drv.Close()
	if err != nil {
		srv.logger.WithError(err).Debugf("failed to close driver")
	}

	srv.logger.Debugf("service closed")
}

func NewSwitcherService(opt opt_helper.Option) (*metathingsSwitcherService, error) {
	opt.Set("service_name", "switcher")

	logger, err := log_helper.NewLogger("switcherd", opt.GetString("log.level"))
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

	drv_fty, err := driver_helper.NewDriverFactory(opt.GetString("driver.descriptor"))
	if err != nil {
		return nil, err
	}

	drv_name := opt.GetString("driver.name")
	drv, err := drv_fty.New(drv_name, opt)
	if err != nil {
		return nil, err
	}
	sw_drv, ok := drv.(driver.SwitcherDriver)
	if !ok {
		return nil, driver_helper.ErrUnmatchDriver
	}
	logger.WithField("driver_name", drv_name).Debugf("load switcher driver")

	opt.Set("logger", logger.WithField("#driver", drv_name))
	err = sw_drv.Init(opt)
	if err != nil {
		return nil, err
	}
	logger.Debugf("switcher driver initialized")

	srv := &metathingsSwitcherService{
		opt:     opt,
		logger:  logger,
		cli_fty: cli_fty,
		drv:     sw_drv,

		switcher_st_psr: state_helper.NewSwitcherStateParser(),
	}
	srv.CoreService = mt_plugin.MakeCoreService(srv.opt, srv.logger, srv.cli_fty)

	return srv, nil
}
