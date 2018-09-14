package metathings_streamd_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	app_cred_mgr "github.com/nayotta/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	pb "github.com/nayotta/metathings/pkg/proto/streamd"
	state_helper "github.com/nayotta/metathings/pkg/streamd/state"
	storage "github.com/nayotta/metathings/pkg/streamd/storage"
	stream_manager "github.com/nayotta/metathings/pkg/streamd/stream"
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
}

type ServiceOptions func(*options)

var defaultServiceOptions = options{
	logLevel: "info",
}

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

type metathingsStreamdService struct {
	grpc_helper.AuthorizationTokenParser

	cli_fty       *client_helper.ClientFactory
	stream_st_psr state_helper.StreamStateParser
	app_cred_mgr  app_cred_mgr.ApplicationCredentialManager
	logger        log.FieldLogger
	opts          options
	storage       storage.Storage
	tk_vdr        token_helper.TokenValidator
	stm_mgr       stream_manager.StreamManager
}

func (self *metathingsStreamdService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	token_str, err := self.GetTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}

	token, err := self.tk_vdr.Validate(token_str)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to validate token via identityd")
		return nil, err
	}

	ctx = context.WithValue(ctx, "token", token_str)
	ctx = context.WithValue(ctx, "credential", token)

	self.logger.WithFields(log.Fields{
		"method":   fullMethodName,
		"user_id":  token.User.Id,
		"username": token.User.Name,
	}).Debugf("validate token")

	return ctx, nil
}

func (self *metathingsStreamdService) copyStream(x storage.Stream) *pb.Stream {
	y := &pb.Stream{
		Id:      *x.Id,
		Name:    *x.Name,
		OwnerId: *x.OwnerId,
		State:   self.stream_st_psr.ToValue(*x.State),
		Sources: self.copySources(x.Sources),
		Groups:  self.copyGroups(x.Groups),
	}
	return y
}

func (self *metathingsStreamdService) copySources(xs []storage.Source) []*pb.Source {
	ys := []*pb.Source{}
	for _, x := range xs {
		ys = append(ys, self.copySource(x))
	}
	return ys
}

func (self *metathingsStreamdService) copySource(x storage.Source) *pb.Source {
	y := &pb.Source{
		Id:       *x.Id,
		Upstream: self.copyUpstream(x.Upstream),
	}
	return y
}

func (self *metathingsStreamdService) copyGroups(xs []storage.Group) []*pb.Group {
	ys := []*pb.Group{}
	for _, x := range xs {
		ys = append(ys, self.copyGroup(x))
	}
	return ys
}

func (self *metathingsStreamdService) copyGroup(x storage.Group) *pb.Group {
	y := &pb.Group{
		Id:      *x.Id,
		Inputs:  self.copyInputs(x.Inputs),
		Outputs: self.copyOutputs(x.Outputs),
	}
	return y
}

func (self *metathingsStreamdService) copyUpstream(x storage.Upstream) *pb.Upstream {
	y := &pb.Upstream{
		Id:     *x.Id,
		Name:   *x.Name,
		Alias:  *x.Alias,
		Config: must_decode_json_string_to_config(*x.Config),
	}
	return y
}

func (self *metathingsStreamdService) copyInputs(xs []storage.Input) []*pb.Input {
	ys := []*pb.Input{}
	for _, x := range xs {
		ys = append(ys, self.copyInput(x))
	}
	return ys
}

func (self *metathingsStreamdService) copyInput(x storage.Input) *pb.Input {
	y := &pb.Input{
		Id:     *x.Id,
		Name:   *x.Name,
		Alias:  *x.Alias,
		Config: must_decode_json_string_to_config(*x.Config),
	}
	return y
}

func (self *metathingsStreamdService) copyOutputs(xs []storage.Output) []*pb.Output {
	ys := []*pb.Output{}
	for _, x := range xs {
		ys = append(ys, self.copyOutput(x))
	}
	return ys
}

func (self *metathingsStreamdService) copyOutput(x storage.Output) *pb.Output {
	y := &pb.Output{
		Id:     *x.Id,
		Name:   *x.Name,
		Alias:  *x.Alias,
		Config: must_decode_json_string_to_config(*x.Config),
	}
	return y
}

func (self *metathingsStreamdService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	err := req.Validate()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	self.reinit_create_request(req)
	stm, err := self.parse_storage_stream(ctx, req)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to parse request to storage stream")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	stm, err = self.storage.CreateStream(stm)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to create stream in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	stm_opt := encode_create_request_to_stream_option(req)
	extra := map[string]interface{}{
		"application_credential_manager": self.app_cred_mgr,
		"client_factory":                 self.cli_fty,
		"logger":                         self.logger,
	}

	_, err = self.stm_mgr.NewStream(stm_opt, extra)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to new stream")
		self.storage.DeleteStream(*stm.Id)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CreateResponse{
		Stream: self.copyStream(stm),
	}

	self.logger.WithFields(log.Fields{
		"id":       *stm.Id,
		"name":     *stm.Name,
		"owner_id": *stm.OwnerId,
		"state":    *stm.State,
	})

	return res, nil
}

