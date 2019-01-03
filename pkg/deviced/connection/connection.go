package metathings_deviced_connection

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	log "github.com/sirupsen/logrus"

	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type Connection interface {
	Err(err ...error) error
	Wait() chan bool
	Close()
}

type connection struct {
	err      error
	c        chan bool
	close_cb func()
}

func (self *connection) Err(err ...error) error {
	if len(err) > 0 && self.err != nil {
		self.err = err[0]
	}
	return self.err
}

func (self *connection) Wait() chan bool {
	return self.c
}

func (self *connection) Close() {
	self.close_cb()
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
	logger  log.FieldLogger
	brfty   BridgeFactory
	storage Storage
}

func (self *connectionCenter) get_session_from_context(ctx context.Context) int32 {
	var x int
	var err error

	x, err = strconv.Atoi(metautils.ExtractIncoming(ctx).Get("session"))
	if err != nil {
		return 0
	}

	return int32(x)
}

func (self *connectionCenter) connection_loop(dev *storage.Device, conn Connection, br Bridge, stm pb.DevicedService_ConnectServer) {
	var err error
	dev_id := *dev.Id
	br_id := br.Id()

	logger := self.logger.WithFields(log.Fields{
		"device": dev_id,
		"bridge": br_id,
	})

	wg := &sync.WaitGroup{}
	wg.Add(2)

	south_to_bridge_quit := make(chan bool)
	south_from_bridge_quit := make(chan bool)
	defer close(south_to_bridge_quit)
	defer close(south_from_bridge_quit)

	south_to_bridge_wait := self.south_to_bridge(dev, conn, br, stm, south_to_bridge_quit, wg)
	south_from_bridge_wait := self.south_from_bridge(dev, conn, br, stm, south_from_bridge_quit, wg)

	select {
	case <-south_to_bridge_wait:
		south_from_bridge_quit <- false
	case <-south_from_bridge_wait:
		south_to_bridge_quit <- false
	}

	logger.Debugf("waiting for disconnect")
	wg.Wait()

	if err = self.storage.RemoveBridgeFromDevice(dev_id, br_id); err != nil {
		self.logger.WithError(err).Errorf("failed to remove bridge from device")
	}
	logger.Debugf("remove bridge from device")

	logger.Debugf("quit connection loop")
}

func (self *connectionCenter) south_to_bridge(dev *storage.Device, conn Connection, br Bridge, stm pb.DevicedService_ConnectServer, quit chan bool, wg *sync.WaitGroup) chan struct{} {
	wait := make(chan struct{})

	go func() {
		var buf []byte
		var res *pb.ConnectResponse
		var sending_bridge Bridge
		var err error

		logger := self.logger.WithFields(log.Fields{
			"dir": fmt.Sprintf("bridge(%v)<-south(%v)", br.Id(), *dev.Id),
		})

		defer func() {
			if err != nil {
				conn.Err(err)
			}

			close(wait)
			wg.Done()
			logger.Debugf("loop closed")
		}()

		for epoch := uint64(0); ; epoch++ {
			logger = logger.WithField("epoch", epoch)

			// TODO(Peer): catch quit signal
			logger.Debugf("waiting")
			if res, err = stm.Recv(); err != nil {
				logger.WithError(err).Debugf("failed to recv dev res")
				return
			}
			logger.Debugf("recv dev res")

			if buf, err = proto.Marshal(res); err != nil {
				logger.WithError(err).Warningf("failed to marshal request data")
				continue
			}

			if res_br_id := parse_bridge_id(*dev.Id, res.GetSessionId()); res_br_id != br.Id() {
				if sending_bridge, err = self.brfty.GetBridge(res_br_id); err != nil {
					logger.WithError(err).Debugf("failed to build bridge for unary call response")
					return
				}
			} else {
				sending_bridge = br
			}

			if err = sending_bridge.South().Send(buf); err != nil {
				logger.WithError(err).Debugf("failed to send msg")
				return
			}
			self.logger.WithField("dir", fmt.Sprintf("bridge(%v)<-south(%v)", sending_bridge.Id(), *dev.Id))
		}
	}()

	return wait
}

func (self *connectionCenter) south_from_bridge(dev *storage.Device, conn Connection, br Bridge, stm pb.DevicedService_ConnectServer, quit chan bool, wg *sync.WaitGroup) chan struct{} {
	wait := make(chan struct{})

	go func() {
		var buf []byte

		var err error

		logger := self.logger.WithFields(log.Fields{
			"dir": fmt.Sprintf("bridge(%v)->south(%v)", br.Id(), *dev.Id),
		})

		defer func() {
			if err != nil {
				conn.Err(err)
			}

			close(wait)
			wg.Done()
			logger.Debugf("loop closed")
		}()

		for epoch := uint64(0); ; epoch++ {
			var req pb.ConnectRequest

			logger = logger.WithField("epoch", epoch)

			// TODO(Peer): catch receiving error
			logger.Debugf("waiting")
			select {
			case buf = <-br.South().AsyncRecv():
				logger.Debugf("recv msg")
			case <-quit:
				logger.Debugf("catch quit signal")
				return
			}

			if err = proto.Unmarshal(buf, &req); err != nil {
				logger.WithError(err).Warningf("failed to unmarshal response data")
				continue
			}

			if err = stm.Send(&req); err != nil {
				logger.WithError(err).Debugf("failed to send msg")
				return
			}
			logger.Debugf("send dev req")
		}
	}()

	return wait
}

