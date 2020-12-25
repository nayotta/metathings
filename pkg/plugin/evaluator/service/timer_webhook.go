package metathings_plugin_evaluator_service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/stretchr/objx"

	context_helper "github.com/nayotta/metathings/pkg/common/context"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	hst "github.com/nayotta/metathings/pkg/common/http/status"
	dvd_pb "github.com/nayotta/metathings/proto/deviced"
	evltr_pb "github.com/nayotta/metathings/proto/evaluatord"
	esdk "github.com/nayotta/metathings/sdk/evaluatord"
)

const (
	TIMER_DATA_CODEC = "json"
)

func (srv *EvaluatorPluginService) get_timer_by_id(ctx context.Context, cli evltr_pb.EvaluatordServiceClient, id string) (*evltr_pb.Timer, error) {
	new_ctx := context_helper.WithToken(ctx, srv.tknr.GetToken())
	gt_req := &evltr_pb.GetTimerRequest{
		Timer: &evltr_pb.OpTimer{
			Id: &wrappers.StringValue{Value: id},
		},
	}

	gt_res, err := cli.GetTimer(new_ctx, gt_req)
	if err != nil {
		return nil, err
	}

	return gt_res.GetTimer(), nil
}

// Timer Default Config v1:
// Fields:
// - name: version
//   description: default config vesrion string
//   type: string
//   default: v1
// - name: device
//   description: device object
//   type: object
//   required: true
//   fields:
//   - name: id
//     description: device id
//     type: string
//     required: true
// - name: data
//   description: launch data
//   type: object
//   default: {}
// - name: data_tags
//   description: launch data tags
//   type: object
//   default: {}
//
// Examples:
//   version: v1
//   device:
//     id: d469d424a64511eaac2f6c4008bb5d9a
//   data:
//     text: "hello, world"
func (srv *EvaluatorPluginService) get_timer_default_config(ctx context.Context, dvd_cli dvd_pb.DevicedServiceClient, timer *evltr_pb.Timer) (*dvd_pb.Config, error) {
	new_ctx := context_helper.WithToken(ctx, srv.tknr.GetToken())
	for _, cfg := range timer.GetConfigs() {
		gc_req := &dvd_pb.GetConfigRequest{
			Config: &dvd_pb.OpConfig{
				Id: &wrappers.StringValue{Value: cfg.GetId()},
			},
		}
		gc_res, err := dvd_cli.GetConfig(new_ctx, gc_req)
		if err != nil {
			return nil, err
		}

		if cfg = gc_res.GetConfig(); cfg.GetAlias() == esdk.TIMER_DEFAULT_CONFIG {
			return cfg, nil
		}
	}

	return nil, ErrTimerDefaultConfigNotFound
}

func (srv *EvaluatorPluginService) TimerWebhook(w http.ResponseWriter, r *http.Request) {
	var err error
	var hs *hst.HttpStatus
	var log_msg string

	ctx := r.Context()
	id := srv.get_param(r, "id")

	src_id := id
	src_typ := "timer"

	logger := srv.get_logger().WithFields(logrus.Fields{
		"#method":     "webhook/timer",
		"source":      src_id,
		"source_type": src_typ,
	})

	defer func() {
		if err != nil {
			logger.WithError(err).Errorf(log_msg)
			srv.HandleResponse(w, r, hs)
		} else {
			logger.Debugf("timer webhook")
			srv.HandleResponse(w, r, hst.NewHttpStatus(http.StatusNoContent, nil))
		}
	}()

	evltr_cli, evltr_cfn, err := srv.cli_fty.NewEvaluatordServiceClient()
	if err != nil {
		log_msg = "failed to new evaluatord service client"
		hs = hst.WrapErrorHttpStatus(http.StatusInternalServerError, err)
		return
	}
	defer evltr_cfn()

	dvd_cli, dvd_cfn, err := srv.cli_fty.NewDevicedServiceClient()
	if err != nil {
		log_msg = "failed to new deviced service client"
		hs = hst.WrapErrorHttpStatus(http.StatusInternalServerError, err)
		return
	}
	defer dvd_cfn()

	tmr, err := srv.get_timer_by_id(ctx, evltr_cli, id)
	if err != nil {
		log_msg = "failed to get timer by id"
		hs = hst.WrapErrorHttpStatus(http.StatusInternalServerError, err)
		return
	}

	cfg, err := srv.get_timer_default_config(ctx, dvd_cli, tmr)
	if err != nil {
		log_msg = "failed to get timer default config"
		hs = hst.WrapErrorHttpStatus(http.StatusInternalServerError, err)
		return
	}

	cfg_str, err := grpc_helper.JSONPBMarshaler.MarshalToString(cfg)
	if err != nil {
		log_msg = "failed to marshal config to json string"
		hs = hst.WrapErrorHttpStatus(http.StatusInternalServerError, err)
		return
	}

	var cfgm map[string]interface{}
	err = json.Unmarshal([]byte(cfg_str), &cfgm)
	if err != nil {
		log_msg = "failed to new config mapping from map string"
		hs = hst.WrapErrorHttpStatus(http.StatusInternalServerError, err)
		return
	}

	cfgx := objx.New(objx.New(cfgm).Get("body").Data())
	switch cfgx.Get("version").String() {
	case "v1":
		err, log_msg, hs = srv.launch_data_by_timer_v1(ctx, evltr_cli, src_id, src_typ, cfgx)
		return
	}

	logger.WithField("version", cfgx.Get("version")).Warningf("unsupported default config version")
}

