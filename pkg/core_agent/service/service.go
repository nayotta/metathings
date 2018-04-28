package meatathings_core_agent_service

import (
	"context"
	"io"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	gpb "github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	app_cred_mgr "github.com/bigdatagz/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/bigdatagz/metathings/pkg/common/client"
	context_helper "github.com/bigdatagz/metathings/pkg/common/context"
	log_helper "github.com/bigdatagz/metathings/pkg/common/log"
	state_helper "github.com/bigdatagz/metathings/pkg/common/state"
	state_pb "github.com/bigdatagz/metathings/pkg/proto/common/state"
	core_pb "github.com/bigdatagz/metathings/pkg/proto/core"
	cored_pb "github.com/bigdatagz/metathings/pkg/proto/core"
	pb "github.com/bigdatagz/metathings/pkg/proto/core_agent"
	echo_pb "github.com/bigdatagz/metathings/pkg/proto/echo"
)

type options struct {
	metathings_addr string
	logLevel        string

	core_agent_home               string
	core_id                       string
	application_credential_id     string
	application_credential_secret string
}

var defaultServiceOptions = options{
	logLevel: "info",
}

type ServiceOptions func(*options)

func SetMetathingsAddr(addr string) ServiceOptions {
	return func(o *options) {
		o.metathings_addr = addr
	}
}

func SetLogLevel(lvl string) ServiceOptions {
	return func(o *options) {
		o.logLevel = lvl
	}
}

func SetCoreAgentHome(path string) ServiceOptions {
	return func(o *options) {
		var usr *user.User
		var err error

		if usr, err = user.Current(); err != nil {
			return
		}
		if path[:2] == "~/" {
			path = filepath.Join(usr.HomeDir, path[2:])
		}
		o.core_agent_home = path
	}
}

func SetCoreId(id string) ServiceOptions {
	return func(o *options) {
		o.core_id = id
	}
}

func SetApplicationCredential(id, secret string) ServiceOptions {
	return func(o *options) {
		o.application_credential_id = id
		o.application_credential_secret = secret
	}
}

type coreAgentService struct {
	app_cred_mgr  app_cred_mgr.ApplicationCredentialManager
	cli_fty       *client_helper.ClientFactory
	entity_st_psr state_helper.EntityStateParser

	logger log.FieldLogger
	opts   options
}

func (srv *coreAgentService) copyEntity(e *core_pb.Entity) *pb.Entity {
	return &pb.Entity{
		Id:          e.Id,
		Name:        e.Name,
		ServiceName: e.ServiceName,
		Endpoint:    e.Endpoint,
		State:       e.State,
	}
}

func (srv *coreAgentService) CreateEntity(ctx context.Context, req *pb.CreateEntityRequest) (*pb.CreateEntityResponse, error) {
	ctx = context_helper.WithToken(ctx, srv.app_cred_mgr.GetToken())
	cli, closeFn, err := srv.cli_fty.NewCoreServiceClient()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to new core service client")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer closeFn()

	r := &core_pb.CreateEntityRequest{
		CoreId:      &gpb.StringValue{srv.opts.core_id},
		Name:        req.Name,
		ServiceName: req.ServiceName,
		Endpoint:    req.Endpoint,
	}

	res, err := cli.CreateEntity(ctx, r)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to create entity")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	srv.logger.WithFields(log.Fields{
		"id":           res.Entity.Id,
		"name":         res.Entity.Name,
		"service_name": res.Entity.ServiceName,
	}).Infof("create entity")

	return &pb.CreateEntityResponse{srv.copyEntity(res.Entity)}, nil
}

func (srv *coreAgentService) DeleteEntity(ctx context.Context, req *pb.DeleteEntityRequest) (*empty.Empty, error) {
	ctx = context_helper.WithToken(ctx, srv.app_cred_mgr.GetToken())
	cli, closeFn, err := srv.cli_fty.NewCoreServiceClient()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to new core service client")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer closeFn()

	r := &core_pb.DeleteEntityRequest{req.Id}
	_, err = cli.DeleteEntity(ctx, r)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to delete entity")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	srv.logger.WithField("id", req.Id).Infof("delete entity")

	return &empty.Empty{}, nil
}

