package metathings_core_service

import (
	"context"
	"time"

	empty "github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/nayotta/metathings/pkg/common"
	app_cred_mgr "github.com/nayotta/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	protobuf_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	state_helper "github.com/nayotta/metathings/pkg/common/state"
	stm_mgr "github.com/nayotta/metathings/pkg/common/stream_manager"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	storage "github.com/nayotta/metathings/pkg/cored/storage"
	state_pb "github.com/nayotta/metathings/pkg/proto/common/state"
	pb "github.com/nayotta/metathings/pkg/proto/cored"
)

type options struct {
	logLevel                      string
	identityd_addr                string
	application_credential_id     string
	application_credential_secret string
	storage_driver                string
	storage_uri                   string
	core_alive_timeout            time.Duration
	entity_alive_timeout          time.Duration
}

var defaultServiceOptions = options{
	logLevel:             "info",
	core_alive_timeout:   30 * time.Second,
	entity_alive_timeout: 30 * time.Second,
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

func SetCoreAliveTimeout(timeout int) ServiceOptions {
	return func(o *options) {
		o.core_alive_timeout = time.Duration(timeout) * time.Second
	}
}

func SetEntityAliveTimeout(timeout int) ServiceOptions {
	return func(o *options) {
		o.entity_alive_timeout = time.Duration(timeout) * time.Second
	}
}

type metathingsCoredService struct {
	grpc_helper.AuthorizationTokenParser

	cli_fty             *client_helper.ClientFactory
	core_st_psr         state_helper.CoreStateParser
	entity_st_psr       state_helper.EntityStateParser
	app_cred_mgr        app_cred_mgr.ApplicationCredentialManager
	stm_mgr             stm_mgr.StreamManager
	logger              log.FieldLogger
	opts                options
	storage             storage.Storage
	tk_vdr              token_helper.TokenValidator
	core_maintain_chans map[string]chan interface{}
}

func (srv *metathingsCoredService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	token_str, err := srv.GetTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}

	token, err := srv.tk_vdr.Validate(token_str)
	if err != nil {
		srv.logger.
			WithField("error", err).
			Errorf("failed to validate token via identityd")
		return nil, err
	}

	ctx = context.WithValue(ctx, "token", token_str)
	ctx = context.WithValue(ctx, "credential", token)

	srv.logger.WithFields(log.Fields{
		"method":   fullMethodName,
		"user_id":  token.User.Id,
		"username": token.User.Name,
	}).Debugf("validate token")

	return ctx, nil
}

func (srv *metathingsCoredService) maintain_core_once(core_id string) {
	if _, ok := srv.core_maintain_chans[core_id]; !ok {
		srv.core_maintain_chans[core_id] = make(chan interface{})
		srv.logger.WithField("core_id", core_id).Debugf("create core maintain channel")
		go srv.maintain_core_loop(core_id)
	}
	srv.core_maintain_chans[core_id] <- nil
	srv.logger.WithField("core_id", core_id).Debugf("send heartbeat signal to core maintain channel")
}

func (srv *metathingsCoredService) maintain_core_loop(core_id string) {
	ch, ok := srv.core_maintain_chans[core_id]
	if !ok {
		srv.logger.WithField("core_id", core_id).Errorf("core maintain channel not found")
		return
	}
	srv.logger.WithField("core_id", core_id).Debugf("start core maintain loop")

	defer func() {
		delete(srv.core_maintain_chans, core_id)
		srv.logger.WithField("core_id", core_id).Debugf("quit core maintain loop")
	}()

	for {
		select {
		case <-ch:
			srv.logger.WithField("core_id", core_id).Debugf("receive heartbeat signal")
			continue
		case <-time.After(srv.opts.core_alive_timeout):
			srv.logger.WithField("core_id", core_id).Warningf("core heartbeat timeout")
			srv.maintain_core(core_id)
			return
		}
	}
}

func (srv *metathingsCoredService) maintain_core(core_id string) {
	state_str := "offline"
	pc := storage.Core{State: &state_str}
	_, err := srv.storage.PatchCore(core_id, pc)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to patch core")
		return
	}

	es, err := srv.storage.ListEntitiesForCore(core_id, storage.Entity{})
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to list entities for core")
		return
	}

	entity_ids := []string{}
	for _, e := range es {
		if *e.State == "offline" {
			continue
		}
		pe := storage.Entity{State: &state_str}
		_, err := srv.storage.PatchEntity(*e.Id, pe)
		entity_ids = append(entity_ids, *e.Id)
		if err != nil {
			srv.logger.WithError(err).Errorf("failed to patch entity")
		}
	}

	srv.logger.WithFields(log.Fields{
		"core_id":    core_id,
		"entity_ids": entity_ids,
	}).Infof("core agent offline")
}

func (srv *metathingsCoredService) CreateCore(ctx context.Context, req *pb.CreateCoreRequest) (*pb.CreateCoreResponse, error) {
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
		Core: srv.copyCore(cc),
	}

	return res, nil
}

