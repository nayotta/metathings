package metathings_device_cloud_service

import (
	"context"
	"math"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/goiiot/libmqtt"
	"github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	protobuf_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	session_helper "github.com/nayotta/metathings/pkg/common/session"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	component "github.com/nayotta/metathings/pkg/component"
	storage "github.com/nayotta/metathings/pkg/device_cloud/storage"
	mosquitto_service "github.com/nayotta/metathings/pkg/plugin/mosquitto/service"
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

func (s *MetathingsDeviceCloudService) try_to_build_device_connection(dev *pb.Device) {
	dev_id := dev.Id

	cur_sess := s.get_session_id()
	sess, err := s.storage.GetDeviceConnectSession(dev_id)
	if err == nil {
		if sess == cur_sess {
			if err = s.storage.SetDeviceConnectSession(dev_id, cur_sess); err != nil {
				s.get_logger().WithError(err).Warningf("failed to refresh device connect session")
			}
		}
		// else {
		//   other device cloud is maintaining device connection, ignore
		// }

	} else if err == storage.ErrNotConnected {
		// try to build device connection in current instance

		// mark down instance session for the device
		err = s.storage.SetDeviceConnectSession(dev_id, s.get_session_id())
		if err != nil {
			s.get_logger().WithError(err).Debugf("failed to lock connection in current instance, maybe locked by other instance")
			return
		}

		err = s.build_device_connection(dev)
		if err != nil {
			// unmark instance session on failed
			s.storage.UnsetDeviceConnectSession(dev_id, s.get_session_id())
			s.get_logger().WithError(err).Errorf("failed to build device connection")
			return
		}

		s.get_logger().WithFields(log.Fields{
			"device": dev_id,
		}).Infof("build device connection")
	} else {
		s.get_logger().WithError(err).Debugf("failed to get device connection status")
	}
}

func (s *MetathingsDeviceCloudService) build_device_connection(dev *pb.Device) error {
	dc, err := NewDeviceConnection(
		"device", dev,
		"storage", s.storage,
		"client_factory", s.cli_fty,
		"logger", s.logger,
		"tokener", s.tknr,
		"mqtt_address", s.opt.Connection.Mqtt.Address,
		"mqtt_username", s.opt.Credential.Id,
		"mqtt_password", mosquitto_service.ParseMosquittoPluginPassword(s.opt.Credential.Id, s.opt.Credential.Secret),
		"device_cloud_session", s.opt.Session.Id,
	)
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
		Connection struct {
			Mqtt struct {
				Address  string
				Username string
				Password string
			}
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

		PingInterval time.Duration
	}
}

type DeviceConnection struct {
	info    *pb.Device
	opt     *DeviceConnectionOption
	storage storage.Storage
	cli_fty *client_helper.ClientFactory
	logger  log.FieldLogger
	tknr    token_helper.Tokener
	stm     pb.DevicedService_ConnectClient
	close   client_helper.CloseFn

	closed         bool
	stm_wg         sync.WaitGroup
	stm_wg_once    sync.Once
	stm_rwmtx      sync.RWMutex
	mdl_info_cache map[string]*pb.Module
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

func (dc *DeviceConnection) get_stream() (pb.DevicedService_ConnectClient, func()) {
	dc.stm_rwmtx.RLock()
	return dc.stm, dc.stm_rwmtx.RUnlock
}

func (dc *DeviceConnection) get_module_info_by_name(name string) *pb.Module {
	var mdl *pb.Module
	var ok bool

	if mdl, ok = dc.mdl_info_cache[name]; ok {
		return mdl
	}

	for _, mdl = range dc.info.Modules {
		if mdl.Name == name {
			dc.mdl_info_cache[name] = mdl
			return mdl
		}
	}

	return nil
}

func (dc *DeviceConnection) Start() error {
	dc.stm_wg.Add(1)
	go dc.main_loop()
	go dc.heartbeat_loop()
	go dc.ping_loop()

	dc.logger.WithField("device", dc.opt.Device.Id).Debugf("device connected")

	return nil
}

func (dc *DeviceConnection) Stop() error {
	dc.clear()
	dc.closed = true

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

	err = dc.storage.UnsetDeviceConnectSession(dc.opt.Device.Id, dc.opt.DeviceCloud.Session.Id)
	if err != nil {
		dc.logger.WithError(err).Warningf("failed to unconnect device in storage")
	}
}

func (dc *DeviceConnection) is_closed() bool {
	return dc.closed
}

func (dc *DeviceConnection) main_loop() {
	rc := 0
	rc_tvl := dc.opt.Config.MinReconnectInterval
	defer func() {
		dc.Stop()
		dc.logger.WithFields(log.Fields{
			"device":  dc.opt.Device.Id,
			"session": dc.opt.Session.Connection,
		}).Debugf("device connection main loop exited")
	}()

	for {
		if dc.is_closed() {
			return
		}

		cur_sess, err := dc.storage.GetDeviceConnectSession(dc.opt.Device.Id)
		if err != nil || cur_sess != dc.opt.DeviceCloud.Session.Id {
			dc.logger.WithFields(log.Fields{
				"device_cloud_session": dc.opt.DeviceCloud.Session.Id,
				"current_crrent":       cur_sess,
			}).WithError(err).Warningf("device connection is not maintaining by this instance")
			return
		}

		if rc > dc.opt.Config.MaxReconnect {
			dc.logger.Warningf("max reconnect to connect deviced")
			return
		}

		err = dc.internal_main_loop()
		if err != nil {
			rc_tvl = time.Duration(math.Min(float64(rc_tvl*2), float64(dc.opt.Config.MaxReconnectInterval)))
			rc++
			dc.logger.WithError(err).Debugf("internal main loop break")
		} else {
			rc_tvl = dc.opt.Config.MinReconnectInterval
			rc = 0
		}
		dc.logger.WithFields(log.Fields{
			"retry":          rc,
			"retry_interval": rc_tvl,
		}).Debugf("restart main loop")
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
	dc.stm_rwmtx.Lock()
	dc.stm, err = cli.Connect(ctx)
	dc.stm_rwmtx.Unlock()
	if err != nil {
		return err
	}
	dc.stm_wg_once.Do(dc.stm_wg.Done)
	dc.logger.Debugf("internal main loop started")

	for {
		stm, unlock := dc.get_stream()
		req, err = stm.Recv()
		unlock()
		if err != nil {
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
	defer dc.logger.WithFields(log.Fields{
		"device":  dc.opt.Device.Id,
		"session": dc.opt.Session.Connection,
	}).Debugf("device connection heartbeat loop exited")

	for {
		if dc.is_closed() {
			return
		}

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

	any_module_alive := false
	pb_mdls := []*pb.OpModule{}
	for _, mdl := range dc.opt.Device.Modules {
		var stat state_pb.ModuleState
		mdl_id := mdl.Id
		hbt, err := dc.storage.GetHeartbeatAt(mdl_id)
		if err != nil {
			hbt = time.Unix(0, 0)
			dc.logger.WithError(err).Warningf("failed to get heartbeat time in storage")
		}
		pb_hbt := protobuf_helper.FromTime(hbt)

		if time.Now().Sub(hbt) < dc.opt.Config.HeartbeatTimeout {
			stat = state_pb.ModuleState_MODULE_STATE_ONLINE
			any_module_alive = true
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

	if !any_module_alive {
		defer dc.Stop()
		dc.logger.Debugf("all modules offline")
		return
	}

	now := time.Now()
	pb_now := protobuf_helper.FromTime(now)

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
		defer dc.Stop()
		dc.logger.WithError(err).Debugf("failed to heartbeat")
		return
	}

	dc.logger.WithFields(log.Fields{
		"device":       dc.opt.Device.Id,
		"heartbeat_at": now,
	}).Debugf("heartbeat")

}

func (dc *DeviceConnection) ping_loop() {
	dc.stm_wg.Wait()
	defer dc.logger.WithFields(log.Fields{
		"device":  dc.opt.Device.Id,
		"session": dc.opt.Session.Connection,
	}).Debugf("device connection ping loop exited")

	for {
		if dc.is_closed() {
			return
		}

		go dc.ping_once()
		time.Sleep(dc.opt.Config.PingInterval)
	}
}

func (dc *DeviceConnection) ping_once() {
	ping_pkt := &pb.ConnectResponse{
		SessionId: 0,
		Kind:      pb.ConnectMessageKind_CONNECT_MESSAGE_KIND_SYSTEM,
		Union: &pb.ConnectResponse_UnaryCall{
			UnaryCall: &pb.UnaryCallValue{
				Name:      "system",
				Component: "system",
				Method:    "ping",
				Value:     nil,
			},
		},
	}

	stm, unlock := dc.get_stream()
	err := stm.Send(ping_pkt)
	unlock()
	if err != nil {
		defer dc.Stop()

		dc.logger.WithError(err).Warningf("failed to send ping request")
		return
	}

	dc.logger.Debugf("sending ping request")
}

func (dc *DeviceConnection) build_mqtt_module_proxy(mdl *pb.Module) (component.ModuleProxy, error) {
	mdl_id := mdl.GetId()

	mdl_sess, err := dc.storage.GetModuleSession(mdl_id)
	if err != nil {
		dc.logger.WithError(err).Debugf("failed to get module session in storage")
		return nil, err
	}

	mqtt_cli, err := libmqtt.NewClient(
		libmqtt.WithVersion(libmqtt.V5, true),
		libmqtt.WithServer(dc.opt.DeviceCloud.Connection.Mqtt.Address),
		libmqtt.WithIdentity(dc.opt.DeviceCloud.Connection.Mqtt.Username, dc.opt.DeviceCloud.Connection.Mqtt.Password),
		libmqtt.WithKeepalive(10, 1.2),
		libmqtt.WithAutoReconnect(true),
		libmqtt.WithBackoffStrategy(time.Second, 5*time.Second, 1.2),
	)
	if err != nil {
		dc.logger.WithError(err).Debugf("failed to new mqtt client")
		return nil, err
	}

	prx, err := component.NewModuleProxy(
		"mqtt",
		"logger", dc.logger,
		"module_id", mdl_id,
		"session_id", mdl_sess,
		"mqtt_client", mqtt_cli,
	)
	if err != nil {
		dc.logger.WithError(err).Debugf("failed to new module proxy")
		return nil, err
	}

	return prx, nil
}

func (dc *DeviceConnection) get_module_proxy(name string) (component.ModuleProxy, error) {
	var err error
	var mdl_prx component.ModuleProxy

	mdl := dc.get_module_info_by_name(name)
	if mdl == nil {
		err = ErrModuleNotFound
		dc.logger.WithError(err).Debugf("failed to get module by name")
		return nil, err
	}

	ep := mdl.GetEndpoint()
	prx_drv := "mqtt"
	u, err := url.Parse(ep)
	if err != nil {
		dc.logger.WithError(err).Debugf("bad module endpoint")
		return nil, ErrBadModuleEndpoint
	}
	scheme := strings.ToLower(u.Scheme)
	if !strings.HasPrefix(scheme, "mtp") {
		return nil, ErrBadModuleEndpoint
	}

	if strings.HasPrefix(scheme, "mtp+") {
		prx_drv = strings.TrimPrefix(scheme, "mtp+")
	}

	switch prx_drv {
	case "mqtt":
		if mdl_prx, err = dc.build_mqtt_module_proxy(mdl); err != nil {
			dc.logger.WithError(err).Debugf("failed to build mqtt module proxy")
			return nil, err
		}
	default:
		dc.logger.Debugf("unsupported module proxy driver")
		return nil, ErrUnsupportedModuleProxyDriver
	}

	return mdl_prx, nil
}

func NewDeviceConnection(args ...interface{}) (*DeviceConnection, error) {
	var ok bool
	opt := &DeviceConnectionOption{}

	opt.Session.Startup = session_helper.GenerateStartupSession()
	opt.Session.Major = session_helper.GenerateMajorSession()
	opt.Session.Connection = session_helper.NewSession(opt.Session.Startup, opt.Session.Major)
	opt.Config.HeartbeatInterval = 13 * time.Second
	opt.Config.HeartbeatTimeout = 57 * time.Second
	opt.Config.MaxReconnect = 7
	opt.Config.MinReconnectInterval = 3 * time.Second
	opt.Config.MaxReconnectInterval = 17 * time.Second
	opt.Config.PingInterval = 27 * time.Second

	dc := &DeviceConnection{
		opt: opt,

		mdl_info_cache: make(map[string]*pb.Module),
	}

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger": opt_helper.ToLogger(&dc.logger),
		"device": func(key string, val interface{}) error {
			if dc.info, ok = val.(*pb.Device); !ok {
				return opt_helper.InvalidArgument(key)
			}
			return nil
		},
		"storage": func(key string, val interface{}) error {
			if dc.storage, ok = val.(storage.Storage); !ok {
				return opt_helper.InvalidArgument(key)
			}
			return nil
		},
		"client_factory": func(key string, val interface{}) error {
			if dc.cli_fty, ok = val.(*client_helper.ClientFactory); !ok {
				return opt_helper.InvalidArgument(key)
			}
			return nil
		},
		"tokener": func(key string, val interface{}) error {
			if dc.tknr, ok = val.(token_helper.Tokener); !ok {
				return opt_helper.InvalidArgument(key)
			}
			return nil
		},
		"mqtt_address":         opt_helper.ToString(&dc.opt.DeviceCloud.Connection.Mqtt.Address),
		"mqtt_username":        opt_helper.ToString(&dc.opt.DeviceCloud.Connection.Mqtt.Username),
		"mqtt_password":        opt_helper.ToString(&dc.opt.DeviceCloud.Connection.Mqtt.Password),
		"device_cloud_session": opt_helper.ToString(&dc.opt.DeviceCloud.Session.Id),
	})(args...); err != nil {
		return nil, err
	}

	dc.opt.Device.Id = dc.info.Id
	for _, mdl := range dc.info.Modules {
		dc.opt.Device.Modules = append(dc.opt.Device.Modules, struct{ Id string }{Id: mdl.Id})
	}

	return dc, nil
}
