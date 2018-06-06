package metathings_motor_service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	log_helper "github.com/nayotta/metathings/pkg/common/log"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	mt_plugin "github.com/nayotta/metathings/pkg/core/plugin"
	driver "github.com/nayotta/metathings/pkg/motor/driver"
	state_helper "github.com/nayotta/metathings/pkg/motor/state"
	pb "github.com/nayotta/metathings/pkg/proto/motor"
)

type metathingsMotorService struct {
	mt_plugin.CoreService
	opt     opt_helper.Option
	logger  log.FieldLogger
	cli_fty *client_helper.ClientFactory
	mtr_mgr *MotorManager

	motor_st_psr  state_helper.MotorStateParser
	motor_dir_psr state_helper.MotorDirectionParser
}

func (srv *metathingsMotorService) copyMotor(m Motor) *pb.Motor {
	mtr := m.Driver.Show()
	return &pb.Motor{
		Name:      m.Name,
		State:     srv.motor_st_psr.ToValue(mtr.State.ToString()),
		Direction: srv.motor_dir_psr.ToValue(mtr.Direction.ToString()),
		Speed:     mtr.Speed,
	}
}

func (srv *metathingsMotorService) copyMotors(ms []Motor) []*pb.Motor {
	mtrs := make([]*pb.Motor, 0, len(ms))
	for _, m := range ms {
		mtrs = append(mtrs, srv.copyMotor(m))
	}
	return mtrs
}

func (srv *metathingsMotorService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	mtrs := srv.copyMotors(srv.mtr_mgr.ListMotors())
	res := &pb.ListResponse{
		Motors: mtrs,
	}
	srv.logger.WithField("motors", mtrs).Debugf("list motors")

	return res, nil
}

func (srv *metathingsMotorService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	m, err := srv.mtr_mgr.GetMotor(req.Name.Value)
	if err != nil {
		srv.logger.WithError(err).WithField("name", req.Name.Value).Errorf("failed to get motor")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	res := &pb.GetResponse{
		Motor: srv.copyMotor(m),
	}
	srv.logger.WithField("motor", m).Debugf("get motor")

	return res, nil
}

func (srv *metathingsMotorService) Stream(stream pb.MotorService_StreamServer) error {
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

	<-quit
	srv.logger.Infof("stream closed")

	return nil
}

func (srv *metathingsMotorService) handleStreamRequest(stream pb.MotorService_StreamServer, req *pb.StreamRequest) {
	switch req.Payload.(type) {
	case *pb.StreamRequest_Ping:
		srv.handleStreamRequest_ping(stream, req)
	case *pb.StreamRequest_SetState:
		srv.handleStreamRequest_setState(stream, req)
	case *pb.StreamRequest_SetDirection:
		srv.handleStreamRequest_setDirection(stream, req)
	case *pb.StreamRequest_SetSpeed:
		srv.handleStreamRequest_setSpeed(stream, req)
	default:
		srv.logger.Warningf("unsupported stream request type")
	}
}

func (srv *metathingsMotorService) handleStreamRequest_ping(stream pb.MotorService_StreamServer, req *pb.StreamRequest) {
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

func (srv *metathingsMotorService) handleStreamRequest_setState(stream pb.MotorService_StreamServer, req *pb.StreamRequest) {
	req_setState := req.GetSetState()
	m, err := srv.mtr_mgr.GetMotor(req_setState.Name.Value)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to get motor")
		return
	}

	st := driver.StateFromValue(int32(req_setState.State))
	mtr, err := m.Driver.Turn(st)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to turn motor state")
		return
	}
	srv.logger.WithField("motor", mtr).Debugf("motor state turning")
}

func (srv *metathingsMotorService) handleStreamRequest_setDirection(stream pb.MotorService_StreamServer, req *pb.StreamRequest) {
	req_setDirection := req.GetSetDirection()
	m, err := srv.mtr_mgr.GetMotor(req_setDirection.Name.Value)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to get motor")
		return
	}

	dir := driver.DirectionFromValue(int32(req_setDirection.Direction))
	mtr, err := m.Driver.SetDirection(dir)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to set motor direction")
		return
	}
	srv.logger.WithField("motor", mtr).Debugf("motor set direction")
}

func (srv *metathingsMotorService) handleStreamRequest_setSpeed(stream pb.MotorService_StreamServer, req *pb.StreamRequest) {
	req_setSpeed := req.GetSetSpeed()
	m, err := srv.mtr_mgr.GetMotor(req_setSpeed.Name.Value)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to get motor")
		return
	}

	mtr, err := m.Driver.SetSpeed(req_setSpeed.Speed.Value)
	if err != nil {
		srv.logger.WithError(err).Errorf("failed to set motor speed")
		return
	}
	srv.logger.WithField("motor", mtr).Debugf("motor set speed")
}

func NewMotorService(opt opt_helper.Option) (*metathingsMotorService, error) {
	opt.Set("service_name", "motor")

	logger, err := log_helper.NewLogger("motord", opt.GetString("log.level"))
	if err != nil {
		return nil, err
	}

	cli_fty_cfgs := client_helper.NewDefaultServiceConfigs(opt.GetString("metathings.address"))
	cli_fty_cfgs[client_helper.AGENTD_CONFIG] = client_helper.ServiceConfig{opt.GetString("agent.address")}
	cli_fty, err := client_helper.NewClientFactory(
		cli_fty_cfgs,
		client_helper.WithInsecureOptionFunc(),
	)
	if err != nil {
		return nil, err
	}

	opt.Set("logger", logger)
	mtr_mgr, err := NewMotorManager(opt)
	if err != nil {
		return nil, err
	}
	logger.Debugf("new motor manager")

	srv := &metathingsMotorService{
		opt:     opt,
		logger:  logger,
		cli_fty: cli_fty,
		mtr_mgr: mtr_mgr,

		motor_st_psr:  state_helper.MOTOR_STATE_PARSER,
		motor_dir_psr: state_helper.MOTOR_DIRECTION_PARSER,
	}

	srv.CoreService = mt_plugin.MakeCoreService(srv.opt, srv.logger, srv.cli_fty)

	return srv, nil
}
