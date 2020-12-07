package metathings_device_service

import (
	"time"

	deviced_pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDeviceServiceImpl) ping_loop() {
	self.conn_stm_wg.Wait()
	for {
		go self.ping_once()
		time.Sleep(self.opt.PingInterval)
	}
}

func (self *MetathingsDeviceServiceImpl) ping_once() {
	self.conn_stm_rwmtx.Lock()
	stm := self.connection_stream()
	self.conn_stm_rwmtx.Unlock()

	ping_pkt := &deviced_pb.ConnectResponse{
		SessionId: 0,
		Kind:      deviced_pb.ConnectMessageKind_CONNECT_MESSAGE_KIND_SYSTEM,
		Union: &deviced_pb.ConnectResponse_UnaryCall{
			UnaryCall: &deviced_pb.UnaryCallValue{
				Name:      "system",
				Component: "system",
				Method:    "ping",
				Value:     nil,
			},
		},
	}

	err := stm.Send(ping_pkt)
	if err != nil {
		// TODO(Peer): reconnect streaming, not stop device and restart
		defer self.Stop()

		self.logger.WithError(err).Warningf("failed to send ping request")
		return
	}

	self.logger.Debugf("sending ping request")
}
