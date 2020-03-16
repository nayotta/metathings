package metathings_plugin_evaluator_service

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/wrappers"
	log "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	context_helper "github.com/nayotta/metathings/pkg/common/context"
	hst "github.com/nayotta/metathings/pkg/common/http/status"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	evltr_plg "github.com/nayotta/metathings/pkg/plugin/evaluator"
	evltr_pb "github.com/nayotta/metathings/pkg/proto/evaluatord"
	dssdk "github.com/nayotta/metathings/sdk/data_storage"
	esdk "github.com/nayotta/metathings/sdk/evaluatord"
)

type EvaluatorPluginServiceOption struct {
	Evaluator struct {
		Endpoint string
	}
}

type EvaluatorPluginService struct {
	opt      *EvaluatorPluginServiceOption
	logger   log.FieldLogger
	tknr     token_helper.Tokener
	dat_stor dssdk.DataStorage
	cli_fty  *client_helper.ClientFactory
}

func (srv *EvaluatorPluginService) get_logger() log.FieldLogger {
	return srv.logger
}

func (srv *EvaluatorPluginService) evaluator_info_string_map_from_evaluator(info *evltr_pb.Evaluator) (map[string]interface{}, error) {
	inf := map[string]interface{}{
		"id": info.GetId(),
	}

	return inf, nil
}

func (srv *EvaluatorPluginService) evaluator_config_string_map_from_evaluator(info *evltr_pb.Evaluator) (map[string]interface{}, error) {
	var cfg map[string]interface{}
	buf, err := new(jsonpb.Marshaler).MarshalToString(info.GetConfig())
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(buf), &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (srv *EvaluatorPluginService) operator_string_map_form_evaluator(info *evltr_pb.Evaluator) (map[string]interface{}, error) {
	var opt map[string]interface{}
	op := info.GetOperator()

	marshaler := new(jsonpb.Marshaler)

	buf, err := marshaler.MarshalToString(op)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(buf), &opt)
	if err != nil {
		return nil, err
	}

	// SYM:REFACTOR:lua_operator
	switch op.GetDriver() {
	case "lua":
		fallthrough
	case "default":
		buf, err := marshaler.MarshalToString(op.GetLua())
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(buf), &opt)
		if err != nil {
			return nil, err
		}
	}

	return opt, nil
}