func (srv *metathingsCoredService) DeleteCore(ctx context.Context, req *pb.DeleteCoreRequest) (*empty.Empty, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

func (srv *metathingsCoredService) PatchCore(context.Context, *pb.PatchCoreRequest) (*pb.PatchCoreResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "unimplement")
}

func (srv *metathingsCoredService) GetCore(ctx context.Context, req *pb.GetCoreRequest) (*pb.GetCoreResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	c, err := srv.storage.GetCore(req.GetId().GetValue())
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to get core")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	srv.logger.WithField("id", *c.Id).Debugf("get core")

	res := &pb.GetCoreResponse{
		Core: srv.copyCore(c),
	}

	return res, nil
}

func (srv *metathingsCoredService) ListCores(ctx context.Context, req *pb.ListCoresRequest) (*pb.ListCoresResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	c := storage.Core{}

	name := req.GetName()
	if name != nil {
		c.Name = &name.Value
	}
	projectId := req.GetProjectId()
	if projectId != nil {
		c.ProjectId = &projectId.Value
	}
	ownerId := req.GetOwnerId()
	if ownerId != nil {
		c.OwnerId = &ownerId.Value
	}
	state := req.GetState()
	if state != state_pb.CoreState_CORE_STATE_UNKNOWN {
		state_str := srv.core_st_psr.ToString(state)
		c.State = &state_str
	}

	cs, err := srv.storage.ListCores(c)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to list cores")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListCoresResponse{
		Cores: []*pb.Core{},
	}
	for _, c := range cs {
		res.Cores = append(res.Cores, srv.copyCore(c))
	}

	srv.logger.Debugf("list cores")
	return res, nil
}

func (srv *metathingsCoredService) ShowCore(ctx context.Context, req *empty.Empty) (*pb.ShowCoreResponse, error) {
	cred := context_helper.Credential(ctx)

	c, err := srv.storage.GetAssignedCoreFromApplicationCredential(cred.ApplicationCredential.Id)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to get assigned core from application credential")
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ShowCoreResponse{
		Core: srv.copyCore(c),
	}

	srv.logger.WithField("application_credential_id", cred.ApplicationCredential.Id).Debugf("show core")
	return res, nil
}

func (srv *metathingsCoredService) copyCore(c storage.Core) *pb.Core {
	return &pb.Core{
		Id:        *c.Id,
		Name:      *c.Name,
		ProjectId: *c.ProjectId,
		OwnerId:   *c.OwnerId,
		State:     srv.core_st_psr.ToValue(*c.State),
	}
}

func (srv *metathingsCoredService) copyEntity(e storage.Entity) *pb.Entity {
	return &pb.Entity{
		Id:          *e.Id,
		CoreId:      *e.CoreId,
		Name:        *e.Name,
		ServiceName: *e.ServiceName,
		Endpoint:    *e.Endpoint,
		State:       srv.entity_st_psr.ToValue(*e.State),
	}
}

func (srv *metathingsCoredService) CreateEntity(ctx context.Context, req *pb.CreateEntityRequest) (*pb.CreateEntityResponse, error) {
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

	srv.logger.WithFields(log.Fields{
		"entity_id": *ce.Id,
		"core_id":   *ce.CoreId,
	}).Infof("create entity")
	return res, nil
}

