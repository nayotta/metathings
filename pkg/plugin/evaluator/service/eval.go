package metathings_plugin_evaluator_service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/wrappers"
	opentracing "github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"

	context_helper "github.com/nayotta/metathings/pkg/common/context"
	hst "github.com/nayotta/metathings/pkg/common/http/status"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	evltr_helper "github.com/nayotta/metathings/pkg/evaluatord/helper"
	evltr_stor "github.com/nayotta/metathings/pkg/evaluatord/storage"
	evltr_plg "github.com/nayotta/metathings/pkg/plugin/evaluator"
	state_pb "github.com/nayotta/metathings/pkg/proto/constant/state"
	evltr_pb "github.com/nayotta/metathings/pkg/proto/evaluatord"
	esdk "github.com/nayotta/metathings/sdk/evaluatord"
)

func (srv *EvaluatorPluginService) get_codec_from_request(r *http.Request) esdk.DataDecoder {
	dat_codec := r.Header.Get(esdk.HTTP_HEADER_DATA_CODEC)
	if dat_codec == "" {
		dat_codec = srv.opt.Codec
	}

	codec, _ := esdk.GetDataDecoder(dat_codec)

	return codec
}

func (srv *EvaluatorPluginService) build_evaluator_context(r *http.Request, info *evltr_pb.Evaluator) (map[string]interface{}, error) {
	cfg, err := srv.evaluator_config_string_map_from_evaluator(info)
	if err != nil {
		return nil, err
	}

	ctx := map[string]interface{}{
		"config": cfg,
		"source": map[string]interface{}{
			"id":   r.Header.Get(esdk.HTTP_HEADER_SOURCE_ID),
			"type": r.Header.Get(esdk.HTTP_HEADER_SOURCE_TYPE),
		},
		"token":     context_helper.AuthorizationToToken(r.Header.Get("Authorization")),
		"timestamp": srv.extract_data_timestamp(r),
		"tags":      srv.extract_data_tags(r),
	}

	if dev_id := r.Header.Get(esdk.HTTP_HEADER_DEVICE_ID); dev_id != "" {
		ctx["device"] = map[string]interface{}{
			"id": dev_id,
		}
	}

	return ctx, nil
}

