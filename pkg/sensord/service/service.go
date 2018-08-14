package metathings_sensord_service

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
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	protobuf_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	sensor_pb "github.com/nayotta/metathings/pkg/proto/sensor"
	pb "github.com/nayotta/metathings/pkg/proto/sensord"
	state_helper "github.com/nayotta/metathings/pkg/sensor/state"
	"github.com/nayotta/metathings/pkg/sensord/service/hub"
	_ "github.com/nayotta/metathings/pkg/sensord/service/hub/default"
	_ "github.com/nayotta/metathings/pkg/sensord/service/hub/kafka"
	storage "github.com/nayotta/metathings/pkg/sensord/storage"
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
	hub                           opt_helper.Option
}

var defaultServiceOptions = options{
	logLevel: "info",
	hub:      opt_helper.NewOption("name", "default"),
}

type ServiceOptions func(*options)

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

func SetHub(opt opt_helper.Option) ServiceOptions {
	return func(o *options) {
		o.hub = opt
	}
}

type metathingsSensordService struct {
	grpc_helper.AuthorizationTokenParser

	cli_fty       *client_helper.ClientFactory
	sensor_st_psr state_helper.SensorStateParser
	app_cred_mgr  app_cred_mgr.ApplicationCredentialManager
	logger        log.FieldLogger
	opts          options
	storage       storage.Storage
	tk_vdr        token_helper.TokenValidator

	hub hub.Hub
}

func (srv *metathingsSensordService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	token_str, err := srv.GetTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}

	token, err := srv.tk_vdr.Validate(token_str)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate token via identityd")
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

func (srv *metathingsSensordService) copySensor(snr storage.Sensor) *pb.Sensor {
	s := &pb.Sensor{
		Id:                      *snr.Id,
		Name:                    *snr.Name,
		CoreId:                  *snr.CoreId,
		EntityName:              *snr.EntityName,
		OwnerId:                 *snr.OwnerId,
		ApplicationCredentialId: *snr.ApplicationCredentialId,
		State: srv.sensor_st_psr.ToValue(*snr.State),
	}

	s.Tags = []string{}
	for _, t := range snr.Tags {
		s.Tags = append(s.Tags, *t.Tag)
	}

	return s
}

func (srv *metathingsSensordService) copySensors(snrs []storage.Sensor) []*pb.Sensor {
	ss := []*pb.Sensor{}
	for _, snr := range snrs {
		ss = append(ss, srv.copySensor(snr))
	}
	return ss
}

func (srv *metathingsSensordService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	cred := context_helper.Credential(ctx)
	snr_id := id_helper.NewId()
	var name_str string
	if name := req.GetName(); name != nil {
		name_str = name.GetValue()
	} else {
		name_str = snr_id
	}
	core_id := req.GetCoreId().GetValue()
	entity_name := req.GetEntityName().GetValue()
	app_cred_id := req.GetApplicationCredentialId().GetValue()
	state := "unknown"

	snr := storage.Sensor{
		Id:                      &snr_id,
		Name:                    &name_str,
		CoreId:                  &core_id,
		EntityName:              &entity_name,
		OwnerId:                 &cred.User.Id,
		ApplicationCredentialId: &app_cred_id,
		State: &state,
	}

	cs, err := srv.storage.CreateSensor(snr)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to create sensor")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.CreateResponse{
		Sensor: srv.copySensor(cs),
	}

	srv.logger.WithFields(log.Fields{
		"id":          *cs.Id,
		"name":        *cs.Name,
		"core_id":     *cs.CoreId,
		"entity_name": *cs.EntityName,
		"owner_id":    *cs.OwnerId,
		"state":       *cs.State,
	})

	return res, nil
}

func (srv *metathingsSensordService) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	snr_id := req.GetId().GetValue()
	err = srv.storage.DeleteSensor(snr_id)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to delete sensor")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	srv.logger.WithField("id", snr_id).Infof("delete sensor")

	return &empty.Empty{}, nil
}

func (srv *metathingsSensordService) Patch(ctx context.Context, req *pb.PatchRequest) (*pb.PatchResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	updated := false
	s := storage.Sensor{}
	snr_id := req.GetId().GetValue()

	if name := req.GetName(); name != nil {
		s.Name = &name.Value
		updated = true
	}

	if !updated {
		return nil, status.Errorf(codes.InvalidArgument, "empty patch request")
	}

	ps, err := srv.storage.PatchSensor(snr_id, s)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to patch sensor")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.PatchResponse{
		Sensor: srv.copySensor(ps),
	}

	srv.logger.WithField("snr_id", snr_id).Infof("patch sensor")

	return res, nil
}

