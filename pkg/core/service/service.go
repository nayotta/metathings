package metathings_core_service

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"
	gpb "github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/bigdatagz/metathings/pkg/common"
	grpc_helper "github.com/bigdatagz/metathings/pkg/common/grpc"
	log_helper "github.com/bigdatagz/metathings/pkg/common/log"
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

type ApplicationCredentialManager struct {
	identityd_addr                string
	application_credential_id     string
	application_credential_secret string
	application_credential_token  string
}

func (mgr *ApplicationCredentialManager) GetToken() string {
	return "mt " + mgr.application_credential_token
}

func NewApplicationCredentialManager(identityd_addr, application_credential_id, application_credential_secret string) (*ApplicationCredentialManager, error) {
	log.WithFields(log.Fields{
		"application_credential_id":     application_credential_id,
		"application_credential_secret": application_credential_secret[0:12] + "...",
	}).Debugf("login via application credential")

	var header metadata.MD
	opts := []grpc.DialOption{grpc.WithInsecure()}
	ctx := context.Background()

	req := &identityd_pb.IssueTokenRequest{}
	req.Method = identityd_pb.AUTH_METHOD_APPLICATION_CREDENTIAL
	req.Payload = &identityd_pb.IssueTokenRequest_ApplicationCredential{
		&identityd_pb.ApplicationCredentialPayload{
			Id:     &gpb.StringValue{application_credential_id},
			Secret: &gpb.StringValue{application_credential_secret},
		},
	}

	conn, err := grpc.Dial(identityd_addr, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cli := identityd_pb.NewIdentityServiceClient(conn)

	_, err = cli.IssueToken(ctx, req, grpc.Header(&header))
	if err != nil {
		return nil, err
	}

	application_credential_token := header["authorization"][0]
	application_credential_token = application_credential_token[3:len(application_credential_token)]

	mgr := &ApplicationCredentialManager{
		identityd_addr,
		application_credential_id,
		application_credential_secret,
		application_credential_token,
	}

	return mgr, nil
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

	appCredMgr *ApplicationCredentialManager
	logger     log.FieldLogger
	opts       options
	storage    Storage
}

func (srv *metathingsCoreService) validateTokenViaIdentityd(token string) (*identityd_pb.Token, error) {
	ctx := context.Background()
	md := metadata.Pairs(
		"authorization-subject", "mt "+token,
		"authorization", srv.appCredMgr.GetToken(),
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

	core_id := common.NewId()

	var name_str string
	name := req.GetName()
	if name != nil {
		name_str = name.Value
	} else {
		name_str = core_id
	}

	c := Core{
		Id:        core_id,
		Name:      name_str,
		ProjectId: "",
		OwnerId:   "",
		State:     "offline",
	}

	cc, err := srv.storage.CreateCore(&c)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	srv.logger.WithFields(log.Fields{
		"id":         cc.Id,
		"name":       cc.Name,
		"project_id": cc.ProjectId,
		"owner_id":   cc.OwnerId,
		"state":      cc.State,
	}).Infof("create core")

	res := &pb.CreateCoreResponse{
		Core: &pb.Core{
			Id:        cc.Id,
			Name:      cc.Name,
			ProjectId: cc.ProjectId,
			OwnerId:   cc.OwnerId,
			State:     coreState2PbCoreState(cc.State),
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
func (srv *metathingsCoreService) Pipeline(pb.CoreService_PipelineServer) error {
	return grpc.Errorf(codes.Unimplemented, "unimplement")
}

func (srv *metathingsCoreService) ListCoresForUser(context.Context, *pb.ListCoresForUserRequest) (*pb.ListCoresForUserResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

func (srv *metathingsCoreService) SendUnaryCall(ctx context.Context, req *pb.SendUnaryCallRequest) (*pb.SendUnaryCallResponse, error) {
	srv.logger.WithFields(log.Fields{
		"service":      req.Payload.ServiceName,
		"method":       req.Payload.MethodName,
		"request-type": req.Payload.Payload.TypeUrl,
	}).Infof("receive unary call request")
	return nil, status.Errorf(codes.Unimplemented, "unimplement")
}

func NewCoreService(opt ...ServiceOptions) (*metathingsCoreService, error) {
	opts := defaultServiceOptions
	for _, o := range opt {
		o(&opts)
	}

	logger, err := log_helper.NewLogger("cored", opts.logLevel)
	if err != nil {
		log.Errorf("failed to new logger: %v", err)
		return nil, err
	}

	storage, err := NewStorage()
	if err != nil {
		log.Errorf("failed to connect storage: %v", err)
		return nil, err
	}

	appCredMgr, err := NewApplicationCredentialManager(
		opts.identityd_addr,
		opts.application_credential_id,
		opts.application_credential_secret,
	)
	if err != nil {
		log.Errorf("failed to new application credential manager")
		return nil, err
	}

	srv := &metathingsCoreService{
		appCredMgr: appCredMgr,
		opts:       opts,
		logger:     logger,
		storage:    storage,
	}
	return srv, nil
}
