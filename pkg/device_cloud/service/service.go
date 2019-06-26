package metathings_device_cloud_service

import (
	"context"

	log "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	storage "github.com/nayotta/metathings/pkg/device_cloud/storage"
)

type MetathingsDeviceCloudServiceOption struct {
	Session struct {
		Id string
	}
}

type MetathingsDeviceCloudService struct {
	opt     *MetathingsDeviceCloudServiceOption
	logger  log.FieldLogger
	storage storage.Storage
	cli_fty *client_helper.ClientFactory
	tknr    token_helper.Tokener
}

func (s *MetathingsDeviceCloudService) get_logger() log.FieldLogger {
	return s.logger
}

func (s *MetathingsDeviceCloudService) context() context.Context {
	return context_helper.WithToken(context.TODO(), s.tknr.GetToken())
}

func (s *MetathingsDeviceCloudService) get_session_id() string {
	return s.opt.Session.Id
}
