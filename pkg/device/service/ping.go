package metathings_device_service

import (
	"time"

	deviced_pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDeviceServiceImpl) ping_loop() {
	for {
		go self.ping_once()
		time.Sleep(self.opt.PingInterval)
	}
}

var (
	PING_PKT = &deviced_pb.ConnectResponse{
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
)

func (self *MetathingsDeviceServiceImpl) ping_once() {
	logger := self.get_logger().WithField("method", "ping_once")

	sessions := self.list_connection_sessions()
	for _, sess := range sessions {
		stm := self.get_connection(sess)
		if stm == nil {
			continue
		}

		if err := stm.Send(PING_PKT); err != nil {
			logger.WithError(err).Warningf("failed to send ping request")
			go self.try_close_connection(sess)
		}
		logger.WithField("session", sess).Debugf("ping")
	}
}
