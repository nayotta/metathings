package metathings_callback_sdk

import (
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/levigross/grequests"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type WebhookCallbackOption struct {
	AllowPlainText bool
	Insecure       bool
	UserAgent      string
	CustomHeaders  map[string]string
	TagPrefix      string
	Url            string
	Timeout        time.Duration
}

func NewWebhookCallbackOption() *WebhookCallbackOption {
	return &WebhookCallbackOption{
		AllowPlainText: false,
		Insecure:       false,
		UserAgent:      "Metathings/beta.v1 EvaluatorPluginWebhookClient",
		TagPrefix:      "X-MTE-Tag-",
		Timeout:        5 * time.Second,
	}
}

type WebhookCallback struct {
	opt *WebhookCallbackOption
}

func (cb *WebhookCallback) parse_http_header_key(x string) string {
	return strings.ReplaceAll(x, "_", "-")
}

func (cb *WebhookCallback) Emit(data interface{}, tags map[string]string) error {
	u, err := url.Parse(cb.opt.Url)
	if err != nil {
		return err
	}

	headers := map[string]string{}
	for k, v := range cb.opt.CustomHeaders {
		headers[cb.parse_http_header_key(k)] = v
	}
	for k, v := range tags {
		headers[cb.parse_http_header_key(cb.opt.TagPrefix+k)] = v
	}

	opt := &grequests.RequestOptions{
		JSON:                data,
		UserAgent:           cb.opt.UserAgent,
		Headers:             headers,
		DialTimeout:         cb.opt.Timeout,
		TLSHandshakeTimeout: cb.opt.Timeout,
		RequestTimeout:      cb.opt.Timeout,
	}

	switch u.Scheme {
	case "http":
		if !cb.opt.AllowPlainText {
			err = errors.New("webhook url not using https")
			return err
		}
	case "https":
		if cb.opt.Insecure {
			opt.InsecureSkipVerify = true
		}
	}

	res, err := grequests.Post(u.String(), opt)
	if err != nil {
		return err
	}

	if !res.Ok {
		return res.Error
	}

	return nil
}

func NewWebhookCallback(args ...interface{}) (Callback, error) {
	var err error
	opt := NewWebhookCallbackOption()

	if err = opt_helper.Setopt(map[string]func(string, interface{}) error{
		"allow_plain_text": opt_helper.ToBool(&opt.AllowPlainText),
		"insecure":         opt_helper.ToBool(&opt.Insecure),
		"useragent":        opt_helper.ToString(&opt.UserAgent),
		"custom_headers":   opt_helper.ToStringMapString(&opt.CustomHeaders),
		"tag_prefix":       opt_helper.ToString(&opt.TagPrefix),
		"url":              opt_helper.ToString(&opt.Url),
		"timeout":          opt_helper.ToDuration(&opt.Timeout),
	})(args...); err != nil {
		return nil, err
	}

	return &WebhookCallback{
		opt: opt,
	}, nil
}

func init() {
	register_callback_factory("default", NewWebhookCallback)
	register_callback_factory("webhook", NewWebhookCallback)
}
