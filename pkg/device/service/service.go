package metathings_device_service

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type MetathingsDeviceService interface {
	Start() error
	Stop() error
	Wait() chan bool
	Err() error
}

type MetathingsDeviceServiceOption struct {
	ModuleAliveTimeout time.Duration
	HeartbeatInterval  time.Duration
}

type MetathingsDeviceServiceImpl struct {
	tknr    token_helper.Tokener
	cli_fty *client_helper.ClientFactory
	logger  log.FieldLogger
	opt     *MetathingsDeviceServiceOption

	info     *deviced_pb.Device
	mdl_db   ModuleDatabase
	conn_stm deviced_pb.DevicedService_ConnectClient
	conn_cfn client_helper.CloseFn
}

func (self *MetathingsDeviceServiceImpl) context_with_token() context.Context {
	return context_helper.WithToken(context.Background(), self.tknr.GetToken())
}

func (self *MetathingsDeviceServiceImpl) Wait() chan bool {
	panic("unimplemented")
}

func (self *MetathingsDeviceServiceImpl) Err() error {
	panic("unimplemented")
}

func NewMetathingsDeviceService(
	tknr token_helper.Tokener,
	cli_fty *client_helper.ClientFactory,
	logger log.FieldLogger,
	opt *MetathingsDeviceServiceOption,
) (MetathingsDeviceService, error) {
	return &MetathingsDeviceServiceImpl{
		tknr:    tknr,
		cli_fty: cli_fty,
		logger:  logger,
		opt:     opt,
	}, nil
}
