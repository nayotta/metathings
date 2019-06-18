package webhook_helper

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"
)

type Webhook struct {
	Id          *string
	Url         *string
	ContentType *string
	Secret      *string
}

type WebhookService interface {
	Add(*Webhook) (*Webhook, error)
	Remove(id string) error
	List(*Webhook) ([]*Webhook, error)
	Get(id string) (*Webhook, error)
	Update(id string, wh *Webhook) (*Webhook, error)
	Trigger(evt interface{}) error
}

type WebhookServiceOption struct {
	ContentType string
}

type webhookService struct {
	opt     *WebhookServiceOption
	storage Storage

	logger log.FieldLogger
}

func (s *webhookService) get_logger() log.FieldLogger {
	return s.logger
}

func (s *webhookService) Add(wh *Webhook) (*Webhook, error) {
	if wh.Id == nil {
		id := id_helper.NewId()
		wh.Id = &id
	}

	if wh.ContentType == nil || *wh.ContentType == "" {
		wh.ContentType = &s.opt.ContentType
	}
	*wh.Secret = base64.StdEncoding.EncodeToString([]byte(*wh.Secret))

	wh, err := s.storage.CreateWebhook(wh)
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to add webhook to storage")
		return nil, err
	}

	s.get_logger().Debugf("add webhook")
	return wh, nil
}

func (s *webhookService) Remove(id string) error {
	err := s.storage.DeleteWebhook(id)
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to remove webhook from storage")
		return err
	}

	s.get_logger().Debugf("remove webhook")
	return nil
}

func (s *webhookService) List(wh *Webhook) ([]*Webhook, error) {
	whs, err := s.storage.ListWebhooks(wh)
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to list webhooks from storage")
		return nil, err
	}

	s.get_logger().Debugf("list webhooks")
	return whs, nil
}

func (s *webhookService) Get(id string) (*Webhook, error) {
	wh, err := s.storage.GetWebhook(id)
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to get webhook from storage")
		return nil, err
	}

	s.get_logger().Debugf("get webhook")
	return wh, nil
}

func (s *webhookService) Update(id string, wh *Webhook) (*Webhook, error) {
	wh, err := s.storage.UpdateWebhook(id, wh)
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to update webhook in storage")
		return nil, err
	}

	s.get_logger().Debugf("update webhook")
	return wh, nil
}

func (s *webhookService) Trigger(evt interface{}) error {
	whs, err := s.storage.ListWebhooks(nil)
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to list webhooks in storage")
		return err
	}

	buf, err := json.Marshal(evt)
	if err != nil {
		s.get_logger().WithError(err).Debugf("failed to marshal event to json")
		return err
	}

	for _, wh := range whs {
		go func(wh *Webhook) {
			req, err := http.NewRequest("POST", *wh.Url, bytes.NewReader(buf))
			if err != nil {
				s.get_logger().WithError(err).Warningf("failed to new http request")
				return
			}

			ct := s.opt.ContentType
			if wh.ContentType != nil {
				ct = *wh.ContentType
			}
			req.Header.Set("Content-Type", ct)
			s.set_webhook_request_header(wh, req)

			if _, err := http.DefaultClient.Do(req); err != nil {
				s.get_logger().WithError(err).Warningf("failed to trigger event")
			}
		}(wh)
	}

	s.get_logger().Debugf("trigger event")
	return nil
}

func (s *webhookService) set_webhook_request_header(wh *Webhook, req *http.Request) {
	ts := time.Now()
	nonce := rand.Int63()
	hmac := passwd_helper.MustParseHmac(*wh.Secret, *wh.Id, ts, nonce)

	req.Header.Set("MT-Webhook-Id", *wh.Id)
	req.Header.Set("MT-Webhook-Timestamp", fmt.Sprintf("%v", ts.UnixNano()))
	req.Header.Set("MT-Webhook-Nonce", fmt.Sprintf("%v", nonce))
	req.Header.Set("MT-Webhook-HMAC", hmac)
}

func NewWebhookService(opt *WebhookServiceOption, args ...interface{}) (WebhookService, error) {
	var logger log.FieldLogger
	var storage Storage

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger": opt_helper.ToLogger(&logger),
		"storage": func(_ string, val interface{}) error {
			var ok bool
			storage, ok = val.(Storage)
			if !ok {
				return opt_helper.InvalidArgument("storage")
			}
			return nil
		},
	})(args...); err != nil {
		return nil, err
	}

	return &webhookService{
		opt:     opt,
		storage: storage,
		logger:  logger,
	}, nil
}
