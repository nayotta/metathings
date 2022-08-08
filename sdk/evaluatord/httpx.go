package metathings_evaluatord_sdk

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type HttpxDataLauncherOption struct {
	Endpoint      string
	HeadersTplStr string
	DataCodec     string
}

func defaultHttpxDataLauncherOption() *HttpxDataLauncherOption {
	return &HttpxDataLauncherOption{
		DataCodec: "json",
	}
}

type HttpxDataLauncher struct {
	logger      log.FieldLogger
	opt         *HttpxDataLauncherOption
	dataEncoder DataEncoder
	headersTpl  *template.Template
}

func NewHttpxDataLauncher(args ...any) (DataLauncher, error) {
	var logger log.FieldLogger

	opt := defaultHttpxDataLauncherOption()

	if err := opt_helper.Setopt(map[string]func(string, interface{}) error{
		"endpoint":         opt_helper.ToString(&opt.Endpoint),
		"data_codec":       opt_helper.ToString(&opt.DataCodec),
		"headers_template": opt_helper.ToString(&opt.HeadersTplStr),
		"logger":           opt_helper.ToLogger(&logger),
	})(args...); err != nil {
		return nil, err
	}

	enc, err := GetDataEncoder(opt.DataCodec)
	if err != nil {
		return nil, err
	}

	tpl, err := template.New("headers").Parse(opt.HeadersTplStr)
	if err != nil {
		return nil, err
	}

	return &HttpxDataLauncher{
		logger:      logger,
		opt:         opt,
		dataEncoder: enc,
		headersTpl:  tpl,
	}, nil
}

func (x *HttpxDataLauncher) Launch(ctx context.Context, src Resource, dat Data) error {
	logger := x.GetLogger().WithFields(log.Fields{
		"#method":     "Launch",
		"source.id":   src.GetId(),
		"source.type": src.GetType(),
	})

	buf, err := x.dataEncoder.Encode(dat)
	if err != nil {
		logger.WithError(err).Debugf("failed to encode data")
		return err
	}

	req, err := http.NewRequest("POST", x.opt.Endpoint, bytes.NewReader(buf))
	if err != nil {
		logger.WithError(err).Debugf("failed to new http request")
		return err
	}

	tplDat, err := x.parseTemplateData(ctx, src, dat)
	if err != nil {
		logger.WithError(err).Debugf("failed to parse template data")
		return err
	}

	var sb strings.Builder
	if err = x.headersTpl.Execute(&sb, tplDat); err != nil {
		logger.WithError(err).Debugf("failed to execute headers template")
		return err
	}

	hdrMap := map[string]string{}
	if err = json.Unmarshal([]byte(sb.String()), &hdrMap); err != nil {
		logger.WithError(err).Debugf("failed to unmarshal header string to string map")
		return err
	}

	for k, v := range hdrMap {
		req.Header.Set(k, v)
	}

	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.WithError(err).Debugf("failed to send http request")
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		buf, err := ioutil.ReadAll(res.Body)
		if err != nil {
			logger.WithError(err).Debugf("failed to read response body")
			return err
		}

		err = fmt.Errorf("%s", string(buf))
		logger.WithError(err).Debugf("failed to launch data")

		return err
	}

	logger.Debugf("launch data")

	return nil
}

func (x *HttpxDataLauncher) GetLogger() log.FieldLogger {
	return x.logger.WithFields(log.Fields{
		"#instance": "HttpxDataLauncher",
	})
}

func (x *HttpxDataLauncher) parseTemplateData(ctx context.Context, src Resource, dat Data) (map[string]any, error) {
	out := map[string]any{
		"contentType":   x.content_type(),
		"authorization": x.extractAuthorization(ctx),
		"deviceID":      ExtractDevice(ctx),
		"sourceID":      src.GetId(),
		"sourceType":    src.GetType(),
		"codec":         x.opt.DataCodec,
		"timestamp":     x.extractTimestamp(ctx),
		"tags":          x.extractTags(ctx),
	}

	return out, nil
}

func (x *HttpxDataLauncher) extractAuthorization(ctx context.Context) string {
	return "Bearer " + ExtractToken(ctx)
}

func (x *HttpxDataLauncher) content_type() string {
	switch x.opt.DataCodec {
	case "json":
		return "application/json"
	default:
		panic("unsupported content type")
	}
}

func (x *HttpxDataLauncher) extractTimestamp(ctx context.Context) string {
	buf, _ := ExtractTimestamp(ctx).MarshalText()
	return string(buf)
}

func (x *HttpxDataLauncher) extractTags(ctx context.Context) string {
	tags := ExtractTags(ctx)
	tags_dat, _ := DataFromMap(cast.ToStringMap(tags))
	tags_jsbuf, _ := x.dataEncoder.Encode(tags_dat)
	return base64.StdEncoding.EncodeToString(tags_jsbuf)
}

func init() {
	registry_data_launcher("httpx", NewHttpxDataLauncher)
}
