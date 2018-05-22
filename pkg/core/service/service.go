package metathings_core_service

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/nayotta/metathings/pkg/common"
	app_cred_mgr "github.com/nayotta/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	state_helper "github.com/nayotta/metathings/pkg/common/state"
	stm_mgr "github.com/nayotta/metathings/pkg/common/stream_manager"
	storage "github.com/nayotta/metathings/pkg/core/storage"
	state_pb "github.com/nayotta/metathings/pkg/proto/common/state"
	pb "github.com/nayotta/metathings/pkg/proto/core"
	identityd_pb "github.com/nayotta/metathings/pkg/proto/identity"
)

type options struct {
	logLevel                      string
	identityd_addr                string
	application_credential_id     string
	application_credential_secret string
	storage_driver                string
	storage_uri                   string
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

func SetStorage(driver, uri string) ServiceOptions {
	return func(o *options) {
		o.storage_driver = driver
		o.storage_uri = uri
	}
}

type metathingsCoreService struct {
	grpc_helper.AuthorizationTokenParser

	cli_fty       *client_helper.ClientFactory
	core_st_psr   state_helper.CoreStateParser
	entity_st_psr state_helper.EntityStateParser
	app_cred_mgr  app_cred_mgr.ApplicationCredentialManager
	stm_mgr       stm_mgr.StreamManager
	logger        log.FieldLogger
	opts          options
	storage       storage.Storage
}

func (srv *metathingsCoreService) validateTokenViaIdentityd(token string) (*identityd_pb.Token, error) {
	ctx := context.Background()
	md := metadata.Pairs(
		"authorization-subject", "mt "+token,
		"authorization", srv.app_cred_mgr.GetToken(),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	cli, closeFn, err := srv.cli_fty.NewIdentityServiceClient()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to new core service client")
		return nil, err
	}
	defer closeFn()

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
			State:     srv.core_st_psr.ToValue(*cc.State),
		},
	}

	return res, nil
}

func (srv *metathingsCoreService) DeleteCore(ctx context.Context, req *pb.DeleteCoreRequest) (*empty.Empty, error) {
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

func (srv *metathingsCoreService) copyEntity(e storage.Entity) *pb.Entity {
	return &pb.Entity{
		Id:          *e.Id,
		CoreId:      *e.CoreId,
		Name:        *e.Name,
		ServiceName: *e.ServiceName,
		Endpoint:    *e.Endpoint,
		State:       srv.entity_st_psr.ToValue(*e.State),
	}
}

func (srv *metathingsCoreService) CreateEntity(ctx context.Context, req *pb.CreateEntityRequest) (*pb.CreateEntityResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	cred := context_helper.Credential(ctx)
	core, err := srv.storage.GetAssignedCoreFromApplicationCredential(cred.ApplicationCredential.Id)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to get assgined core from application credential")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	entity_id := common.NewId()

	var name_str string
	name := req.GetName()
	if name != nil {
		name_str = name.Value
	} else {
		name_str = req.GetServiceName().Value
	}
	state := "offline"

	e := storage.Entity{
		Id:          &entity_id,
		Name:        &name_str,
		ServiceName: &req.ServiceName.Value,
		Endpoint:    &req.Endpoint.Value,
		CoreId:      core.Id,
		State:       &state,
	}

	ce, err := srv.storage.CreateEntity(e)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to create entity")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	srv.logger.WithFields(log.Fields{
		"entity_id": *ce.Id,
		"core_id":   *ce.CoreId,
	}).Infof("create entity")

	res := &pb.CreateEntityResponse{
		Entity: &pb.Entity{
			Id:          *ce.Id,
			CoreId:      *ce.CoreId,
			Name:        *ce.Name,
			ServiceName: *ce.ServiceName,
			Endpoint:    *ce.Endpoint,
			State:       srv.entity_st_psr.ToValue(*ce.State),
		},
	}

	return res, nil
}

func (srv *metathingsCoreService) DeleteEntity(ctx context.Context, req *pb.DeleteEntityRequest) (*empty.Empty, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	err = srv.storage.DeleteEntity(req.Id.Value)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to delete entity")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	srv.logger.WithField("id", req.Id.Value).Infof("delete entity")

	return &empty.Empty{}, nil
}

func (srv *metathingsCoreService) PatchEntity(ctx context.Context, req *pb.PatchEntityRequest) (*pb.PatchEntityResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	fields := log.Fields{"id": req.Id.Value}
	entity := storage.Entity{}
	if req.State != state_pb.EntityState_ENTITY_STATE_UNKNOWN {
		state_str := srv.entity_st_psr.ToString(req.State)
		entity.State = &state_str
		fields["state"] = state_str
	}

	entity, err = srv.storage.PatchEntity(req.Id.Value, entity)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to patch entity")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	srv.logger.WithFields(fields).Infof("patch entity")

	return &pb.PatchEntityResponse{Entity: srv.copyEntity(entity)}, nil
}

func (srv *metathingsCoreService) GetEntity(ctx context.Context, req *pb.GetEntityRequest) (*pb.GetEntityResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	entity, err := srv.storage.GetEntity(req.Id.Value)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to get entity")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	srv.logger.WithField("id", req.Id.Value).Debugf("get entity")

	return &pb.GetEntityResponse{Entity: srv.copyEntity(entity)}, nil
}

