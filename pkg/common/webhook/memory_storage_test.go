package webhook_helper

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
)

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

type memoryStorageTestSuite struct {
	logger  log.FieldLogger
	storage *MemoryStorage
	suite.Suite
}

func (s *memoryStorageTestSuite) SetupTest() {
	logger, err := log_helper.NewLogger("test", "debug")
	s.Nil(err)

	storage, err := NewStorage("memory", "logger", logger)
	s.Nil(err)

	s.storage = storage.(*MemoryStorage)

	_, err = s.storage.CreateWebhook(test_default_webhook)
	s.Nil(err)
}

func (s *memoryStorageTestSuite) TestCreateWebhook() {
	_, err := s.storage.CreateWebhook(test_webhook)
	s.Nil(err)

	wh, err := s.storage.GetWebhook(test_webhook_id)
	s.Nil(err)
	equal_webhook(s.Suite, test_webhook, wh)
}

func (s *memoryStorageTestSuite) TestDeleteWebhook() {
	err := s.storage.DeleteWebhook(test_default_webhook_id)
	s.Nil(err)

	_, err = s.storage.GetWebhook(test_default_webhook_id)
	s.NotNil(err)
}

func (s *memoryStorageTestSuite) TestGetWebhook() {
	wh, err := s.storage.GetWebhook(test_default_webhook_id)
	s.Nil(err)
	equal_webhook(s.Suite, test_default_webhook, wh)
}

func (s *memoryStorageTestSuite) TestListWebhooks() {
	whs, err := s.storage.ListWebhooks(nil)
	s.Nil(err)
	s.Len(whs, 1)
}

func TestMemoryStorageTestSuite(t *testing.T) {
	suite.Run(t, new(memoryStorageTestSuite))
}
