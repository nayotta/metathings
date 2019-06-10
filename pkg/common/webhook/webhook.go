package webhook_helper

import (
	"bytes"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type Webhook struct {
	Id          *string
	Url         *string
	ContentType *string
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

	if wh.ContentType == nil {
		wh.ContentType = &s.opt.ContentType
	}

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
			if _, err := http.Post(*wh.Url, s.opt.ContentType, bytes.NewReader(buf)); err != nil {
				s.get_logger().WithError(err).Warningf("failed to trigger event")
			}
		}(wh)
	}

	s.get_logger().Debugf("trigger event")
	return nil
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
