package webhook_helper

import "github.com/stretchr/testify/suite"

var (
	test_default_webhook_id           = "default"
	test_default_webhook_url          = "http://www.example.com/webhook"
	test_default_webhook_content_type = "application/json"
	test_default_webhook              = &Webhook{
		Id:          &test_default_webhook_id,
		Url:         &test_default_webhook_url,
		ContentType: &test_default_webhook_content_type,
	}

	test_webhook_id           = "test"
	test_webhook_url          = "http://www.example1.com/webhook"
	test_webhook_content_type = "application/json"
	test_webhook              = &Webhook{
		Id:          &test_webhook_id,
		Url:         &test_webhook_url,
		ContentType: &test_webhook_content_type,
	}
)

func equal_webhook(s suite.Suite, expect, actual *Webhook) {
	s.Equal(*expect.Id, *actual.Id)
	s.Equal(*expect.Url, *actual.Url)
	s.Equal(*expect.ContentType, *actual.ContentType)
}
