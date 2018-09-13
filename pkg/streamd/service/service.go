package metathings_streamd_service

import (
	"context"
	"encoding/json"
	"fmt"

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

func (self *metathingsStreamdService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	err := req.Validate()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	stm, err := self.parse_storage_stream(ctx, req)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to parse request to storage stream")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	_, err = self.storage.CreateStream(stm)
	if err != nil {
		self.logger.WithError(err).Errorf("failed to create stream in storage")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	panic("unimplemented")
}

func encode_config_to_json_string(x map[string]*pb.ConfigValue) (string, error) {
	y := map[string]string{}
	for k, v := range x {
		switch v.GetValue().(type) {
		case *pb.ConfigValue_Double:
			y[k] = fmt.Sprintf("%v", v.GetDouble())
		case *pb.ConfigValue_Int64:
			y[k] = fmt.Sprintf("%v", v.GetInt64())
		case *pb.ConfigValue_Uint64:
			y[k] = fmt.Sprintf("%v", v.GetUint64())
		case *pb.ConfigValue_String_:
			y[k] = fmt.Sprintf("%v", v.GetString_())
		}
	}

	buf, err := json.Marshal(y)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func (self *metathingsStreamdService) parse_storage_stream(ctx context.Context, req *pb.CreateRequest) (storage.Stream, error) {
	cred := context_helper.Credential(ctx)
	stm_id := id_helper.NewId()
	stm_name := req.GetName().GetValue()
	stm_state := "stop"

	sources := []storage.Source{}
	for _, req_src := range req.GetSources() {
		src_id := id_helper.NewId()
		upstm_id := id_helper.NewId()
		req_upstm := req_src.GetUpstream()
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
		grp_id := id_helper.NewId()

		inputs := []storage.Input{}
		for _, req_in := range req_grp.GetInputs() {
			in_id := id_helper.NewId()
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
			out_id := id_helper.NewId()
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

func (self *metathingsStreamdService) Delete(context.Context, *pb.DeleteRequest) (*empty.Empty, error) {
	panic("unimplemented")
}

func (self *metathingsStreamdService) Start(context.Context, *pb.StartRequest) (*pb.StartResponse, error) {
	panic("unimplemented")
}

func (self *metathingsStreamdService) Stop(context.Context, *pb.StopRequest) (*pb.StopResponse, error) {
	panic("unimplemented")
}

func (self *metathingsStreamdService) Get(context.Context, *pb.GetRequest) (*pb.GetResponse, error) {
	panic("unimplemented")
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
