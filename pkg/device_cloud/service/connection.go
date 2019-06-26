package metathings_device_cloud_service

import (
	"context"
	"math"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	protobuf_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	session_helper "github.com/nayotta/metathings/pkg/common/session"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	storage "github.com/nayotta/metathings/pkg/device_cloud/storage"
	state_pb "github.com/nayotta/metathings/pkg/proto/constant/state"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (s *MetathingsDeviceCloudService) get_device_by_module_id(mdl_id string) (*pb.Device, error) {
	cli, cfn, err := s.cli_fty.NewDevicedServiceClient()
	if err != nil {
		return nil, err
	}
	defer cfn()

	req := &pb.GetDeviceByModuleRequest{
		Module: &pb.OpModule{
			Id: &wrappers.StringValue{Value: mdl_id},
		},
	}

	res, err := cli.GetDeviceByModule(s.context(), req)
	if err != nil {
		return nil, err
	}

	return res.GetDevice(), nil
}

func (s *MetathingsDeviceCloudService) try_to_build_device_connection_by_module_id(mdl_id string) {
	dev, err := s.get_device_by_module_id(mdl_id)
	if err != nil {
		s.get_logger().WithError(err).Errorf("failed to get device in deviced")
		return
	}
	dev_id := dev.Id

	err = s.storage.IsConnected(s.get_session_id(), dev_id)
	switch err {
	case nil:
		// this instance is maintaining device connection, ignore
	case storage.ErrConnectedByOtherDeviceCloud:
		// other instance is maintaining device connection, ignore
	case storage.ErrNotConnected:
		// try to build device connection in current instance

		// mark down instance session for the device
		err = s.storage.ConnectDevice(s.get_session_id(), dev_id)
		if err != nil {
			s.get_logger().WithError(err).Debugf("failed to lock connection in current instance, maybe locked by other instance")
			return
		}

		err = s.build_device_connection(dev)
		if err != nil {
			// unmark instance session on failed
			s.storage.UnconnectDevice(s.get_session_id(), dev_id)
			s.get_logger().WithError(err).Errorf("failed to build device connection")
			return
		}

		s.get_logger().WithFields(log.Fields{
			"module": mdl_id,
			"device": dev_id,
		}).Infof("build device connection")
	default:
		s.get_logger().WithError(err).Debugf("failed to get device connection status")
	}
}

func (s *MetathingsDeviceCloudService) build_device_connection(dev *pb.Device) error {
	dc, err := NewDeviceConnection(dev, s.storage, s.cli_fty, s.logger)
	if err != nil {
		return err
	}

	err = dc.Start()
	if err != nil {
		return err
	}

	return nil
}

type DeviceConnectionOption struct {
	DeviceCloud struct {
		Session struct {
			Id string
		}
	}
	Session struct {
		Startup    int32
		Major      int32
		Connection int64
	}
	Device struct {
		Id      string
		Modules []struct {
			Id string
		}
	}

	Config struct {
		HeartbeatInterval time.Duration
		HeartbeatTimeout  time.Duration

		MaxReconnect         int
		MinReconnectInterval time.Duration
		MaxReconnectInterval time.Duration
	}
}

type DeviceConnection struct {
	opt     *DeviceConnectionOption
	storage storage.Storage
	cli_fty *client_helper.ClientFactory
	logger  log.FieldLogger
	tknr    token_helper.Tokener
	stm     pb.DevicedService_ConnectClient
	close   client_helper.CloseFn

	stm_wg      sync.WaitGroup
	stm_wg_once sync.Once
}

func (dc *DeviceConnection) context_with_session_and_device() context.Context {
	return context_helper.NewOutgoingContext(
		context.TODO(),
		context_helper.WithTokenOp(dc.tknr.GetToken()),
		context_helper.WithSessionOp(dc.opt.Session.Connection),
		context_helper.WithDeviceOp(dc.opt.Device.Id),
	)
}

func (dc *DeviceConnection) context() context.Context {
	return context_helper.WithToken(context.TODO(), dc.tknr.GetToken())
}

func (dc *DeviceConnection) Start() error {
	go dc.main_loop()
	go dc.heartbeat_loop()
	go dc.ping_loop()

	dc.logger.WithField("device", dc.opt.Device.Id).Debugf("device connected")

	return nil
}

func (dc *DeviceConnection) Stop() error {
	dc.clear()

	return nil
}

func (dc *DeviceConnection) clear() {
	var err error

	if dc.close != nil {
		err = dc.close()
		if err != nil {
			dc.logger.WithError(err).Warningf("failed to close deviced client connection")
		}
	}

	err = dc.storage.UnconnectDevice(dc.opt.DeviceCloud.Session.Id, dc.opt.Device.Id)
	if err != nil {
		dc.logger.WithError(err).Warningf("failed to unconnect device in storage")
	}
}

func (dc *DeviceConnection) main_loop() {
	rc := 0
	rc_tvl := dc.opt.Config.MinReconnectInterval
	defer dc.clear()

	for {
		err := dc.storage.IsConnected(dc.opt.DeviceCloud.Session.Id, dc.opt.Device.Id)
		if err != nil {
			dc.logger.WithError(err).Warningf("device connection is not maintaining by this instance")
			return
		}

		if rc > dc.opt.Config.MaxReconnect {
			dc.logger.Warningf("max reconnect to connect deviced")
			return
		}

		err = dc.internal_main_loop()
		if err != nil {
			rc_tvl = time.Duration(math.Min(float64(rc_tvl*2), float64(dc.opt.Config.MaxReconnectInterval)))
			rc += 1
		} else {
			rc_tvl = dc.opt.Config.MinReconnectInterval
			rc = 0
		}
		time.Sleep(rc_tvl)

	}
}

func (dc *DeviceConnection) internal_main_loop() error {
	var cli pb.DevicedServiceClient
	var req *pb.ConnectRequest
	var err error

	cli, dc.close, err = dc.cli_fty.NewDevicedServiceClient()
	if err != nil {
		return err
	}
	defer func() {
		if dc.close != nil {
			dc.close()
		}
		dc.close = nil
	}()

	ctx := dc.context_with_session_and_device()
	dc.stm, err = cli.Connect(ctx)
	if err != nil {
		return err
	}

	for {
		if req, err = dc.stm.Recv(); err != nil {
			dc.logger.WithError(err).Warningf("failed to recv message from connection stream")
			return nil
		}

		dc.logger.WithFields(log.Fields{
			"session": req.GetSessionId().GetValue(),
			"kind":    req.GetKind(),
		}).Debugf("rcev msg")

		go dc.handle(req)
	}
}

func (dc *DeviceConnection) heartbeat_loop() {
	dc.stm_wg.Wait()
	for {
		go dc.heartbeat_loop_once()
		time.Sleep(dc.opt.Config.HeartbeatInterval)
	}
}

func (dc *DeviceConnection) heartbeat_loop_once() {
	cli, cfn, err := dc.cli_fty.NewDevicedServiceClient()
	if err != nil {
		dc.logger.WithError(err).Warningf("failed to connect to deviced service")
		return
	}
	defer cfn()

	now := time.Now()
	pb_now := protobuf_helper.FromTime(now)
	pb_mdls := []*pb.OpModule{}
	for _, mdl := range dc.opt.Device.Modules {
		var stat state_pb.ModuleState
		mdl_id := mdl.Id
		hbt, err := dc.storage.GetHeartbeatAt(mdl_id)
		if err != nil {
			hbt = time.Time{}
			dc.logger.WithError(err).Warningf("failed to get heartbeat time in storage")
		}
		pb_hbt := protobuf_helper.FromTime(hbt)

		if time.Now().Sub(hbt) < dc.opt.Config.HeartbeatTimeout {
			stat = state_pb.ModuleState_MODULE_STATE_ONLINE
		} else {
			stat = state_pb.ModuleState_MODULE_STATE_OFFLINE
		}

		pb_mdl := &pb.OpModule{
			Id:          &wrappers.StringValue{Value: mdl_id},
			HeartbeatAt: &pb_hbt,
			State:       stat,
		}

		pb_mdls = append(pb_mdls, pb_mdl)
	}

	req := &pb.HeartbeatRequest{
		Device: &pb.OpDevice{
			Id:          &wrappers.StringValue{Value: dc.opt.Device.Id},
			HeartbeatAt: &pb_now,
			Modules:     pb_mdls,
		},
		StartupSession: &wrappers.Int32Value{Value: dc.opt.Session.Startup},
	}

	_, err = cli.Heartbeat(dc.context(), req)
	// TODO(Peer): should stop connect after failed to heartbeat
	if err != nil {
		dc.logger.WithError(err).Debugf("failed to heartbeat")
		return
	}

	dc.logger.WithFields(log.Fields{
		"device":       dc.opt.Device.Id,
		"heartbeat_at": now,
	}).Debugf("heartbeat")

	panic("unimplemented")
}

func (dc *DeviceConnection) ping_loop() {
	panic("unimplemented")
}

func NewDeviceConnection(dev *pb.Device, storage storage.Storage, cli_fty *client_helper.ClientFactory, logger log.FieldLogger) (*DeviceConnection, error) {
	opt := &DeviceConnectionOption{}

	opt.Session.Startup = session_helper.GenerateStartupSession()
	opt.Session.Major = session_helper.GenerateMajorSession()
	opt.Session.Connection = session_helper.NewSession(opt.Session.Startup, opt.Session.Major)
	opt.Device.Id = dev.Id
	for _, mdl := range dev.Modules {
		opt.Device.Modules = append(opt.Device.Modules, struct{ Id string }{Id: mdl.Id})
	}
	opt.Config.HeartbeatInterval = 13 * time.Second
	opt.Config.MaxReconnect = 7
	opt.Config.MinReconnectInterval = 3 * time.Second
	opt.Config.MaxReconnectInterval = 17 * time.Second

	return &DeviceConnection{
		opt:     opt,
		storage: storage,
		cli_fty: cli_fty,
		logger:  logger,
	}, nil
}
