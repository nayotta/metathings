package metathings_deviced_connection

import (
	"context"
	"strconv"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	log "github.com/sirupsen/logrus"

	storage "github.com/nayotta/metathings/pkg/deviced/storage"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type Connection interface {
	Err(err ...error) error
	Wait() chan bool
}

type connection struct {
	err error
	c   chan bool
}

func (self *connection) Err(err ...error) error {
	if len(err) > 0 {
		self.err = err[0]
	}
	return self.err
}

func (self *connection) Wait() chan bool {
	return self.c
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
	wg := &sync.WaitGroup{}
	wg.Add(2)

	br2stm_quit := make(chan bool)
	stm2br_quit := make(chan bool)
	defer close(br2stm_quit)
	defer close(stm2br_quit)

	br2stm_wait := self.br2stm(dev, conn, br, stm, br2stm_quit, wg)
	stm2br_wait := self.stm2br(dev, conn, br, stm, stm2br_quit, wg)
	defer close(br2stm_wait)
	defer close(stm2br_wait)

	select {
	case <-br2stm_wait:
		stm2br_quit <- false
	case <-stm2br_quit:
		br2stm_quit <- false
	}

	wg.Wait()
}

func (self *connectionCenter) br2stm(dev *storage.Device, conn Connection, br Bridge, stm pb.DevicedService_ConnectServer, quit chan bool, wg *sync.WaitGroup) chan bool {
	wait := make(chan bool)

	go func() {
		var buf []byte
		var req pb.ConnectRequest
		var err error

		defer wg.Done()
		defer func() {
			if err != nil {
				conn.Err(err)
			}

			wait <- false
		}()

		for {
			if buf, err = br.Recv(); err != nil {
				return
			}

			if err = proto.Unmarshal(buf, &req); err != nil {
				return
			}

			if err = stm.Send(&req); err != nil {
				return
			}
		}
	}()

	return wait
}

func (self *connectionCenter) stm2br(dev *storage.Device, conn Connection, br Bridge, stm pb.DevicedService_ConnectServer, quit chan bool, wg *sync.WaitGroup) chan bool {
	wait := make(chan bool)

	go func() {
		var buf []byte
		var res *pb.ConnectResponse
		var res_br Bridge
		var err error

		defer wg.Done()
		defer func() {
			if err != nil {
				conn.Err(err)
			}

			wait <- false
		}()

		for {
			if res, err = stm.Recv(); err != nil {
				return
			}

			if res.GetUnaryCall() != nil {
				if res_br, err = self.brfty.BuildBridge(*dev.Id, res.SessionId); err != nil {
					return
				}
			} else {
				res_br = br
			}

			if buf, err = proto.Marshal(res); err != nil {
				return
			}

			if err = res_br.Send(buf); err != nil {
				return
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

	err = self.storage.AddBridgeToDevice(dev_id, br_id)
	if err != nil {
		return nil, err
	}
	defer self.storage.RemoveBridgeFromDevice(dev_id, br_id)

	conn := &connection{
		c: make(chan bool),
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

	if conn_br, err = self.brfty.GetBridge(br_ids[0]); err != nil {
		return nil, err
	}

	if sess_br, err = self.brfty.BuildBridge(dev_id, conn_req_sess); err != nil {
		return nil, err
	}

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

	if err = conn_br.Send(buf); err != nil {
		return nil, err
	}

	if buf, err = sess_br.Recv(); err != nil {
		return nil, err
	}

	if err = proto.Unmarshal(buf, &conn_res); err != nil {
		return nil, err
	}

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
