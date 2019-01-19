package metathings_device_service

import (
	"time"

	log "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	pb "github.com/nayotta/metathings/pkg/proto/device"
	deviced_pb "github.com/nayotta/metathings/pkg/proto/deviced"
)

type MetathingsDeviceService interface {
	pb.DeviceServiceServer
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

	startup_session int32
}

func (self *MetathingsDeviceServiceImpl) Stop() error {
	panic("unimplemented")
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
		tknr:            tknr,
		cli_fty:         cli_fty,
		logger:          logger,
		opt:             opt,
		startup_session: rand_helper.Int63(),
	}, nil
}
