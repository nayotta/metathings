package metathings_deviced_service

import (
	"os"
	"os/signal"
	"syscall"
)

func (self *MetathingsDevicedService) HandleSignal() {
	logger := self.get_logger().WithField("method", "HandleSignal")

	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGTERM)
	logger.Debugf("register signal handler")

	<-c
	logger.Infof("recv signal: SIGTERM")

	if err := self.doTerminating(); err != nil {
		logger.WithError(err).Errorf("failed to do terminating")
	}

	close(self.quit_chan)
}
