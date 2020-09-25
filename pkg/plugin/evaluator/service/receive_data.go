package metathings_plugin_evaluator_service

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	opentracing "github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"

	context_helper "github.com/nayotta/metathings/pkg/common/context"
	hst "github.com/nayotta/metathings/pkg/common/http/status"
	opentracing_helper "github.com/nayotta/metathings/pkg/common/opentracing"
	evltr_pb "github.com/nayotta/metathings/proto/evaluatord"
	esdk "github.com/nayotta/metathings/sdk/evaluatord"
)

func (srv *EvaluatorPluginService) new_request(options ...nethttp.ClientOption) func(context.Context, string, string, io.Reader) (*http.Request, *nethttp.Tracer, error) {
	return opentracing_helper.NewRequester(srv)(options...)
}

func (srv *EvaluatorPluginService) list_evaluators_by_source(ctx context.Context, cli evltr_pb.EvaluatordServiceClient, src_id, src_typ string) ([]*evltr_pb.Evaluator, error) {
	new_ctx := context_helper.WithToken(ctx, srv.tknr.GetToken())
	lebs_req := &evltr_pb.ListEvaluatorsBySourceRequest{
		Source: &evltr_pb.OpResource{
			Id:   &wrappers.StringValue{Value: src_id},
			Type: &wrappers.StringValue{Value: src_typ},
		},
	}

	lebs_res, err := cli.ListEvaluatorsBySource(new_ctx, lebs_req)
	if err != nil {
		return nil, err
	}

	return lebs_res.GetEvaluators(), nil
}

func (srv *EvaluatorPluginService) ReceiveData(w http.ResponseWriter, r *http.Request) {
	var err error

	ctx := r.Context()
	src_id := r.Header.Get(esdk.HTTP_HEADER_SOURCE_ID)
	src_typ := r.Header.Get(esdk.HTTP_HEADER_SOURCE_TYPE)

	logger := srv.get_logger().WithFields(log.Fields{
		"#method":     "receive_data",
		"source":      src_id,
		"source_type": src_typ,
	})

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WithError(err).Errorf("failed to read request body")
		srv.HandleResponse(w, r, hst.WrapErrorHttpStatus(http.StatusBadRequest, err))
		return
	}

	// TODO(Peer): client pool
	// TODO(Peer): cache evaluator info
	evltr_cli, cfn, err := srv.cli_fty.NewEvaluatordServiceClient()
	if err != nil {
		logger.WithError(err).Errorf("failed to new evaluatord service client")
		srv.HandleResponse(w, r, hst.WrapErrorHttpStatus(http.StatusInternalServerError, err))
		return
	}
	defer cfn()

	evltrs, err := srv.list_evaluators_by_source(ctx, evltr_cli, src_id, src_typ)
	if err != nil {
		logger.WithError(err).Errorf("failed to list evaluators by source from evaluatord")
		srv.HandleResponse(w, r, hst.WrapErrorHttpStatus(http.StatusInternalServerError, err))
		return
	}

	for _, evltr := range evltrs {
		dev_id := r.Header.Get(esdk.HTTP_HEADER_DEVICE_ID)
		dat_ts := r.Header.Get(esdk.HTTP_HEADER_DATA_TIMESTAMP)

		// TODO(Peer): tls support
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
			logger.WithError(err).Errorf("failed to new http request")
			srv.HandleResponse(w, r, hst.WrapErrorHttpStatus(http.StatusInternalServerError, err))
			return
		}
		if hr != nil {
			defer hr.Finish()
		}

		req.Header.Set("Content-Type", r.Header.Get("Content-Type"))
		req.Header.Set("Authorization", r.Header.Get("Authorization"))
		req.Header.Set("X-Evaluator-ID", evltr.GetId())
		req.Header.Set(esdk.HTTP_HEADER_SOURCE_ID, r.Header.Get(esdk.HTTP_HEADER_SOURCE_ID))
		req.Header.Set(esdk.HTTP_HEADER_SOURCE_TYPE, r.Header.Get(esdk.HTTP_HEADER_SOURCE_TYPE))
		if dev_id != "" {
			req.Header.Set(esdk.HTTP_HEADER_DEVICE_ID, dev_id)
		}
		req.Header.Set(esdk.HTTP_HEADER_DATA_TIMESTAMP, dat_ts)
		if buf := r.Header.Get(esdk.HTTP_HEADER_DATA_TAGS); buf != "" {
			req.Header.Set(esdk.HTTP_HEADER_DATA_TAGS, buf)
		}

		res, err := srv.get_http_client().Do(req)
		if err != nil {
			logger.WithError(err).Errorf("failed to send http request")
			srv.HandleResponse(w, r, hst.WrapErrorHttpStatus(http.StatusInternalServerError, err))
			return
		}
		defer res.Body.Close()

		if res.StatusCode >= 400 {
			buf, err := ioutil.ReadAll(res.Body)
			if err != nil {
				logger.WithError(err).Errorf("failed to read response body")
				srv.HandleResponse(w, r, hst.WrapErrorHttpStatus(http.StatusInternalServerError, err))
				return
			}

			var res_body struct {
				Message string
			}
			json.Unmarshal(buf, &res_body)
			logger.WithField("error", res_body.Message).Errorf("failed to eval in evaluator plugin")
			srv.HandleResponse(w, r, hst.NewErrorHttpStatus(res.StatusCode, res_body.Message))
			return
		}
	}

	srv.HandleResponse(w, r, hst.NewHttpStatus(http.StatusNoContent, nil))
	logger.Debugf("receive data")
}
