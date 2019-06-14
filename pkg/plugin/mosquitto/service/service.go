package metathings_mosquitto_plugin_service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	webhook_helper "github.com/nayotta/metathings/pkg/common/webhook"
	storage "github.com/nayotta/metathings/pkg/plugin/mosquitto/storage"
)

type MosquittoPluginServiceOption struct {
	Webhook struct {
		Secret string
	}
}

type MosquittoPluginService struct {
	logger  log.FieldLogger
	opt     *MosquittoPluginServiceOption
	storage storage.Storage
}

func (s *MosquittoPluginService) get_logger() log.FieldLogger {
	return s.logger
}

func (s *MosquittoPluginService) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if !webhook_helper.ValidateHmac(s.opt.Webhook.Secret, r) {
		s.get_logger().Warningf("failed to validate hmac")
		return
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to read request body")
		return
	}

	evt := new(webhook_helper.Event)
	err = json.Unmarshal(buf, evt)
	if err != nil {
		s.get_logger().WithError(err).Debugf("bad webhook format")
		return
	}

	fmt.Println(evt)

	switch evt.Action() {
	case "create_credential":
	case "delete_credential":
	default:
		s.get_logger().WithField("action", evt.Action).Warningf("unexpected action")
	}
}

func NewMosquittoPluginService(
	opt *MosquittoPluginServiceOption,
	logger log.FieldLogger,
	storage storage.Storage,
) (*MosquittoPluginService, error) {
	return &MosquittoPluginService{
		opt:     opt,
		logger:  logger,
		storage: storage,
	}, nil
}
