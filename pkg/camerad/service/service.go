package metathings_camerad_service

import (
	"context"
	"fmt"
	"math/rand"
	"net/url"
	"path"

	"github.com/golang/protobuf/ptypes/empty"
	gpb "github.com/golang/protobuf/ptypes/wrappers"
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
)

type options struct {
	logLevel                      string
	metathingsd_addr              string
	identityd_addr                string
	cored_addr                    string
	application_credential_id     string
	application_credential_secret string
	storage_driver                string
	storage_uri                   string
	rtmp_url                      string
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

func SetMetathingsdAddr(addr string) ServiceOptions {
	return func(o *options) {
		o.metathingsd_addr = addr
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

func SetRtmpUrl(url string) ServiceOptions {
	return func(o *options) {
		o.rtmp_url = url
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

func (srv *metathingsCameradService) ContextWithToken(ctxs ...context.Context) context.Context {
	ctx := context.Background()
	if len(ctxs) > 0 {
		ctx = ctxs[0]
	}
	token_str := srv.app_cred_mgr.GetToken()
	ctx = context_helper.WithToken(ctx, token_str)
	return ctx
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
	cfg := &camera_pb.CameraConfig{}
	if c.Url != nil {
		cfg.Url = *c.Url
	}
	if c.Device != nil {
		cfg.Device = *c.Device
	}
	if c.Width != nil && c.Height != nil {
		cfg.Width = *c.Width
		cfg.Height = *c.Height
	}
	if c.Bitrate != nil {
		cfg.Bitrate = *c.Bitrate
	}
	if c.Framerate != nil {
		cfg.Framerate = *c.Framerate
	}

	return &pb.Camera{
		Id:         *c.Id,
		Name:       *c.Name,
		CoreId:     *c.CoreId,
		OwnerId:    *c.OwnerId,
		EntityName: *c.EntityName,
		State:      srv.camera_st_psr.ToValue(*c.State),
		Config:     cfg,
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
	core_id := req.GetCoreId().GetValue()
	entity_name := req.GetEntityName().GetValue()
	app_cred_id := req.GetApplicationCredentialId().GetValue()
	state := "unknown"
	empty_str := ""
	var zero_int uint32 = 0

	cam := storage.Camera{
		Id:                      &cam_id,
		Name:                    &name_str,
		CoreId:                  &core_id,
		EntityName:              &entity_name,
		OwnerId:                 &cred.User.Id,
		ApplicationCredentialId: &app_cred_id,
		State:     &state,
		Url:       &empty_str,
		Device:    &empty_str,
		Width:     &zero_int,
		Height:    &zero_int,
		Bitrate:   &zero_int,
		Framerate: &zero_int,
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
	go func() {
		cli, cfn, err := srv.cli_fty.NewCoredServiceClient()
		if err != nil {
			srv.logger.WithError(err).Errorf("failed to create cored service client")
			return
		}
		defer cfn()

		req := client_helper.MustNewUnaryCallRequest(
			core_id, entity_name, "camera", "Show", &empty.Empty{},
		)

		res, err := cli.UnaryCall(srv.ContextWithToken(), req)
		if err != nil {
			srv.logger.WithError(err).Errorf("failed to show camera from cored")
			return
		}

		var show_res camera_pb.ShowResponse
		client_helper.DecodeUnaryCallResponse(res, &show_res)

		state := show_res.Camera.State
		state_str := srv.camera_st_psr.ToString(state)

		c := storage.Camera{
			State: &state_str,
		}

		_, err = srv.storage.PatchCamera(cam_id, c)
		if err != nil {
			srv.logger.WithField("id", cam_id).WithError(err).Errorf("failed to patch camera")
			return
		}

		srv.logger.WithFields(log.Fields{
			"id":    cam_id,
			"state": state,
		}).Debugf("update camera state after created")
	}()

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

func (srv *metathingsCameradService) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	cam_id := req.GetId().GetValue()

	err = srv.storage.DeleteCamera(cam_id)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to delete camera")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	srv.logger.WithField("id", cam_id).Infof("delete camera")

	return &empty.Empty{}, nil
}

func (srv *metathingsCameradService) Patch(ctx context.Context, req *pb.PatchRequest) (*pb.PatchResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	updated := false
	c := storage.Camera{}
	cam_id := req.GetId().GetValue()

	name := req.GetName()
	if name != nil {
		c.Name = &name.Value
		updated = true
	}

	cfg := req.GetConfig()
	if cfg != nil {
		device := cfg.GetDevice()
		if device != nil {
			c.Device = &device.Value
			updated = true
		}

		width := cfg.GetWidth()
		height := cfg.GetHeight()
		if width != nil && height != nil {
			c.Width = &width.Value
			c.Height = &height.Value
			updated = true
		}

		bitrate := cfg.GetBitrate()
		if bitrate != nil {
			c.Bitrate = &bitrate.Value
			updated = true
		}

		framerate := cfg.GetFramerate()
		if framerate != nil {
			c.Framerate = &framerate.Value
			updated = true
		}
	}

	if !updated {
		return nil, status.Errorf(codes.InvalidArgument, "empty patch request")
	}

	pc, err := srv.storage.PatchCamera(cam_id, c)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to patch camera")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.PatchResponse{
		Camera: srv.copyCamera(pc),
	}

	srv.logger.WithField("cam_id", cam_id).Infof("patch camera")

	return res, nil
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

	core_id := req.GetCoreId()
	if core_id != nil {
		c.CoreId = &core_id.Value
	}

	entity_name := req.GetEntityName()
	if entity_name != nil {
		c.EntityName = &entity_name.Value
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

	core_id := req.GetCoreId()
	if core_id != nil {
		c.CoreId = &core_id.Value
	}

	entity_name := req.GetEntityName()
	if entity_name != nil {
		c.EntityName = &entity_name.Value
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

const (
	LIVE_ID_LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	LIVE_ID_LENGTH  = 128
)

func (srv *metathingsCameradService) newLiveId() string {
	buf := make([]byte, LIVE_ID_LENGTH)
	for i := 0; i < LIVE_ID_LENGTH; i++ {
		buf[i] = LIVE_ID_LETTERS[rand.Int31n(int32(len(LIVE_ID_LETTERS)))]
	}
	return string(buf)
}

func (srv *metathingsCameradService) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	live, err := url.Parse(srv.opts.rtmp_url)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to parse rtmp address")
		return nil, status.Errorf(codes.Internal, "bad rtmp address")
	}
	live_id := srv.newLiveId()
	live.Path = path.Join(live.Path, fmt.Sprintf("metathings.live.%v", live_id))
	live_str := live.String()
	cam_id := req.GetId().GetValue()

	c, err := srv.storage.GetCamera(cam_id)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to get camera")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if *c.State != "stop" {
		srv.logger.WithFields(log.Fields{
			"cam_id": cam_id,
			"state":  *c.State,
		}).Errorf("failed to start camera with unstartable state")
		return nil, status.Errorf(codes.OutOfRange, "unstartable state")
	}

	start_req := &camera_pb.StartRequest{
		Config: &camera_pb.StartConfig{
			Url: &gpb.StringValue{Value: live_str},
		},
	}

	if c.Device != nil {
		start_req.Config.Device = &gpb.StringValue{Value: *c.Device}
	}
	if c.Width != nil && c.Height != nil {
		start_req.Config.Width = &gpb.UInt32Value{Value: *c.Width}
		start_req.Config.Height = &gpb.UInt32Value{Value: *c.Height}
	}
	if c.Bitrate != nil {
		start_req.Config.Bitrate = &gpb.UInt32Value{Value: *c.Bitrate}
	}
	if c.Framerate != nil {
		start_req.Config.Framerate = &gpb.UInt32Value{Value: *c.Framerate}
	}

	call_req := client_helper.MustNewUnaryCallRequest(*c.CoreId, *c.EntityName, "camera", "Start", start_req)

	cli, cfn, err := srv.cli_fty.NewCoredServiceClient()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to new core service client")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer cfn()

	var start_res camera_pb.StartResponse
	call_res, err := cli.UnaryCall(srv.ContextWithToken(), call_req)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to call start on entity")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	err = client_helper.DecodeUnaryCallResponse(call_res, &start_res)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to decode response")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	state_str := "starting"
	cfg := start_res.Camera.Config
	c = storage.Camera{
		Url:       &cfg.Url,
		Device:    &cfg.Device,
		Width:     &cfg.Width,
		Height:    &cfg.Height,
		Bitrate:   &cfg.Bitrate,
		Framerate: &cfg.Framerate,
		State:     &state_str,
	}

	c, err = srv.storage.PatchCamera(cam_id, c)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to update camera")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.StartResponse{
		Camera: srv.copyCamera(c),
	}

	srv.logger.WithField("cam_id", cam_id).Infof("camera starting")

	return res, nil
}

func (srv *metathingsCameradService) Stop(ctx context.Context, req *pb.StopRequest) (*pb.StopResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	cam_id := req.GetId().GetValue()
	c, err := srv.storage.GetCamera(cam_id)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to get camera")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	call_req := client_helper.MustNewUnaryCallRequest(*c.CoreId, *c.EntityName, "camera", "Stop", &empty.Empty{})

	cli, cfn, err := srv.cli_fty.NewCoredServiceClient()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to new core service client")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer cfn()

	call_res, err := cli.UnaryCall(srv.ContextWithToken(), call_req)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to call stop on entity")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var stop_res camera_pb.StopResponse
	err = client_helper.DecodeUnaryCallResponse(call_res, &stop_res)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to decode response")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	state_str := "terminating"
	c = storage.Camera{
		State: &state_str,
	}

	c, err = srv.storage.PatchCamera(cam_id, c)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to update camera")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.StopResponse{
		Camera: srv.copyCamera(c),
	}

	srv.logger.WithField("cam_id", cam_id).Infof("camera terminating")

	return res, nil
}

func (srv *metathingsCameradService) Callback(ctx context.Context, req *pb.CallbackRequest) (*empty.Empty, error) {
	pc := storage.Camera{}
	updated := false

	// TODO(Peer): update config, blablabla
	state := req.GetState()
	if state != camera_pb.CameraState_CAMERA_STATE_UNKNOWN {
		state_str := srv.camera_st_psr.ToString(state)
		pc.State = &state_str
		updated = true
	}

	if !updated {
		srv.logger.Debugf("not change")
		return &empty.Empty{}, nil
	}

	cred := context_helper.Credential(ctx)
	app_cred_id := cred.ApplicationCredential.Id

	c := storage.Camera{
		ApplicationCredentialId: &app_cred_id,
	}

	cs, err := srv.storage.ListCameras(c)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to list cameras for callback")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if len(cs) == 0 {
		srv.logger.Debugf("unknown application credential id")
		return nil, status.Errorf(codes.NotFound, "unknown application credential id")
	}

	cam_id := *cs[0].Id

	_, err = srv.storage.PatchCamera(cam_id, pc)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to patch camera")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	srv.logger.WithFields(log.Fields{
		"cam_id":                    cam_id,
		"application_credential_id": app_cred_id,
	}).Debugf("camera callback")

	return &empty.Empty{}, nil
}

func (srv *metathingsCameradService) ShowToEntity(ctx context.Context, req *empty.Empty) (*pb.ShowToEntityResponse, error) {
	cred := context_helper.Credential(ctx)
	app_cred_id := cred.ApplicationCredential.Id

	c := storage.Camera{
		ApplicationCredentialId: &app_cred_id,
	}
	cs, err := srv.storage.ListCameras(c)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to list cameras for show to entity")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if len(cs) == 0 {
		srv.logger.WithField("application_credential_id", app_cred_id).Errorf("unknown application credential id")
		return nil, status.Errorf(codes.NotFound, "unknown application credential id")
	}

	res := &pb.ShowToEntityResponse{
		Camera: srv.copyCamera(cs[0]),
	}

	srv.logger.WithField("application_credential_id", app_cred_id).Debugf("show to entity")

	return res, nil
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

	cli_fty_cfgs := client_helper.NewDefaultServiceConfigs(opts.metathingsd_addr)
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