func (srv *EvaluatorPluginService) HandleResponse(w http.ResponseWriter, r *http.Request, hs *hst.HttpStatus) {
	code := hs.Code()
	w.WriteHeader(code)
	if code == http.StatusNoContent {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(hs.Message()); err != nil {
		srv.get_logger().WithError(err).Errorf("failed to write http response")
		return
	}
}

func (srv *EvaluatorPluginService) decode_eval_request(r *http.Request) (esdk.Data, error) {
	dat_codec := r.Header.Get(esdk.HTTP_HEADER_DATA_CODEC)

	// TODO(Peer): initial default codec
	if dat_codec == "" {
		dat_codec = "json"
	}

	dec, err := esdk.GetDataDecoder(dat_codec)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	dat, err := dec.Decode(buf)
	if err != nil {
		return nil, err
	}

	return dat, nil
}

func (srv *EvaluatorPluginService) get_evaluator(ctx context.Context, cli evltr_pb.EvaluatordServiceClient, evltr_id string) (*evltr_pb.Evaluator, error) {
	new_ctx := context_helper.WithToken(ctx, srv.tknr.GetToken())
	get_evltr_req := &evltr_pb.GetEvaluatorRequest{
		Evaluator: &evltr_pb.OpEvaluator{
			Id: &wrappers.StringValue{Value: evltr_id},
		},
	}
	get_evltr_res, err := cli.GetEvaluator(new_ctx, get_evltr_req)
	if err != nil {
		return nil, err
	}

	return get_evltr_res.GetEvaluator(), nil
}

func (srv *EvaluatorPluginService) Eval(w http.ResponseWriter, r *http.Request) {
	// TODO(Peer): wrap opentracing tags
	ctx := r.Context()
	evltr_id := r.Header.Get("X-Evaluator-ID")
	src_id := r.Header.Get(esdk.HTTP_HEADER_SOURCE_ID)
	src_typ := r.Header.Get(esdk.HTTP_HEADER_SOURCE_TYPE)

	logger := srv.get_logger().WithFields(log.Fields{
		"#method":     "eval",
		"source":      src_id,
		"source_type": src_typ,
		"evaluator":   evltr_id,
	})

	dat, err := srv.decode_eval_request(r)
	if err != nil {
		logger.WithError(err).Errorf("failed to decode eval request")
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

	evltr_info, err := srv.get_evaluator(ctx, evltr_cli, evltr_id)
	if err != nil {
		logger.WithError(err).Errorf("failed to get evaluator from evaluatord")
		srv.HandleResponse(w, r, hst.WrapErrorHttpStatus(http.StatusInternalServerError, err))
		return
	}

	evltr_inf, err := srv.evaluator_info_string_map_from_evaluator(evltr_info)
	if err != nil {
		logger.WithError(err).Errorf("failed to parse evaluator info")
		srv.HandleResponse(w, r, hst.WrapErrorHttpStatus(http.StatusInternalServerError, err))
		return
	}

	evltr_cfg, err := srv.evaluator_config_string_map_from_evaluator(evltr_info)
	if err != nil {
		logger.WithError(err).Errorf("failed to parse evaluator config")
		srv.HandleResponse(w, r, hst.WrapErrorHttpStatus(http.StatusInternalServerError, err))
		return
	}
	evltr_cfg["source"] = map[string]interface{}{
		"id":   src_id,
		"type": src_typ,
	}

	op_opt, err := srv.operator_string_map_form_evaluator(evltr_info)
	if err != nil {
		logger.WithError(err).Errorf("failed to parse operator option")
		srv.HandleResponse(w, r, hst.WrapErrorHttpStatus(http.StatusInternalServerError, err))
		return
	}

	evltr, err := evltr_plg.NewEvaluator(
		"info", evltr_inf,
		"config", evltr_cfg,
		"operator", op_opt,
		"logger", srv.get_logger(),
	)
	if err != nil {
		logger.WithError(err).Errorf("failed to new evaluator instance")
		srv.HandleResponse(w, r, hst.WrapErrorHttpStatus(http.StatusInternalServerError, err))
		return
	}

	err = evltr.Eval(ctx, dat)
	if err != nil {
		logger.WithError(err).Errorf("failed to eval")
		srv.HandleResponse(w, r, hst.WrapErrorHttpStatus(http.StatusInternalServerError, err))
		return
	}

	srv.HandleResponse(w, r, hst.NewHttpStatus(http.StatusNoContent, nil))
	logger.Debugf("eval")
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
		// TODO(Peer): tls support
		req, err := http.NewRequest("POST", srv.opt.Evaluator.Endpoint, bytes.NewReader(body))
		if err != nil {
			logger.WithError(err).Errorf("failed to new http request")
			srv.HandleResponse(w, r, hst.WrapErrorHttpStatus(http.StatusInternalServerError, err))
			return
		}

		req.Header.Set("Content-Type", r.Header.Get("Content-Type"))
		req.Header.Set("Authorization", r.Header.Get("Authorization"))
		req.Header.Set("X-Evaluator-ID", evltr.GetId())
		req.Header.Set(esdk.HTTP_HEADER_SOURCE_ID, r.Header.Get(esdk.HTTP_HEADER_SOURCE_ID))
		req.Header.Set(esdk.HTTP_HEADER_SOURCE_TYPE, r.Header.Get(esdk.HTTP_HEADER_SOURCE_TYPE))
		req = req.WithContext(ctx)

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			logger.WithError(err).Errorf("failed to send http request")
			srv.HandleResponse(w, r, hst.WrapErrorHttpStatus(http.StatusInternalServerError, err))
			return
		}

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

func NewEvaluatorPluginService(
	opt *EvaluatorPluginServiceOption,
	logger log.FieldLogger,
	tknr token_helper.Tokener,
	dat_stor dssdk.DataStorage,
	cli_fty *client_helper.ClientFactory,
) (*EvaluatorPluginService, error) {
	return &EvaluatorPluginService{
		opt:      opt,
		logger:   logger,
		tknr:     tknr,
		dat_stor: dat_stor,
		cli_fty:  cli_fty,
	}, nil
}
