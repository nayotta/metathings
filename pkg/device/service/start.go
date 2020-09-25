package metathings_device_service

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"

	deviced_pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDeviceServiceImpl) set_device_version_to_simple_storage(cli deviced_pb.DevicedServiceClient) error {
	_, err := cli.PutObject(self.context(), &deviced_pb.PutObjectRequest{
		Object: &deviced_pb.OpObject{
			Device: &deviced_pb.OpDevice{
				Id: &wrappers.StringValue{Value: self.info.Id},
			},
			Prefix: &wrappers.StringValue{Value: "/sys/firmware/device/version"},
			Name:   &wrappers.StringValue{Value: "current"},
		},
		Content: &wrappers.BytesValue{Value: []byte(self.GetVersion())},
	})
	return err
}

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

	if err = self.set_device_version_to_simple_storage(cli); err != nil {
		return err
	}

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