func (srv *EvaluatorPluginService) extract_data_tags(r *http.Request) map[string]interface{} {
	dec := srv.get_codec_from_request(r)

	tags_b64 := r.Header.Get(esdk.HTTP_HEADER_DATA_TAGS)
	if tags_b64 == "" {
		return map[string]interface{}{}
	}

	tags_buf, err := base64.StdEncoding.DecodeString(tags_b64)
	if err != nil {
		return map[string]interface{}{}
	}

	tags_dat, err := dec.Decode(tags_buf)
	if err != nil {
		return map[string]interface{}{}
	}

	return cast.ToStringMap(tags_dat.Iter())
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

func (srv *EvaluatorPluginService) decode_eval_request(r *http.Request) (esdk.Data, error) {
	dec := srv.get_codec_from_request(r)

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

func (srv *EvaluatorPluginService) extract_data_timestamp(r *http.Request) int64 {
	if s := r.Header.Get(esdk.HTTP_HEADER_DATA_TIMESTAMP); s == "" {
		return time.Now().Unix()
	} else {
		return cast.ToTime(s).Unix()
	}
}

func (srv *EvaluatorPluginService) operator_string_map_from_evaluator(info *evltr_pb.Evaluator) (map[string]interface{}, error) {
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

func (srv *EvaluatorPluginService) Eval(w http.ResponseWriter, r *http.Request) {
	var err error
	var hs *hst.HttpStatus
	var log_msg string
	var ret esdk.Data

	ctx := r.Context()
	evltr_id := r.Header.Get("X-Evaluator-ID")
	src_id := r.Header.Get(esdk.HTTP_HEADER_SOURCE_ID)
	src_typ := r.Header.Get(esdk.HTTP_HEADER_SOURCE_TYPE)
	tsk_id := id_helper.NewId()
	tsk := &evltr_stor.Task{
		Id: &tsk_id,
		Source: &evltr_stor.Resource{
			Id:   &src_id,
			Type: &src_typ,
		},
	}

	logger := srv.get_logger().WithFields(log.Fields{
		"#method":     "eval",
		"source":      src_id,
		"source_type": src_typ,
		"evaluator":   evltr_id,
	})
	defer func() {
		if err != nil {
			logger.WithError(err).Errorf(log_msg)
			srv.HandleResponse(w, r, hs)
			if err := srv.task_stor.PatchTask(ctx, tsk, &evltr_stor.TaskState{
				State: evltr_helper.TASK_STATE_ENUMER.ToStringP(state_pb.TaskState_TASK_STATE_ERROR),
				Tags: map[string]interface{}{
					"error": err.Error(),
				},
			}); err != nil {
				logger.WithError(err).Warning("failed to patch task")
			}
		} else {
			srv.HandleResponse(w, r, hst.NewHttpStatus(http.StatusNoContent, nil))
			if err := srv.task_stor.PatchTask(ctx, tsk, &evltr_stor.TaskState{
				State: evltr_helper.TASK_STATE_ENUMER.ToStringP(state_pb.TaskState_TASK_STATE_DONE),
				Tags:  ret.Iter(),
			}); err != nil {
				logger.WithError(err).Warning("failed to patch task")
			}
		}
	}()

	if err := srv.task_stor.PatchTask(ctx, tsk, &evltr_stor.TaskState{
		State: evltr_helper.TASK_STATE_ENUMER.ToStringP(state_pb.TaskState_TASK_STATE_CREATED),
		Tags: map[string]interface{}{
			"evaluator.id": evltr_id,
		},
	}); err != nil {
		logger.WithError(err).Warning("failed to patch task")
		err = nil
	}

	dat, err := srv.decode_eval_request(r)
	if err != nil {
		log_msg = "failed to decode eval request"
		hs = hst.WrapErrorHttpStatus(http.StatusBadRequest, err)
		return
	}

	// TODO(Peer): client pool
	// TODO(Peer): cache evaluator info
	evltr_cli, cfn, err := srv.cli_fty.NewEvaluatordServiceClient()
	if err != nil {
		log_msg = "failed to new evaluatord service client"
		hs = hst.WrapErrorHttpStatus(http.StatusInternalServerError, err)
		return
	}
	defer cfn()

	evltr_info, err := srv.get_evaluator(ctx, evltr_cli, evltr_id)
	if err != nil {
		log_msg = "failed to get evaluator from evaluatord"
		hs = hst.WrapErrorHttpStatus(http.StatusInternalServerError, err)
		return
	}

	evltr_inf, err := srv.evaluator_info_string_map_from_evaluator(evltr_info)
	if err != nil {
		log_msg = "failed to parse evaluator info"
		hs = hst.WrapErrorHttpStatus(http.StatusInternalServerError, err)
		return
	}

	evltr_ctx, err := srv.build_evaluator_context(r, evltr_info)
	if err != nil {
		log_msg = "failed to build evaluator context"
		hs = hst.WrapErrorHttpStatus(http.StatusInternalServerError, err)
		return
	}

	op_opt, err := srv.operator_string_map_from_evaluator(evltr_info)
	if err != nil {
		log_msg = "failed to parse operator option"
		hs = hst.WrapErrorHttpStatus(http.StatusInternalServerError, err)
		return
	}

	evltr, err := evltr_plg.NewEvaluator(
		"info", evltr_inf,
		"context", evltr_ctx,
		"operator", op_opt,
		"caller", srv.caller,
		"sms_sender", srv.sms_sender,
		"logger", srv.get_logger(),
		"data_storage", srv.dat_stor,
		"simple_storage", srv.smpl_stor,
		"client_factory", srv.cli_fty,
	)
	if err != nil {
		log_msg = "failed to new evaluator instance"
		hs = hst.WrapErrorHttpStatus(http.StatusInternalServerError, err)
		return
	}

	if err := srv.task_stor.PatchTask(ctx, tsk, &evltr_stor.TaskState{
		State: evltr_helper.TASK_STATE_ENUMER.ToStringP(state_pb.TaskState_TASK_STATE_RUNNING),
	}); err != nil {
		logger.WithError(err).Warning("failed to patch task")
	}
	if srv.IsTraced(r) {
		var sp opentracing.Span
		sp, ctx = opentracing.StartSpanFromContext(ctx, "Evaluator.Eval")
		defer sp.Finish()
	}

	// TODO(Peer): retry to eval
	ret, err = evltr.Eval(ctx, dat)
	if err != nil {
		log_msg = "failed to eval"
		hs = hst.WrapErrorHttpStatus(http.StatusInternalServerError, err)
		return
	}

	logger.Debugf("eval")
}