func (srv *coreAgentService) PatchEntity(ctx context.Context, req *pb.PatchEntityRequest) (*pb.PatchEntityResponse, error) {
	ctx = context_helper.WithToken(ctx, srv.app_cred_mgr.GetToken())
	cli, closeFn, err := srv.cli_fty.NewCoreServiceClient()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to new core service client")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer closeFn()

	fields := log.Fields{"id": req.Id.Value}
	r := &core_pb.PatchEntityRequest{Id: req.Id}
	if req.Name != nil {
		r.Name = req.Name
		fields["name"] = req.Name.Value
	}
	if req.State != state_pb.EntityState_ENTITY_STATE_UNKNOWN {
		r.State = req.State
		fields["state"] = srv.entity_st_psr.ToString(req.State)
	}

	res, err := cli.PatchEntity(ctx, r)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to patch entity")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	srv.logger.WithFields(fields).Infof("patch entity")

	return &pb.PatchEntityResponse{srv.copyEntity(res.Entity)}, nil
}

func (srv *coreAgentService) GetEntity(ctx context.Context, req *pb.GetEntityRequest) (*pb.GetEntityResponse, error) {
	ctx = context_helper.WithToken(ctx, srv.app_cred_mgr.GetToken())
	cli, closeFn, err := srv.cli_fty.NewCoreServiceClient()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to new core service client")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer closeFn()

	r := &core_pb.GetEntityRequest{req.Id}

	res, err := cli.GetEntity(ctx, r)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to get entity")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	srv.logger.WithField("id", req.Id.Value).Debugf("get entity")

	return &pb.GetEntityResponse{srv.copyEntity(res.Entity)}, nil
}

func (srv *coreAgentService) ListEntities(ctx context.Context, req *pb.ListEntitiesRequest) (*pb.ListEntitiesResponse, error) {
	ctx = context_helper.WithToken(ctx, srv.app_cred_mgr.GetToken())
	cli, closeFn, err := srv.cli_fty.NewCoreServiceClient()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to new core service client")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer closeFn()

	r := &core_pb.ListEntitiesForCoreRequest{
		CoreId: &gpb.StringValue{srv.opts.core_id},
	}

	res, err := cli.ListEntitiesForCore(ctx, r)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to list entities for core")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	var entities []*pb.Entity
	for _, e := range res.Entities {
		entities = append(entities, srv.copyEntity(e))
	}
	srv.logger.Debugf("list entity")

	return &pb.ListEntitiesResponse{entities}, nil
}

func (srv *coreAgentService) CreateOrGetEntity(ctx context.Context, req *pb.CreateOrGetEntityRequest) (*pb.CreateOrGetEntityResponse, error) {
	ctx = context_helper.WithToken(ctx, srv.app_cred_mgr.GetToken())
	cli, closeFn, err := srv.cli_fty.NewCoreServiceClient()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to new core service client")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer closeFn()

	r := &core_pb.ListEntitiesForCoreRequest{
		CoreId: &gpb.StringValue{srv.opts.core_id},
	}

	res, err := cli.ListEntitiesForCore(ctx, r)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to list entities for core")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for _, e := range res.Entities {
		if e.Name == req.Name.Value {
			return &pb.CreateOrGetEntityResponse{srv.copyEntity(e)}, nil
		}
	}

	r1 := &core_pb.CreateEntityRequest{
		CoreId:      &gpb.StringValue{srv.opts.core_id},
		Name:        req.Name,
		ServiceName: req.ServiceName,
		Endpoint:    req.Endpoint,
	}

	res1, err := cli.CreateEntity(ctx, r1)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to create entity")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.CreateOrGetEntityResponse{srv.copyEntity(res1.Entity)}, nil
}

func (srv *coreAgentService) Heartbeat(ctx context.Context, req *pb.HeartbeatRequest) (*empty.Empty, error) {
	ctx = context_helper.WithToken(ctx, srv.app_cred_mgr.GetToken())
	cli, closeFn, err := srv.cli_fty.NewCoreServiceClient()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to new core service client")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer closeFn()

	r := &core_pb.GetEntityRequest{req.EntityId}
	res, err := cli.GetEntity(ctx, r)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to get entity")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if res.Entity.State == state_pb.EntityState_ENTITY_STATE_ONLINE {
		return &empty.Empty{}, nil
	}

	r1 := &core_pb.PatchEntityRequest{
		Id:    req.EntityId,
		State: state_pb.EntityState_ENTITY_STATE_ONLINE,
	}

	_, err = cli.PatchEntity(ctx, r1)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to patch entity")
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (srv *coreAgentService) ServeOnStream() error {
	token := srv.app_cred_mgr.GetToken()
	ctx := context_helper.WithToken(context.Background(), token)

	grpc_opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                2 * time.Second,
			Timeout:             20 * time.Second,
			PermitWithoutStream: true,
		}),
	}
	conn, err := grpc.Dial(srv.opts.metathings_addr, grpc_opts...)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to dial to metathings service")
		return err
	}
	defer conn.Close()

	cli := cored_pb.NewCoreServiceClient(conn)

	stream, err := cli.Stream(ctx)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to stream to core service")
		return err
	}

	return srv.serveOnStream(stream)
}

