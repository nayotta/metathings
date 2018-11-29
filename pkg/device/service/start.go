package metathings_device_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

func (self *MetathingsDeviceServiceImpl) start() error {
	var cli deviced_pb.DevicedServiceClient
	var show_device_res *deviced_pb.ShowDeviceResponse
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

	ctx = context_helper.WithToken(context.Background(), tkn_txt_str)
	show_device_res, err = cli.ShowDevice(ctx, &empty.Empty{})
	if err != nil {
		return err
	}

	self.mdl_db = NewModuleDatabase(show_device_res.GetDevice().GetModules(), self.opt.ModuleAliveTimeout)

	go self.heartbeat_loop()
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