func (srv *metathingsCoreService) ListEntities(ctx context.Context, req *pb.ListEntitiesRequest) (*pb.ListEntitiesResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	entities, err := srv.storage.ListEntities(storage.Entity{})
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to list entities")
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	es := []*pb.Entity{}
	for _, e := range entities {
		es = append(es, srv.copyEntity(e))
	}

	return &pb.ListEntitiesResponse{Entities: es}, nil
}

func (srv *metathingsCoreService) ListEntitiesForCore(ctx context.Context, req *pb.ListEntitiesForCoreRequest) (*pb.ListEntitiesForCoreResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	entites, err := srv.storage.ListEntitiesForCore(req.CoreId.Value, storage.Entity{})
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to list entities for core")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	es := []*pb.Entity{}
	for _, e := range entites {
		es = append(es, srv.copyEntity(e))
	}

	return &pb.ListEntitiesForCoreResponse{Entities: es}, nil
}

func (srv *metathingsCoreService) Heartbeat(ctx context.Context, req *pb.HeartbeatRequest) (*empty.Empty, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	cred := context_helper.Credential(ctx)
	core, err := srv.storage.GetAssignedCoreFromApplicationCredential(cred.ApplicationCredential.Id)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to get assigned core from application credential")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if *core.State != "online" {
		state_str := "online"
		_, err = srv.storage.PatchCore(*core.Id, storage.Core{State: &state_str})
		if err != nil {
			srv.logger.WithError(err).Errorf("failed to patch core")
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	return &empty.Empty{}, nil
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
			WithField("core_id", *core.Id).
			WithError(err).
			Errorf("failed to register stream")
	}
	srv.logger.WithField("core_id", *core.Id).Infof("register stream")

	<-close_chan
	srv.logger.WithField("core_id", *core.Id).Infof("stream closed")

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
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.UnaryCallResponse{Payload: res}, nil
}

func (srv *metathingsCoreService) StreamCall(cstm pb.CoreService_StreamCallServer) error {
	req, err := cstm.Recv()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to recv config")
		return status.Errorf(codes.Internal, err.Error())
	}

	if !isStreamCallConfigRequestPayload(req.Payload) {
		srv.logger.Errorf("not stream call config request")
		return status.Errorf(codes.Internal, "unconfiged stream call request")
	}

	agstm, close_fn, err := srv.stm_mgr.StreamCall(req.CoreId.Value, req.Payload)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to stream call to core agent")
		return status.Errorf(codes.Internal, err.Error())
	}
	defer close_fn()
	srv.logger.Debugf("stream call config done")

	errs := make(chan error)
	quit := make(chan interface{})

	go func() {
		var err error
		var req *pb.StreamCallRequest

		defer func() {
			errs <- err
			quit <- nil
		}()

		for {
			select {
			case <-quit:
				srv.logger.Debugf("receive quit signal, quit core side stream")
				return
			default:
			}

			req, err = cstm.Recv()
			if err != nil {
				srv.logger.WithError(err).Errorf("failed to recv data from client")
				return
			}
			srv.logger.Debugf("recv data from client")

			if !isStreamCallDataRequestPayload(req.Payload) {
				srv.logger.Warningf("not stream call data request")
				continue
			}

			err = agstm.Send(&pb.StreamRequest{
				MessageType: pb.StreamMessageType_STREAM_MESSAGE_TYPE_USER,
				Payload:     &pb.StreamRequest_StreamCall{StreamCall: req.Payload},
			})
			if err != nil {
				srv.logger.WithError(err).Errorf("failed to send data to client")
				return
			}
			srv.logger.Debugf("send data to agent")

		}
	}()

	go func() {
		var err error
		var res *pb.StreamResponse

		defer func() {
			errs <- err
			quit <- nil
		}()

		for {
			select {
			case <-quit:
				srv.logger.Debugf("receive quit signal, quit agent side stream")
				return
			default:
			}

			res, err = agstm.Recv()
			if err != nil {
				srv.logger.WithError(err).Errorf("failed to recv data from agent")
				return
			}
			srv.logger.Debugf("recv data from agent")

			if !isStreamCallDataResponsePayload(res) {
				srv.logger.Warningf("not stream call data response")
				continue
			}

			err = cstm.Send(&pb.StreamCallResponse{
				Payload: &pb.StreamCallResponsePayload{
					Payload: res.Payload.(*pb.StreamResponse_StreamCall).StreamCall.Payload,
				},
			})
			if err != nil {
				srv.logger.WithError(err).Errorf("failed to send data to core")
			}
			srv.logger.Debugf("send data to client")
		}
	}()

	err = <-errs
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to stream call")
		return err
	}

	srv.logger.Debugf("stream call done")

	return nil
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

	cli_fty, err := client_helper.NewClientFactory(
		client_helper.NewDefaultServiceConfigs(opts.identityd_addr),
		client_helper.WithInsecureOptionFunc(),
	)
	if err != nil {
		log.WithError(err).Errorf("failed to new client factory")
		return nil, err
	}

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

	stm_mgr, err := stm_mgr.NewStreamManager(logger)
	if err != nil {
		log.WithError(err).Errorf("failed to new stream manager")
		return nil, err
	}

	srv := &metathingsCoreService{
		cli_fty:       cli_fty,
		core_st_psr:   state_helper.NewCoreStateParser(),
		entity_st_psr: state_helper.NewEntityStateParser(),
		app_cred_mgr:  app_cred_mgr,
		stm_mgr:       stm_mgr,
		opts:          opts,
		logger:        logger,
		storage:       storage,
	}
	return srv, nil
}
