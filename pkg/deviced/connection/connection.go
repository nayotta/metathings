package metathings_deviced_connection

import (
	"context"
	"strconv"
	"sync"

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

type ConnectionCenter interface {
	BuildConnection(*storage.Device, pb.DevicedService_ConnectServer) (Connection, error)
	UnaryCall(*storage.Device, *pb.OpUnaryCallValue) (*pb.UnaryCallValue, error)
	StreamCall(pb.DevicedService_StreamCallServer) (StreamConnection, error)
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
		"devid": dev_id,
		"brid":  br_id,
	})

	wg := &sync.WaitGroup{}
	wg.Add(2)

	br2stm_quit := make(chan bool)
	stm2br_quit := make(chan bool)
	defer close(br2stm_quit)
	defer close(stm2br_quit)

	br2stm_wait := self.br2stm(dev, conn, br, stm, br2stm_quit, wg)
	stm2br_wait := self.stm2br(dev, conn, br, stm, stm2br_quit, wg)

	select {
	case <-br2stm_wait:
		stm2br_quit <- false
	case <-stm2br_wait:
		br2stm_quit <- false
	}

	logger.Debugf("waiting for disconnect")
	wg.Wait()

	if err = self.storage.RemoveBridgeFromDevice(dev_id, br_id); err != nil {
		self.logger.WithError(err).Errorf("failed to remove bridge from device")
	}
	logger.Debugf("remove bridge from device")

	logger.Debugf("quit connection loop")

}

func (self *connectionCenter) br2stm(dev *storage.Device, conn Connection, br Bridge, stm pb.DevicedService_ConnectServer, quit chan bool, wg *sync.WaitGroup) chan bool {
	wait := make(chan bool)

	go func() {
		var err error

		logger := self.logger.WithFields(log.Fields{
			"#from": "bridge",
			"#to":   "stream",
			"devid": *dev.Id,
		})

		defer func() {
			if err != nil {
				conn.Err(err)
			}

			close(wait)
			wg.Done()
			logger.Debugf("connection closed")
		}()

		abr := NewAsyncBridgeWrapper(br)
		for {
			var buf []byte
			var req pb.ConnectRequest

			select {
			case <-quit:
				logger.Debugf("quit signal from stm2br")
				return
			case buf = <-abr.Recv():
				logger.Debugf("recv msg")
			}

			if err = proto.Unmarshal(buf, &req); err != nil {
				logger.WithError(err).Debugf("failed to unmarshal response data")
				return
			}

			if err = stm.Send(&req); err != nil {
				logger.WithError(err).Debugf("failed to send msg")
				return
			}
			logger.Debugf("send msg")
		}
	}()

	return wait
}

func (self *connectionCenter) stm2br(dev *storage.Device, conn Connection, br Bridge, stm pb.DevicedService_ConnectServer, quit chan bool, wg *sync.WaitGroup) chan bool {
	wait := make(chan bool)

	go func() {
		var err error

		logger := self.logger.WithFields(log.Fields{
			"#from": "stream",
			"#to":   "bridge",
			"devid": *dev.Id,
		})

		defer func() {
			if err != nil {
				conn.Err(err)
			}

			close(wait)
			wg.Done()
			logger.Debugf("connection closed")
		}()

		for {
			var buf []byte
			var res *pb.ConnectResponse
			var res_br Bridge

			if res, err = stm.Recv(); err != nil {
				logger.WithError(err).Debugf("failed to recv msg")
				return
			}
			logger.Debugf("recv msg")

			if res.GetUnaryCall() != nil {
				if res_br, err = self.brfty.BuildBridge(*dev.Id, res.SessionId); err != nil {
					logger.WithError(err).Debugf("failed to build bridge for unary call")
					return
				}
			} else {
				res_br = br
			}

			if buf, err = proto.Marshal(res); err != nil {
				logger.WithError(err).Debugf("failed to marshal request data")
				return
			}

			abr := NewAsyncBridgeWrapper(res_br)
			select {
			case <-quit:
				logger.Debugf("quit signal from br2stm")
			case abr.Send() <- buf:
				logger.Debugf("send msg")
			}
		}
	}()

	return wait
}

func (self *connectionCenter) BuildConnection(dev *storage.Device, stm pb.DevicedService_ConnectServer) (Connection, error) {
	ctx := stm.Context()
	sess := self.get_session_from_context(ctx)
	dev_id := *dev.Id

	br, err := self.brfty.BuildBridge(dev_id, sess)
	if err != nil {
		return nil, err
	}
	br_id := br.Id()
	self.logger.WithField("brid", br_id).Debugf("build bridge")

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
	self.logger.WithField("bridge", br_ids[0]).Debugf("get bridge")

	if sess_br, err = self.brfty.BuildBridge(dev_id, conn_req_sess); err != nil {
		return nil, err
	}
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

	if err = conn_br.Send(buf); err != nil {
		return nil, err
	}
	self.logger.Debugf("send request")

	if buf, err = sess_br.Recv(); err != nil {
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

func (self *connectionCenter) StreamCall(pb.DevicedService_StreamCallServer) (StreamConnection, error) {
	panic("unimplemented")
}

func NewConnectionCenter(brfty BridgeFactory, storage Storage, logger log.FieldLogger) (ConnectionCenter, error) {
	return &connectionCenter{
		logger:  logger,
		brfty:   brfty,
		storage: storage,
	}, nil
}
