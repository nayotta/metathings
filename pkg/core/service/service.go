package metathings_core_service

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/bigdatagz/metathings/pkg/common"
	log_helper "github.com/bigdatagz/metathings/pkg/common/log"
	pb "github.com/bigdatagz/metathings/pkg/proto/core"
)

type options struct {
	logLevel       string
	identityd_addr string
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

type metathingsCoreService struct {
	logger  log.FieldLogger
	opts    options
	storage Storage
}

func (srv *metathingsCoreService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "mt")
	if err != nil {
		return nil, err
	}
	srv.logger.WithFields(log.Fields{
		"method": fullMethodName,
		"token":  token,
	}).Debugf("authenticate...")

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

func NewCoreService(opt ...ServiceOptions) *metathingsCoreService {
	opts := defaultServiceOptions
	for _, o := range opt {
		o(&opts)
	}

	logger, err := log_helper.NewLogger("cored", opts.logLevel)
	if err != nil {
		log.Fatalf("failed to new logger: %v", err)
	}

	storage, err := NewStorage()
	if err != nil {
		log.Fatalf("failed to connect storage: %v", err)
	}

	srv := &metathingsCoreService{
		opts:    opts,
		logger:  logger,
		storage: storage,
	}
	return srv
}
