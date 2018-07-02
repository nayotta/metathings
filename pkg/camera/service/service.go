package metathings_camera_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	driver "github.com/nayotta/metathings/pkg/camera/driver"
	state_helper "github.com/nayotta/metathings/pkg/camera/state"
	app_cred_mgr "github.com/nayotta/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	driver_helper "github.com/nayotta/metathings/pkg/common/driver"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mt_plugin "github.com/nayotta/metathings/pkg/cored/plugin"
	pb "github.com/nayotta/metathings/pkg/proto/camera"
	camerad_pb "github.com/nayotta/metathings/pkg/proto/camerad"
)

type metathingsCameraService struct {
	mt_plugin.CoreService
	opt          opt_helper.Option
	logger       log.FieldLogger
	cli_fty      *client_helper.ClientFactory
	drv          driver.CameraDriver
	app_cred_mgr app_cred_mgr.ApplicationCredentialManager

	camera_st_psr              state_helper.CameraStateParser
	state_notification_channel chan driver.CameraState
}

func (srv *metathingsCameraService) copyCamera(cam driver.Camera) *pb.Camera {

	return &pb.Camera{
		State: srv.camera_st_psr.ToValue(cam.State.ToString()),
		Config: &pb.CameraConfig{
			Url:       cam.Config.Url,
			Device:    cam.Config.Device,
			Width:     cam.Config.Width,
			Height:    cam.Config.Height,
			Bitrate:   cam.Config.Bitrate,
			Framerate: cam.Config.Framerate,
		},
	}
}

func (srv *metathingsCameraService) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	cam, err := srv.drv.Show()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if cam.State != driver.STATE_STOP {
		return nil, status.Errorf(codes.FailedPrecondition, "camera not startable")
	}

	cfg := driver.CameraConfig{
		Url: req.Config.GetUrl().GetValue(),
	}
	srv.logger.WithField("url", cfg.Url).Debugf("set camera url")
	if dev := req.Config.GetDevice(); dev != nil {
		cfg.Device = dev.GetValue()
		srv.logger.WithField("device", cfg.Device).Debugf("set camera device")
	}
	if w := req.Config.GetWidth(); w != nil {
		cfg.Width = w.GetValue()
		srv.logger.WithField("width", cfg.Width).Debugf("set camera width")
	}
	if h := req.Config.GetHeight(); h != nil {
		cfg.Height = h.GetValue()
		srv.logger.WithField("height", cfg.Height).Debugf("set camera height")
	}
	if br := req.Config.GetBitrate(); br != nil {
		cfg.Bitrate = br.GetValue()
		srv.logger.WithField("bitrate", cfg.Bitrate).Debugf("set camera bitrate")
	}
	if fr := req.Config.GetFramerate(); fr != nil {
		cfg.Framerate = fr.GetValue()
		srv.logger.WithField("framerate", cfg.Framerate).Debugf("set camera framerate")
	}

	cam, err = srv.drv.Start(cfg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	srv.logger.Infof("start camera")

	return &pb.StartResponse{Camera: srv.copyCamera(cam)}, nil
}

func (srv *metathingsCameraService) Stop(ctx context.Context, req *empty.Empty) (*pb.StopResponse, error) {
	cam, err := srv.drv.Show()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if cam.State != driver.STATE_RUNNING {
		return nil, status.Errorf(codes.FailedPrecondition, "camera not stopable")
	}

	cam, err = srv.drv.Stop()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	srv.logger.Infof("stop camera")

	return &pb.StopResponse{Camera: srv.copyCamera(cam)}, nil
}

func (srv *metathingsCameraService) Show(ctx context.Context, req *empty.Empty) (*pb.ShowResponse, error) {
	cam, err := srv.drv.Show()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	srv.logger.WithField("camera", cam).Debugf("show camera")

	return &pb.ShowResponse{Camera: srv.copyCamera(cam)}, nil
}

func (srv *metathingsCameraService) Close() {
	err := srv.drv.Close()
	if err != nil {
		srv.logger.WithError(err).Debugf("failed to close driver")
	}

	if srv.state_notification_channel != nil {
		nc := srv.drv.(driver.NotificationCenter)
		nc.CloseStateNotificationChannel(srv.state_notification_channel)
	}

	srv.logger.Debugf("service closed")
}

func (srv *metathingsCameraService) state_notification_handler() {
	for {
		s := <-srv.state_notification_channel
		srv.logger.WithField("state", s.ToString()).Debugf("recv state notification")
		go func() {
			srv.update_camerad_state(s)
		}()
	}
}

func (srv *metathingsCameraService) update_camerad_state(s driver.CameraState) {
	cli, cfn, err := srv.cli_fty.NewCameradServiceClient()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to new camerad service client")
	}
	defer cfn()

	token_str := srv.app_cred_mgr.GetToken()
	ctx := context_helper.WithToken(context.Background(), token_str)

	req := &camerad_pb.CallbackRequest{
		State: srv.camera_st_psr.ToValue(s.ToString()),
	}

	_, err = cli.Callback(ctx, req)
	if err != nil {
		srv.logger.WithError(err).Debugf("failed to update camera state")
		return
	}
	srv.logger.WithField("state", s.ToString()).Debugf("update state in camerad")
}

func NewCameraService(opt opt_helper.Option) (*metathingsCameraService, error) {
	opt.Set("service_name", "camera")

	logger, err := log_helper.NewLogger("camera", opt.GetString("log.level"))
	if err != nil {
		return nil, err
	}

	cli_fty_cfgs := client_helper.NewDefaultServiceConfigs(opt.GetString("metathings.address"))
	cli_fty_cfgs[client_helper.AGENT_CONFIG] = client_helper.ServiceConfig{opt.GetString("agent.address")}
	cli_fty, err := client_helper.NewClientFactory(
		cli_fty_cfgs,
		client_helper.WithInsecureOptionFunc(),
	)
	if err != nil {
		return nil, err
	}

	app_cred_mgr, err := app_cred_mgr.NewApplicationCredentialManager(
		cli_fty,
		opt.GetString("application_credential.id"),
		opt.GetString("application_credential.secret"),
	)
	if err != nil {
		log.WithError(err).Errorf("failed to new application credential managaer")
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
		logger:       logger,
		cli_fty:      cli_fty,
		opt:          opt,
		drv:          cam_drv,
		app_cred_mgr: app_cred_mgr,

		camera_st_psr: state_helper.NewCameraStateParser(),
	}

	srv.CoreService = mt_plugin.MakeCoreService(srv.opt, srv.logger, srv.cli_fty)

	nc, ok := srv.drv.(driver.NotificationCenter)
	if ok {
		srv.state_notification_channel = nc.GetStateNotificationChannel()
		go srv.state_notification_handler()
	}

	return srv, nil
}
