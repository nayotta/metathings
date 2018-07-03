package metathings_servo_service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mt_plugin "github.com/nayotta/metathings/pkg/cored/plugin"
	pb "github.com/nayotta/metathings/pkg/proto/servo"
	driver "github.com/nayotta/metathings/pkg/servo/driver"
	state_helper "github.com/nayotta/metathings/pkg/servo/state"
)

type metathingsServoService struct {
	mt_plugin.CoreService
	opt     opt_helper.Option
	logger  log.FieldLogger
	cli_fty *client_helper.ClientFactory
	srv_mgr *ServoManager

	servo_st_psr state_helper.ServoStateParser
}

func (srv *metathingsServoService) copyServo(s Servo) *pb.Servo {
	sv := s.Driver.Show()
	return &pb.Servo{
		Name:  s.Name,
		State: srv.servo_st_psr.ToValue(sv.State.ToString()),
		Angle: sv.Angle,
	}
}

func (srv *metathingsServoService) copyServos(ss []Servo) []*pb.Servo {
	svs := make([]*pb.Servo, 0, len(ss))
	for _, s := range ss {
		svs = append(svs, srv.copyServo(s))
	}
	return svs
}

func (srv *metathingsServoService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	svs := srv.copyServos(srv.srv_mgr.ListServos())
	res := &pb.ListResponse{
		Servos: svs,
	}
	srv.logger.WithField("servos", svs).Debugf("list servos")

	return res, nil
}

func (srv *metathingsServoService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	err := req.Validate()
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to validate request data")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	s, err := srv.srv_mgr.GetServo(req.Name.Value)
	if err != nil {
		srv.logger.WithError(err).WithField("name", req.Name.Value).Errorf("failed to get servo")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	res := &pb.GetResponse{
		Servo: srv.copyServo(s),
	}
	srv.logger.WithField("servo", s).Debugf("get servo")

	return res, nil
}

func (srv *metathingsServoService) Stream(stream pb.ServoService_StreamServer) error {
	quit := make(chan interface{})

	srv.logger.Infof("stream begin")

	go func() {
		defer func() { quit <- nil }()
		for {
			reqs, err := stream.Recv()
			if err != nil {
				srv.handleGRPCError(err, "failed to recv data from agent")
				return
			}

			for _, req := range reqs.Requests {
				go srv.handleStreamRequest(stream, req)
			}
		}
	}()

	return status.Errorf(codes.Unimplemented, "unimplemented")
}

func (srv *metathingsServoService) handleStreamRequest(stream pb.ServoService_StreamServer, req *pb.StreamRequest) {
	switch req.Payload.(type) {
	case *pb.StreamRequest_Ping:
		srv.handleStreamRequest_ping(stream, req)
	case *pb.StreamRequest_SetAngle:
		srv.handleStreamRequest_setAngle(stream, req)
	default:
		srv.logger.Warningf("unsupported stream request type")
	}
}

func (srv *metathingsServoService) handleStreamRequest_ping(stream pb.ServoService_StreamServer, req *pb.StreamRequest) {
	req_ping := req.GetPing()
	res := &pb.StreamResponse{
		Session: req.Session.Value,
		Payload: &pb.StreamResponse_Ping{
			Ping: &pb.StreamPingResponse{
				Timestamp: req_ping.Timestamp,
			},
		},
	}

	err := stream.Send(res)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to send ping response to agent")
		return
	}

	srv.logger.WithFields(log.Fields{
		"session":   req.Session.Value,
		"timestamp": req_ping.Timestamp,
	}).Debugf("send ping response to agent")
}

func (srv *metathingsServoService) handleStreamRequest_setState(stream pb.ServoService_StreamServer, req *pb.StreamRequest) {
	req_setState := req.GetSetState()
	s, err := srv.srv_mgr.GetServo(req_setState.Name.Value)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to get servo")
		return
	}

	st := driver.StateFromValue(int32(req_setState.State))
	sv, err := s.Driver.Turn(st)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to turn servo state")
		return
	}
	srv.logger.WithField("servo", sv).Debugf("servo state turning")
}

func (srv *metathingsServoService) handleStreamRequest_setAngle(stream pb.ServoService_StreamServer, req *pb.StreamRequest) {
	req_setAngle := req.GetSetAngle()
	s, err := srv.srv_mgr.GetServo(req_setAngle.Name.Value)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to get servo")
		return
	}

	sv, err := s.Driver.SetAngle(req_setAngle.Angle.Value)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to set servo angle")
		return
	}
	srv.logger.WithField("servo", sv).Debugf("servo set angle")
}

func (srv *metathingsServoService) Close() {
	var err error
	for _, sv := range srv.srv_mgr.ListServos() {
		if err = sv.Driver.Close(); err != nil {
			srv.logger.WithField("name", sv.Name).WithError(err).Debugf("failed to close servo driver")
		}
	}
	srv.logger.Debugf("service closed")
}

func NewServoService(opt opt_helper.Option) (*metathingsServoService, error) {
	opt.Set("service_name", "servo")

	logger, err := log_helper.NewLogger("servo", opt.GetString("log.level"))
	if err != nil {
		return nil, err
	}

	cli_fty_cfgs := client_helper.NewDefaultServiceConfigs(opt.GetString("metathings.address"))
	cli_fty_cfgs[client_helper.AGENT_CONFIG] = client_helper.ServiceConfig{opt.GetString("agent.address")}
	cli_fty, err := client_helper.NewClientFactory(
		cli_fty_cfgs,
		client_helper.WithInsecureOptionFunc(),
	)
	if err != nil {
		return nil, err
	}

	opt.Set("logger", logger)
	srv_mgr, err := NewServoManager(opt)
	if err != nil {
		return nil, err
	}
	logger.Debugf("new servo manager")

	srv := &metathingsServoService{
		opt:     opt,
		logger:  logger,
		cli_fty: cli_fty,
		srv_mgr: srv_mgr,

		servo_st_psr: state_helper.SERVO_STATE_PARSER,
	}
	return srv, nil
}
