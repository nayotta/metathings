package metathings_core_agent_service

import (
	"context"
	"os"
	"path"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	gpb "github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"

	helper "github.com/nayotta/metathings/pkg/common"
	app_cred_mgr "github.com/nayotta/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	protobuf_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	state_helper "github.com/nayotta/metathings/pkg/common/state"
	mt_plugin "github.com/nayotta/metathings/pkg/core/plugin"
	state_pb "github.com/nayotta/metathings/pkg/proto/common/state"
	core_pb "github.com/nayotta/metathings/pkg/proto/core"
	cored_pb "github.com/nayotta/metathings/pkg/proto/core"
	pb "github.com/nayotta/metathings/pkg/proto/core_agent"
)

type options struct {
	metathings_addr string
	logLevel        string

	core_agent_home               string
	core_id                       string
	application_credential_id     string
	application_credential_secret string
	service_descriptor            string

	heartbeat_interval int
}

var (
	GRPC_KEEPALIVE = grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                2 * time.Second,
		Timeout:             10 * time.Second,
		PermitWithoutStream: true,
	})
)

var defaultServiceOptions = options{
	logLevel:           "info",
	heartbeat_interval: 5,
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

func SetServiceDescriptor(path string) ServiceOptions {
	return func(o *options) {
		path = helper.ExpendHomePath(path)
		o.service_descriptor = path
	}
}

func SetHeartbeatInterval(interval int) ServiceOptions {
	return func(o *options) {
		o.heartbeat_interval = interval
	}
}

type coreAgentService struct {
	app_cred_mgr  app_cred_mgr.ApplicationCredentialManager
	cli_fty       *client_helper.ClientFactory
	entity_st_psr state_helper.EntityStateParser
	serv_desc     *mt_plugin.ServiceDescriptor

	mtx_dp_op   *sync.Mutex
	dispatchers map[string]mt_plugin.DispatcherPlugin

	heartbeat_entities map[string]time.Time

	logger log.FieldLogger
	opts   options
}

func (srv *coreAgentService) HeartbeatLoop() error {
	interval := time.Duration(srv.opts.heartbeat_interval) * time.Second
	for {
		go func() {
			err := srv.HeartbeatOnce()
			if err != nil {
				srv.logger.WithError(err).Errorf("failed to heartbeat")
			}
		}()
		<-time.After(interval)
	}
}

func (srv *coreAgentService) HeartbeatOnce() error {
	ctx := context_helper.WithToken(context.Background(), srv.app_cred_mgr.GetToken())
	cli, cfn, err := srv.cli_fty.NewCoreServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	entities := []*core_pb.HeartbeatEntity{}
	for id, _ := range srv.heartbeat_entities {
		if srv.heartbeat_entities[id].Equal(UNAVAILABLE_TIME) {
			continue
		}
		ts := protobuf_helper.FromTime(srv.heartbeat_entities[id])
		entities = append(entities, &core_pb.HeartbeatEntity{
			Id:          &gpb.StringValue{Value: id},
			HeartbeatAt: &ts,
		})
		srv.heartbeat_entities[id] = UNAVAILABLE_TIME
	}

	req := &core_pb.HeartbeatRequest{
		Entities: entities,
	}
	_, err = cli.Heartbeat(ctx, req)
	if err != nil {
		return err
	}

	return nil
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
		CoreId:      &gpb.StringValue{Value: srv.opts.core_id},
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

	return &pb.CreateEntityResponse{Entity: srv.copyEntity(res.Entity)}, nil
}

func (srv *coreAgentService) DeleteEntity(ctx context.Context, req *pb.DeleteEntityRequest) (*empty.Empty, error) {
	ctx = context_helper.WithToken(ctx, srv.app_cred_mgr.GetToken())
	cli, closeFn, err := srv.cli_fty.NewCoreServiceClient()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to new core service client")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer closeFn()

	r := &core_pb.DeleteEntityRequest{Id: req.Id}
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

	return &pb.PatchEntityResponse{Entity: srv.copyEntity(res.Entity)}, nil
}

func (srv *coreAgentService) GetEntity(ctx context.Context, req *pb.GetEntityRequest) (*pb.GetEntityResponse, error) {
	ctx = context_helper.WithToken(ctx, srv.app_cred_mgr.GetToken())
	cli, closeFn, err := srv.cli_fty.NewCoreServiceClient()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to new core service client")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer closeFn()

	r := &core_pb.GetEntityRequest{Id: req.Id}

	res, err := cli.GetEntity(ctx, r)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to get entity")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	srv.logger.WithField("id", req.Id.Value).Debugf("get entity")

	return &pb.GetEntityResponse{Entity: srv.copyEntity(res.Entity)}, nil
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
		Name:        req.Name,
		ServiceName: req.ServiceName,
		State:       req.State,
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

	return &pb.ListEntitiesResponse{Entities: entities}, nil
}

func (srv *coreAgentService) LoadDispatcherPlugin(e *pb.Entity) {
	srv.mtx_dp_op.Lock()
	defer srv.mtx_dp_op.Unlock()

	if srv.isLoadedDispatcherPlugin(e) {
		return
	}

	srv.loadDispatcherPlugin(e)
}

func (srv *coreAgentService) loadDispatcherPlugin(e *pb.Entity) {
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

	err = dp.Init(opt_helper.Option{
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

func (srv *coreAgentService) IsLoadedDispatcherPlugin(e *pb.Entity) bool {
	srv.mtx_dp_op.Lock()
	defer srv.mtx_dp_op.Unlock()

	return srv.isLoadedDispatcherPlugin(e)
}

func (srv *coreAgentService) isLoadedDispatcherPlugin(e *pb.Entity) bool {
	if id := e.GetId(); id != "" {
		_, ok := srv.dispatchers[id]
		if ok {
			return true
		}
	}

	if name := e.GetName(); name != "" {
		_, ok := srv.dispatchers[name]
		if ok {
			return true
		}
	}

	return false
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

	r := &core_pb.ListEntitiesForCoreRequest{}
	res, err := cli.ListEntitiesForCore(ctx, r)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to list entities for core")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for _, e := range res.Entities {
		if e.Name == req.Name.Value {
			e1 := srv.copyEntity(e)
			defer srv.LoadDispatcherPlugin(e1)
			return &pb.CreateOrGetEntityResponse{Entity: e1}, nil
		}
	}

	r1 := &core_pb.CreateEntityRequest{
		CoreId:      &gpb.StringValue{Value: srv.opts.core_id},
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
	return &pb.CreateOrGetEntityResponse{Entity: e1}, nil
}

func (srv *coreAgentService) Heartbeat(ctx context.Context, req *pb.HeartbeatRequest) (*empty.Empty, error) {
	entity_id := req.GetEntityId().GetValue()
	if _, ok := srv.heartbeat_entities[entity_id]; !ok {
		ctx = context_helper.WithToken(ctx, srv.app_cred_mgr.GetToken())
		cli, closeFn, err := srv.cli_fty.NewCoreServiceClient()
		if err != nil {
			srv.logger.WithError(err).Errorf("failed to new core service client")
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		defer closeFn()

		r := &core_pb.GetEntityRequest{Id: req.EntityId}
		res, err := cli.GetEntity(ctx, r)
		if err != nil {
			srv.logger.WithError(err).Errorf("failed to get entity")
			return nil, status.Errorf(codes.Internal, err.Error())
		}

		e := &pb.Entity{
			Id:          res.Entity.Id,
			Name:        res.Entity.Name,
			ServiceName: res.Entity.ServiceName,
		}

		if !srv.IsLoadedDispatcherPlugin(e) {
			srv.LoadDispatcherPlugin(e)
		}
	}
	srv.heartbeat_entities[entity_id] = time.Now()
	srv.logger.WithField("id", req.EntityId.Value).Debugf("entity heartbeat")

	return &empty.Empty{}, nil
}

func (srv *coreAgentService) ServeOnStream() error {
	token := srv.app_cred_mgr.GetToken()
	ctx := context_helper.WithToken(context.Background(), token)

	cli, cfn, err := srv.cli_fty.NewCoreServiceClient(GRPC_KEEPALIVE)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to dial to metathings service")
		return err
	}
	defer cfn()

	stream, err := cli.Stream(ctx)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to stream to core service")
		return err
	}
	srv.logger.Debugf("connect to core service on streaming")

	return srv.serveOnStream(stream)
}

func (srv *coreAgentService) serveOnStream(stream core_pb.CoreService_StreamClient) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			err = srv.handleGRPCError(err, "failed to recv data from core")
			return err
		}

		ctx := stream.Context()
		res, err := srv.dispatch(ctx, req)
		if err != nil {
			srv.logger.WithError(err).Errorf("failed to dispatch")
			continue
		}

		err = stream.Send(res)
		if err != nil {
			err = srv.handleGRPCError(err, "failed to send data to entity")
			return err
		}
	}
}

