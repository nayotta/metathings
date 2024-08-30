package metathings_plugin_vernemq_service

import (
	mqtt_helper "github.com/nayotta/metathings/pkg/common/mqtt"
	passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"
	webhook_helper "github.com/nayotta/metathings/pkg/common/webhook"
	storage "github.com/nayotta/metathings/pkg/plugin/vernemq/storage"
)

func (s *VernemqPluginService) handleCreateCredentialEvent(evt *webhook_helper.Event) error {
	logger := s.getLogger().WithField("#method", "handleCreateCredentialEvent")

	credId := evt.GetString("credential.id")
	if credId == "" {
		logger.Warningf("invalid argument: credential id")
		return ErrBadRequest
	}

	credSecret := evt.GetString("credential.secret")
	if credSecret == "" {
		logger.Warningf("invalid argument: credential secret")
		return ErrBadRequest
	}

	logger = logger.WithField("username", credId)
	passwd := mqtt_helper.ParseMqttPassword(credId, credSecret)
	passhash := passwd_helper.MustParseBcrypt(passwd)
	topic := "mt/#"
	usr := &storage.User{
		ClientID: credId,
		Username: credId,
		Passhash: passhash,
		PublishAcls: []storage.PublishAcl{
			{Pattern: topic},
		},
		SubscribeAcls: []storage.SubscribeAcl{
			{Pattern: topic},
		},
	}

	err := s.storage.CreateOrUpdateUser(usr)
	if err != nil {
		logger.WithError(err).Debugf("failed to add user to storage")
		return err
	}

	logger.Tracef("handle create credential event")

	return nil
}
