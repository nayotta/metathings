package metathings_deviced_connection

import (
	"context"
	"strconv"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	session_helper "github.com/nayotta/metathings/pkg/common/session"
	session_storage "github.com/nayotta/metathings/pkg/deviced/session_storage"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type Connection interface {
	Err(err ...error) error
	Wait() chan bool
	Done()
	Cleanup()
}

type connection struct {
	err        error
	c          chan bool
	done_once  *sync.Once
	err_once   *sync.Once
	cleanup_cb func()
}

func (self *connection) Err(err ...error) error {
	if len(err) > 0 {
		self.err_once.Do(func() {
			self.err = err[0]
		})
	}

	return self.err
}

func (self *connection) Wait() chan bool {
	return self.c
}

func (self *connection) Cleanup() {
	if self.cleanup_cb != nil {
		self.cleanup_cb()
	}
}

func (self *connection) Done() {
	self.done_once.Do(func() { close(self.c) })
}

func NewConnection(cleanup_cb func()) Connection {
	return &connection{
		c:          make(chan bool),
		done_once:  new(sync.Once),
		err_once:   new(sync.Once),
		cleanup_cb: cleanup_cb,
	}
}

type StreamConnection interface {
	Connection
}

type streamConnection struct {
	connection
}

type ConnectionCenter interface {
	BuildConnection(*storage.Device, pb.DevicedService_ConnectServer) (Connection, error)
	UnaryCall(*storage.Device, *pb.OpUnaryCallValue) (*pb.UnaryCallValue, error)
	StreamCall(*storage.Device, *pb.OpStreamCallConfig, pb.DevicedService_StreamCallServer) error
}

type connectionCenter struct {
	logger          log.FieldLogger
	brfty           BridgeFactory
	storage         Storage
	session_storage session_storage.SessionStorage
}

func (self *connectionCenter) get_session_from_context(ctx context.Context) int64 {
	var x int64
	var err error

	x, err = strconv.ParseInt(metautils.ExtractIncoming(ctx).Get("session"), 0, 64)
	if err != nil {
		return 0
	}

	return x
}

func (self *connectionCenter) connection_loop(dev *storage.Device, conn Connection, br Bridge, stm pb.DevicedService_ConnectServer) {
	defer conn.Done()

	dev_id := *dev.Id
	br_id := br.Id()

	logger := self.logger.WithFields(log.Fields{
		"device": dev_id,
		"bridge": br_id,
		"side":   "south",
	})

	wg := &sync.WaitGroup{}
	wg.Add(2)

	south_to_bridge_quit := make(chan bool)
	south_from_bridge_quit := make(chan bool)
	defer close(south_to_bridge_quit)
	defer close(south_from_bridge_quit)

	south_to_bridge_wait := self.south_to_bridge(dev, conn, br, stm, south_to_bridge_quit, wg, logger.WithField("#dir", "SB"))
	south_from_bridge_wait := self.south_from_bridge(dev, conn, br, stm, south_from_bridge_quit, wg, logger.WithField("#dir", "BS"))

	select {
	case <-south_to_bridge_wait:
		south_from_bridge_quit <- false
	case <-south_from_bridge_wait:
		south_to_bridge_quit <- false
	}

	logger.Debugf("waiting for disconnect")
	wg.Wait()

	logger.Debugf("connection loop closed")
}

func (self *connectionCenter) south_to_bridge(
	dev *storage.Device,
	conn Connection,
	br Bridge,
	south pb.DevicedService_ConnectServer,
	quit chan bool,
	wg *sync.WaitGroup,
	logger log.FieldLogger,
) chan struct{} {
	wait := make(chan struct{})

	go func() {
		var buf []byte
		var res *pb.ConnectResponse
		var sending_bridge Bridge
		var ok bool
		var err error

		defer func() {
			if err != nil {
				conn.Err(err)
			}

			close(wait)
			logger.Debugf("close waiting channel")

			wg.Done()
			logger.Debugf("wg done")

			logger.Debugf("loop closed")
		}()

		south_recv_chan := make(chan *pb.ConnectResponse)
		go func(ch chan *pb.ConnectResponse, stm pb.DevicedService_ConnectServer) {
			defer close(ch)
			for {
				res, err := stm.Recv()
				if err != nil {
					logger.WithError(err).Debugf("failed to recv msg")
					return
				}
				ch <- res
			}
		}(south_recv_chan, south)

		for epoch := uint64(0); ; epoch++ {
			logger = logger.WithField("epoch", epoch)
			is_temp_sess := false

			select {
			case res, ok = <-south_recv_chan:
				if !ok {
					logger.Debugf("south recv channel closed")
					return
				}
				logger.Debugf("recv dev res")
			case <-quit:
				logger.Debugf("catch quit signal")
				return
			}

			if buf, err = proto.Marshal(res); err != nil {
				logger.WithError(err).Warningf("failed to marshal request data")
				continue
			}

			if res_br_id := parse_bridge_id(*dev.Id, res.GetSessionId()); res_br_id != br.Id() {
				if sending_bridge, err = self.brfty.GetBridge(res_br_id); err != nil {
					logger.WithError(err).Debugf("failed to build bridge for unary call response")
					return
				}
				is_temp_sess = true
			} else {
				sending_bridge = br
			}

			if err = sending_bridge.South().Send(buf); err != nil {
				logger.WithError(err).Debugf("failed to send msg")
				return
			}
			if is_temp_sess {
				sending_bridge.Close()
			}

			logger.Debugf("send msg")
		}
	}()

	return wait
}

func (self *connectionCenter) south_from_bridge(
	dev *storage.Device,
	conn Connection,
	br Bridge,
	stm pb.DevicedService_ConnectServer,
	quit chan bool,
	wg *sync.WaitGroup,
	logger log.FieldLogger,
) chan struct{} {
	wait := make(chan struct{})

	go func() {
		var buf []byte
		var ok bool
		var err error

		defer func() {
			if err != nil {
				conn.Err(err)
			}

			br.South().Send(must_marshal_message(new_exit_response_message(0)))
			logger.Debugf("send exit msg to north")

			close(wait)
			logger.Debugf("close waiting channel")

			wg.Done()
			logger.Debugf("wg done")

			logger.Debugf("loop closed")
		}()

		handler := self.new_south_from_bridge_handler(dev, stm, br)
		for epoch := uint64(0); ; epoch++ {
			var req pb.ConnectRequest
			logger = logger.WithField("epoch", epoch)

			// TODO(Peer): catch receiving error
			select {
			case buf, ok = <-br.South().AsyncRecv():
				if !ok {
					logger.Warningf("bridge disconnected")
					return
				}
				logger.Debugf("recv msg")
			case <-quit:
				logger.Debugf("catch quit signal")
				return
			}

			if err = proto.Unmarshal(buf, &req); err != nil {
				logger.WithError(err).Warningf("failed to unmarshal response data")
				continue
			}

			if err = handler(&req, logger); err != nil {
				return
			}

		}
	}()

	return wait
}

func (self *connectionCenter) new_south_from_bridge_handler(dev *storage.Device, south pb.DevicedService_ConnectServer, bridge Bridge) func(*pb.ConnectRequest, log.FieldLogger) error {
	return func(req *pb.ConnectRequest, logger log.FieldLogger) error {
		var err error

		stm_req := req.GetStreamCall()
		if stm_req != nil {
			switch stm_req.Union.(type) {
			case *pb.OpStreamCallValue_Exit:
				logger.Debugf("recv exit msg")
				return context.Canceled
			}
		}

		if err = south.Send(req); err != nil {
			logger.WithError(err).Debugf("failed to send msg")
			return err
		}

		logger.Debugf("send dev req")

		return nil
	}
}

func (self *connectionCenter) BuildConnection(dev *storage.Device, stm pb.DevicedService_ConnectServer) (Connection, error) {
	var cleanup_cb func()
	ctx := stm.Context()
	sess := self.get_session_from_context(ctx)
	dev_id := *dev.Id

	logger := self.logger.WithFields(log.Fields{
		"session": sess,
		"devie":   dev_id,
	})

	self.printSessionInfo(sess)

	br, err := self.brfty.BuildBridge(dev_id, sess)
	if err != nil {
		logger.WithError(err).Debugf("failed to build bridge")
		return nil, err
	}
	br_id := br.Id()

	if session_helper.IsMajorSession(sess) {
		cur_sess, err := self.session_storage.GetStartupSession(dev_id)
		if err != nil {
			logger.WithError(err).Debugf("failed to get startup session")
			return nil, err
		}

		if cur_sess != 0 {
			err = ErrDuplicatedDeviceInstance
			logger.WithError(err).Debugf("duplicated major connection for device")
			return nil, err
		}

		startup_sess := session_helper.GetStartupSession(sess)
		if err = self.session_storage.SetStartupSessionIfNotExists(dev_id, startup_sess, session_helper.STARTUP_SESSION_EXPIRE); err != nil {
			logger.WithError(err).Debugf("failed to set startup session")
			return nil, err
		}

		err = self.storage.AddBridgeToDevice(dev_id, startup_sess, br_id)
		if err != nil {
			return nil, err
		}
		cleanup_cb = func() {
			logger = logger.WithField("bridge", br_id)

			if err = self.storage.RemoveBridgeFromDevice(dev_id, startup_sess, br_id); err != nil {
				logger.WithError(err).Warningf("failed to remove bridge from device")
			} else {
				logger.Debugf("remove bridge from device")
			}

			if err = self.session_storage.UnsetStartupSession(dev_id); err != nil {
				logger.WithError(err).Warningf("failed to unset startup session")
			} else {
				logger.Debugf("unset startup session")
			}

			if err = br.Close(); err != nil {
				logger.WithError(err).Warningf("failed to close bridge")
			}

			logger.Debugf("connection cleanup")
		}
		logger.WithField("bridge", br_id).Debugf("add bridge to device")
	} else {
		cleanup_cb = func() {
			logger = logger.WithField("bridge", br_id)

			if err = br.Close(); err != nil {
				logger.WithError(err).Warningf("failed to close bridge")
			}

			logger.Debugf("connection cleanup")
		}
	}

	conn := NewConnection(cleanup_cb)
	go self.connection_loop(dev, conn, br, stm)

	return conn, nil
}

func (self *connectionCenter) UnaryCall(dev *storage.Device, req *pb.OpUnaryCallValue) (*pb.UnaryCallValue, error) {
	var startup_sess int32
	var br_ids []string
	var br_id string
	var conn_br Bridge
	var sess_br Bridge
	var buf []byte
	var conn_res pb.ConnectResponse
	var ucv *pb.UnaryCallValue
	var err error
	var crerr *pb.ErrorValue
	dev_id := *dev.Id

	if startup_sess, err = self.session_storage.GetStartupSession(dev_id); err != nil {
		self.logger.WithError(err).Debugf("failed to get startup session")
	}

	conn_sess := session_helper.GenerateTempSession()
	sess := session_helper.NewSession(startup_sess, conn_sess)

	if br_ids, err = self.storage.ListBridgesFromDevice(dev_id, startup_sess); err != nil {
		return nil, err
	}
	self.logger.WithFields(log.Fields{
		"bridges": br_ids,
		"device":  dev_id,
	}).Debugf("list bridges from device")

	if len(br_ids) == 0 {
		err = ErrDeviceOffline
		self.logger.WithError(err).Debugf("device bridges is empty")
		return nil, err
	}
	br_id = br_ids[0]

	if conn_br, err = self.brfty.GetBridge(br_id); err != nil {
		return nil, err
	}
	defer conn_br.Close()
	self.logger.WithField("bridge", br_id).Debugf("get connection bridge")

	if sess_br, err = self.brfty.BuildBridge(dev_id, sess); err != nil {
		return nil, err
	}
	defer sess_br.Close()
	self.logger.WithField("bridge", sess_br.Id()).Debugf("build session bridge")

	conn_req := &pb.ConnectRequest{
		SessionId: &wrappers.Int64Value{Value: sess},
		Kind:      pb.ConnectMessageKind_CONNECT_MESSAGE_KIND_USER,
		Union: &pb.ConnectRequest_UnaryCall{
			UnaryCall: req,
		},
	}

	if buf, err = proto.Marshal(conn_req); err != nil {
		return nil, err
	}
	self.logger.Debugf("marshal request")

	if err = conn_br.North().Send(buf); err != nil {
		return nil, err
	}
	self.logger.Debugf("send request")

	if buf, err = sess_br.North().Recv(); err != nil {
		return nil, err
	}
	self.logger.Debugf("recv response")

	if err = proto.Unmarshal(buf, &conn_res); err != nil {
		return nil, err
	}
	self.logger.Debugf("unmarshal response")

	if crerr = conn_res.GetErr(); crerr != nil {
		return nil, status.Errorf(codes.Code(crerr.GetCode()), crerr.GetMessage())
	}

	if ucv = conn_res.GetUnaryCall(); ucv == nil {
		return nil, ErrUnexpectedResponse
	}

	return ucv, nil
}

func (self *connectionCenter) StreamCall(dev *storage.Device, cfg *pb.OpStreamCallConfig, stm pb.DevicedService_StreamCallServer) error {
	var startup_sess int32
	var br_ids []string
	var br_id string
	var conn_br Bridge
	var sess_br Bridge
	var buf []byte
	var err error
	dev_id := *dev.Id

	logger := self.logger.WithFields(log.Fields{
		"device": dev_id,
		"side":   "north",
	})

	if startup_sess, err = self.session_storage.GetStartupSession(dev_id); err != nil {
		self.logger.WithError(err).Debugf("failed to get startup session")
		return err
	}

	conn_sess := session_helper.GenerateMinorSession()
	sess := session_helper.NewSession(startup_sess, conn_sess)

	logger = logger.WithField("session", sess)
	self.printSessionInfo(sess)

	if br_ids, err = self.storage.ListBridgesFromDevice(dev_id, startup_sess); err != nil {
		logger.WithError(err).WithField("device_id", dev_id).Debugf("failed to get bridge")
		return err
	}
	logger.WithFields(log.Fields{
		"bridges": br_ids,
		"device":  dev_id,
	}).Debugf("list bridges from device")

	if len(br_ids) == 0 {
		return ErrDeviceOffline
	}
	br_id = br_ids[0]

	if conn_br, err = self.brfty.GetBridge(br_id); err != nil {
		logger.WithError(err).Debugf("failed to get bridge from bridge factory")
		return err
	}
	logger.WithField("bridge", br_id).Debugf("pick connection bridge")

	if sess_br, err = self.brfty.BuildBridge(dev_id, sess); err != nil {
		logger.WithError(err).Debugf("failed to build bridge")
		return err
	}
	defer sess_br.Close()
	logger.WithField("bridge", sess_br.Id()).Debugf("build session bridge")

	go func() {
		defer conn_br.Close()
		cfg_req := &pb.ConnectRequest{
			SessionId: &wrappers.Int64Value{Value: sess},
			Kind:      pb.ConnectMessageKind_CONNECT_MESSAGE_KIND_USER,
			Union: &pb.ConnectRequest_StreamCall{
				StreamCall: &pb.OpStreamCallValue{
					Union: &pb.OpStreamCallValue_Config{
						Config: cfg,
					},
				},
			},
		}

		if buf, err = proto.Marshal(cfg_req); err != nil {
			logger.WithError(err).Debugf("failed to marshal config msg")
			return
		}

		if err = conn_br.North().Send(buf); err != nil {
			logger.WithError(err).Debugf("failed to send config request")
			return
		}
		logger.Debugf("send config request to device")
	}()

	wg := &sync.WaitGroup{}
	wg.Add(2)
	north_to_bridge_quit := make(chan struct{})
	north_from_bridge_quit := make(chan struct{})
	defer close(north_to_bridge_quit)
	defer close(north_from_bridge_quit)

	loop_logger := self.logger.WithFields(log.Fields{
		"#method":    cfg.GetMethod().GetValue(),
		"#component": cfg.GetComponent().GetValue(),
		"#name":      cfg.GetName().GetValue(),
	})
	north_to_bridge_wait := self.north_to_bridge(dev, sess, stm, sess_br, &err, north_to_bridge_quit, wg, loop_logger.WithFields(log.Fields{"#dir": "NB", "bridge": sess_br.Id()}))
	north_from_bridge_wait := self.north_from_bridge(dev, sess, stm, sess_br, &err, north_from_bridge_quit, wg, loop_logger.WithFields(log.Fields{"#dir": "BN", "bridge": sess_br.Id()}))

	select {
	case <-north_to_bridge_wait:
		north_from_bridge_quit <- struct{}{}
	case <-north_from_bridge_wait:
		north_to_bridge_quit <- struct{}{}
	}

	logger.Debugf("waiting for disconnect")
	wg.Wait()

	if err != nil {
		logger.WithError(err).Debugf("streaming error")
		return err
	}
	logger.Debugf("streaming closed")

	return nil
}

func (self *connectionCenter) north_to_bridge(
	dev *storage.Device,
	sess int64,
	north pb.DevicedService_StreamCallServer,
	bridge Bridge,
	perr *error,
	quit chan struct{},
	wg *sync.WaitGroup,
	logger log.FieldLogger,
) chan struct{} {
	wait := make(chan struct{})
	go func() {
		var buf []byte
		var req *pb.StreamCallRequest
		var ok bool
		var err error

		defer func() {
			if *perr == nil && err != nil {
				*perr = err
			}

			close(wait)
			logger.Debugf("close waiting channel")

			wg.Done()
			logger.Debugf("wg done")

			logger.Debugf("loop closed")
		}()

		north_recv_chan := make(chan *pb.StreamCallRequest)
		go func(ch chan *pb.StreamCallRequest, stm pb.DevicedService_StreamCallServer) {
			defer close(ch)
			for {
				req, err := stm.Recv()
				if err != nil {
					logger.WithError(err).Debugf("failed to recv msg")
					return
				}
				ch <- req
			}
		}(north_recv_chan, north)

		for epoch := uint64(0); ; epoch++ {
			logger = logger.WithField("epoch", epoch)

			select {
			case req, ok = <-north_recv_chan:
				if !ok {
					logger.Debugf("north recv channel closed")
					return
				}
				logger.Debugf("recv cli req")
			case <-quit:
				logger.Debugf("catch quit signal")
				return
			}

			conn_req := &pb.ConnectRequest{
				SessionId: &wrappers.Int64Value{Value: sess},
				Kind:      pb.ConnectMessageKind_CONNECT_MESSAGE_KIND_USER,
				Union: &pb.ConnectRequest_StreamCall{
					StreamCall: req.GetValue(),
				},
			}

			if buf, err = proto.Marshal(conn_req); err != nil {
				logger.WithError(err).Debugf("failed to marshal request")
				return
			}

			if err = bridge.North().Send(buf); err != nil {
				logger.WithError(err).Debugf("failed to send msg")
				return
			}
			logger.Debugf("send msg")
		}
	}()

	return wait
}

func (self *connectionCenter) north_from_bridge(
	dev *storage.Device,
	sess int64,
	north pb.DevicedService_StreamCallServer,
	bridge Bridge,
	perr *error,
	quit chan struct{},
	wg *sync.WaitGroup,
	logger log.FieldLogger,
) chan bool {
	wait := make(chan bool)
	go func() {
		var buf []byte
		var ok bool
		var err error

		defer func() {
			if *perr == nil && err != nil {
				*perr = err
			}

			bridge.North().Send(must_marshal_message(new_exit_request_message(sess)))
			logger.Debugf("send exit request to south")

			close(wait)
			logger.Debugf("close waiting channel")

			wg.Done()
			logger.Debugf("wg done")

			logger.Debugf("loop closed")
		}()

		handler := self.new_north_from_bridge_handler(dev, north, bridge)
		for epoch := uint64(0); ; epoch++ {
			var res pb.ConnectResponse

			select {
			case buf, ok = <-bridge.North().AsyncRecv():
				if !ok {
					return
				}
				logger.Debugf("recv msg")
			case <-quit:
				logger.Debugf("catch quit signal")
				return
			}

			if err = proto.Unmarshal(buf, &res); err != nil {
				logger.WithError(err).Debugf("failed to unmarshal response")
				return
			}

			if err = handler(&res, logger); err != nil {
				return
			}
		}
	}()

	return wait
}

func (self *connectionCenter) new_north_from_bridge_handler(dev *storage.Device, north pb.DevicedService_StreamCallServer, bridge Bridge) func(*pb.ConnectResponse, log.FieldLogger) error {
	acked := false
	acked_once := new(sync.Once)

	return func(res *pb.ConnectResponse, logger log.FieldLogger) error {
		var err error

		stm_res := res.GetStreamCall()
		switch stm_res.Union.(type) {
		case *pb.StreamCallValue_Value:
			if acked != true {
				logger.Warningf("recv msg but not acked, drop it")
				return nil
			}

			if err = north.Send(&pb.StreamCallResponse{
				Device: &pb.Device{Id: *dev.Id},
				Value:  res.GetStreamCall(),
			}); err != nil {
				logger.WithError(err).Debugf("failed to send response")
				return err
			}
			logger.Debugf("send cli res")
		case *pb.StreamCallValue_Config:
			// TODO(Peer): catch error when send ack failed

			// aviod to resend ack msg
			acked_once.Do(func() {
				if err = bridge.North().Send(must_marshal_message(new_config_ack_request_message(res.GetSessionId()))); err != nil {
					logger.WithError(err).Debugf("failed to send ack msg")
					return
				}

				if err = north.Send(new_config_ack_response_message_for_north(*dev.Id)); err != nil {
					logger.WithError(err).Debugf("failed to send ack msg to north")
					return
				}

				acked = true
				logger.Debugf("send ack msg")
			})
		case *pb.StreamCallValue_ConfigAck:
			logger.Warningf("should not reach here")
		case *pb.StreamCallValue_Exit:
			logger.Debugf("recv exit msg")
			return context.Canceled
		default:
			logger.Debugf("unexpected response")
		}

		return nil
	}
}

func NewConnectionCenter(brfty BridgeFactory, stor Storage, sess_stor session_storage.SessionStorage, logger log.FieldLogger) (ConnectionCenter, error) {
	return &connectionCenter{
		logger:          logger,
		brfty:           brfty,
		storage:         stor,
		session_storage: sess_stor,
	}, nil
}
