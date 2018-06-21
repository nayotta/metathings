package metathings_camerad_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	state_helper "github.com/nayotta/metathings/pkg/camera/state"
	storage "github.com/nayotta/metathings/pkg/camerad/storage"
	"github.com/nayotta/metathings/pkg/common"
	app_cred_mgr "github.com/nayotta/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	camera_pb "github.com/nayotta/metathings/pkg/proto/camera"
	pb "github.com/nayotta/metathings/pkg/proto/camerad"
	cored_pb "github.com/nayotta/metathings/pkg/proto/cored"
)

type options struct {
	logLevel                      string
	identityd_addr                string
	cored_addr                    string
	application_credential_id     string
	application_credential_secret string
	storage_driver                string
	storage_uri                   string
	rtmp_addr                     string
}

var defaultServiceOptions = options{
	logLevel: "info",
}

type ServiceOptions func(*options)

func SetLogLevel(lvl string) ServiceOptions {
	return func(o *options) {
		o.logLevel = lvl
	}
}

func SetIdentitydAddr(addr string) ServiceOptions {
	return func(o *options) {
		o.identityd_addr = addr
	}
}

func SetCoredAddr(addr string) ServiceOptions {
	return func(o *options) {
		o.cored_addr = addr
	}
}

func SetApplicationCredential(id, secret string) ServiceOptions {
	return func(o *options) {
		o.application_credential_id = id
		o.application_credential_secret = secret
	}
}

func SetStorage(driver, uri string) ServiceOptions {
	return func(o *options) {
		o.storage_driver = driver
		o.storage_uri = uri
	}
}

func SetRtmpAddr(addr string) ServiceOptions {
	return func(o *options) {
		o.rtmp_addr = addr
	}
}

type metathingsCameradService struct {
	grpc_helper.AuthorizationTokenParser

	cli_fty       *client_helper.ClientFactory
	camera_st_psr state_helper.CameraStateParser
	app_cred_mgr  app_cred_mgr.ApplicationCredentialManager
	logger        log.FieldLogger
	opts          options
	storage       storage.Storage
	tk_vdr        token_helper.TokenValidator
}

func (srv *metathingsCameradService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	token_str, err := srv.GetTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}

	token, err := srv.tk_vdr.Validate(token_str)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate token via identityd")
		return nil, err
	}

	ctx = context.WithValue(ctx, "token", token_str)
	ctx = context.WithValue(ctx, "credential", token)

	srv.logger.WithFields(log.Fields{
		"method":   fullMethodName,
		"user_id":  token.User.Id,
		"username": token.User.Name,
	}).Debugf("validator token")

	return ctx, nil
}

func (srv *metathingsCameradService) copyCamera(c storage.Camera) *pb.Camera {
	return &pb.Camera{
		Id:   *c.Id,
		Name: *c.Name,
		Core: &cored_pb.Core{
			Id:      *c.CoreId,
			OwnerId: *c.OwnerId,
		},
		Entity: &cored_pb.Entity{
			Name: *c.EntityName,
		},
		State: srv.camera_st_psr.ToValue(*c.State),
		Config: &camera_pb.CameraConfig{
			Url:       *c.Url,
			Device:    *c.Device,
			Width:     *c.Width,
			Height:    *c.Height,
			Bitrate:   *c.Bitrate,
			Framerate: *c.Framerate,
		},
	}
}

func (srv *metathingsCameradService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	cred := context_helper.Credential(ctx)
	cam_id := common.NewId()
	var name_str string
	name := req.GetName()
	if name != nil {
		name_str = name.GetValue()
	} else {
		name_str = cam_id
	}
	core_id := req.GetCore().GetId().GetValue()
	entity_name := req.GetEntity().GetName().GetValue()
	state := "unknown"
	empty_str := ""
	var zero_int uint32 = 0

	cam := storage.Camera{
		Id:         &cam_id,
		Name:       &name_str,
		CoreId:     &core_id,
		EntityName: &entity_name,
		OwnerId:    &cred.User.Id,
		State:      &state,
		Url:        &empty_str,
		Device:     &empty_str,
		Width:      &zero_int,
		Height:     &zero_int,
		Bitrate:    &zero_int,
		Framerate:  &zero_int,
	}

	cfg := req.GetConfig()
	if cfg != nil {
		device := cfg.GetDevice()
		if device != nil {
			cam.Device = &device.Value
		}

		width := cfg.GetWidth()
		height := cfg.GetHeight()
		if width != nil && height != nil {
			cam.Width = &width.Value
			cam.Height = &height.Value
		}

		bitrate := cfg.GetBitrate()
		if bitrate != nil {
			cam.Bitrate = &bitrate.Value
		}

		framerate := cfg.GetFramerate()
		if framerate != nil {
			cam.Framerate = &framerate.Value
		}
	}

	cc, err := srv.storage.CreateCamera(cam)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to create camera")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	srv.logger.WithFields(log.Fields{
		"id":          *cc.Id,
		"name":        *cc.Name,
		"core_id":     *cc.CoreId,
		"entity_name": *cc.EntityName,
		"owner_id":    *cc.OwnerId,
		"state":       *cc.State,
	}).Infof("create camera")

	res := &pb.CreateResponse{
		Camera: srv.copyCamera(cc),
	}

	return res, nil
}

