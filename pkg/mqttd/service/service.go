package metathingsmqttdservice

import (
	"context"
	"math/rand"
	"time"

	app_cred_mgr "github.com/nayotta/metathings/pkg/common/application_credential_manager"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	policy_helper "github.com/nayotta/metathings/pkg/common/policy"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	mt_plugin "github.com/nayotta/metathings/pkg/cored/plugin"
	identityd_policy "github.com/nayotta/metathings/pkg/identityd2/policy"
	connection "github.com/nayotta/metathings/pkg/mqttd/connection"
	storage "github.com/nayotta/metathings/pkg/mqttd/storage"
	cored_pb "github.com/nayotta/metathings/pkg/proto/cored"
	identityd_pb "github.com/nayotta/metathings/pkg/proto/identityd2"
	pb "github.com/nayotta/metathings/pkg/proto/mqttd"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
)

// MetathingsDevicedServiceOption MetathingsMqttdServiceOption struct
type MetathingsMqttdServiceOption struct {
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
	// GRPCKEEPALIVE GRPCKEEPALIVE
	GRPCKEEPALIVE = grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                2 * time.Second,
		Timeout:             10 * time.Second,
		PermitWithoutStream: true,
	})
)

// MetathingsMqttdService MetathingsMqttdService struct
type MetathingsMqttdService struct {
	grpc_helper.AuthorizationTokenParser
	appCredMgr app_cred_mgr.ApplicationCredentialManager

	tknr     token_helper.Tokener
	cliFty   *client_helper.ClientFactory
	opt      *MetathingsMqttdServiceOption
	logger   log.FieldLogger
	storage  storage.Storage
	enforcer identityd_policy.Enforcer
	vdr      token_helper.TokenValidator
	cc       connection.MqttBridge

	heartbeatSession uint64
	dispatchers      map[string]mt_plugin.DispatcherPlugin
}

func (sev *MetathingsMqttdService) getDeviceByContext(ctx context.Context) (*storage.Device, error) {
	var tkn *identityd_pb.Token
	var dev *storage.Device
	var err error

	tkn = context_helper.ExtractToken(ctx)

	if dev, err = sev.storage.GetDevice(tkn.Entity.Id); err != nil {
		return nil, err
	}

	return dev, nil
}

func (sev *MetathingsMqttdService) enforce(ctx context.Context, obj, act string) error {
	var err error

	tkn := context_helper.ExtractToken(ctx)

	var groups []string
	for _, g := range tkn.Groups {
		groups = append(groups, g.Id)
	}

	if err = sev.enforcer.Enforce(tkn.Domain.Id, groups, tkn.Entity.Id, obj, act); err != nil {
		if err == identityd_policy.ErrPermissionDenied {
			sev.logger.WithFields(log.Fields{
				"subject": tkn.Entity.Id,
				"domain":  tkn.Domain.Id,
				"groups":  groups,
				"object":  obj,
				"action":  act,
			}).Warningf("denied to do #action")
			return status.Errorf(codes.PermissionDenied, err.Error())
		}
		sev.logger.WithError(err).Errorf("failed to enforce")
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}

func (sev *MetathingsMqttdService) validateChain(providers []interface{}, invokers []interface{}) error {
	defaultInvokers := []interface{}{policy_helper.ValidateValidator}
	invokers = append(defaultInvokers, invokers...)
	if err := policy_helper.ValidateChain(
		providers,
		invokers,
	); err != nil {
		sev.logger.WithError(err).Warningf("failed to validate request data")
		return status.Errorf(codes.InvalidArgument, err.Error())
	}

	return nil
}

func (sev *MetathingsMqttdService) isIgnoreMethod(md *grpc_helper.MethodDescription) bool {
	return false
}

// AuthFuncOverride AuthFuncOverride
func (sev *MetathingsMqttdService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	//debug here
	return ctx, nil

	var tkn *identityd_pb.Token
	var tknTxt string
	var newCtx context.Context
	var err error
	var md *grpc_helper.MethodDescription

	if md, err = grpc_helper.ParseMethodDescription(fullMethodName); err != nil {
		sev.logger.WithError(err).Warningf("failed to parse method description")
		return ctx, err
	}

	if sev.isIgnoreMethod(md) {
		return ctx, nil
	}

	if tknTxt, err = sev.GetTokenFromContext(ctx); err != nil {
		sev.logger.WithError(err).Warningf("failed to get token from context")
		return ctx, err
	}

	if tkn, err = sev.vdr.Validate(tknTxt); err != nil {
		sev.logger.WithError(err).Warningf("failed to validate token in identity service")
		return ctx, err
	}

	newCtx = context.WithValue(ctx, "token", tkn)

	sev.logger.WithFields(log.Fields{
		"method":    md.Method,
		"entity_id": tkn.Entity.Id,
		"domain_id": tkn.Domain.Id,
	}).Debugf("authorize token")

	return newCtx, nil
}

func (sev *MetathingsMqttdService) getDispatcherPlugin(name string, serviceName string) (mt_plugin.DispatcherPlugin, bool) {
	dp, ok := sev.dispatchers[name]
	return dp, ok
}

// ServeOnStream ServeOnStream
func (sev *MetathingsMqttdService) ServeOnStream() error {
	token := sev.appCredMgr.GetToken()
	ctx := context_helper.WithToken(context.Background(), token)

	cli, cfn, err := sev.cliFty.NewCoredServiceClient(GRPCKEEPALIVE)
	if err != nil {
		sev.logger.WithError(err).Errorf("failed to dial to metathings service")
		return err
	}
	defer cfn()

	stream, err := cli.Stream(ctx)
	if err != nil {
		sev.logger.WithError(err).Errorf("failed to stream to core service")
		return err
	}
	sev.logger.Debugf("connect to core service on streaming")

	return sev.serveOnStream(stream)
}

func (sev *MetathingsMqttdService) serveOnStream(stream cored_pb.CoredService_StreamClient) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			err = sev.handleGRPCError(err, "failed to recv data from core")
			return err
		}

		ctx := stream.Context()
		res, err := sev.dispatch(ctx, req)
		if err != nil {
			sev.logger.WithError(err).Errorf("failed to dispatch")
			continue
		}

		err = stream.Send(res)
		if err != nil {
			err = sev.handleGRPCError(err, "failed to send data to entity")
			return err
		}
	}
}

