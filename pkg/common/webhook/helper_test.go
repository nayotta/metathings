package webhook_helper

import "github.com/stretchr/testify/suite"

func equal_webhook(s suite.Suite, expect, actual *Webhook) {
	s.Equal(*expect.Id, *actual.Id)
	s.Equal(*expect.Url, *actual.Url)
	s.Equal(*expect.ContentType, *actual.ContentType)
}