func (srv *metathingsCameradService) Delete(context.Context, *pb.DeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsCameradService) Patch(context.Context, *pb.PatchRequest) (*pb.PatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsCameradService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	c, err := srv.storage.GetCamera(req.GetId().GetValue())
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to get core")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	srv.logger.WithField("id", *c.Id).Debugf("get camera")

	res := &pb.GetResponse{
		Camera: srv.copyCamera(c),
	}

	return res, nil
}

func (srv *metathingsCameradService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	c := storage.Camera{}

	name := req.GetName()
	if name != nil {
		c.Name = &name.Value
	}

	core := req.GetCore()
	if core != nil {
		c.CoreId = &core.Id.Value
	}

	entity := req.GetEntity()
	if entity != nil {
		c.EntityName = &entity.Name.Value
	}

	state := req.GetState()
	if state != camera_pb.CameraState_CAMERA_STATE_UNKNOWN {
		state_str := srv.camera_st_psr.ToString(state)
		c.State = &state_str
	}

	cs, err := srv.storage.ListCameras(c)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to list cameras")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListResponse{
		Cameras: []*pb.Camera{},
	}

	for _, c := range cs {
		res.Cameras = append(res.Cameras, srv.copyCamera(c))
	}

	srv.logger.Debugf("list cameras")
	return res, nil
}

func (srv *metathingsCameradService) ListForUser(ctx context.Context, req *pb.ListForUserRequest) (*pb.ListForUserResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	cred := context_helper.Credential(ctx)
	user_id := cred.User.Id
	c := storage.Camera{}

	name := req.GetName()
	if name != nil {
		c.Name = &name.Value
	}

	core := req.GetCore()
	if core != nil {
		c.CoreId = &core.Id.Value
	}

	entity := req.GetEntity()
	if entity != nil {
		c.EntityName = &entity.Name.Value
	}

	state := req.GetState()
	if state != camera_pb.CameraState_CAMERA_STATE_UNKNOWN {
		state_str := srv.camera_st_psr.ToString(state)
		c.State = &state_str
	}

	cs, err := srv.storage.ListCamerasForUser(user_id, c)
	if err != nil {
		srv.logger.WithField("user_id", user_id).WithError(err).Errorf("failed to list cameras for user")
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListForUserResponse{
		Cameras: []*pb.Camera{},
	}
	for _, c := range cs {
		res.Cameras = append(res.Cameras, srv.copyCamera(c))
	}

	srv.logger.WithField("user_id", user_id).Debugf("list cameras for user")

	return res, nil
}

func (srv *metathingsCameradService) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsCameradService) Stop(ctx context.Context, req *pb.StopRequest) (*pb.StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsCameradService) Callback(ctx context.Context, req *pb.CallbackRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "unimplemented")
}

func NewCameradService(opt ...ServiceOptions) (*metathingsCameradService, error) {
	opts := defaultServiceOptions
	for _, o := range opt {
		o(&opts)
	}

	logger, err := log_helper.NewLogger("camerad", opts.logLevel)
	if err != nil {
		log.WithError(err).Errorf("failed to new logger")
		return nil, err
	}

	cli_fty_cfgs := client_helper.NewDefaultServiceConfigs(opts.identityd_addr)
	cli_fty_cfgs[client_helper.CORED_CONFIG] = client_helper.ServiceConfig{Address: opts.cored_addr}
	cli_fty_cfgs[client_helper.IDENTITYD_CONFIG] = client_helper.ServiceConfig{Address: opts.identityd_addr}
	cli_fty, err := client_helper.NewClientFactory(
		cli_fty_cfgs,
		client_helper.WithInsecureOptionFunc(),
	)

	storage, err := storage.NewStorage(opts.storage_driver, opts.storage_uri, logger)
	if err != nil {
		log.WithError(err).Errorf("failed to connect storage")
		return nil, err
	}

	app_cred_mgr, err := app_cred_mgr.NewApplicationCredentialManager(
		cli_fty,
		opts.application_credential_id,
		opts.application_credential_secret,
	)
	if err != nil {
		log.WithError(err).Errorf("failed to new application credential manager")
		return nil, err
	}

	tk_vdr := token_helper.NewTokenValidator(app_cred_mgr, cli_fty, logger)

	srv := &metathingsCameradService{
		cli_fty:       cli_fty,
		camera_st_psr: state_helper.NewCameraStateParser(),
		app_cred_mgr:  app_cred_mgr,
		opts:          opts,
		logger:        logger,
		storage:       storage,
		tk_vdr:        tk_vdr,
	}
	return srv, nil
}
