package metathings_core_service

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/bigdatagz/metathings/pkg/common"
	app_cred_mgr "github.com/bigdatagz/metathings/pkg/common/application_credential_manager"
	context_helper "github.com/bigdatagz/metathings/pkg/common/context"
	grpc_helper "github.com/bigdatagz/metathings/pkg/common/grpc"
	log_helper "github.com/bigdatagz/metathings/pkg/common/log"
	stm_mgr "github.com/bigdatagz/metathings/pkg/common/stream_manager"
	storage "github.com/bigdatagz/metathings/pkg/core/storage"
	pb "github.com/bigdatagz/metathings/pkg/proto/core"
	identityd_pb "github.com/bigdatagz/metathings/pkg/proto/identity"
)

type options struct {
	logLevel                      string
	identityd_addr                string
	application_credential_id     string
	application_credential_secret string
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

func SetApplicationCredential(id, secret string) ServiceOptions {
	return func(o *options) {
		o.application_credential_id = id
		o.application_credential_secret = secret
	}
}

type metathingsCoreService struct {
	grpc_helper.AuthorizationTokenParser

	app_cred_mgr app_cred_mgr.ApplicationCredentialManager
	stm_mgr      stm_mgr.StreamManager
	logger       log.FieldLogger
	opts         options
	storage      storage.Storage
}

func (srv *metathingsCoreService) validateTokenViaIdentityd(token string) (*identityd_pb.Token, error) {
	ctx := context.Background()
	md := metadata.Pairs(
		"authorization-subject", "mt "+token,
		"authorization", srv.app_cred_mgr.GetToken(),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(srv.opts.identityd_addr, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cli := identityd_pb.NewIdentityServiceClient(conn)
	res, err := cli.ValidateToken(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}

	return res.Token, nil
}

func (srv *metathingsCoreService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	token_str, err := srv.GetTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}

	token, err := srv.validateTokenViaIdentityd(token_str)
	if err != nil {
		srv.logger.
			WithField("error", err).
			Errorf("failed to validate token via metathings identity service")
		return nil, err
	}

	ctx = context.WithValue(ctx, "token", token_str)
	ctx = context.WithValue(ctx, "credential", token)

	srv.logger.WithFields(log.Fields{
		"method":   fullMethodName,
		"user_id":  token.User.Id,
		"username": token.User.Name,
	}).Debugf("validate token via metathings identity service")

	return ctx, nil
}

func pbCoreState2CoreState(s pb.CoreState) string {
	if t, ok := map[pb.CoreState]string{
		pb.CoreState_CORE_STATE_UNKNOWN: "unknown",
		pb.CoreState_CORE_STATE_ONLINE:  "online",
		pb.CoreState_CORE_STATE_OFFLINE: "offline",
	}[s]; ok {
		return t
	}
	return "unknown"
}

func coreState2PbCoreState(s string) pb.CoreState {
	if t, ok := map[string]pb.CoreState{
		"unknown": pb.CoreState_CORE_STATE_UNKNOWN,
		"online":  pb.CoreState_CORE_STATE_ONLINE,
		"offline": pb.CoreState_CORE_STATE_OFFLINE,
	}[s]; ok {
		return t
	}
	return pb.CoreState_CORE_STATE_UNKNOWN
}

func (srv *metathingsCoreService) CreateCore(ctx context.Context, req *pb.CreateCoreRequest) (*pb.CreateCoreResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	cred := context_helper.Credential(ctx)

	core_id := common.NewId()

	var name_str string
	name := req.GetName()
	if name != nil {
		name_str = name.Value
	} else {
		name_str = core_id
	}
	state := "offline"

	c := storage.Core{
		Id:        &core_id,
		Name:      &name_str,
		ProjectId: &cred.Project.Id,
		OwnerId:   &cred.User.Id,
		State:     &state,
	}

	cc, err := srv.storage.CreateCore(c)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	srv.logger.WithFields(log.Fields{
		"id":         *cc.Id,
		"name":       *cc.Name,
		"project_id": *cc.ProjectId,
		"owner_id":   *cc.OwnerId,
		"state":      *cc.State,
	}).Infof("create core")

	err = srv.storage.AssignCoreToApplicationCredential(*cc.Id, cred.ApplicationCredential.Id)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to assign core to application credential")
		srv.storage.DeleteCore(*cc.Id)
		return nil, err
	}
	srv.logger.WithFields(log.Fields{
		"core_id":     *cc.Id,
		"app_cred_id": cred.ApplicationCredential.Id,
	}).Infof("assign core to application credential")

	res := &pb.CreateCoreResponse{
		Core: &pb.Core{
			Id:        *cc.Id,
			Name:      *cc.Name,
			ProjectId: *cc.ProjectId,
			OwnerId:   *cc.OwnerId,
			State:     coreState2PbCoreState(*cc.State),
		},
	}

	return res, nil
}

func (srv *metathingsCoreService) DeleteCore(context.Context, *pb.DeleteCoreRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

func (srv *metathingsCoreService) PatchCore(context.Context, *pb.PatchCoreRequest) (*pb.PatchCoreResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

func (srv *metathingsCoreService) GetCore(context.Context, *pb.GetCoreRequest) (*pb.GetCoreResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

func (srv *metathingsCoreService) ListCores(context.Context, *pb.ListCoresRequest) (*pb.ListCoresResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

func (srv *metathingsCoreService) Heartbeat(context.Context, *pb.HeartbeatRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

func (srv *metathingsCoreService) Stream(stream pb.CoreService_StreamServer) error {
	ctx := stream.Context()
	cred := context_helper.Credential(ctx)
	if cred == nil || cred.ApplicationCredential == nil {
		srv.logger.Errorf("token dont created by application credential")
		return status.Errorf(codes.Internal, "token dont created by application credential")
	}

	core, err := srv.storage.GetAssignedCoreFromApplicationCredential(cred.ApplicationCredential.Id)
	if err != nil {
		srv.logger.WithError(err).Errorf("not core assigend to application credential, should not be here, may be hacked.")
		return status.Errorf(codes.Internal, "not core assigned to application credential")
	}

	close_chan, err := srv.stm_mgr.Register(*core.Id, stream)
	if err != nil {
		srv.logger.
			WithField("core_id", core.Id).
			WithError(err).
			Errorf("failed to register stream")
	}
	srv.logger.WithField("core_id", core.Id).Infof("register stream")

	<-close_chan
	srv.logger.WithField("core_id", core.Id).Infof("stream closed")

	return nil
}

func (srv *metathingsCoreService) ListCoresForUser(context.Context, *pb.ListCoresForUserRequest) (*pb.ListCoresForUserResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

func (srv *metathingsCoreService) UnaryCall(ctx context.Context, req *pb.UnaryCallRequest) (*pb.UnaryCallResponse, error) {
	res, err := srv.stm_mgr.UnaryCall(req.CoreId.Value, req.Payload)
	if err != nil {
		srv.logger.
			WithField("core_id", req.CoreId.Value).
			WithError(err).
			Errorf("failed to unary call")
		return nil, err
	}
	return &pb.UnaryCallResponse{res}, nil
}

func NewCoreService(opt ...ServiceOptions) (*metathingsCoreService, error) {
	opts := defaultServiceOptions
	for _, o := range opt {
		o(&opts)
	}

	logger, err := log_helper.NewLogger("cored", opts.logLevel)
	if err != nil {
		log.WithError(err).Errorf("failed to new logger")
		return nil, err
	}

	storage, err := storage.NewStorage(":memory:", logger)
	if err != nil {
		log.WithError(err).Errorf("failed to connect storage")
		return nil, err
	}

	app_cred_mgr, err := app_cred_mgr.NewApplicationCredentialManager(
		opts.identityd_addr,
		opts.application_credential_id,
		opts.application_credential_secret,
	)
	if err != nil {
		log.WithError(err).Errorf("failed to new application credential manager")
		return nil, err
	}

	stm_mgr, err := stm_mgr.NewStreamManager(logger)
	if err != nil {
		log.WithError(err).Errorf("failed to new stream manager")
		return nil, err
	}

	srv := &metathingsCoreService{
		app_cred_mgr: app_cred_mgr,
		stm_mgr:      stm_mgr,
		opts:         opts,
		logger:       logger,
		storage:      storage,
	}
	return srv, nil
}
