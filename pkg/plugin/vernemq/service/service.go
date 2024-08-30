package metathings_plugin_vernemq_service

import (
	"encoding/base64"

	"github.com/PeerXu/option-go"
	"github.com/sirupsen/logrus"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
	storage "github.com/nayotta/metathings/pkg/plugin/vernemq/storage"
)

type VernemqPluginService struct {
	webhookSecret string
	logger        logrus.FieldLogger
	storage       storage.Storage
}

func NewVernemqPluginService(opts ...option.ApplyOption) (*VernemqPluginService, error) {
	o := option.Apply(opts...)

	webhookSecret, err := GetWebhookSecret(o)
	if err != nil {
		return nil, err
	}

	logger, err := log_helper.GetLogger(o)
	if err != nil {
		return nil, err
	}

	storage, err := storage.GetStorage(o)
	if err != nil {
		return nil, err
	}

	return &VernemqPluginService{
		webhookSecret: base64.StdEncoding.EncodeToString([]byte(webhookSecret)),
		logger:        logger,
		storage:       storage,
	}, nil
}
