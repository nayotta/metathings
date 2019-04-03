package metathings_camerad_service

import (
	"context"
	"fmt"
	"math/rand"
	"net/url"
	"path"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	state_helper "github.com/nayotta/metathings/pkg/camera/state"
	storage "github.com/nayotta/metathings/pkg/camerad/storage"
	app_cred_mgr "github.com/nayotta/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	camera_pb "github.com/nayotta/metathings/pkg/proto/camera"
	pb "github.com/nayotta/metathings/pkg/proto/camerad"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type MetathingsCameradServiceOption struct {
	RtmpUrl string
}

type MetathingsCameradService struct {
	cli_fty       *client_helper.ClientFactory
	camera_st_psr state_helper.CameraStateParser
	app_cred_mgr  app_cred_mgr.ApplicationCredentialManager
	logger        log.FieldLogger
	opt           *MetathingsCameradServiceOption
	storage       storage.Storage
	tk_vdr        token_helper.TokenValidator
}

func NewMetathingsCameradService(
	opt *MetathingsCameradServiceOption,
	cli_fty *client_helper.ClientFactory,
	camera_st_psr state_helper.CameraStateParser,
	app_cred_mgr app_cred_mgr.ApplicationCredentialManager,
	logger log.FieldLogger,
	storage storage.Storage,
	tk_vdr token_helper.TokenValidator,
) (pb.CameradServiceServer, error) {
	return &MetathingsCameradService{
		opt:           opt,
		cli_fty:       cli_fty,
		camera_st_psr: camera_st_psr,
		app_cred_mgr:  app_cred_mgr,
		logger:        logger,
		storage:       storage,
		tk_vdr:        tk_vdr,
	}, nil
}

func (srv *MetathingsCameradService) ContextWithToken(ctxs ...context.Context) context.Context {
	ctx := context.Background()
	if len(ctxs) > 0 {
		ctx = ctxs[0]
	}
	token_str := srv.app_cred_mgr.GetToken()
	ctx = context_helper.WithToken(ctx, token_str)
	return ctx
}

func (srv *MetathingsCameradService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	token_str, err := grpc_helper.GetTokenFromContext(ctx)
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
		"method":    fullMethodName,
		"entity_id": token.Entity.Id,
	}).Debugf("validator token")

	return ctx, nil
}