func (self *metathingsStreamdService) reinit_create_request(req *pb.CreateRequest) {
	req.Id.Value = id_helper.NewId()

	for _, src := range req.GetSources() {
		src.Id.Value = id_helper.NewId()
		src.Upstream.Id.Value = id_helper.NewId()
	}

	for _, grp := range req.GetGroups() {
		grp.Id.Value = id_helper.NewId()
		for _, in := range grp.GetInputs() {
			in.Id.Value = id_helper.NewId()
		}
		for _, out := range grp.GetOutputs() {
			out.Id.Value = id_helper.NewId()
		}
	}
}

func (self *metathingsStreamdService) parse_storage_stream(ctx context.Context, req *pb.CreateRequest) (storage.Stream, error) {
	cred := context_helper.Credential(ctx)
	stm_id := req.GetId().GetValue()
	stm_name := req.GetName().GetValue()
	stm_state := "stop"

	sources := []storage.Source{}
	for _, req_src := range req.GetSources() {
		src_id := req_src.GetId().GetValue()
		req_upstm := req_src.GetUpstream()
		upstm_id := req_upstm.GetId().GetValue()
		upstm_name := req_upstm.GetName().GetValue()
		upstm_alias := req_upstm.GetAlias().GetValue()
		upstm_config, err := encode_config_to_json_string(req_upstm.GetConfig())
		if err != nil {
			return storage.Stream{}, err
		}

		upstream := storage.Upstream{
			Id:       &upstm_id,
			SourceId: &src_id,
			Name:     &upstm_name,
			Alias:    &upstm_alias,
			Config:   &upstm_config,
		}

		source := storage.Source{
			Id:       &src_id,
			StreamId: &stm_id,
			Upstream: upstream,
		}

		sources = append(sources, source)
	}

	groups := []storage.Group{}
	for _, req_grp := range req.GetGroups() {
		grp_id := req_grp.GetId().GetValue()

		inputs := []storage.Input{}
		for _, req_in := range req_grp.GetInputs() {
			in_id := req_in.GetId().GetValue()
			in_name := req_in.GetName().GetValue()
			in_alias := req_in.GetAlias().GetValue()
			in_config, err := encode_config_to_json_string(req_in.GetConfig())
			if err != nil {
				return storage.Stream{}, err
			}

			input := storage.Input{
				Id:      &in_id,
				GroupId: &grp_id,
				Name:    &in_name,
				Alias:   &in_alias,
				Config:  &in_config,
			}

			inputs = append(inputs, input)
		}

		outputs := []storage.Output{}
		for _, req_out := range req_grp.GetOutputs() {
			out_id := req_out.GetId().GetValue()
			out_name := req_out.GetName().GetValue()
			out_alias := req_out.GetAlias().GetValue()
			out_config, err := encode_config_to_json_string(req_out.GetConfig())
			if err != nil {
				return storage.Stream{}, nil
			}

			output := storage.Output{
				Id:      &out_id,
				GroupId: &grp_id,
				Name:    &out_name,
				Alias:   &out_alias,
				Config:  &out_config,
			}

			outputs = append(outputs, output)
		}

		group := storage.Group{
			Id:      &grp_id,
			Inputs:  inputs,
			Outputs: outputs,
		}

		groups = append(groups, group)
	}

	stm := storage.Stream{
		Id:      &stm_id,
		Name:    &stm_name,
		OwnerId: &cred.User.Id,
		State:   &stm_state,
		Sources: sources,
		Groups:  groups,
	}

	return stm, nil
}

