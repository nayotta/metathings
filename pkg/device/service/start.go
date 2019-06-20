package metathings_device_service

import (
	"github.com/golang/protobuf/ptypes/empty"
)

func (self *MetathingsDeviceServiceImpl) pre_start() error {
	cli, cfn, err := self.cli_fty.NewDevicedServiceClient()
	if err != nil {
		return err
	}
	defer cfn()

	res, err := cli.ShowDevice(self.context(), &empty.Empty{})
	if err != nil {
		return err
	}

	self.mdl_db = NewModuleDatabase(res.GetDevice().GetModules(), self.opt.ModuleAliveTimeout, self.logger)
	self.info = res.GetDevice()

	return nil
}

func (self *MetathingsDeviceServiceImpl) start() error {
	var err error

	if err = self.pre_start(); err != nil {
		return err
	}

	self.conn_stm_wg.Add(1)
	go self.main_loop()
	go self.heartbeat_loop()
	go self.ping_loop()

	return nil
}

func (self *MetathingsDeviceServiceImpl) Start() error {
	var err error

	if err = self.start(); err != nil {
		self.logger.WithError(err).Errorf("failed to start device")
		return err
	}

	self.logger.Infof("device start")

	return nil
}
