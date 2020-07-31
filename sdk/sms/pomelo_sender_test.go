package metathings_sms_sdk_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/objx"
	"github.com/stretchr/testify/suite"

	smssdk "github.com/nayotta/metathings/sdk/sms"
)

type PomeloSmsSenderTestSute struct {
	suite.Suite
}

func (s *PomeloSmsSenderTestSute) TestSendSms() {
	smsid := "test-id"
	numbers := []string{"12345678901"}
	args := map[string]string{
		"a": "1",
		"b": "2",
	}

	rr := http.NewServeMux()
	rr.HandleFunc("/sms/"+smsid, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf, err := ioutil.ReadAll(r.Body)
		s.Require().Nil(err)
		defer r.Body.Close()

		body := map[string]interface{}{}
		err = json.Unmarshal(buf, &body)
		s.Require().Nil(err)

		bodyx := objx.New(body)
		s.Equal(numbers[0], bodyx.Get("phoneNumberSet[0]").String())
		s.Equal("1", bodyx.Get("arguments.a").String())
		s.Equal("2", bodyx.Get("arguments.b").String())

		w.WriteHeader(http.StatusOK)
	}))

	ts := httptest.NewServer(rr)
	defer ts.Close()

	ss, err := smssdk.NewSmsSender("pomelo", "base_url", ts.URL)
	s.Require().Nil(err)

	err = ss.SendSms(context.TODO(), smsid, numbers, args)
	s.Require().Nil(err)
}

func TestPomeloSmsSenderTestSuite(t *testing.T) {
	suite.Run(t, new(PomeloSmsSenderTestSute))
}
