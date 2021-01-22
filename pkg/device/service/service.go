package metathings_device_service

import (
	"context"
	"sync"
	"time"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	log "github.com/sirupsen/logrus"

	afo_helper "github.com/nayotta/metathings/pkg/common/auth_func_overrider"
	"github.com/nayotta/metathings/pkg/common/binary_synchronizer"
	bin_sync "github.com/nayotta/metathings/pkg/common/binary_synchronizer"
	client_helper "github.com/nayotta/metathings/pkg/common/client"
	fx_helper "github.com/nayotta/metathings/pkg/common/fx"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	session_helper "github.com/nayotta/metathings/pkg/common/session"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	version_helper "github.com/nayotta/metathings/pkg/common/version"
	pb "github.com/nayotta/metathings/proto/device"
	deviced_pb "github.com/nayotta/metathings/proto/deviced"
)

type MetathingsDeviceService interface {
	version_helper.Versioner
	pb.DeviceServiceServer
	Start() error
	Stop() error
}

type MetathingsDeviceServiceOption struct {
	ModuleAliveTimeout   time.Duration
	HeartbeatInterval    time.Duration
	HeartbeatMaxRetry    int
	MaxReconnectInterval time.Duration
	MinReconnectInterval time.Duration
	PingInterval         time.Duration
}

type MetathingsDeviceServiceImpl struct {
	grpc_auth.ServiceAuthFuncOverride
	version_helper.Versioner
	*grpc_helper.ErrorParser
	tknr       token_helper.Tokener
	cli_fty    *client_helper.ClientFactory
	logger     log.FieldLogger
	opt        *MetathingsDeviceServiceOption
	app_getter *fx_helper.FxAppGetter
	bs         bin_sync.BinarySynchronizer

	info             *deviced_pb.Device
	mdl_db           ModuleDatabase
	conn_stm         deviced_pb.DevicedService_ConnectClient
	conn_stm_rwmtx   sync.RWMutex
	conn_stm_wg      sync.WaitGroup
	conn_stm_wg_once sync.Once
	conn_cfn         client_helper.CloseFn

	startup_session int32

	stats_heartbeat_fails int

	synchronizing_firmware_mtx   sync.Mutex
	stats_synchronizing_firmware bool
}

var (
	ignore_methods = []string{
		"IssueModuleToken",
	}
)

func (self *MetathingsDeviceServiceImpl) IsIgnoreMethod(md *grpc_helper.MethodDescription) bool {
	for _, m := range ignore_methods {
		if md.Method == m {
			return true
		}
	}

	return false
}

func (self *MetathingsDeviceServiceImpl) Stop() error {
	return self.app_getter.Get().Stop(context.TODO())
}

func (self *MetathingsDeviceServiceImpl) connection_stream() deviced_pb.DevicedService_ConnectClient {
	return self.conn_stm
}

func (self *MetathingsDeviceServiceImpl) get_module_info(id string) (*deviced_pb.Module, error) {
	for _, m := range self.info.Modules {
		if m.Id == id {
			return m, nil
		}
	}

	return nil, ErrModuleNotFound
}

func (self *MetathingsDeviceServiceImpl) get_logger() log.FieldLogger {
	return self.logger
}

func NewMetathingsDeviceService(
	ver version_helper.Versioner,
	tknr token_helper.Tokener,
	cli_fty *client_helper.ClientFactory,
	logger log.FieldLogger,
	tkvdr token_helper.TokenValidator,
	opt *MetathingsDeviceServiceOption,
	app_getter *fx_helper.FxAppGetter,
	bs binary_synchronizer.BinarySynchronizer,
) (MetathingsDeviceService, error) {
	srv := &MetathingsDeviceServiceImpl{
		ErrorParser:     grpc_helper.NewErrorParser(em),
		Versioner:       ver,
		bs:              bs,
		tknr:            tknr,
		cli_fty:         cli_fty,
		logger:          logger,
		opt:             opt,
		startup_session: session_helper.GenerateStartupSession(),
		app_getter:      app_getter,

		stats_synchronizing_firmware: false,
	}
	srv.ServiceAuthFuncOverride = afo_helper.NewAuthFuncOverrider(tkvdr, srv, logger)

	return srv, nil
}