func (self *connectionCenter) BuildConnection(dev *storage.Device, stm pb.DevicedService_ConnectServer) (Connection, error) {
	ctx := stm.Context()
	sess := self.get_session_from_context(ctx)
	dev_id := *dev.Id

	self.logger.WithFields(log.Fields{
		"session": sess,
		"stage":   "begin",
	}).Debugf("build connection")

	br, err := self.brfty.BuildBridge(dev_id, sess)
	if err != nil {
		return nil, err
	}
	br_id := br.Id()
	self.logger.WithField("bridge", br_id).Debugf("build bridge")

	err = self.storage.AddBridgeToDevice(dev_id, br_id)
	if err != nil {
		return nil, err
	}
	self.logger.WithFields(log.Fields{
		"brid":  br_id,
		"devid": *dev.Id,
	}).Debugf("add bridge to device")

	conn := &connection{
		c: make(chan bool),
		close_cb: func() {
			self.storage.RemoveBridgeFromDevice(dev_id, br_id)
		},
	}

	go self.connection_loop(dev, conn, br, stm)

	self.logger.WithFields(log.Fields{
		"session": sess,
		"stage":   "end",
	}).Debugf("build connection")

	return conn, nil
}

func (self *connectionCenter) UnaryCall(dev *storage.Device, req *pb.OpUnaryCallValue) (*pb.UnaryCallValue, error) {
	var br_ids []string
	var conn_br Bridge
	var sess_br Bridge
	var buf []byte
	var conn_res pb.ConnectResponse
	var ucv *pb.UnaryCallValue
	var err error
	dev_id := *dev.Id
	conn_req_sess := generate_session()

	if br_ids, err = self.storage.ListBridgesFromDevice(dev_id); err != nil {
		return nil, err
	}
	self.logger.WithFields(log.Fields{
		"bridges": br_ids,
		"device":  dev_id,
	}).Debugf("list bridges from device")

	if conn_br, err = self.brfty.GetBridge(br_ids[0]); err != nil {
		return nil, err
	}
	defer conn_br.Close()
	self.logger.WithField("bridge", br_ids[0]).Debugf("get bridge")

	if sess_br, err = self.brfty.BuildBridge(dev_id, conn_req_sess); err != nil {
		return nil, err
	}
	defer sess_br.Close()
	self.logger.WithField("bridge", sess_br.Id()).Debugf("build recv bridge")

	conn_req := &pb.ConnectRequest{
		SessionId: &wrappers.Int32Value{Value: conn_req_sess},
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

	if ucv = conn_res.GetUnaryCall(); ucv == nil {
		return nil, ErrUnexpectedResponse
	}

	return ucv, nil
}

func (self *connectionCenter) StreamCall(dev *storage.Device, cfg *pb.OpStreamCallConfig, stm pb.DevicedService_StreamCallServer) error {
	var br_ids []string
	var conn_br Bridge
	var sess_br Bridge
	var buf []byte
	var cfg_res pb.ConnectResponse
	var err error
	dev_id := *dev.Id
	sess := generate_session()

	if br_ids, err = self.storage.ListBridgesFromDevice(dev_id); err != nil {
		self.logger.WithError(err).WithField("device_id", dev_id).Debugf("failed to get bridge")
		return err
	}
	self.logger.WithFields(log.Fields{
		"device":  dev_id,
		"bridges": br_ids,
	}).Debugf("list bridges from device")

	if conn_br, err = self.brfty.GetBridge(br_ids[0]); err != nil {
		self.logger.WithError(err).Debugf("failed to get bridge from bridge factory")
		return err
	}
	self.logger.WithField("bridge", br_ids[0]).Debugf("pick bridge")

	if sess_br, err = self.brfty.BuildBridge(dev_id, sess); err != nil {
		self.logger.WithError(err).Debugf("failed to build bridge")
		return err
	}
	self.logger.WithFields(log.Fields{
		"device":  dev_id,
		"session": sess,
	}).Debugf("build bridge")

	cfg_req := &pb.ConnectRequest{
		SessionId: &wrappers.Int32Value{Value: sess},
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
		self.logger.WithError(err).Debugf("failed to marshal config msg")
		return err
	}

	go func() {
		// TODO(Peer): design protocol to avoid to send message delay

		// wait for north receiving channel initialized
		time.Sleep(500 * time.Millisecond)
		if err = conn_br.North().Send(buf); err != nil {
			self.logger.WithError(err).Debugf("failed to send config request")
			return
		}
		self.logger.Debugf("send config request to device")
	}()

	if buf, err = sess_br.North().Recv(); err != nil {
		self.logger.WithError(err).Debugf("failed to recv config response")
		return err
	}

	if err = proto.Unmarshal(buf, &cfg_res); err != nil {
		self.logger.WithError(err).Debugf("failed to unmarshal config msg")
		return err
	}

	cfg_val := cfg_res.GetStreamCall().GetConfig()
	if cfg_res.GetSessionId() != sess ||
		cfg_val.GetName() != cfg.GetName().GetValue() ||
		cfg_val.GetComponent() != cfg.GetComponent().GetValue() ||
		cfg_val.GetMethod() != cfg.GetMethod().GetValue() {
		err = ErrUnexpectedResponse
		self.logger.WithError(err).WithFields(log.Fields{}).Debugf("unexpected config response")
		return err
	}
	self.logger.Debugf("recv config response from device")

	wg := &sync.WaitGroup{}
	wg.Add(2)
	north_to_bridge_quit := make(chan struct{})
	north_from_bridge_quit := make(chan struct{})
	defer close(north_to_bridge_quit)
	defer close(north_from_bridge_quit)

	north_to_bridge_wait := self.north_to_bridge(dev, cfg_val, sess, stm, sess_br, &err, north_to_bridge_quit, wg)
	north_from_bridge_wait := self.north_from_bridge(dev, cfg_val, sess, stm, sess_br, &err, north_from_bridge_quit, wg)

	select {
	case <-north_to_bridge_wait:
		north_from_bridge_quit <- struct{}{}
	case <-north_from_bridge_wait:
		north_to_bridge_quit <- struct{}{}
	}

	self.logger.Debugf("waiting for disconnect")
	wg.Wait()

	if err != nil {
		self.logger.WithError(err).Debugf("streaming error")
		return err
	}
	self.logger.Debugf("streaming closed")

	return nil
}

func (self *connectionCenter) north_to_bridge(dev *storage.Device, cfg *pb.StreamCallConfig, sess int32, north pb.DevicedService_StreamCallServer, bridge Bridge, perr *error, quit chan struct{}, wg *sync.WaitGroup) chan struct{} {
	wait := make(chan struct{})
	go func() {
		var buf []byte
		var req *pb.StreamCallRequest
		var err error

		logger := self.logger.WithFields(log.Fields{
			"dir":        fmt.Sprintf("north(%v)->bridge(%v)", *dev.Id, bridge.Id()),
			"#method":    cfg.GetMethod(),
			"#component": cfg.GetComponent(),
			"#name":      cfg.GetName(),
		})

		defer func() {
			if *perr == nil && err != nil {
				*perr = err
			}
			close(wait)
			logger.Debugf("loop closed")
		}()
		for epoch := uint64(0); ; epoch++ {
			logger = logger.WithField("epoch", epoch)

			logger.Debugf("waiting")
			if req, err = north.Recv(); err != nil {
				logger.WithError(err).Debugf("failed to recv msg")
				return
			}
			logger.Debugf("recv cli req")

			conn_req := &pb.ConnectRequest{
				SessionId: &wrappers.Int32Value{Value: sess},
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

func (self *connectionCenter) north_from_bridge(dev *storage.Device, cfg *pb.StreamCallConfig, sess int32, north pb.DevicedService_StreamCallServer, bridge Bridge, perr *error, quit chan struct{}, wg *sync.WaitGroup) chan bool {
	wait := make(chan bool)
	go func() {
		var buf []byte
		var err error

		logger := self.logger.WithFields(log.Fields{
			"dir":        fmt.Sprintf("north(%v)<-bridge(%v)", *dev.Id, bridge.Id()),
			"#method":    cfg.GetMethod(),
			"#component": cfg.GetComponent(),
			"#name":      cfg.GetName(),
		})

		defer func() {
			if *perr == nil && err != nil {
				*perr = err
			}
			close(wait)
			logger.Debugf("loop closed")
		}()
		for epoch := uint64(0); ; epoch++ {
			var res pb.ConnectResponse

			// TODO(Peer): catch receiving error
			logger.Debugf("waiting")
			select {
			case buf = <-bridge.North().AsyncRecv():
				logger.Debugf("recv msg")
			case <-quit:
				logger.Debugf("catch quit signal")
				return
			}

			if err = proto.Unmarshal(buf, &res); err != nil {
				logger.WithError(err).Debugf("failed to unmarshal response")
				return
			}

			stm_res := &pb.StreamCallResponse{
				Device: &pb.Device{Id: *dev.Id},
				Value:  res.GetStreamCall(),
			}

			if err = north.Send(stm_res); err != nil {
				logger.WithError(err).Debugf("failed to send response")
				return
			}
			logger.Debugf("send cli res")
		}
	}()

	return wait
}

func NewConnectionCenter(brfty BridgeFactory, stor Storage, logger log.FieldLogger) (ConnectionCenter, error) {
	return &connectionCenter{
		logger:  logger,
		brfty:   brfty,
		storage: stor,
	}, nil
}
