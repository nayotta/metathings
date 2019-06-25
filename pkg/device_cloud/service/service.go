package metathings_device_cloud_service

import (
	"context"

	log "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	storage "github.com/nayotta/metathings/pkg/device_cloud/storage"
)

type MetathingsDeviceCloudServiceOption struct {
	SessionID string
}

type MetathingsDeviceCloudService struct {
	opt     *MetathingsDeviceCloudServiceOption
	logger  log.FieldLogger
	storage storage.Storage
	cli_fty client_helper.ClientFactory
}

func (s *MetathingsDeviceCloudService) get_logger() log.FieldLogger {
	return s.logger
}

func (s *MetathingsDeviceCloudService) context() context.Context {
	panic("unimplemented")
}

func (s *MetathingsDeviceCloudService) get_session_id() string {
	return s.opt.SessionID
}