func (srv *EvaluatorPluginService) launch_data_by_timer_v1(
	ctx context.Context,
	evltr_cli evltr_pb.EvaluatordServiceClient,
	src_id, src_typ string,
	cfgx objx.Map,
) (error, string, *hst.HttpStatus) {
	dev_id := cfgx.Get("device.id").String()
	buf, _ := time.Now().MarshalText()
	dat_ts := string(buf)
	dat := cast.ToStringMap(cfgx.Get("data").Data())
	dat_tags := cast.ToStringMap(cfgx.Get("data_tags").Data())
	body, err := json.Marshal(dat)
	if err != nil {
		return err, "failed to marshal launch data to json string", hst.WrapErrorHttpStatus(http.StatusInternalServerError, err)
	}

	evltrs, err := srv.list_evaluators_by_source(ctx, evltr_cli, src_id, src_typ)
	if err != nil {
		return err, "failed to list evaluators by timer", hst.WrapErrorHttpStatus(http.StatusInternalServerError, err)
	}

	for _, evltr := range evltrs {
		evltr_id := evltr.GetId()

		// TODO(Peer): support more data codec
		req, hr, err := srv.new_request(
			nethttp.ClientSpanObserver(func(span opentracing.Span, r *http.Request) {
				span.SetTag("source", src_id)
				span.SetTag("source_type", src_typ)
				if dev_id != "" {
					span.SetTag("device", dev_id)
				}
				if dat_ts != "" {
					span.SetTag("data_timestamp", dat_ts)
				}
			}),
		)(ctx, "POST", srv.opt.Evaluator.Endpoint, bytes.NewReader(body))
		if err != nil {
			return err, "failed to new launch data request", hst.WrapErrorHttpStatus(http.StatusInternalServerError, err)
		}
		if hr != nil {
			defer hr.Finish()
		}

		req.Header.Set(esdk.HTTP_HEADER_DATA_CODEC, TIMER_DATA_CODEC)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", srv.tknr.GetToken())
		req.Header.Set("X-Evaluator-ID", evltr_id)
		req.Header.Set(esdk.HTTP_HEADER_SOURCE_ID, src_id)
		req.Header.Set(esdk.HTTP_HEADER_SOURCE_TYPE, src_typ)
		if dev_id != "" {
			req.Header.Set(esdk.HTTP_HEADER_DEVICE_ID, dev_id)
		}
		req.Header.Set(esdk.HTTP_HEADER_DATA_TIMESTAMP, dat_ts)
		if len(dat_tags) > 0 {
			datenc, _ := esdk.GetDataEncoder(TIMER_DATA_CODEC)
			dat_tags_dat, _ := esdk.DataFromMap(dat_tags)
			dat_tags_buf, _ := datenc.Encode(dat_tags_dat)
			req.Header.Set(esdk.HTTP_HEADER_DATA_TAGS, string(dat_tags_buf))
		}

		res, err := srv.get_http_client().Do(req)
		if err != nil {
			return err, "failed to launch data", hst.WrapErrorHttpStatus(http.StatusInternalServerError, err)
		}
		defer res.Body.Close()

		if res.StatusCode >= 400 {
			buf, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return err, "failed to read launch data response", hst.WrapErrorHttpStatus(http.StatusInternalServerError, err)
			}

			bodyx, _ := objx.FromJSON(string(buf))
			err_msg := bodyx.Get("message").String()
			return errors.New(err_msg), err_msg, hst.NewErrorHttpStatus(res.StatusCode, err_msg)
		}

	}

	return nil, "timer triggered", hst.NewHttpStatus(http.StatusNoContent, nil)
}
