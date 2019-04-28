package metathings_device_service

import (
	"context"
	"math"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func parse_error_to_connect_error_response(name, service, method string, err error) *deviced_pb.ConnectResponse_Err {
	s := status.Convert(err)

	return &deviced_pb.ConnectResponse_Err{
		Err: &deviced_pb.ErrorValue{
			Name:      name,
			Component: service,
			Method:    method,
			Code:      uint32(s.Code()),
			Message:   s.Message(),
		},
	}
}

func (self *MetathingsDeviceServiceImpl) handle(req *deviced_pb.ConnectRequest) error {
	switch req.Kind {
	case deviced_pb.ConnectMessageKind_CONNECT_MESSAGE_KIND_SYSTEM:
		return self.handle_system_request(req)
	case deviced_pb.ConnectMessageKind_CONNECT_MESSAGE_KIND_USER:
		return self.handle_user_request(req)
	default:
		self.logger.Warningf("unexpected request data")
		return nil
	}
}

func (self *MetathingsDeviceServiceImpl) handle_system_request(req *deviced_pb.ConnectRequest) error {
	panic("unimplemented")
}

func (self *MetathingsDeviceServiceImpl) handle_user_request(req *deviced_pb.ConnectRequest) error {
	switch req.Union.(type) {
	case *deviced_pb.ConnectRequest_UnaryCall:
		return self.handle_user_unary_request(req)
	case *deviced_pb.ConnectRequest_StreamCall:
		return self.handle_user_stream_request(req)
	default:
		self.logger.WithField("union", req.Union).Debugf("unsupported union type")
		panic("unimplemented")
	}
}

func (self *MetathingsDeviceServiceImpl) handle_user_unary_request(req *deviced_pb.ConnectRequest) error {
	req_val := req.GetUnaryCall()
	sess := req.GetSessionId().GetValue()
	kind := req.GetKind()
	component := req_val.GetComponent().GetValue()
	name := req_val.GetName().GetValue()
	method := req_val.GetName().GetValue()

	logger := self.logger.WithFields(log.Fields{
		"#session":   sess,
		"#component": component,
		"#name":      name,
		"#method":    method,
	})

	mdl, err := self.mdl_db.Lookup(name)
	if err != nil {
		logger.WithError(err).Debugf("failed to lookup module in database")
		return err
	}
	logger.Debugf("lookup module in storage")

	res_val, err := mdl.UnaryCall(context.Background(), req_val)
	logger.Debugf("unary call in module")

	var res *deviced_pb.ConnectResponse

	if err != nil {
		res = &deviced_pb.ConnectResponse{
			SessionId: sess,
			Kind:      kind,
			Union:     parse_error_to_connect_error_response(name, component, method, err),
		}
	} else {
		res = &deviced_pb.ConnectResponse{
			SessionId: sess,
			Kind:      kind,
			Union: &deviced_pb.ConnectResponse_UnaryCall{
				UnaryCall: res_val,
			},
		}
	}

	err = self.conn_stm.Send(res)
	if err != nil {
		logger.Debugf("failed to send msg")
		return err
	}
	logger.Debugf("send msg")

	return nil
}

func (self *MetathingsDeviceServiceImpl) handle_user_stream_request(req *deviced_pb.ConnectRequest) error {
	var cli deviced_pb.DevicedServiceClient
	var cfn client_helper.CloseFn
	var stm deviced_pb.DevicedService_ConnectClient
	var err error

	req_val := req.GetStreamCall()
	sess := req.GetSessionId().GetValue()
	cfg := req_val.GetConfig()
	name := cfg.GetName().GetValue()
	component := cfg.GetComponent().GetValue()
	method := cfg.GetMethod().GetValue()

	logger := self.logger.WithFields(log.Fields{
		"#session":   sess,
		"#component": component,
		"#name":      name,
		"#method":    method,
	})

	if cli, cfn, err = self.cli_fty.NewDevicedServiceClient(); err != nil {
		logger.WithError(err).Debugf("failed to new deviced service client")
		return err
	}
	defer cfn()
	logger.Debugf("new serviced service client")

	ctx := context_helper.NewOutgoingContext(
		context.Background(),
		context_helper.WithTokenOp(self.tknr.GetToken()),
		context_helper.WithSessionOp(sess),
	)
	if stm, err = cli.Connect(ctx); err != nil {
		logger.WithError(err).Debugf("failed to connect to deviced service")
		return err
	}
	logger.Debugf("create deviced stream")

	acked := make(chan struct{})
	acked_once := new(sync.Once)
	go func() {
		// TODO(Peer): make SEND_RES_MAX_RETRY configurable.
		for cnt := 0; cnt < SEND_RES_MAX_RETRY; cnt++ {
			select {
			case _, ok := <-acked:
				if !ok {
					return
				}
			default:
			}

			cfg_res := &deviced_pb.ConnectResponse{
				SessionId: sess,
				Kind:      req.GetKind(),
				Union: &deviced_pb.ConnectResponse_StreamCall{
					StreamCall: &deviced_pb.StreamCallValue{
						Union: &deviced_pb.StreamCallValue_Config{
							Config: &deviced_pb.StreamCallConfig{
								Name:      name,
								Component: component,
								Method:    method,
							},
						},
					},
				},
			}
			if err = stm.Send(cfg_res); err != nil {
				logger.WithError(err).Debugf("failed to send config response")
				return
			}
			logger.Debugf("send config response")
			// TODO(Peer): make SEND_RES_INTERVAL configurable.
			time.Sleep(time.Duration(math.Min(B_SEND_RES_INTERVAL+(float64(cnt)*A_SEND_RES_INTERVAL), MAX_SEND_RES_INTERVAL)) * time.Millisecond)
		}
		select {
		case _, ok := <-acked:
			if !ok {
				return
			}
		default:
			logger.Warningf("failed to recv config ack")
		}
	}()

	mdl, err := self.mdl_db.Lookup(name)
	if err != nil {
		logger.WithError(err).Debugf("failed to lookup module in database")
		return err
	}
	logger.Debugf("lookup module in storage")

	stm = newHijackStream(stm, func(self_ *hijackStream, req *deviced_pb.ConnectRequest) error {
		switch req.GetStreamCall().Union.(type) {
		case *deviced_pb.OpStreamCallValue_Value:
			self_.RecvChan() <- req
		case *deviced_pb.OpStreamCallValue_Config:
		case *deviced_pb.OpStreamCallValue_ConfigAck:
			logger.Debugf("recv config ack")
			acked_once.Do(func() { close(acked) })
		case *deviced_pb.OpStreamCallValue_Exit:
			logger.Debugf("recv exit")
			return context.Canceled
		}

		return nil
	})

	err = mdl.StreamCall(context.Background(), req, stm)
	if err != nil {
		logger.WithError(err).Debugf("failed to stream call")
		return err
	}
	logger.Debugf("stream closed")

	return nil
}

type hijackStream struct {
	deviced_pb.DevicedService_ConnectClient

	on_recv       func(*hijackStream, *deviced_pb.ConnectRequest) error
	recv_chan     chan *deviced_pb.ConnectRequest
	recv_err      error
	recv_err_once *sync.Once
}

func (self *hijackStream) Recv() (*deviced_pb.ConnectRequest, error) {
	select {
	case req, ok := <-self.recv_chan:
		if !ok {
			self.recv_err_once.Do(func() { self.recv_err = ErrFailedToRecvMessage })
			return nil, self.recv_err
		}
		return req, nil
	}
}

func (self *hijackStream) RecvChan() chan *deviced_pb.ConnectRequest {
	return self.recv_chan
}

func (self *hijackStream) start() {
	go func() {
		defer close(self.recv_chan)
		for {
			req, err := self.DevicedService_ConnectClient.Recv()
			if err != nil {
				self.recv_err_once.Do(func() { self.recv_err = err })
				break
			}
			err = self.on_recv(self, req)
			if err != nil {
				self.recv_err_once.Do(func() { self.recv_err = err })
				break
			}
		}
	}()
}

func newHijackStream(stm deviced_pb.DevicedService_ConnectClient, on_recv func(*hijackStream, *deviced_pb.ConnectRequest) error) *hijackStream {
	hjstm := &hijackStream{
		DevicedService_ConnectClient: stm,
		on_recv:       on_recv,
		recv_chan:     make(chan *deviced_pb.ConnectRequest, 128),
		recv_err_once: new(sync.Once),
	}

	hjstm.start()

	return hjstm
}
