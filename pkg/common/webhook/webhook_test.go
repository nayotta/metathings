package webhook_helper

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
	passwd_helper "github.com/nayotta/metathings/pkg/common/passwd"
)

var (
	TEST_TRIGGER_TIMEOUT = 50 * time.Millisecond
)

type WebhookServiceTestSuite struct {
	suite.Suite

	ts        *httptest.Server
	whs       *webhookService
	triggered chan struct{}
	wh        *Webhook
}

func (s *WebhookServiceTestSuite) SetupTest() {
	logger, err := log_helper.NewLogger("test", "debug")
	s.Nil(err)

	storage, err := NewStorage("memory", "logger", logger)
	s.Nil(err)

	opt := &WebhookServiceOption{
		ContentType: "application/json",
	}

	whs, err := NewWebhookService(opt, "logger", logger, "storage", storage)
	s.whs = whs.(*webhookService)

	s.triggered = make(chan struct{})
	s.ts = httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("MT-Webhook-Id")
		ts_s := r.Header.Get("MT-Webhook-Timestamp")
		nonce_s := r.Header.Get("MT-Webhook-Nonce")
		hmac := r.Header.Get("MT-Webhook-HMAC")

		nsec, err := strconv.ParseInt(ts_s, 10, 64)
		s.Nil(err)
		ts := time.Unix(0, nsec)
		nonce, err := strconv.ParseInt(nonce_s, 10, 64)
		s.Nil(err)

		if passwd_helper.ValidateHmac(hmac, *s.wh.Secret, id, ts, nonce) {
			close(s.triggered)
		}
	}))

	secret := "c2VjcmV0" // base64("secret")
	wh := &Webhook{
		Url:    &s.ts.URL,
		Secret: &secret,
	}
	wh, err = whs.Add(wh)
	s.Nil(err)

	s.wh = &Webhook{
		Id:          wh.Id,
		Url:         &s.ts.URL,
		ContentType: &s.whs.opt.ContentType,
		Secret:      wh.Secret,
	}
}

func (s *WebhookServiceTestSuite) TestAdd() {
	tmp_triggered := make(chan struct{})
	tmp_ts := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("MT-Webhook-Id")
		ts_s := r.Header.Get("MT-Webhook-Timestamp")
		nonce_s := r.Header.Get("MT-Webhook-Nonce")
		hmac := r.Header.Get("MT-Webhook-HMAC")

		nsec, err := strconv.ParseInt(ts_s, 10, 64)
		s.Nil(err)
		ts := time.Unix(0, nsec)
		nonce, err := strconv.ParseInt(nonce_s, 10, 64)
		s.Nil(err)

		if passwd_helper.ValidateHmac(hmac, *s.wh.Secret, id, ts, nonce) {
			close(tmp_triggered)
		}
	}))

	secret := "c2VjcmV0" // base64("secret")
	wh, err := s.whs.Add(&Webhook{
		Url:    &tmp_ts.URL,
		Secret: &secret,
	})
	s.Nil(err)

	s.Equal(tmp_ts.URL, *wh.Url)
	s.NotNil(wh.Id)
	s.NotEqual("", *wh.Id)

	err = s.whs.Trigger(map[string]interface{}{})
	s.Nil(err)

	select {
	case <-s.triggered:
	case <-time.After(TEST_TRIGGER_TIMEOUT):
		s.Fail("default webhook should triggered")
	}

	select {
	case <-tmp_triggered:
	case <-time.After(TEST_TRIGGER_TIMEOUT):
		s.Fail("temp webhook should triggered")
	}
}

func (s *WebhookServiceTestSuite) TestRemove() {
	err := s.whs.Remove(*s.wh.Id)
	s.Nil(err)

	_, err = s.whs.Get(*s.wh.Id)
	s.NotNil(err)

	select {
	case <-s.triggered:
		s.Fail("default webhook should not triggered")
	case <-time.After(TEST_TRIGGER_TIMEOUT):
	}
}

func (s *WebhookServiceTestSuite) TestGet() {
	wh, err := s.whs.Get(*s.wh.Id)
	s.Nil(err)

	equal_webhook(s.Suite, s.wh, wh)
}

func (s *WebhookServiceTestSuite) TestList() {
	whs, err := s.whs.List(nil)
	s.Nil(err)
	s.Len(whs, 1)
}

func (s *WebhookServiceTestSuite) TestUpdate() {
	tmp_url := "http://www.example.com/webhook"
	wh, err := s.whs.Update(*s.wh.Id, &Webhook{
		Url: &tmp_url,
	})
	s.Nil(err)
	s.Equal(tmp_url, *wh.Url)

	wh, err = s.whs.Get(*s.wh.Id)
	s.Nil(err)
	s.Equal(tmp_url, *wh.Url)
}

func (s *WebhookServiceTestSuite) TestTrigger() {
	err := s.whs.Trigger(map[string]interface{}{})
	s.Nil(err)

	select {
	case <-s.triggered:
	case <-time.After(TEST_TRIGGER_TIMEOUT):
		s.Fail("default webhook should triggered")
	}
}

func TestWebhookServiceTestSuite(t *testing.T) {
	suite.Run(t, new(WebhookServiceTestSuite))
}
