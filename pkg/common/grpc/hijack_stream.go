package grpc_helper

import (
	"errors"
	"sync"

	deviced_pb "github.com/nayotta/metathings/proto/deviced"
)

var (
	ErrFailedToRecvMessage = errors.New("failed to recv message")
)

type HijackStream struct {
	deviced_pb.DevicedService_ConnectClient

	on_recv       func(*HijackStream, *deviced_pb.ConnectRequest) error
	recv_chan     chan *deviced_pb.ConnectRequest
	recv_err      error
	recv_err_once *sync.Once
}

func (self *HijackStream) Recv() (*deviced_pb.ConnectRequest, error) {
	select {
	case req, ok := <-self.recv_chan:
		if !ok {
			self.recv_err_once.Do(func() { self.recv_err = ErrFailedToRecvMessage })
			return nil, self.recv_err
		}
		return req, nil
	}
}

func (self *HijackStream) RecvChan() chan *deviced_pb.ConnectRequest {
	return self.recv_chan
}

func (self *HijackStream) start() {
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

func NewHijackStream(stm deviced_pb.DevicedService_ConnectClient, on_recv func(*HijackStream, *deviced_pb.ConnectRequest) error) *HijackStream {
	hjstm := &HijackStream{
		DevicedService_ConnectClient: stm,
		on_recv:                      on_recv,
		recv_chan:                    make(chan *deviced_pb.ConnectRequest, 128),
		recv_err_once:                new(sync.Once),
	}

	hjstm.start()

	return hjstm
}