func (srv *coreAgentService) dispatch_system(ctx context.Context, req *core_pb.StreamRequest) (*core_pb.StreamResponse, error) {
	switch req.Payload.(type) {
	case *core_pb.StreamRequest_StreamCall:
		return srv.dispatch_system_stream(ctx, req)
	default:
		return nil, ErrUnsupportMessageType
	}

}

func (srv *coreAgentService) dispatch_system_stream(ctx context.Context, req *core_pb.StreamRequest) (*core_pb.StreamResponse, error) {
	payload := req.Payload.(*core_pb.StreamRequest_StreamCall)

	switch payload.StreamCall.Payload.(type) {
	case *core_pb.StreamCallRequestPayload_Config:
		return srv.dispatch_system_stream_config(ctx, req)
	default:
		return nil, ErrUnsupportMessageType
	}
}

func (srv *coreAgentService) dispatch_system_stream_config(ctx context.Context, req *core_pb.StreamRequest) (*core_pb.StreamResponse, error) {
	payload := req.Payload.(*core_pb.StreamRequest_StreamCall)
	config := payload.StreamCall.Payload.(*core_pb.StreamCallRequestPayload_Config)
	name := config.Config.Name.Value
	service_name := config.Config.ServiceName.Value
	method_name := config.Config.MethodName.Value

	dp, ok := srv.getDispatcherPlugin(name, service_name)
	if !ok {
		return nil, ErrPluginNotFound
	}

	estm, err := dp.StreamCall(method_name, ctx)
	if err != nil {
		return nil, err
	}
	srv.logger.Debugf("build streaming to entity for stream call")

	cli, cfn, err := srv.cli_fty.NewCoreServiceClient(GRPC_KEEPALIVE)
	if err != nil {
		return nil, err
	}
	// dont close connect here, cause it will be used by other goroutines.

	cstm_ctx := context_helper.NewOutgoingContext(
		context.Background(),
		context_helper.WithTokenOp(srv.app_cred_mgr.GetToken()),
		context_helper.WithSessionIdOp(req.SessionId.Value))
	cstm, err := cli.Stream(cstm_ctx)
	if err != nil {
		return nil, err
	}
	srv.logger.Debugf("build streaming to core for stream call")

	clear := func() {
		cfn()
		estm.Close()
	}

	// TODO(Peer): pass context to entity.
	// core -> entity
	go func() {
		defer clear()
		for {
			creq, err := cstm.Recv()
			if err != nil {
				srv.handleGRPCError(err, "failed to recv stream data from core")
				return
			}
			srv.logger.Debugf("recv data from core")

			stm_call, ok := creq.Payload.(*core_pb.StreamRequest_StreamCall)
			if !ok {
				srv.logger.WithField("stage", "StreamRequest_StreamCall").Errorf("failed to convert request type")
				continue
			}

			dat, ok := stm_call.StreamCall.Payload.(*core_pb.StreamCallRequestPayload_Data)
			if !ok {
				srv.logger.WithField("stage", "StreamCallRequestPayload_Data").Errorf("failed to convert request type")
				continue
			}

			req := dat.Data.Value
			err = estm.Send(req)
			if err != nil {
				srv.handleGRPCError(err, "failed to send data to entity")
				return
			}
			srv.logger.Debugf("send data to entity")
		}
	}()

	// entity -> core
	go func() {
		defer clear()
		for {
			ereq, err := estm.Recv()
			if err != nil {
				srv.handleGRPCError(err, "failed to recv data from entity")
				return
			}
			srv.logger.Debugf("recv data from entity")

			res := &core_pb.StreamResponse{
				MessageType: core_pb.StreamMessageType_STREAM_MESSAGE_TYPE_USER,
				Payload: &core_pb.StreamResponse_StreamCall{
					StreamCall: &core_pb.StreamCallResponsePayload{
						Payload: &core_pb.StreamCallResponsePayload_Data{
							Data: &core_pb.StreamCallDataResponse{
								Value: ereq,
							},
						},
					},
				},
			}

			err = cstm.Send(res)
			if err != nil {
				srv.handleGRPCError(err, "failed to send data to core")
				return
			}
			srv.logger.Debugf("send data to core")
		}
	}()

	return &core_pb.StreamResponse{
		SessionId:   req.SessionId.Value,
		MessageType: req.MessageType,
		Payload: &core_pb.StreamResponse_StreamCall{
			StreamCall: &core_pb.StreamCallResponsePayload{
				Payload: &core_pb.StreamCallResponsePayload_Config{
					Config: &core_pb.StreamCallConfigResponse{
						Name:        name,
						ServiceName: service_name,
						MethodName:  method_name,
					},
				},
			},
		},
	}, nil
}

