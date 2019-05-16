package metathings_device_service

import (
	"time"

	log "github.com/sirupsen/logrus"

	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDeviceServiceImpl) main_loop() {
	for {
		self.internal_main_loop()
		time.Sleep(self.opt.ReconnectInterval)
	}
}

func (self *MetathingsDeviceServiceImpl) internal_main_loop() {
	var err error
	var req *deviced_pb.ConnectRequest

	// build connection
	cli, cfn, err := self.cli_fty.NewDevicedServiceClient()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to connect to deviced service")
		return
	}
	defer cfn()
	self.conn_stm_wg_once.Do(func() { self.conn_stm_wg.Done() })

	ctx := self.context_with_sesion()
	self.conn_stm_rwmtx.Lock()
	self.conn_stm, err = cli.Connect(ctx)
	self.conn_stm_rwmtx.Unlock()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to build connection to deviced")
		return
	}

	// handle message loop
	for {
		self.conn_stm_rwmtx.RLock()
		if req, err = self.connection_stream().Recv(); err != nil {
			self.logger.WithError(err).Errorf("failed to recv message from connection stream")
			self.conn_stm_rwmtx.RUnlock()
			return
		}
		self.conn_stm_rwmtx.RUnlock()

		self.logger.WithFields(log.Fields{
			"session": req.GetSessionId().GetValue(),
			"kind":    req.GetKind(),
		}).Debugf("recv msg")

		go self.handle(req)
	}
}
