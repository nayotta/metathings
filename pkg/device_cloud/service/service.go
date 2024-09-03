package metathings_device_cloud_service

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	profile_storage_interface "github.com/nayotta/metathings/pkg/device_cloud/profile_storage/interface"
	storage "github.com/nayotta/metathings/pkg/device_cloud/storage"
)

type MetathingsDeviceCloudServiceOption struct {
	Session struct {
		Id string
	}
	Connection map[string]interface{}
	Credential struct {
		Id     string
		Secret string
	}
	Methods struct {
		PushFrameToFlow struct {
			PushAckTimeout     time.Duration
			PushAckFailedLimit int
		}
	}
}

type MetathingsDeviceCloudService struct {
	opt             *MetathingsDeviceCloudServiceOption
	logger          log.FieldLogger
	storage         storage.Storage
	profile_storage profile_storage_interface.ProfileStorage
	cli_fty         *client_helper.ClientFactory
	tknr            token_helper.Tokener
	tkvdr           token_helper.TokenValidator
}

func (s *MetathingsDeviceCloudService) get_logger() log.FieldLogger {
	return s.logger.WithField("#instance", "MetathingsDeviceCloudService")
}

func (s *MetathingsDeviceCloudService) context() context.Context {
	return context.Background()
}

func (s *MetathingsDeviceCloudService) context_with_token(ctx context.Context) context.Context {
	return context_helper.WithToken(ctx, s.tknr.GetToken())
}

func (s *MetathingsDeviceCloudService) context_with_device(dev_id string) context.Context {
	return context_helper.NewOutgoingContext(
		context.TODO(),
		context_helper.WithTokenOp(s.tknr.GetToken()),
		context_helper.WithDeviceOp(dev_id),
	)
}

func (s *MetathingsDeviceCloudService) get_session_id() string {
	return s.opt.Session.Id
}

func NewMetathingsDeviceCloudService(
	opt *MetathingsDeviceCloudServiceOption,
	logger log.FieldLogger,
	storage storage.Storage,
	profile_storage profile_storage_interface.ProfileStorage,
	cli_fty *client_helper.ClientFactory,
	tknr token_helper.Tokener,
	tkvdr token_helper.TokenValidator,
) (*MetathingsDeviceCloudService, error) {
	return &MetathingsDeviceCloudService{
		opt:             opt,
		logger:          logger,
		storage:         storage,
		profile_storage: profile_storage,
		cli_fty:         cli_fty,
		tknr:            tknr,
		tkvdr:           tkvdr,
	}, nil
}
