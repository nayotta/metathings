package metathings_callback_sdk_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/cast"
	"github.com/stretchr/objx"
	"github.com/stretchr/testify/suite"

	cbsdk "github.com/nayotta/metathings/sdk/callback"
)

type WebhookCallbackTestSuite struct {
	suite.Suite
}

func (s *WebhookCallbackTestSuite) TestEmit() {
	rr := http.NewServeMux()
	rr.HandleFunc("/webhook", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Equal("hello, world!", r.Header.Get("X-Text"))
		s.Equal("a", r.Header.Get("X-MTE-Tag-A"))
		s.Equal("b", r.Header.Get("X-MTE-Tag-B"))

		buf, err := ioutil.ReadAll(r.Body)
		s.Require().Nil(err)
		defer r.Body.Close()

		body := map[string]interface{}{}
		err = json.Unmarshal(buf, &body)
		s.Require().Nil(err)

		bodyx := objx.New(body)
		s.Equal("hello, world!", bodyx.Get("text").String())
		s.Equal(42, cast.ToInt(bodyx.Get("answer").Data()))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{}"))
	}))

	ts := httptest.NewServer(rr)
	defer ts.Close()

	custom_headers := map[string]string{
		"X-Text": "hello, world!",
	}

	cb, err := cbsdk.NewCallback("default",
		"allow_plain_text", true,
		"custom_headers", custom_headers,
		"url", ts.URL+"/webhook",
	)
	s.Require().Nil(err)

	data := map[string]interface{}{
		"text":   "hello, world!",
		"answer": 42,
	}
	tags := map[string]string{
		"a": "a",
		"b": "b",
	}

	err = cb.Emit(data, tags)
	s.Nil(err)
}

func TestWebhookCallbackTestSuite(t *testing.T) {
	suite.Run(t, new(WebhookCallbackTestSuite))
}
