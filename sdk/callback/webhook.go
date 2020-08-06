package metathings_callback_sdk

import (
	"errors"
	"net/url"
	"time"

	"github.com/levigross/grequests"
	"github.com/sirupsen/logrus"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type WebhookCallbackOption struct {
	AllowPlainText bool
	Insecure       bool
	UserAgent      string
	CustomHeaders  map[string]string
	Url            string
	Timeout        time.Duration
}

func NewWebhookCallbackOption() *WebhookCallbackOption {
	return &WebhookCallbackOption{
		AllowPlainText: false,
		Insecure:       false,
		UserAgent:      "Metathings/beta.v1 EvaluatorPluginWebhookClient",
		Timeout:        5 * time.Second,
	}
}

type WebhookCallback struct {
	opt    *WebhookCallbackOption
	logger logrus.FieldLogger
}

func (cb *WebhookCallback) get_logger() logrus.FieldLogger {
	return cb.logger.WithFields(logrus.Fields{
		"custom_headers": cb.opt.CustomHeaders,
		"url":            cb.opt.Url,
	})
}

func (cb *WebhookCallback) Emit(data interface{}) error {
	logger := cb.get_logger()

	u, err := url.Parse(cb.opt.Url)
	if err != nil {
		logger.WithError(err).Debugf("failed to parse webhook url")
		return err
	}

	opt := &grequests.RequestOptions{
		JSON:                data,
		UserAgent:           cb.opt.UserAgent,
		Headers:             cb.opt.CustomHeaders,
		DialTimeout:         cb.opt.Timeout,
		TLSHandshakeTimeout: cb.opt.Timeout,
		RequestTimeout:      cb.opt.Timeout,
	}

	switch u.Scheme {
	case "http":
		if !cb.opt.AllowPlainText {
			err = errors.New("webhook url not using https")
			logger.WithError(err).Debugf(err.Error())
			return err
		}
	case "https":
		if cb.opt.Insecure {
			opt.InsecureSkipVerify = true
		}
	}

	res, err := grequests.Post(u.String(), opt)
	if err != nil {
		logger.WithError(err).Debugf("failed to send webhook callback")
		return err
	}

	if !res.Ok {
		err = res.Error
		logger.WithError(err).Debugf("failed to send webhook callback")
		return err
	}

	return nil
}

func NewWebhookCallback(args ...interface{}) (Callback, error) {
	var logger logrus.FieldLogger
	var err error
	opt := NewWebhookCallbackOption()

	if err = opt_helper.Setopt(map[string]func(string, interface{}) error{
		"allow_plain_text": opt_helper.ToBool(&opt.AllowPlainText),
		"insecure":         opt_helper.ToBool(&opt.Insecure),
		"useragent":        opt_helper.ToString(&opt.UserAgent),
		"custom_headers":   opt_helper.ToStringMapString(&opt.CustomHeaders),
		"url":              opt_helper.ToString(&opt.Url),
		"timeout":          opt_helper.ToDuration(&opt.Timeout),
		"logger":           opt_helper.ToLogger(&logger),
	})(args...); err != nil {
		return nil, err
	}

	return &WebhookCallback{
		opt:    opt,
		logger: logger,
	}, nil
}

func init() {
	register_callback_factory("default", NewWebhookCallback)
	register_callback_factory("webhook", NewWebhookCallback)
}
