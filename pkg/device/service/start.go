package metathings_device_service

import (
	"context"

	context_helper "github.com/nayotta/metathings/pkg/common/context"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDeviceServiceImpl) start() error {
	var cli deviced_pb.DevicedServiceClient
	var err error

	tkn_txt_str := self.tknr.GetToken()
	cli, self.conn_cfn, err = self.cli_fty.NewDevicedServiceClient()
	if err != nil {
		return err
	}

	ctx := context_helper.WithToken(context.Background(), tkn_txt_str)
	self.conn_stm, err = cli.Connect(ctx)
	if err != nil {
		return err
	}

	go self.main_loop()

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