func (srv *MetathingsCameradService) copyCamera(c storage.Camera) *pb.Camera {
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

func (srv *MetathingsCameradService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	cred := context_helper.Credential(ctx)
	cam_id := id_helper.NewId()
	var name_str string
	if name := req.GetName(); name != nil {
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
		// cli, cfn, err := srv.cli_fty.NewCoredServiceClient()
		// if err != nil {
		// 	srv.logger.WithError(err).Errorf("failed to create cored service client")
		// 	return
		// }
		// defer cfn()

		// req := client_helper.MustNewUnaryCallRequest(
		// 	core_id, entity_name, "camera", "Show", &empty.Empty{},
		// )

		// res, err := cli.UnaryCall(srv.ContextWithToken(), req)
		// if err != nil {
		// 	srv.logger.WithError(err).Errorf("failed to show camera from cored")
		// 	return
		// }

		// var show_res camera_pb.ShowResponse
		// client_helper.DecodeUnaryCallResponse(res, &show_res)

		// state := show_res.Camera.State
		// state_str := srv.camera_st_psr.ToString(state)

		// _, err = srv.set_camera_state(cam_id, state_str)
		// if err != nil {
		// 	srv.logger.WithField("id", cam_id).WithError(err).Errorf("failed to patch camera")
		// 	return
		// }

		// srv.logger.WithFields(log.Fields{
		// 	"id":    cam_id,
		// 	"state": state,
		// }).Debugf("update camera state after created")
		panic("unimplemented")
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

func (srv *MetathingsCameradService) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
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

func (srv *MetathingsCameradService) Patch(ctx context.Context, req *pb.PatchRequest) (*pb.PatchResponse, error) {
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

func (srv *MetathingsCameradService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
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

func (srv *MetathingsCameradService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
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

	owner_id := req.GetOwnerId()
	if owner_id != nil {
		c.OwnerId = &owner_id.Value
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

func (srv *MetathingsCameradService) ListForUser(ctx context.Context, req *pb.ListForUserRequest) (*pb.ListForUserResponse, error) {
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
		return nil, status.Errorf(codes.Internal, err.Error())
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

func (srv *MetathingsCameradService) newLiveId() string {
	buf := make([]byte, LIVE_ID_LENGTH)
	for i := 0; i < LIVE_ID_LENGTH; i++ {
		buf[i] = LIVE_ID_LETTERS[rand.Int31n(int32(len(LIVE_ID_LETTERS)))]
	}
	return string(buf)
}

func (srv *MetathingsCameradService) set_camera_state(cam_id string, state string) (storage.Camera, error) {
	c, err := srv.storage.PatchCamera(cam_id, storage.Camera{State: &state})
	if err != nil {
		return storage.Camera{}, err
	}

	srv.logger.WithFields(log.Fields{
		"state": state,
	}).Debugf("update camera state")
	return c, nil
}

func (srv *MetathingsCameradService) sync_camera_state(cli deviced_pb.DevicedServiceClient, cam_id string, core_id string, entity_name string) (camera_pb.CameraState, error) {
	// call_req := client_helper.MustNewUnaryCallRequest(core_id, entity_name, "camera", "Show", &empty.Empty{})
	// call_res, err := cli.UnaryCall(srv.ContextWithToken(), call_req)
	// if err != nil {
	// 	return camera_pb.CameraState_CAMERA_STATE_UNKNOWN, err
	// }

	// var show_res camera_pb.ShowResponse
	// err = client_helper.DecodeUnaryCallResponse(call_res, &show_res)
	// if err != nil {
	// 	return camera_pb.CameraState_CAMERA_STATE_UNKNOWN, err
	// }

	// cam_st := show_res.Camera.State
	// srv.set_camera_state(cam_id, srv.camera_st_psr.ToString(cam_st))
	// srv.logger.WithFields(log.Fields{
	// 	"cam_id": cam_id,
	// 	"state":  cam_st,
	// }).Debugf("sync camera state")

	// return cam_st, nil
	panic("unimplemented")
}

func (srv *MetathingsCameradService) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	live, err := url.Parse(srv.opt.RtmpUrl)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to parse rtmp address")
		return nil, status.Errorf(codes.Internal, "bad rtmp address")
	}
	live_id := srv.newLiveId()
	live.Path = path.Join(live.Path, fmt.Sprintf("metathings.live.%v", live_id))
	// live_str := live.String()
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

	srv.set_camera_state(cam_id, "starting")
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to update camera")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	go func() {
		// var err error
		// defer func() {
		// 	if err != nil {
		// 		srv.set_camera_state(cam_id, "stop")
		// 	}
		// }()

		// cli, cfn, err := srv.cli_fty.NewCoredServiceClient()
		// if err != nil {
		// 	srv.logger.WithError(err).Errorf("failed to new core service client")
		// 	return
		// }
		// defer cfn()

		// st, err := srv.sync_camera_state(cli, cam_id, *c.CoreId, *c.EntityName)
		// if err != nil {
		// 	srv.logger.WithField("cam_id", cam_id).WithError(err).Errorf("failed to sync camera state from agent")
		// 	return
		// }

		// if st == camera_pb.CameraState_CAMERA_STATE_RUNNING {
		// 	srv.logger.WithField("cam_id", cam_id).Debugf("camera already running")
		// 	return
		// }

		// start_req := &camera_pb.StartRequest{
		// 	Config: &camera_pb.StartConfig{
		// 		Url: &gpb.StringValue{Value: live_str},
		// 	},
		// }

		// if c.Device != nil {
		// 	start_req.Config.Device = &gpb.StringValue{Value: *c.Device}
		// }
		// if c.Width != nil && c.Height != nil {
		// 	start_req.Config.Width = &gpb.UInt32Value{Value: *c.Width}
		// 	start_req.Config.Height = &gpb.UInt32Value{Value: *c.Height}
		// }
		// if c.Bitrate != nil {
		// 	start_req.Config.Bitrate = &gpb.UInt32Value{Value: *c.Bitrate}
		// }
		// if c.Framerate != nil {
		// 	start_req.Config.Framerate = &gpb.UInt32Value{Value: *c.Framerate}
		// }

		// call_req := client_helper.MustNewUnaryCallRequest(*c.CoreId, *c.EntityName, "camera", "Start", start_req)

		// call_res, err := cli.UnaryCall(srv.ContextWithToken(), call_req)
		// if err != nil {
		// 	srv.logger.WithError(err).Errorf("failed to call start on entity")
		// 	return
		// }

		// var start_res camera_pb.StartResponse
		// err = client_helper.DecodeUnaryCallResponse(call_res, &start_res)
		// if err != nil {
		// 	srv.logger.WithError(err).Errorf("failed to decode start camera response")
		// 	return
		// }

		// cfg := start_res.Camera.Config
		// c = storage.Camera{
		// 	Url:       &cfg.Url,
		// 	Device:    &cfg.Device,
		// 	Width:     &cfg.Width,
		// 	Height:    &cfg.Height,
		// 	Bitrate:   &cfg.Bitrate,
		// 	Framerate: &cfg.Framerate,
		// }

		// c, err = srv.storage.PatchCamera(cam_id, c)
		// if err != nil {
		// 	srv.logger.WithError(err).Errorf("failed to update camera")
		// 	return
		// }

		// srv.logger.WithField("cam_id", cam_id).Debugf("send start command to camera")
		panic("unimplemented")
	}()

	res := &pb.StartResponse{
		Camera: srv.copyCamera(c),
	}

	srv.logger.WithField("cam_id", cam_id).Infof("camera starting")

	return res, nil
}

func (srv *MetathingsCameradService) Stop(ctx context.Context, req *pb.StopRequest) (*pb.StopResponse, error) {
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

	if *c.State != "running" {
		srv.logger.WithFields(log.Fields{
			"cam_id": cam_id,
			"state":  *c.State,
		}).Errorf("failed to stop camera with unstopable state")
		return nil, status.Errorf(codes.OutOfRange, "unstopable state")
	}

	srv.set_camera_state(cam_id, "terminating")
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to update camera")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	go func() {
		// var err error
		// defer func() {
		// 	if err != nil {
		// 		srv.set_camera_state(cam_id, "running")
		// 	}
		// }()

		// cli, cfn, err := srv.cli_fty.NewCoredServiceClient()
		// if err != nil {
		// 	srv.logger.WithError(err).Errorf("failed to new core service client")
		// 	return
		// }
		// defer cfn()

		// st, err := srv.sync_camera_state(cli, cam_id, *c.CoreId, *c.EntityName)
		// if err != nil {
		// 	srv.logger.WithField("cam_id", cam_id).WithError(err).Errorf("failed to sync camera state from agent")
		// 	return
		// }

		// if st == camera_pb.CameraState_CAMERA_STATE_STOP {
		// 	srv.logger.WithField("cam_id", cam_id).Debugf("camera already stop")
		// 	return
		// }

		// call_req := client_helper.MustNewUnaryCallRequest(*c.CoreId, *c.EntityName, "camera", "Stop", &empty.Empty{})

		// call_res, err := cli.UnaryCall(srv.ContextWithToken(), call_req)
		// if err != nil {
		// 	srv.logger.WithError(err).Errorf("failed to call stop on entity")
		// 	return
		// }

		// var stop_res camera_pb.StopResponse
		// err = client_helper.DecodeUnaryCallResponse(call_res, &stop_res)
		// if err != nil {
		// 	srv.logger.WithError(err).Errorf("failed to decode response")
		// 	return
		// }

		// srv.logger.WithField("cam_id", cam_id).Debugf("send stop command to camera")
		panic("unimplemented")
	}()

	res := &pb.StopResponse{
		Camera: srv.copyCamera(c),
	}

	srv.logger.WithField("cam_id", cam_id).Infof("camera terminating")

	return res, nil
}

func (srv *MetathingsCameradService) Callback(ctx context.Context, req *pb.CallbackRequest) (*empty.Empty, error) {
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

func (srv *MetathingsCameradService) ShowToEntity(ctx context.Context, req *empty.Empty) (*pb.ShowToEntityResponse, error) {
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
