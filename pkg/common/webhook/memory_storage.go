package webhook_helper

import (
	log "github.com/sirupsen/logrus"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type MemoryStorage struct {
	logger log.FieldLogger

	whs map[string]*Webhook
}

func (s *MemoryStorage) CreateWebhook(wh *Webhook) (*Webhook, error) {
	s.whs[*wh.Id] = deepcopy_webhook(wh)

	s.logger.Debugf("create webhook")

	return wh, nil
}

func (s *MemoryStorage) DeleteWebhook(id string) error {
	_, ok := s.whs[id]
	if !ok {
		return ErrWebhookNotFound
	}

	delete(s.whs, id)

	s.logger.WithField("webhook", id).Debugf("delete webhook")

	return nil
}

func (s *MemoryStorage) ListWebhooks(wh *Webhook) ([]*Webhook, error) {
	// TODO(Peer): apply webhook filter

	whs := []*Webhook{}
	for _, wh := range s.whs {
		whs = append(whs, deepcopy_webhook(wh))
	}

	s.logger.Debugf("list webhooks")

	return whs, nil
}

func (s *MemoryStorage) GetWebhook(id string) (*Webhook, error) {
	wh, err := s.get_webhook(id)
	if err != nil {
		return nil, err
	}

	s.logger.WithField("webhook", id).Debugf("get webhook")

	return wh, nil
}

func (s *MemoryStorage) get_webhook(id string) (*Webhook, error) {
	wh, ok := s.whs[id]
	if !ok {
		return nil, ErrWebhookNotFound
	}

	wh = deepcopy_webhook(wh)
	// erase secret field
	wh.Secret = nil

	return wh, nil
}

func (s *MemoryStorage) UpdateWebhook(id string, wh *Webhook) (*Webhook, error) {
	twh, ok := s.whs[id]
	if !ok {
		return nil, ErrWebhookNotFound
	}

	if wh.Url != nil {
		twh.Url = wh.Url
	}

	if wh.ContentType != nil {
		twh.ContentType = wh.ContentType
	}

	if wh.Secret != nil {
		twh.Secret = wh.Secret
	}

	s.whs[id] = deepcopy_webhook(twh)

	s.logger.WithField("webhook", id).Debugf("update webhook")

	return s.get_webhook(id)
}

type MemoryStorageFactory struct{}

func (*MemoryStorageFactory) New(args ...interface{}) (Storage, error) {
	var logger log.FieldLogger

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"logger": opt_helper.ToLogger(&logger),
	})(args...); err != nil {
		return nil, err
	}

	return &MemoryStorage{
		whs:    make(map[string]*Webhook),
		logger: logger,
	}, nil
}

func init() {
	register_storage_factory("memory", new(MemoryStorageFactory))
}