func (srv *metathingsCoredService) DeleteEntity(ctx context.Context, req *pb.DeleteEntityRequest) (*empty.Empty, error) {
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

func (srv *metathingsCoredService) PatchEntity(ctx context.Context, req *pb.PatchEntityRequest) (*pb.PatchEntityResponse, error) {
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

func (srv *metathingsCoredService) GetEntity(ctx context.Context, req *pb.GetEntityRequest) (*pb.GetEntityResponse, error) {
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

func (srv *metathingsCoredService) ListEntities(ctx context.Context, req *pb.ListEntitiesRequest) (*pb.ListEntitiesResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	e := storage.Entity{}

	core_id := req.GetCoreId()
	if core_id != nil {
		e.CoreId = &core_id.Value
	}

	name := req.GetName()
	if name != nil {
		e.Name = &name.Value
	}

	service_name := req.GetServiceName()
	if service_name != nil {
		e.ServiceName = &service_name.Value
	}

	state := req.GetState()
	if state != state_pb.EntityState_ENTITY_STATE_UNKNOWN {
		state_str := srv.entity_st_psr.ToString(state)
		e.State = &state_str
	}

	entities, err := srv.storage.ListEntities(e)
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

func (srv *metathingsCoredService) ListEntitiesForCore(ctx context.Context, req *pb.ListEntitiesForCoreRequest) (*pb.ListEntitiesForCoreResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	cred := context_helper.Credential(ctx)
	core, err := srv.storage.GetAssignedCoreFromApplicationCredential(cred.ApplicationCredential.Id)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to get assigned core from application credential")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	e := storage.Entity{}

	name := req.GetName()
	if name != nil {
		e.Name = &name.Value
	}

	service_name := req.GetServiceName()
	if service_name != nil {
		e.ServiceName = &service_name.Value
	}

	state := req.GetState()
	if state != state_pb.EntityState_ENTITY_STATE_UNKNOWN {
		state_str := srv.entity_st_psr.ToString(state)
		e.State = &state_str
	}

	entites, err := srv.storage.ListEntitiesForCore(*core.Id, e)
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

func (srv *metathingsCoredService) Heartbeat(ctx context.Context, req *pb.HeartbeatRequest) (*empty.Empty, error) {
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

	now := time.Now()
	pc := storage.Core{
		HeartbeatAt: &now,
	}

	if *core.State != "online" {
		state_str := "online"
		pc.State = &state_str
	}
	_, err = srv.storage.PatchCore(*core.Id, pc)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to patch core")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for _, hbe := range req.Entities {
		e, err := srv.storage.GetEntity(hbe.Id.Value)
		if err != nil {
			srv.logger.WithError(err).Warningf("failed to get entity on heartbeat")
			continue
		}

		if *e.CoreId != *core.Id {
			srv.logger.Warningf("entity not belong to core")
			continue
		}

		patch_flag := false
		var pe storage.Entity
		var state_str string
		hbt := protobuf_helper.ToTime(*hbe.HeartbeatAt)
		if e.HeartbeatAt == nil || !e.HeartbeatAt.Equal(hbt) {
			pe.HeartbeatAt = &hbt
			patch_flag = true
		}
		if time.Now().Sub(hbt) > srv.opts.entity_alive_timeout {
			state_str = "offline"
		} else {
			state_str = "online"
		}
		if *e.State != state_str {
			pe.State = &state_str
			patch_flag = true
		}
		if patch_flag {
			_, err = srv.storage.PatchEntity(hbe.Id.Value, pe)
			if err != nil {
				srv.logger.WithError(err).Errorf("failed to patch entity")
				return nil, status.Errorf(codes.Internal, err.Error())
			}
		}
	}
	srv.maintain_core_once(*core.Id)

	return &empty.Empty{}, nil
}

func (srv *metathingsCoredService) Stream(stream pb.CoredService_StreamServer) error {
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
		return err
	}
	srv.logger.WithField("core_id", *core.Id).Infof("register stream")

	<-close_chan
	srv.logger.WithField("core_id", *core.Id).Infof("stream closed")

	return nil
}

func (srv *metathingsCoredService) ListCoresForUser(ctx context.Context, req *pb.ListCoresForUserRequest) (*pb.ListCoresForUserResponse, error) {
	cred := context_helper.Credential(ctx)
	user_id := cred.User.Id
	c := storage.Core{}

	name := req.GetName()
	if name != nil {
		c.Name = &name.Value
	}

	projectId := req.GetProjectId()
	if projectId != nil {
		c.ProjectId = &projectId.Value
	}

	state := req.GetState()
	if state != state_pb.CoreState_CORE_STATE_UNKNOWN {
		state_str := srv.core_st_psr.ToString(state)
		c.State = &state_str
	}

	cs, err := srv.storage.ListCoresForUser(user_id, c)
	if err != nil {
		srv.logger.WithField("user_id", user_id).WithError(err).Errorf("failed to list cores for user")
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListCoresForUserResponse{
		Cores: []*pb.Core{},
	}
	for _, c := range cs {
		res.Cores = append(res.Cores, srv.copyCore(c))
	}

	srv.logger.WithField("user_id", user_id).Debugf("list cores for user")

	return res, nil
}

func (srv *metathingsCoredService) UnaryCall(ctx context.Context, req *pb.UnaryCallRequest) (*pb.UnaryCallResponse, error) {
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

func (srv *metathingsCoredService) StreamCall(cstm pb.CoredService_StreamCallServer) error {
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
			req, err = cstm.Recv()
			select {
			case <-quit:
				srv.logger.Debugf("receive quit signal, quit core side stream")
				return
			default:
			}
			if err != nil {
				err = srv.handleGRPCError(err, "failed to recv data from client")
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
				err = srv.handleGRPCError(err, "failed to send data to client")
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
			res, err = agstm.Recv()
			select {
			case <-quit:
				srv.logger.Debugf("receive quit signal, quit agent side stream")
				return
			default:
			}

			if err != nil {
				err = srv.handleGRPCError(err, "failed to recv data from agent")
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
				err = srv.handleGRPCError(err, "failed to send data to core")
				return
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

func NewCoredService(opt ...ServiceOptions) (*metathingsCoredService, error) {
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

	srv := &metathingsCoredService{
		cli_fty:             cli_fty,
		core_st_psr:         state_helper.NewCoreStateParser(),
		entity_st_psr:       state_helper.NewEntityStateParser(),
		app_cred_mgr:        app_cred_mgr,
		stm_mgr:             stm_mgr,
		opts:                opts,
		logger:              logger,
		storage:             storage,
		tk_vdr:              token_helper.NewTokenValidator(app_cred_mgr, cli_fty, logger),
		core_maintain_chans: map[string]chan interface{}{},
	}
	return srv, nil
}
