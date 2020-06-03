package metathings_evaluatord_timer

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/objx"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

func must_url_join(url_str string, paths ...string) string {
	u, _ := url.Parse(url_str)
	u.Path = path.Join(append([]string{u.Path}, paths...)...)
	return u.String()
}

type DkronTimerApi struct {
	call   func(ctx context.Context, method string, path string, body map[string]interface{}) (map[string]interface{}, error)
	body   objx.Map
	logger logrus.FieldLogger
}

func (t *DkronTimerApi) get_logger() logrus.FieldLogger {
	return t.logger
}

func (t *DkronTimerApi) Id() string {
	return t.body.Get("name").String()
}

func (t *DkronTimerApi) Schedule() string {
	return t.body.Get("schedule").String()
}

func (t *DkronTimerApi) Timezone() string {
	return t.body.Get("timezone").String()
}

func (t *DkronTimerApi) Enabled() bool {
	return !t.body.Get("disabled").Bool()
}

func (t *DkronTimerApi) Set(ctx context.Context, opts ...TimerOption) error {
	logger := t.get_logger().WithField("timer", t.Id())

	o := map[string]interface{}{}
	for _, opt := range opts {
		opt(o)
	}
	ox := objx.New(o)

	body := map[string]interface{}{
		"name": t.Id(),
	}
	if schedule := ox.Get("schedule"); !schedule.IsNil() {
		body["schedule"] = schedule.String()
	} else {
		body["schedule"] = t.Schedule()
	}

	if timezone := ox.Get("timezone"); !timezone.IsNil() {
		body["timezone"] = timezone.String()
	}

	if enabled := ox.Get("enabled"); !enabled.IsNil() {
		body["disabled"] = !enabled.Bool()
	}

	body, err := t.call(ctx, "POST", "/jobs", body)
	if err != nil {
		logger.WithError(err).Debugf("failed to set timer")
		return err
	}

	t.body = objx.New(body)
	logger.Debugf("set timer")

	return nil
}

func (t *DkronTimerApi) Delete(ctx context.Context) error {
	_, err := t.call(ctx, "DELETE", fmt.Sprintf("/jobs/%s", t.Id()), nil)
	if err != nil {
		return err
	}
	t.body = objx.Nil
	return nil
}

func NewDkronTimerApi(
	call func(ctx context.Context, method string, path string, body map[string]interface{}) (map[string]interface{}, error),
	body map[string]interface{},
	logger logrus.FieldLogger,
) (*DkronTimerApi, error) {
	return &DkronTimerApi{
		call:   call,
		body:   objx.New(body),
		logger: logger,
	}, nil
}

type DkronTimerBackendOption struct {
	Timeout time.Duration
	Url     string
	Webhook string
}

func NewDkronTimerBackendOption() *DkronTimerBackendOption {
	opt := &DkronTimerBackendOption{}

	opt.Timeout = 5 * time.Second

	return opt
}

type DkronTimerBackend struct {
	opt    *DkronTimerBackendOption
	logger logrus.FieldLogger
}

func (b *DkronTimerBackend) get_logger() logrus.FieldLogger {
	return b.logger
}

func (b *DkronTimerBackend) http_client() *http.Client {
	return &http.Client{
		Timeout: b.opt.Timeout,
	}
}

func (b *DkronTimerBackend) call(ctx context.Context, method string, endpoint string, payload map[string]interface{}) (map[string]interface{}, error) {
	var buf []byte
	var err error

	url := must_url_join(b.opt.Url, "v1", endpoint)
	if payload != nil {
		buf, err = json.Marshal(payload)
		if err != nil {
			return nil, err
		}
	}

	// TODO(Peer): tls support
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := b.http_client().Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		buf, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(string(buf))
	}

	buf, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	ret := map[string]interface{}{}
	if len(buf) > 0 {
		err = json.Unmarshal(buf, &ret)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (b *DkronTimerBackend) Create(ctx context.Context, opts ...TimerOption) (TimerApi, error) {
	o := map[string]interface{}{}
	for _, opt := range opts {
		opt(o)
	}
	ox := objx.New(o)

	timer_id := ox.Get("id").String()
	if timer_id == "" {
		return nil, ErrTimerIdIsEmpty
	}

	logger := b.get_logger().WithField("timer", timer_id)
	body := map[string]interface{}{
		"name":        timer_id,
		"displayname": timer_id,
		"schedule":    ox.Get("schedule").String(),
		"timezone":    ox.Get("timezone").String(),
		"disabled":    !ox.Get("enabled").Bool(),
		"executor":    "http",
		"executor_config": map[string]interface{}{
			"expectedCode": "200",
			"method":       "POST",
			"timeout":      fmt.Sprintf("%d", int(b.opt.Timeout/time.Second)),
			"url":          must_url_join(b.opt.Webhook, timer_id),
		},
	}
	res, err := b.call(ctx, "POST", "/jobs", body)
	if err != nil {
		logger.WithError(err).Debugf("failed to create timer")
		return nil, err
	}

	logger.Debugf("create timer")

	return NewDkronTimerApi(b.call, res, logger)
}

func (b *DkronTimerBackend) Get(ctx context.Context, id string) (TimerApi, error) {
	logger := b.get_logger().WithField("timer", id)

	res, err := b.call(ctx, "GET", fmt.Sprintf("/jobs/%s", id), nil)
	if err != nil {
		logger.WithError(err).Debugf("failed to get timer")
		return nil, err
	}

	logger.Debugf("get timer")

	return NewDkronTimerApi(b.call, res, logger)
}

func NewDkronTimerBackend(args ...interface{}) (TimerBackend, error) {
	var err error
	var logger logrus.FieldLogger
	opt := NewDkronTimerBackendOption()

	if err = opt_helper.Setopt(map[string]func(string, interface{}) error{
		"timeout": opt_helper.ToDuration(&opt.Timeout),
		"url":     opt_helper.ToString(&opt.Url),
		"webhook": opt_helper.ToString(&opt.Webhook),
		"logger":  opt_helper.ToLogger(&logger),
	})(args...); err != nil {
		return nil, err
	}

	return &DkronTimerBackend{
		opt:    opt,
		logger: logger,
	}, nil
}

func init() {
	register_timer_backend_factory("dkron", NewDkronTimerBackend)
}