func (srv *metathingsSensordService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	s, err := srv.storage.GetSensor(req.GetId().GetValue())
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to get sensor")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.GetResponse{
		Sensor: srv.copySensor(s),
	}

	srv.logger.WithField("id", *s.Id).Debugf("get sensor")

	return res, nil
}

func (srv *metathingsSensordService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	s := storage.Sensor{}

	if name := req.GetName(); name != nil {
		s.Name = &name.Value
	}

	if core_id := req.GetCoreId(); core_id != nil {
		s.CoreId = &core_id.Value
	}

	if entity_name := req.GetEntityName(); entity_name != nil {
		s.EntityName = &entity_name.Value
	}

	if owner_id := req.GetOwnerId(); owner_id != nil {
		s.OwnerId = &owner_id.Value
	}

	if state := req.GetState(); state != sensor_pb.SensorState_SENSOR_STATE_UNKNOWN {
		state_str := srv.sensor_st_psr.ToString(state)
		s.State = &state_str
	}

	ss, err := srv.storage.ListSensors(s)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to list sensors")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListResponse{
		Sensors: srv.copySensors(ss),
	}

	srv.logger.Debugf("list sensors")

	return res, nil
}

func (srv *metathingsSensordService) ListForUser(ctx context.Context, req *pb.ListForUserRequest) (*pb.ListForUserResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	cred := context_helper.Credential(ctx)
	user_id := cred.User.Id

	s := storage.Sensor{}

	if name := req.GetName(); name != nil {
		s.Name = &name.Value
	}

	if core_id := req.GetCoreId(); core_id != nil {
		s.CoreId = &core_id.Value
	}

	if entity_name := req.GetEntityName(); entity_name != nil {
		s.EntityName = &entity_name.Value
	}

	if state := req.GetState(); state != sensor_pb.SensorState_SENSOR_STATE_UNKNOWN {
		state_str := srv.sensor_st_psr.ToString(state)
		s.State = &state_str
	}

	ss, err := srv.storage.ListSensorsForUser(user_id, s)
	if err != nil {
		srv.logger.WithField("user_id", user_id).WithError(err).Errorf("failed to list sensors for user")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &pb.ListForUserResponse{
		Sensors: srv.copySensors(ss),
	}

	srv.logger.WithField("user_id", user_id).Debugf("list sensors for user")

	return res, nil
}

func (srv *metathingsSensordService) subscribe(stm pb.SensordService_SubscribeServer, quit chan interface{}) {
	defer func() {
		quit <- nil
		srv.logger.Debugf("send quit signal to subscribler")
	}()

	subs := make(map[string]hub.Subscriber)

subscrible_loop:
	for {
		reqs, err := stm.Recv()
		if err != nil {
			grpc_helper.HandleGRPCError(srv.logger, err, "failed to recv data from subscriber")
			return
		}

		for _, r := range reqs.Requests {
			switch req := r.Payload.(type) {
			case *pb.SubscribeRequest_SubscribeById:
				err = srv.handle_subscribe_request_subscribe_by_id(stm, subs, req)
			case *pb.SubscribeRequest_UnsubscribeById:
				err = srv.handle_subscribe_request_unsubscribe_by_id(stm, subs, req)
			}

			if err != nil {
				srv.logger.WithError(err).Errorf("failed to handle subscribe request")
				break subscrible_loop
			}
		}
	}

	return
}

func (srv *metathingsSensordService) handle_subscribe_request_subscribe_by_id(stm pb.SensordService_SubscribeServer, subs map[string]hub.Subscriber, req *pb.SubscribeRequest_SubscribeById) error {
	snr_id := req.SubscribeById.GetId().GetValue()
	if _, ok := subs[snr_id]; ok {
		srv.logger.WithField("snr_id", snr_id).Warningf("sensor already in subscribling")
		return nil
	}

	subscriber_opt := opt_helper.NewOption("sensor_id", snr_id)
	sub, err := srv.hub.Subscriber(subscriber_opt)
	if err != nil {
		srv.logger.WithField("snr_id", snr_id).Errorf("failed to get subscribler")
		return nil
	}

	subs[snr_id] = sub
	go func(stm pb.SensordService_SubscribeServer, sub hub.Subscriber) {
		defer func() {
			sub.Close()
			delete(subs, snr_id)
		}()
		for {
			dat, err := sub.Subscribe()
			if err != nil {
				srv.logger.WithError(err).Errorf("failed to subscribe data from subscriber")
				return
			}

			res := &pb.SubscribeResponses{
				Responses: []*pb.SubscribeResponse{
					&pb.SubscribeResponse{Data: dat},
				},
			}

			err = stm.Send(res)
			if err != nil {
				srv.logger.WithError(err).Errorf("failed to send data to subscribe stream")
				return
			}
		}
	}(stm, sub)

	srv.logger.WithField("snr_id", snr_id).Debugf("subscrible data by sensor id")
	return nil
}

func (srv *metathingsSensordService) handle_subscribe_request_unsubscribe_by_id(stm pb.SensordService_SubscribeServer, subs map[string]hub.Subscriber, req *pb.SubscribeRequest_UnsubscribeById) error {
	var sub hub.Subscriber
	var ok bool
	snr_id := req.UnsubscribeById.GetId().GetValue()
	if sub, ok = subs[snr_id]; !ok {
		srv.logger.WithField("snr_id", snr_id).Errorf("subscribing sensor not found")
	}
	sub.Close()

	srv.logger.WithField("snr_id", snr_id).Debugf("unsubscrible sensor by sensor id")
	return nil
}

func (srv *metathingsSensordService) Subscribe(stm pb.SensordService_SubscribeServer) error {
	quit := make(chan interface{})

	go srv.subscribe(stm, quit)

	<-quit

	srv.logger.Infof("subscribe done")
	return nil
}

func (srv *metathingsSensordService) publisher_option(snr storage.Sensor) opt_helper.Option {
	return opt_helper.NewOption(
		"sensor_id", *snr.Id,
		"core_id", *snr.CoreId,
		"entity_name", *snr.EntityName,
		"owner_id", *snr.OwnerId,
	)
}

func (srv *metathingsSensordService) Publish(stm pb.SensordService_PublishServer) error {
	ctx := stm.Context()
	cred := context_helper.Credential(ctx)
	app_cred_id := cred.ApplicationCredential.Id

	s := storage.Sensor{
		ApplicationCredentialId: &app_cred_id,
	}

	ss, err := srv.storage.ListSensors(s)
	if err != nil {
		srv.logger.WithError(err).WithField("application_credential_id", app_cred_id).Errorf("failed to list sensors with application credential id")
		return status.Errorf(codes.Internal, err.Error())
	}

	if len(ss) == 0 {
		srv.logger.WithField("application_credential_id", app_cred_id).Errorf("not registerd sensor")
		return status.Errorf(codes.NotFound, ErrNotRegisteredSensor.Error())
	}

	snr := ss[0]
	publisher, err := srv.hub.Publisher(srv.publisher_option(snr))
	if err != nil {
		srv.logger.WithError(err).WithField("application_credential_id", app_cred_id).Errorf("failed to get publisher")
		return status.Errorf(codes.Internal, err.Error())
	}
	quit := make(chan interface{})

	go func() {
		defer func() {
			publisher.Close()
			srv.logger.WithField("snr_id", *snr.Id).Debugf("close publisher")
			quit <- nil
			srv.logger.WithField("snr_id", *snr.Id).Debugf("send quit signal to publisher")
		}()
		for {
			reqs, err := stm.Recv()
			if err != nil {
				grpc_helper.HandleGRPCError(srv.logger.WithField("snr_id", *snr.Id), err, "failed to recv data from publisher")
				return
			}

			now := protobuf_helper.Now()
			for _, req := range reqs.Requests {
				switch req.Payload.(type) {
				case *pb.PublishRequest_Data:
					dat := req.GetData()
					dat.ArrivedAt = &now
					dat.SensorId = *snr.Id

					if err = publisher.Publish(dat); err != nil {
						srv.logger.WithError(err).Warningf("failed to publish data to hub")
					}
				}
			}
		}
	}()

	<-quit

	srv.logger.WithField("snr_id", *snr.Id).Infof("publish done")
	return nil
}

func NewSensordService(opt ...ServiceOptions) (*metathingsSensordService, error) {
	opts := defaultServiceOptions
	for _, o := range opt {
		o(&opts)
	}

	logger, err := log_helper.NewLogger("sensord", opts.logLevel)
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
		logger.WithError(err).Errorf("failed to new client factory")
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

	opts.hub.Set("logger", logger)
	hub, err := hub.NewHub(opts.hub)
	if err != nil {
		logger.WithError(err).Errorf("failed to new pubsub hub")
		return nil, err
	}

	tk_vdr := token_helper.NewTokenValidator(app_cred_mgr, cli_fty, logger)

	srv := &metathingsSensordService{
		cli_fty:       cli_fty,
		sensor_st_psr: state_helper.NewSensorStateParser(),
		app_cred_mgr:  app_cred_mgr,
		opts:          opts,
		logger:        logger,
		storage:       storage,
		tk_vdr:        tk_vdr,
		hub:           hub,
	}

	logger.Debugf("new sensord service")

	return srv, nil
}
