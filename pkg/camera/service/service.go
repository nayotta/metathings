package metathings_camera_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	driver "github.com/nayotta/metathings/pkg/camera/driver"
	state_helper "github.com/nayotta/metathings/pkg/camera/state"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	driver_helper "github.com/nayotta/metathings/pkg/common/driver"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mt_plugin "github.com/nayotta/metathings/pkg/core/plugin"
	pb "github.com/nayotta/metathings/pkg/proto/camera"
)

type metathingsCameraService struct {
	mt_plugin.CoreService
	opt     opt_helper.Option
	logger  log.FieldLogger
	cli_fty *client_helper.ClientFactory
	drv     driver.CameraDriver

	camera_st_psr state_helper.CameraStateParser
}

func (srv *metathingsCameraService) copyCamera() (*pb.Camera, error) {
	cam, err := srv.drv.Show()
	if err != nil {
		return nil, err
	}

	cfg := cam.Config

	return &pb.Camera{
		State: srv.camera_st_psr.ToValue(cam.State.ToString()),
		Config: &pb.CameraConfig{
			Url:       cfg.Url,
			Device:    cfg.Device,
			Width:     cfg.Width,
			Height:    cfg.Height,
			Bitrate:   cfg.Bitrate,
			Framerate: cfg.Framerate,
		},
	}, nil
}

func (srv *metathingsCameraService) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsCameraService) Stop(ctx context.Context, req *empty.Empty) (*pb.StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsCameraService) Show(ctx context.Context, req *empty.Empty) (*pb.ShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func NewCameraService(opt opt_helper.Option) (*metathingsCameraService, error) {
	opt.Set("service_name", "camera")

	logger, err := log_helper.NewLogger("camerad", opt.GetString("log.level"))
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
	cam_drv, ok := drv.(driver.CameraDriver)
	if !ok {
		return nil, driver_helper.ErrUnmatchDriver
	}
	logger.WithField("driver_name", drv_name).Debugf("load camera driver")

	opt.Set("logger", logger.WithField("#driver", drv_name))
	err = cam_drv.Init(opt)
	if err != nil {
		return nil, err
	}
	logger.Debugf("camera driver initialized")

	srv := &metathingsCameraService{
		logger:  logger,
		cli_fty: cli_fty,
		opt:     opt,
		drv:     cam_drv,

		camera_st_psr: state_helper.NewCameraStateParser(),
	}

	srv.CoreService = mt_plugin.MakeCoreService(srv.opt, srv.logger, srv.cli_fty)

	return srv, nil
}
