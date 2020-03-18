package metathings_evaluatord_sdk

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

const (
	HTTP_HEADER_SOURCE_ID      = "X-Evaluator-Source-ID"
	HTTP_HEADER_SOURCE_TYPE    = "X-Evaluator-Source-Type"
	HTTP_HEADER_DEVICE_ID      = "X-Evaluator-Device-ID"
	HTTP_HEADER_DATA_CODEC     = "X-Evaluator-Data-Codec"
	HTTP_HEADER_DATA_TIMESTAMP = "X-Evaluator-Data-Timestamp"
	HTTP_HEADER_DATA_TAGS      = "X-Evaluator-Data-Tags"
)

type HttpDataLauncherOption struct {
	Endpoint  string
	DataCodec string
}

type HttpDataLauncher struct {
	logger       log.FieldLogger
	opt          *HttpDataLauncherOption
	data_encoder DataEncoder
}

func (hdl *HttpDataLauncher) http_content_type() string {
	switch hdl.opt.DataCodec {
	case "json":
		return "application/json"
	default:
		panic("unsupported content type")
	}
}

func (hdl *HttpDataLauncher) http_authorization(ctx context.Context) string {
	return "Bearer " + ExtractToken(ctx)
}

func (hdl *HttpDataLauncher) http_data_tags(ctx context.Context) string {
	tags := ExtractTags(ctx)
	tags_dat, _ := DataFromMap(cast.ToStringMap(tags))
	tags_jsbuf, _ := hdl.data_encoder.Encode(tags_dat)
	return base64.StdEncoding.EncodeToString(tags_jsbuf)
}

func (hdl *HttpDataLauncher) http_data_timestamp(ctx context.Context) string {
	buf, _ := ExtractTimestamp(ctx).MarshalText()
	return string(buf)
}

func (hdl *HttpDataLauncher) http_device_id(ctx context.Context) string {
	return ExtractDevice(ctx)
}

func (hdl *HttpDataLauncher) Launch(ctx context.Context, src Resource, dat Data) error {
	body, err := hdl.data_encoder.Encode(dat)
	if err != nil {
		return err
	}

	// TODO(Peer): tls support
	req, err := http.NewRequest("POST", hdl.opt.Endpoint, bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", hdl.http_content_type())
	req.Header.Set("Authorization", hdl.http_authorization(ctx))
	req.Header.Set(HTTP_HEADER_SOURCE_ID, src.GetId())
	req.Header.Set(HTTP_HEADER_SOURCE_TYPE, src.GetType())
	if buf := hdl.http_device_id(ctx); buf != "" {
		req.Header.Set(HTTP_HEADER_DEVICE_ID, buf)
	}
	req.Header.Set(HTTP_HEADER_DATA_CODEC, hdl.opt.DataCodec)
	req.Header.Set(HTTP_HEADER_DATA_TIMESTAMP, hdl.http_data_timestamp(ctx))
	if buf := hdl.http_data_tags(ctx); buf != "" {
		req.Header.Set(HTTP_HEADER_DATA_TAGS, buf)
	}
	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		buf, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		var res_body struct {
			Message string
		}
		if err = json.Unmarshal(buf, &res_body); err != nil {
			return err
		}

		return errors.New(res_body.Message)
	}

	return nil
}

func DefaultHttpDataLauncherOption() *HttpDataLauncherOption {
	return &HttpDataLauncherOption{
		DataCodec: "json",
	}
}

func NewHttpDataLauncher(args ...interface{}) (DataLauncher, error) {
	var logger log.FieldLogger

	opt := DefaultHttpDataLauncherOption()

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"endpoint":   opt_helper.ToString(&opt.Endpoint),
		"data_codec": opt_helper.ToString(&opt.DataCodec),
		"logger":     opt_helper.ToLogger(&logger),
	})(args...); err != nil {
		return nil, err
	}

	enc, err := GetDataEncoder(opt.DataCodec)
	if err != nil {
		return nil, err
	}

	hdl := &HttpDataLauncher{
		opt:          opt,
		data_encoder: enc,
		logger:       logger,
	}

	return hdl, nil
}

func init() {
	registry_data_launcher("http", NewHttpDataLauncher)
}