func (srv *coreAgentService) serveOnStream(stream core_pb.CoreService_StreamClient) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			srv.logger.Infof("core service closed")
			return nil
		}

		if err != nil {
			srv.logger.WithError(err).Errorf("failed to recv")
			return err
		}

		ctx := stream.Context()
		res, err := srv.dispatch(ctx, req)
		if err != nil {
			srv.logger.WithError(err).Errorf("failed to dispatch")
			continue
		}

		stream.Send(res)
	}

}

func (srv *coreAgentService) dispatch(ctx context.Context, req *core_pb.StreamRequest) (*core_pb.StreamResponse, error) {
	echo_res := echo_pb.EchoResponse{"hello, world"}
	any_res, err := ptypes.MarshalAny(&echo_res)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to marshal response to Any type")
		return nil, err
	}
	res := &core_pb.StreamResponse{
		SessionId:   req.SessionId.Value,
		MessageType: req.MessageType,
		Payload: &core_pb.StreamResponse_UnaryCall{
			UnaryCall: &core_pb.UnaryCallResponsePayload{
				ServiceName: "echo",
				MethodName:  "echo",
				Payload:     any_res,
			},
		},
	}
	return res, nil
}

func getCoreIdFromFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	data := make([]byte, 64)
	_, err = f.Read(data)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func getCoreIdFromService(opts options, token string) (string, error) {
	ctx := context.Background()
	md := metadata.Pairs("authorization", token)
	ctx = metadata.NewOutgoingContext(ctx, md)
	grpc_opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(opts.metathings_addr, grpc_opts...)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	cli := cored_pb.NewCoreServiceClient(conn)
	req := &cored_pb.CreateCoreRequest{}
	res, err := cli.CreateCore(ctx, req)
	if err != nil {
		return "", err
	}

	return res.Core.Id, nil
}

func saveCoreIdToPath(id, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(id))
	if err != nil {
		return err
	}

	return nil
}

func NewCoreAgentService(opt ...ServiceOptions) (srv *coreAgentService, err error) {

	opts := defaultServiceOptions
	for _, o := range opt {
		o(&opts)
	}

	logger, err := log_helper.NewLogger("core-agent", opts.logLevel)
	if err != nil {
		log.WithField("error", err).Errorf("failed to new logger")
		return nil, err
	}

	app_cred_mgr, err := app_cred_mgr.NewApplicationCredentialManager(
		opts.metathings_addr,
		opts.application_credential_id,
		opts.application_credential_secret,
	)
	if err != nil {
		log.WithField("error", err).Errorf("failed to new application credential manager")
		return nil, err
	}

	if opts.core_id == "" {
		var core_id string
		core_id_path := path.Join(opts.core_agent_home, "core-id")
		if core_id, err := getCoreIdFromFile(core_id_path); err != nil {
			if core_id, err = getCoreIdFromService(opts, app_cred_mgr.GetToken()); err != nil {
				return nil, err
			}
			if err = saveCoreIdToPath(core_id, core_id_path); err != nil {
				return nil, err
			}
		}
		opts.core_id = core_id
	}

	cli_fty, err := client_helper.NewClientFactory(
		client_helper.ServiceConfigs{
			client_helper.DEFAULT_CONFIG: client_helper.ServiceConfig{
				Address: opts.metathings_addr,
			},
		},
		func() []grpc.DialOption {
			return []grpc.DialOption{grpc.WithInsecure()}
		},
	)
	if err != nil {
		log.WithError(err).Errorf("failed to new client factory")
		return nil, err
	}

	srv = &coreAgentService{
		entity_st_psr: state_helper.NewEntityStateParser(),
		app_cred_mgr:  app_cred_mgr,
		cli_fty:       cli_fty,
		logger:        logger,
		opts:          opts,
	}
	return srv, nil
}
