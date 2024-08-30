package metathings_plugin_vernemq_service

import (
	"io"
	"net/http"

	webhook_helper "github.com/nayotta/metathings/pkg/common/webhook"
)

func (s *VernemqPluginService) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	logger := s.getLogger().WithField("#method", "WebhookHandler")
	defer w.WriteHeader(http.StatusOK)

	if !webhook_helper.ValidateHmac(s.webhookSecret, r) {
		logger.Warningf("failed to validate request")
		return
	}

	buf, err := io.ReadAll(r.Body)
	if err != nil {
		logger.WithError(err).Debugf("failed to read request body")
		return
	}
	defer r.Body.Close()

	evt, err := webhook_helper.UnmarshalEvent(buf)
	if err != nil {
		logger.WithError(err).Debugf("bad webhook format")
		return
	}

	act := evt.GetString("action")
	logger = logger.WithField("action", act)
	switch act {
	case "create_credential":
		err = s.handleCreateCredentialEvent(evt)
	case "delete_credential":
		err = s.handleDeleteCredentialEvent(evt)
	default:
		logger.Warningf("unsupported action")
		return
	}

	if err != nil {
		logger.WithError(err).Errorf("failed to handle event")
		return
	}

	logger.Tracef("handle webhook request")
}
