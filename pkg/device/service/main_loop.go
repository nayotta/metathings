package metathings_device_service

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"

	deviced_pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDeviceServiceImpl) current_connection_count() int {
	return len(self.conns)
}

func (self *MetathingsDeviceServiceImpl) new_conn_signal() chan struct{} {
	return self.new_conn_ch
}

func (self *MetathingsDeviceServiceImpl) trigger_new_conn_signal() {
	self.new_conn_signal() <- struct{}{}
}

func (self *MetathingsDeviceServiceImpl) main_loop() {
	logger := self.get_logger().WithFields(log.Fields{
		"method": "main_loop",
	})

	go func() {
		t := time.NewTicker(self.opt.NewConnectionPeriod)
		defer t.Stop()

		for {
			<-t.C
			self.trigger_new_conn_signal()
		}
	}()

	var triggered_at time.Time
	for ; ; _ = <-self.new_conn_signal() {
		if time.Now().Sub(triggered_at) < self.opt.NewConnectionThreshold {
			logger.Debugf("new connection too quick")
			continue
		}

		actuals := self.current_connection_count()
		logger.WithFields(log.Fields{
			"actuals": actuals,
			"expects": self.opt.ExpectedConnections,
		}).Debugf("connection stats")

		if self.current_connection_count() < self.opt.ExpectedConnections {
			self.build_connection()
		}

		triggered_at = time.Now()
	}
}

func (self *MetathingsDeviceServiceImpl) build_connection() (err error) {
	logger := self.get_logger().WithField("method", "build_connection")

	var closer func() error
	cli, cfn, err := self.cli_fty.NewDevicedServiceClient()
	if err != nil {
		logger.WithError(err).Errorf("failed to connect to deviced service")
		return
	}
	closer = cfn
	defer func() {
		if closer != nil {
			closer()
		}
	}()

	sess := self.generator_major_session()
	ctx := self.context_with_sesion(context.TODO(), sess)
	conn, err := cli.Connect(ctx)
	if err != nil {
		logger.WithError(err).Errorf("failed to build connection")
		return
	}

	logger = logger.WithField("session", sess)

	self.add_connection(sess, conn, cfn)
	go self.connloop(sess)
	closer = nil

	logger.Debugf("build connection")

	return nil
}

func (self *MetathingsDeviceServiceImpl) setup_connection_nodename(sess int64, nodename string) error {
	self.conns_mtx.Lock()
	defer self.conns_mtx.Unlock()

	if !self.opt.ConnectToSameNode {
		exists := make(map[string]bool)
		for _, val := range self.nodes {
			exists[val] = true
		}
		if _, found := exists[nodename]; found {
			return ErrConnectToSameNode
		}
	}

	self.nodes[sess] = nodename

	return nil
}

func (self *MetathingsDeviceServiceImpl) remove_connection_nodename_nl(sess int64) {
	delete(self.nodes, sess)
}

func (self *MetathingsDeviceServiceImpl) init_connection(sess int64, conn deviced_pb.DevicedService_ConnectClient) error {
	logger := self.get_logger().WithFields(log.Fields{
		"method":  "init_connection",
		"session": sess,
	})

	// inner error
	var er error

	done := make(chan struct{})
	go func() {
		defer close(done)

		var req *deviced_pb.ConnectRequest

		for {
			req, er = conn.Recv()
			if er != nil {
				logger.WithError(er).Errorf("failed to receive message from conn")
				return
			}
			logger.Debugf("recv msg")

			uc := req.GetUnaryCall()
			component := uc.GetComponent().GetValue()
			name := uc.GetName().GetValue()
			method := uc.GetMethod().GetValue()
			if component != "system" ||
				name != "system" ||
				method != "nodename" {
				logger.WithFields(log.Fields{
					"component": component,
					"name":      name,
					"method":    method,
				}).Warningf("connection uninitialized, drop packet")
				continue
			}

			var nodename wrappers.StringValue
			if err := ptypes.UnmarshalAny(uc.GetValue(), &nodename); err != nil {
				logger.Warningf("failed to unmarshal any message")
				continue
			}
			inner_logger := logger.WithField("node", nodename.GetValue())

			if er = self.setup_connection_nodename(sess, nodename.GetValue()); er != nil {
				logger.WithError(er).Warningf("failed to setup connection nodename")
			}
			inner_logger.Debugf("setup connection nodename")

			return
		}
	}()

	timer := time.NewTimer(self.opt.InitConnectionTimeout)
	defer timer.Stop()

	for {
		select {
		case <-done:
			if er != nil {
				return er
			}

			return nil
		case <-timer.C:
			return ErrInitializeConnectionTimeout
		default:
		}

		hnpkt := &deviced_pb.ConnectResponse{
			Kind: deviced_pb.ConnectMessageKind_CONNECT_MESSAGE_KIND_SYSTEM,
			Union: &deviced_pb.ConnectResponse_UnaryCall{
				UnaryCall: &deviced_pb.UnaryCallValue{
					Component: "system",
					Name:      "system",
					Method:    "nodename",
				},
			},
		}

		if err := conn.Send(hnpkt); err != nil {
			return err
		}
		logger.Debugf("send nodename rqeuest")

		time.Sleep(self.opt.NodenameRequestPeriod)
	}
}

