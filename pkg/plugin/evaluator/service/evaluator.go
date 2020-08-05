package metathings_plugin_evaluator_service

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	client_helper "github.com/nayotta/metathings/pkg/common/client"
	hst "github.com/nayotta/metathings/pkg/common/http/status"
	opentracing_helper "github.com/nayotta/metathings/pkg/common/opentracing"
	token_helper "github.com/nayotta/metathings/pkg/common/token"
	evltr_stor "github.com/nayotta/metathings/pkg/evaluatord/storage"
	dssdk "github.com/nayotta/metathings/sdk/data_storage"
	dsdk "github.com/nayotta/metathings/sdk/deviced"
	smssdk "github.com/nayotta/metathings/sdk/sms"
)

type EvaluatorPluginServiceOption struct {
	Evaluator struct {
		Endpoint string
	}
	Server struct {
		Name string
	}
	Codec    string
	IsTraced bool
}

func NewEvaluatorPluginServiceOption() *EvaluatorPluginServiceOption {
	opt := &EvaluatorPluginServiceOption{}

	opt.Codec = "json"
	opt.Server.Name = "fission"

	return opt
}

type EvaluatorPluginService struct {
	opt        *EvaluatorPluginServiceOption
	logger     log.FieldLogger
	tknr       token_helper.Tokener
	dat_stor   dssdk.DataStorage
	smpl_stor  dsdk.SimpleStorage
	flow       dsdk.Flow
	task_stor  evltr_stor.TaskStorage
	caller     dsdk.Caller
	sms_sender smssdk.SmsSender
	cli_fty    *client_helper.ClientFactory
	get_param  GetParamFunc
}

func (srv *EvaluatorPluginService) get_logger() log.FieldLogger {
	return srv.logger
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

func (srv *EvaluatorPluginService) get_http_client() *http.Client {
	return opentracing_helper.GetHTTPClient(srv)
}

func (srv *EvaluatorPluginService) IsTraced(*http.Request) bool {
	return srv.opt.IsTraced
}

func NewEvaluatorPluginService(
	opt *EvaluatorPluginServiceOption,
	logger log.FieldLogger,
	tknr token_helper.Tokener,
	dat_stor dssdk.DataStorage,
	smpl_stor dsdk.SimpleStorage,
	flow dsdk.Flow,
	task_stor evltr_stor.TaskStorage,
	caller dsdk.Caller,
	sms_sender smssdk.SmsSender,
	cli_fty *client_helper.ClientFactory,
) (*EvaluatorPluginService, error) {
	srv := &EvaluatorPluginService{
		opt:        opt,
		logger:     logger,
		tknr:       tknr,
		dat_stor:   dat_stor,
		smpl_stor:  smpl_stor,
		flow:       flow,
		task_stor:  task_stor,
		caller:     caller,
		sms_sender: sms_sender,
		cli_fty:    cli_fty,
	}

	switch opt.Server.Name {
	case "gorilla":
		srv.get_param = getGorillaParam
	case "fission":
		fallthrough
	default:
		srv.get_param = getFissionParam
	}

	return srv, nil
}
