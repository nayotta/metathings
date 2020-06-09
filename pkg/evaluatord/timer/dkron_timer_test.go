package metathings_evaluatord_timer

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/objx"
	"github.com/stretchr/testify/suite"
)

var (
	test_timer_id              = "test-timer-id"
	test_timer_schedule        = "test-timer-schedule"
	test_timer_timezone        = "test-timer-timezone"
	test_timer_enabled         = true
	test_timer_backend_webhook = "http://timer-webhook"
	test_timer_backend_timeout = 1
	test_timer_new_schedule    = "test-timer-new-schedule"
	test_timer_new_timezone    = "test-timer-new-timezone"
	test_timer_disabled        = true
)

type DkronTimerBackendTestSuite struct {
	suite.Suite

	b     *DkronTimerBackend
	mux   *http.ServeMux
	httpd *httptest.Server
	ctx   context.Context
}

func (s *DkronTimerBackendTestSuite) SetupTest() {
	s.mux = http.NewServeMux()
	s.httpd = httptest.NewServer(s.mux)
	s.ctx = context.TODO()

	tb, err := NewDkronTimerBackend(
		"url", s.httpd.URL,
		"timeout", test_timer_backend_timeout,
		"webhook", test_timer_backend_webhook,
		"logger", logrus.New(),
	)
	s.Require().Nil(err)
	s.b = tb.(*DkronTimerBackend)
}

func (s *DkronTimerBackendTestSuite) BeforeTest(suiteName, methodName string) {
	switch methodName {
	case "TestCreate":
		s.mux.HandleFunc("/v1/jobs", func(w http.ResponseWriter, r *http.Request) {
			s.Equal(r.Method, "POST")

			buf, err := ioutil.ReadAll(r.Body)
			s.Require().Nil(err)
			defer r.Body.Close()

			ox, err := objx.FromJSON(string(buf))
			s.Require().Nil(err)

			s.Equal(test_timer_id, ox.Get("name").String())
			s.Equal(test_timer_id, ox.Get("displayname").String())
			s.Equal(test_timer_schedule, ox.Get("schedule").String())
			s.Equal(test_timer_timezone, ox.Get("timezone").String())
			s.Equal("http", ox.Get("executor").String())
			s.Equal(DKRON_EXECUTOR_CONFIG_EXPECTED_CODE, ox.Get("executor_config.expectedCode").String())
			s.Equal(DKRON_EXECUTOR_CONFIG_METHOD, ox.Get("executor_config.method").String())
			s.Equal(fmt.Sprintf("%v", test_timer_backend_timeout), ox.Get("executor_config.timeout").String())
			s.Equal(must_url_join(test_timer_backend_webhook, ox.Get("name").String()), ox.Get("executor_config.url").String())
			w.WriteHeader(http.StatusOK)
			_, err = w.Write(buf)
			s.Require().Nil(err)
		})
	case "TestGet":
		s.mux.HandleFunc(fmt.Sprintf("/v1/jobs/%s", test_timer_id), func(w http.ResponseWriter, r *http.Request) {
			s.Equal(r.Method, "GET")

			timer_id := r.URL.Path[9:]
			ret := map[string]interface{}{
				"name":        test_timer_id,
				"displayname": test_timer_id,
				"schedule":    test_timer_schedule,
				"timezone":    test_timer_timezone,
				"executor":    "http",
				"executor_config": map[string]interface{}{
					"expectedCode": "200",
					"method":       "POST",
					"timeout":      fmt.Sprintf("%v", test_timer_backend_timeout),
					"url":          must_url_join(test_timer_backend_webhook, timer_id),
				},
			}
			buf, err := json.Marshal(ret)
			s.Require().Nil(err)

			w.WriteHeader(http.StatusOK)
			_, err = w.Write(buf)
			s.Require().Nil(err)
		})
	case "TestDelete":
		s.mux.HandleFunc(fmt.Sprintf("/v1/jobs/%s", test_timer_id), func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				timer_id := r.URL.Path[9:]
				ret := map[string]interface{}{
					"name": timer_id,
				}
				buf, err := json.Marshal(ret)
				s.Require().Nil(err)

				w.WriteHeader(http.StatusOK)
				_, err = w.Write(buf)
			case "DELETE":
				w.WriteHeader(http.StatusNoContent)
			}
		})
	case "TestSet":
		s.mux.HandleFunc(fmt.Sprintf("/v1/jobs/%s", test_timer_id), func(w http.ResponseWriter, r *http.Request) {
			s.Equal(r.Method, "GET")
			timer_id := r.URL.Path[9:]
			ret := map[string]interface{}{
				"name": timer_id,
			}
			buf, err := json.Marshal(ret)
			s.Require().Nil(err)

			w.WriteHeader(http.StatusOK)
			_, err = w.Write(buf)
		})
		s.mux.HandleFunc("/v1/jobs", func(w http.ResponseWriter, r *http.Request) {
			s.Equal(r.Method, "POST")

			buf, err := ioutil.ReadAll(r.Body)
			s.Require().Nil(err)
			defer r.Body.Close()

			ox, err := objx.FromJSON(string(buf))
			s.Require().Nil(err)

			s.Equal(test_timer_new_schedule, ox.Get("schedule").String())
			s.Equal(test_timer_new_timezone, ox.Get("timezone").String())
			s.Equal(true, ox.Get("disabled").Bool())

			ret := map[string]interface{}{
				"name":     test_timer_id,
				"schedule": test_timer_new_schedule,
				"timezone": test_timer_new_timezone,
				"disabled": true,
			}
			buf, err = json.Marshal(ret)
			s.Require().Nil(err)

			w.WriteHeader(http.StatusOK)
			_, err = w.Write(buf)
			s.Require().Nil(err)
		})
	}
}

func (s *DkronTimerBackendTestSuite) TestCreate() {
	timer, err := s.b.Create(
		s.ctx,
		SetId(test_timer_id),
		SetSchedule(test_timer_schedule),
		SetTimezone(test_timer_timezone),
		SetEnabled(test_timer_enabled),
	)
	s.Require().Nil(err)

	s.Equal(test_timer_id, timer.Id())
	s.Equal(test_timer_schedule, timer.Schedule())
	s.Equal(test_timer_timezone, timer.Timezone())
	s.Equal(test_timer_enabled, timer.Enabled())
}

func (s *DkronTimerBackendTestSuite) TestGet() {
	timer, err := s.b.Get(s.ctx, test_timer_id)
	s.Require().Nil(err)

	s.Equal(test_timer_id, timer.Id())
	s.Equal(test_timer_schedule, timer.Schedule())
	s.Equal(test_timer_timezone, timer.Timezone())
	s.Equal(test_timer_enabled, timer.Enabled())
}

func (s *DkronTimerBackendTestSuite) TestDelete() {
	timer, err := s.b.Get(s.ctx, test_timer_id)
	s.Require().Nil(err)

	err = timer.Delete(s.ctx)
	s.Require().Nil(err)
}

func (s *DkronTimerBackendTestSuite) TestSet() {
	timer, err := s.b.Get(s.ctx, test_timer_id)
	s.Require().Nil(err)

	err = timer.Set(s.ctx,
		SetEnabled(false),
		SetSchedule(test_timer_new_schedule),
		SetTimezone(test_timer_new_timezone),
	)
	s.Require().Nil(err)

	s.Equal(test_timer_new_schedule, timer.Schedule())
	s.Equal(test_timer_new_timezone, timer.Timezone())
	s.Equal(false, timer.Enabled())
}

func TestDkronTimerBackendTestSuite(t *testing.T) {
	suite.Run(t, new(DkronTimerBackendTestSuite))
}
