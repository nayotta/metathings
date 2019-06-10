package webhook_helper

import (
	"github.com/stretchr/testify/suite"
)

var (
	test_default_webhook_id           = "default"
	test_default_webhook_url          = "http://www.example.com/webhook"
	test_default_webhook_content_type = "application/json"
	test_default_webhook_secret       = "c2VjcmV0" // base64("secret")
	test_default_webhook              = &Webhook{
		Id:          &test_default_webhook_id,
		Url:         &test_default_webhook_url,
		ContentType: &test_default_webhook_content_type,
		Secret:      &test_default_webhook_secret,
	}

	test_webhook_id           = "test"
	test_webhook_url          = "http://www.example1.com/webhook"
	test_webhook_content_type = "application/json"
	test_webhook_secret       = "c2VjcmV0" // base64("secret")
	test_webhook              = &Webhook{
		Id:          &test_webhook_id,
		Url:         &test_webhook_url,
		ContentType: &test_webhook_content_type,
		Secret:      &test_webhook_secret,
	}
)

func equal_webhook(s suite.Suite, expect, actual *Webhook) {
	s.Equal(*expect.Id, *actual.Id)
	s.Equal(*expect.Url, *actual.Url)
	s.Equal(*expect.ContentType, *actual.ContentType)
	if expect.Secret != nil && actual.Secret != nil {
		s.Equal(*expect.Secret, *actual.Secret)
	}
}