func (sev *MetathingsMqttdService) dispatchUser(ctx context.Context, req *cored_pb.StreamRequest) (*cored_pb.StreamResponse, error) {
	payload, ok := req.Payload.(*cored_pb.StreamRequest_UnaryCall)
	if !ok {
		return nil, ErrUnsupportPayloadType
	}

	name := payload.UnaryCall.Name.Value
	serviceName := payload.UnaryCall.ServiceName.Value
	methodName := payload.UnaryCall.MethodName.Value
	reqValue := payload.UnaryCall.Value

	dp, ok := sev.getDispatcherPlugin(name, serviceName)
	if !ok {
		return nil, ErrPluginNotFound
	}

	res, err := dp.UnaryCall(methodName, ctx, reqValue)
	if err != nil {
		errRes := &cored_pb.StreamResponse{
			SessionId:   req.SessionId.Value,
			MessageType: req.MessageType,
			Payload: &cored_pb.StreamResponse_Err{
				Err: &cored_pb.StreamErrorResponsePayload{
					Name:        name,
					ServiceName: serviceName,
					MethodName:  methodName,
					Context:     err.Error(),
				},
			},
		}
		return errRes, nil
	}

	res1 := &cored_pb.StreamResponse{
		SessionId:   req.SessionId.Value,
		MessageType: req.MessageType,
		Payload: &cored_pb.StreamResponse_UnaryCall{
			UnaryCall: &cored_pb.UnaryCallResponsePayload{
				Name:        name,
				ServiceName: serviceName,
				MethodName:  methodName,
				Value:       res,
			},
		},
	}

	return res1, nil
}

func (sev *MetathingsMqttdService) dispatch(ctx context.Context, req *cored_pb.StreamRequest) (*cored_pb.StreamResponse, error) {
	switch req.MessageType {
	case cored_pb.StreamMessageType_STREAM_MESSAGE_TYPE_USER:
		return sev.dispatchUser(ctx, req)
	default:
		return nil, ErrUnsupportMessageType
	}
}

// NewMetathingsMqttdService NewMetathingsMqttdService
func NewMetathingsMqttdService(
	opt *MetathingsMqttdServiceOption,
	logger log.FieldLogger,
	storage storage.Storage,
	//enforcer identityd_policy.Enforcer,
	//vdr token_helper.TokenValidator,
	//tknr token_helper.Tokener,
	cliFty *client_helper.ClientFactory,
	cc connection.MqttBridge,
) (pb.MqttdServiceServer, error) {
	appCredMgr, err := app_cred_mgr.NewApplicationCredentialManager(
		cliFty,
		opt.application_credential_id,
		opt.application_credential_secret,
	)

	if err != nil {
		log.WithError(err).Errorf("failed to NewApplicationCredentialManager")
	}

	return &MetathingsMqttdService{
		opt:        opt,
		appCredMgr: appCredMgr,
		logger:     logger,
		storage:    storage,
		//debug here
		//enforcer:         enforcer,
		//vdr:              vdr,
		//tknr:             tknr,
		cliFty:           cliFty,
		cc:               cc,
		heartbeatSession: rand.Uint64(),
	}, nil
}