func (self *MetathingsDeviceServiceImpl) connloop(sess int64) {
	var err error

	defer self.try_close_connection(sess)

	conn := self.get_connection_nl(sess)
	logger := self.get_logger().WithFields(log.Fields{
		"method": "connloop",
	})

	if err = self.init_connection(sess, conn); err != nil {
		logger.WithError(err).Errorf("failed to initial connection")
		return
	}
	logger.Debugf("connection initailized")

	for {
		req, err := conn.Recv()
		if err != nil {
			logger.WithError(err).Errorf("failed to recv msg")
			return
		}

		logger.WithFields(log.Fields{
			"session": req.GetSessionId().GetValue(),
			"kind":    req.GetKind(),
		}).Debugf("recv msg")

		go self.handle(req)
	}
}

func (self *MetathingsDeviceServiceImpl) try_close_connection(sess int64) {
	logger := self.get_logger().WithFields(log.Fields{
		"method":  "try_close_connection",
		"session": sess,
	})

	self.conns_mtx.Lock()
	defer self.conns_mtx.Unlock()

	conn := self.get_connection_nl(sess)
	if conn == nil {
		logger.Warningf("connection closed")
		return
	}

	cfn := self.get_connection_close_fn_nl(sess)

	// close Connect streaming
	if err := conn.CloseSend(); err != nil {
		logger.WithError(err).Warningf("failed to close connection")
	}

	// put grpc connection back to pool
	if err := cfn(); err != nil {
		logger.WithError(err).Warningf("failed to close connection")
	}

	self.remove_connection_nl(sess)
	self.remove_connection_close_fn_nl(sess)
	self.remove_connection_nodename_nl(sess)

	logger.Debugf("close connection")

	self.trigger_new_conn_signal()
}

func (self *MetathingsDeviceServiceImpl) add_connection(session int64, conn deviced_pb.DevicedService_ConnectClient, cfn func() error) {
	self.conns_mtx.Lock()
	defer self.conns_mtx.Unlock()

	self.add_connection_nl(session, conn)
	self.add_connection_close_fn_nl(session, cfn)
}

func (self *MetathingsDeviceServiceImpl) add_connection_nl(session int64, conn deviced_pb.DevicedService_ConnectClient) {
	self.conns[session] = conn
}

func (self *MetathingsDeviceServiceImpl) get_connection(session int64) deviced_pb.DevicedService_ConnectClient {
	self.conns_mtx.Lock()
	defer self.conns_mtx.Unlock()

	return self.get_connection_nl(session)
}

func (self *MetathingsDeviceServiceImpl) get_connection_nl(session int64) deviced_pb.DevicedService_ConnectClient {
	return self.conns[session]
}

func (self *MetathingsDeviceServiceImpl) remove_connection_nl(session int64) {
	delete(self.conns, session)
}

func (self *MetathingsDeviceServiceImpl) add_connection_close_fn_nl(session int64, cfn func() error) {
	self.close_fns[session] = cfn
}

func (self *MetathingsDeviceServiceImpl) get_connection_close_fn_nl(session int64) func() error {
	return self.close_fns[session]
}

func (self *MetathingsDeviceServiceImpl) remove_connection_close_fn_nl(session int64) {
	delete(self.close_fns, session)
}

func (self *MetathingsDeviceServiceImpl) list_connection_sessions() []int64 {
	self.conns_mtx.Lock()
	defer self.conns_mtx.Unlock()

	return self.list_connection_sessions_nl()
}

func (self *MetathingsDeviceServiceImpl) list_connection_sessions_nl() []int64 {
	var sessions []int64

	for sess := range self.conns {
		sessions = append(sessions, sess)
	}

	return sessions
}