func (srv *coreAgentService) dispatch_user(ctx context.Context, req *core_pb.StreamRequest) (*core_pb.StreamResponse, error) {
	payload, ok := req.Payload.(*core_pb.StreamRequest_UnaryCall)
	if !ok {
		return nil, ErrUnsupportPayloadType
	}

	name := payload.UnaryCall.Name.Value
	service_name := payload.UnaryCall.ServiceName.Value
	method_name := payload.UnaryCall.MethodName.Value
	req_value := payload.UnaryCall.Value

	dp, ok := srv.getDispatcherPlugin(name, service_name)
	if !ok {
		return nil, ErrPluginNotFound
	}

	res, err := dp.UnaryCall(method_name, ctx, req_value)
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
				Value:       res,
			},
		},
	}

	return res1, nil
}

func (srv *coreAgentService) dispatch(ctx context.Context, req *core_pb.StreamRequest) (*core_pb.StreamResponse, error) {
	switch req.MessageType {
	case core_pb.StreamMessageType_STREAM_MESSAGE_TYPE_USER:
		return srv.dispatch_user(ctx, req)
	case core_pb.StreamMessageType_STREAM_MESSAGE_TYPE_SYSTEM:
		return srv.dispatch_system(ctx, req)
	default:
		return nil, ErrUnsupportMessageType
	}
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

func getCoreIdFromService(cli_fty *client_helper.ClientFactory, token string) (string, error) {
	ctx := context.Background()
	ctx = context_helper.WithToken(ctx, token)
	cli, fn, err := cli_fty.NewCoreServiceClient()
	if err != nil {
		return "", err
	}
	defer fn()

	show_core_res, err := cli.ShowCore(ctx, &empty.Empty{})
	if err == nil {
		return show_core_res.Core.Id, nil
	}

	create_core_req := &cored_pb.CreateCoreRequest{}
	create_core_res, err := cli.CreateCore(ctx, create_core_req)
	if err != nil {
		return "", err
	}

	return create_core_res.Core.Id, nil
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
			if core_id, err = getCoreIdFromService(cli_fty, app_cred_mgr.GetToken()); err != nil {
				return nil, err
			}
			if err = saveCoreIdToPath(core_id, core_id_path); err != nil {
				return nil, err
			}
		}
		opts.core_id = core_id
	}

	serv_desc, err := mt_plugin.LoadServiceDescriptor(opts.service_descriptor)
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

		mtx_dp_op:          new(sync.Mutex),
		dispatchers:        make(map[string]mt_plugin.DispatcherPlugin),
		heartbeat_entities: make(map[string]time.Time),
	}
	return srv, nil
}
