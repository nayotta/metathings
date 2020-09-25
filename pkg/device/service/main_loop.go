package metathings_device_service

import (
	"math"
	"time"

	log "github.com/sirupsen/logrus"

	session_helper "github.com/nayotta/metathings/pkg/common/session"
	deviced_pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDeviceServiceImpl) main_loop() {
	rc_tvl := self.opt.MinReconnectInterval
	for {
		err := self.internal_main_loop()
		if err != nil {
			rc_tvl = time.Duration(math.Min(float64(rc_tvl*2), float64(self.opt.MaxReconnectInterval)))
		} else {
			rc_tvl = self.opt.MinReconnectInterval
		}
		time.Sleep(rc_tvl)
	}
}

func (self *MetathingsDeviceServiceImpl) _refresh_startup_session() {
	self.startup_session = session_helper.GenerateStartupSession()
}

func (self *MetathingsDeviceServiceImpl) internal_main_loop() error {
	var err error
	var req *deviced_pb.ConnectRequest

	// build connection
	cli, cfn, err := self.cli_fty.NewDevicedServiceClient()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to connect to deviced service")
		return err
	}
	defer cfn()

	// TODO(Peer): DONT refresh startup session
	self._refresh_startup_session()

	ctx := self.context_with_sesion()
	self.conn_stm_rwmtx.Lock()
	self.conn_stm, err = cli.Connect(ctx)
	self.conn_stm_rwmtx.Unlock()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to build connection to deviced")
		return err
	}
	self.conn_stm_wg_once.Do(func() {
		time.Sleep(200 * time.Millisecond)
		self.conn_stm_wg.Done()
	})

	// handle message loop
	for {
		self.conn_stm_rwmtx.RLock()
		conn := self.connection_stream()
		self.conn_stm_rwmtx.RUnlock()
		if req, err = conn.Recv(); err != nil {
			self.logger.WithError(err).Errorf("failed to recv message from connection stream")
			return nil
		}

		self.logger.WithFields(log.Fields{
			"session": req.GetSessionId().GetValue(),
			"kind":    req.GetKind(),
		}).Debugf("recv msg")

		go self.handle(req)
	}
}
