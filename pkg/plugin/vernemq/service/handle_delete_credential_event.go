package metathings_plugin_vernemq_service

import (
	"github.com/sirupsen/logrus"

	webhook_helper "github.com/nayotta/metathings/pkg/common/webhook"
)

func (s *VernemqPluginService) handleDeleteCredentialEvent(evt *webhook_helper.Event) error {
	logger := s.getLogger().WithField("#method", "handleDeleteCredentialEvent")

	credId := evt.GetString("credential.id")
	if credId == "" {
		logger.Warningf("invalid argument: credential.id")
		return ErrBadRequest
	}

	logger = logger.WithFields(logrus.Fields{
		"username":  credId,
		"client-id": credId,
	})

	err := s.storage.RemoveUser(credId, credId)
	if err != nil {
		logger.WithError(err).Debugf("failed to remove user from storage")
		return err
	}

	logger.Tracef("handle delete credential event")

	return nil
}
