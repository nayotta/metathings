package metathings_mosquitto_plugin_service

import (
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"
	webhook_helper "github.com/nayotta/metathings/pkg/common/webhook"
	storage "github.com/nayotta/metathings/pkg/plugin/mosquitto/storage"
)

var (
	WEBHOOK_HMAC_TIMESTAMP, _    = time.Parse(time.RFC3339, "2019-01-01T00:00:00Z")
	WEBHOOK_HMAC_TIMESTAMP_INT64 = int64(1546300800000000000) // WEBHOOK_HMAC_TIMESTAMP.UnixNano()
	WEBHOOK_HMAC_NONCE           = int64(1024)
)

func ParseMosquittoPluginPassword(id, secret string) string {
	return passwd_helper.MustParseHmac(secret, id, WEBHOOK_HMAC_TIMESTAMP, WEBHOOK_HMAC_NONCE)
}

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

	evt, err := webhook_helper.UnmarshalEvent(buf)
	if err != nil {
		s.get_logger().WithError(err).Debugf("bad webhook format")
		return
	}

	act := evt.GetString("action")
	if act == "" {
		s.get_logger().WithField("action", act).Warningf("unexpected action")
		return
	}

	logger := s.get_logger().WithField("action", act)

	switch act {
	case "create_credential":
		err = s.handle_create_credential_event(evt)
	case "delete_credential":
		err = s.handle_delete_credential_event(evt)
	}

	if err != nil {
		logger.WithError(err).Errorf("failed to handle event")
		return
	}
}

func (s *MosquittoPluginService) handle_create_credential_event(evt *webhook_helper.Event) error {
	id := evt.GetString("credential.id")
	if id == "" {
		s.get_logger().Warningf("invaild argument: credential.id")
		return ErrBadRequest
	}

	secret := evt.GetString("credential.secret")
	if secret == "" {
		s.get_logger().Warningf("invalid argument: credential.secret")
		return ErrBadRequest
	}

	logger := s.get_logger().WithField("username", id)
	hmac := ParseMosquittoPluginPassword(id, secret)
	passwd := passwd_helper.MustParsePbkdf2(hmac)

	topic := "mt/#"
	mask := "rw"

	usr := &storage.User{
		Username: &id,
		Password: &passwd,
		Permissions: []*storage.Permission{
			&storage.Permission{
				Topic: &topic,
				Mask:  &mask,
			},
		},
	}
	err := s.storage.AddUser(usr)
	if err != nil {
		logger.Errorf("failed to add user in storage")
		return err
	}

	logger.Debugf("handle create credential event")

	return nil
}

func (s *MosquittoPluginService) handle_delete_credential_event(evt *webhook_helper.Event) error {
	id := evt.GetString("credential.id")
	if id == "" {
		s.get_logger().Warningf("invaild argument: crednetial.id")
		return ErrBadRequest
	}

	logger := s.get_logger().WithField("username", id)

	err := s.storage.RemoveUser(id)
	if err != nil {
		logger.Errorf("failed to remove user in storage")
		return err
	}

	logger.Debugf("handle delete credential event")

	return nil
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
