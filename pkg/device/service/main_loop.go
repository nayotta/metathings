package metathings_device_service

import (
	log "github.com/sirupsen/logrus"

	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDeviceServiceImpl) main_loop() {
	var req *deviced_pb.ConnectRequest
	var err error

	defer self.conn_cfn()

	for {
		if req, err = self.conn_stm.Recv(); err != nil {
			self.logger.WithError(err).Errorf("failed to recv msg from connect stream")
			return
		}

		self.logger.WithFields(log.Fields{
			"session": req.SessionId,
			"kind":    req.Kind,
		}).Debugf("recv msg")
	}
}
