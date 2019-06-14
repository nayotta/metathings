package cmd_contrib

import (
	log "github.com/sirupsen/logrus"

	webhook_helper "github.com/nayotta/metathings/pkg/common/webhook"
)

type WebhookOptioner interface {
	GetContentTypeP() *string
	GetContentType() string
	SetContentType(string)

	GetUrlP() *string
	GetUrl() string
	SetUrl(string)

	GetSecretP() *string
	GetSecret() string
	SetSecret(string)
}

type WebhookOption struct {
	Url         string `mapstructure:"url"`
	ContentType string `mapstructure:"content_type"`
	Secret      string `mapstructure:"secret"`
}

func (o *WebhookOption) GetContentTypeP() *string {
	return &o.ContentType
}

func (o *WebhookOption) GetContentType() string {
	return o.ContentType
}

func (o *WebhookOption) SetContentType(v string) {
	o.ContentType = v
}

func (o *WebhookOption) GetUrlP() *string {
	return &o.Url
}

func (o *WebhookOption) GetUrl() string {
	return o.Url
}

func (o *WebhookOption) SetUrl(v string) {
	o.Url = v
}

func (o *WebhookOption) GetSecretP() *string {
	return &o.Secret
}

func (o *WebhookOption) GetSecret() string {
	return o.Secret
}

func (o *WebhookOption) SetSecret(v string) {
	o.Secret = v
}

type WebhookServiceOptioner interface {
	GetWebhook(string) WebhookOptioner
	Keys() []string
}

type WebhookServiceOption struct {
	WebhookService map[string]*WebhookOption `mapstructure:"webhook_service"`
}

func (o *WebhookServiceOption) GetWebhook(v string) WebhookOptioner {
	wh, ok := o.WebhookService[v]
	if !ok {
		return nil
	}
	return wh
}

func (o *WebhookServiceOption) Keys() []string {
	ks := []string{}
	for k, _ := range o.WebhookService {
		ks = append(ks, k)
	}
	return ks
}

func CreateWebhookServiceOption() WebhookServiceOption {
	return WebhookServiceOption{
		WebhookService: make(map[string]*WebhookOption),
	}
}

func NewWebhookService(opt WebhookServiceOptioner, logger log.FieldLogger) (webhook_helper.WebhookService, error) {
	o := &webhook_helper.WebhookServiceOption{}

	dwh := opt.GetWebhook("default")
	if dwh == nil {
		logger.Warningf("missing default webhook config")
	} else {
		o.ContentType = dwh.GetContentType()
	}

	stor, err := webhook_helper.NewStorage("memory", "logger", logger)
	if err != nil {
		return nil, err
	}

	whs, err := webhook_helper.NewWebhookService(o, "logger", logger, "storage", stor)
	if err != nil {
		return nil, err
	}

	for _, k := range opt.Keys() {
		if k != "default" {
			wh := opt.GetWebhook(k)
			wh_url := wh.GetUrl()
			wh_ct := wh.GetContentType()
			wh_srt := wh.GetSecret()

			if _, err = whs.Add(&webhook_helper.Webhook{
				Url:         &wh_url,
				ContentType: &wh_ct,
				Secret:      &wh_srt,
			}); err != nil {
				return nil, err
			}
		}
	}

	return whs, nil
}