func (self *metathingsStreamdService) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	err := req.Validate()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	stm_id := req.GetId().GetValue()
	stm, err := self.stm_mgr.GetStream(stm_id)
	if stm.State() != stream_manager.STREAM_STATE_STOP {
		self.logger.WithField("id", stm_id).Errorf("failed to delete stream cause state not in stop")
		return nil, status.Errorf(codes.FailedPrecondition, "stream state not in stop")
	}

	err = self.stm_mgr.DeleteStream(stm_id)
	if err != nil {
		self.logger.WithError(err).WithField("id", stm_id).Errorf("failed to delete stream in stream manager")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	err = self.storage.DeleteStream(stm_id)
	if err != nil {
		self.logger.WithField("id", stm_id).Errorf("failed to delete stream in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (self *metathingsStreamdService) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	err := req.Validate()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	stm_id := req.GetId().GetValue()
	stm, err := self.stm_mgr.GetStream(stm_id)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get stream from stream manager")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	// TODO(Peer): should be consider about starting timeout.
	stm.Once(stream_manager.START_EVENT, func(stream_manager.Event, interface{}) {
		stm_state := "running"
		_, err := self.storage.PatchStream(stm_id, storage.Stream{State: &stm_state})
		if err != nil {
			self.logger.WithError(err).Errorf("failed to patch stream state")
			return
		}
		self.logger.WithField("id", stm_id).Infof("stream started")
	})

	err = stm.Start()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to start stream")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	stm_state := "starting"
	stm_s, err := self.storage.PatchStream(stm_id, storage.Stream{State: &stm_state})
	if err != nil {
		self.logger.WithError(err).Errorf("failed to patch stream state")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.StartResponse{Stream: self.copyStream(stm_s)}

	self.logger.WithField("id", stm_id).Debugf("start stream")
	return res, nil
}

func (self *metathingsStreamdService) Stop(ctx context.Context, req *pb.StopRequest) (*pb.StopResponse, error) {
	err := req.Validate()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	stm_id := req.GetId().GetValue()
	stm, err := self.stm_mgr.GetStream(stm_id)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get stream from stream manager")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	stm.Once(stream_manager.STOP_EVENT, func(stream_manager.Event, interface{}) {
		stm_state := "stop"
		_, err := self.storage.PatchStream(stm_id, storage.Stream{State: &stm_state})
		if err != nil {
			self.logger.WithError(err).Errorf("failed to patch stream state")
			return
		}
		self.logger.WithField("id", stm_id).Infof("stream terminated")
	})

	err = stm.Stop()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to stop stream")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	stm_state := "termianting"
	stm_s, err := self.storage.PatchStream(stm_id, storage.Stream{State: &stm_state})
	if err != nil {
		self.logger.WithError(err).Errorf("failed to patch state")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.StopResponse{Stream: self.copyStream(stm_s)}

	self.logger.WithField("id", stm_id).Debugf("stop stream")
	return res, nil
}

func (self *metathingsStreamdService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	err := req.Validate()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	stm_id := req.GetId().GetValue()
	stm, err := self.storage.GetStream(stm_id)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to get stream from storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.GetResponse{Stream: self.copyStream(stm)}

	self.logger.WithField("id", stm_id).Debugf("get stream")
	return res, nil
}

func (self *metathingsStreamdService) List(context.Context, *pb.ListRequest) (*pb.ListResponse, error) {
	panic("unimplemented")
}

func (self *metathingsStreamdService) ListForUser(context.Context, *pb.ListForUserRequest) (*pb.ListForUserResponse, error) {
	panic("unimplemented")
}

func NewStreamdService(opt ...ServiceOptions) (*metathingsStreamdService, error) {
	opts := defaultServiceOptions
	for _, o := range opt {
		o(&opts)
	}

	logger, err := log_helper.NewLogger("streamd", opts.logLevel)
	if err != nil {
		return nil, err
	}

	cli_fty_cfgs := client_helper.NewDefaultServiceConfigs(opts.metathingsd_addr)
	cli_fty_cfgs[client_helper.CORED_CONFIG] = client_helper.ServiceConfig{Address: opts.cored_addr}
	cli_fty_cfgs[client_helper.IDENTITYD_CONFIG] = client_helper.ServiceConfig{Address: opts.identityd_addr}
	cli_fty, err := client_helper.NewClientFactory(
		cli_fty_cfgs,
		client_helper.WithInsecureOptionFunc(),
	)
	if err != nil {
		log.WithError(err).Errorf("failed to new client factory")
		return nil, err
	}

	storage, err := storage.NewStorage(opts.storage_driver, opts.storage_uri, logger)
	if err != nil {
		logger.WithError(err).Errorf("failed to connect storage")
		return nil, err
	}

	app_cred_mgr, err := app_cred_mgr.NewApplicationCredentialManager(
		cli_fty,
		opts.application_credential_id,
		opts.application_credential_secret,
	)
	if err != nil {
		logger.WithError(err).Errorf("failed to new application credential manager")
		return nil, err
	}

	stm_mgr_fty := stream_manager.NewDefaultStreamManagerFactory()
	stm_mgr, err := stm_mgr_fty.Set("application_credential_manager", app_cred_mgr).
		Set("client_factory", cli_fty).
		Set("logger", logger).
		New()
	if err != nil {
		logger.WithError(err).Errorf("failed to new stream manager")
		return nil, err
	}

	tk_vdr := token_helper.NewTokenValidator(app_cred_mgr, cli_fty, logger)

	srv := &metathingsStreamdService{
		cli_fty:       cli_fty,
		stream_st_psr: state_helper.NewStreamStateParser(),
		app_cred_mgr:  app_cred_mgr,
		opts:          opts,
		logger:        logger,
		storage:       storage,
		tk_vdr:        tk_vdr,
		stm_mgr:       stm_mgr,
	}

	logger.Debugf("new streamd service")

	return srv, nil
}
