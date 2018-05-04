package meatathings_core_agent_service

import (
	"context"
	"io"
	"os"
	"path"
	"sync"
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

	helper "github.com/bigdatagz/metathings/pkg/common"
	app_cred_mgr "github.com/bigdatagz/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/bigdatagz/metathings/pkg/common/client"
	context_helper "github.com/bigdatagz/metathings/pkg/common/context"
	log_helper "github.com/bigdatagz/metathings/pkg/common/log"
	state_helper "github.com/bigdatagz/metathings/pkg/common/state"
	mt_plugin "github.com/bigdatagz/metathings/pkg/core/plugin"
	state_pb "github.com/bigdatagz/metathings/pkg/proto/common/state"
	core_pb "github.com/bigdatagz/metathings/pkg/proto/core"
	cored_pb "github.com/bigdatagz/metathings/pkg/proto/core"
	pb "github.com/bigdatagz/metathings/pkg/proto/core_agent"
)

type options struct {
	metathings_addr string
	logLevel        string

	core_agent_home               string
	core_id                       string
	application_credential_id     string
	application_credential_secret string
	service_descriptor_path       string
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
		o.core_agent_home = helper.ExpendHomePath(path)
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

func SetServiceDescriptorPath(path string) ServiceOptions {
	return func(o *options) {
		path = helper.ExpendHomePath(path)
		o.service_descriptor_path = path
	}
}

type coreAgentService struct {
	app_cred_mgr  app_cred_mgr.ApplicationCredentialManager
	cli_fty       *client_helper.ClientFactory
	entity_st_psr state_helper.EntityStateParser
	serv_desc     *mt_plugin.ServiceDescriptor

	mtx_dp_op   *sync.Mutex
	dispatchers map[string]mt_plugin.DispatcherPlugin

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

func (srv *coreAgentService) loadDispatcherPlugin(e *pb.Entity) {
	srv.mtx_dp_op.Lock()
	defer srv.mtx_dp_op.Unlock()

	_, ok := srv.dispatchers[e.Name]
	if ok {
		return
	}

	dp, err := srv.serv_desc.GetDispatcherPlugin(e.ServiceName)
	if err != nil {
		srv.logger.WithError(err).
			WithFields(log.Fields{
				"name":         e.Name,
				"service_name": e.ServiceName,
				"id":           e.Id,
			}).
			Errorf("failed to load dispatcher plugin")
		return
	}

	err = dp.Init(mt_plugin.PluginOptions{
		"endpoint": e.Endpoint,
	})
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to init plugin")
		return
	}

	srv.dispatchers[e.Name] = dp
	srv.logger.WithFields(log.Fields{
		"name":         e.Name,
		"service_name": e.ServiceName,
		"id":           e.Id,
	}).Debugf("load dispatcher plugin")
}

func (srv *coreAgentService) getDispatcherPlugin(name string, service_name string) (mt_plugin.DispatcherPlugin, bool) {
	dp, ok := srv.dispatchers[name]
	return dp, ok
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
			e1 := srv.copyEntity(e)
			defer srv.loadDispatcherPlugin(e1)
			return &pb.CreateOrGetEntityResponse{e1}, nil
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

	e1 := srv.copyEntity(res1.Entity)
	defer srv.loadDispatcherPlugin(e1)
	return &pb.CreateOrGetEntityResponse{e1}, nil
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
	if req.MessageType != core_pb.StreamMessageType_STREAM_MESSAGE_TYPE_USER {
		return nil, ErrUnsupportMessageType
	}

	payload, ok := req.Payload.(*core_pb.StreamRequest_UnaryCall)
	if !ok {
		return nil, ErrUnsupportPayloadType
	}

	name := payload.UnaryCall.Name.Value
	service_name := payload.UnaryCall.ServiceName.Value
	method_name := payload.UnaryCall.MethodName.Value
	req_payload := payload.UnaryCall.Payload

	dp, ok := srv.getDispatcherPlugin(name, service_name)
	if !ok {
		return nil, ErrPluginNotFound
	}

	res, err := dp.UnaryCall(method_name, ctx, req_payload)
	if err != nil {
		return nil, err
	}

	res_payload, err := ptypes.MarshalAny(res)
	if err != nil {
		return nil, err
	}

	res1 := &core_pb.StreamResponse{
		SessionId:   req.SessionId.Value,
		MessageType: req.MessageType,
		Payload: &core_pb.StreamResponse_UnaryCall{
			&core_pb.UnaryCallResponsePayload{
				Name:        name,
				ServiceName: service_name,
				MethodName:  method_name,
				Payload:     res_payload,
			},
		},
	}

	return res1, nil
}

func getCoreIdFromFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	data := make([]byte, 64)
	n, err := f.Read(data)
	if err != nil {
		return "", err
	}

	return string(data[:n]), nil
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

	cli_fty, err := client_helper.NewClientFactory(
		client_helper.NewDefaultServiceConfigs(opts.metathings_addr),
		client_helper.WithInsecureOptionFunc(),
	)
	if err != nil {
		log.WithError(err).Errorf("failed to new client factory")
		return nil, err
	}

	app_cred_mgr, err := app_cred_mgr.NewApplicationCredentialManager(
		cli_fty,
		opts.application_credential_id,
		opts.application_credential_secret,
	)
	if err != nil {
		log.WithField("error", err).Errorf("failed to new application credential manager")
		return nil, err
	}

	if opts.core_id == "" {
		core_id_path := path.Join(opts.core_agent_home, "core-id")
		core_id, err := getCoreIdFromFile(core_id_path)
		if err != nil {
			if core_id, err = getCoreIdFromService(opts, app_cred_mgr.GetToken()); err != nil {
				return nil, err
			}
			if err = saveCoreIdToPath(core_id, core_id_path); err != nil {
				return nil, err
			}
		}
		opts.core_id = core_id
	}

	serv_desc, err := mt_plugin.LoadServiceDescriptor(opts.service_descriptor_path)
	if err != nil {
		log.WithError(err).Errorf("failed to load service descriptor")
		return nil, err
	}

	srv = &coreAgentService{
		entity_st_psr: state_helper.NewEntityStateParser(),
		app_cred_mgr:  app_cred_mgr,
		serv_desc:     serv_desc,
		cli_fty:       cli_fty,
		logger:        logger,
		opts:          opts,

		mtx_dp_op:   new(sync.Mutex),
		dispatchers: make(map[string]mt_plugin.DispatcherPlugin),
	}
	return srv, nil
}
