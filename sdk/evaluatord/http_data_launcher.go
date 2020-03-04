package metathings_evaluatord_sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

const (
	HTTP_HEADER_SOURCE_ID   = "X-Evaluator-Source-ID"
	HTTP_HEADER_SOURCE_TYPE = "X-Evaluator-Source-Type"
	HTTP_HEADER_DATA_CODEC  = "X-Evaluator-Data-Codec"
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
	req.Header.Set("Authorization", "Bearer "+ExtractToken(ctx))
	req.Header.Set(HTTP_HEADER_SOURCE_ID, src.GetId())
	req.Header.Set(HTTP_HEADER_SOURCE_TYPE, src.GetType())
	req.Header.Set(HTTP_HEADER_DATA_CODEC, hdl.opt.DataCodec)
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
