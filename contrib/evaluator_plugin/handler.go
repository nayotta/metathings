package main

import (
	"net/http"

	evltr_plg_cmd "github.com/nayotta/metathings/pkg/plugin/evaluator/cmd"
	service "github.com/nayotta/metathings/pkg/plugin/evaluator/service"
)

// initialized by go build flags
// https://docs.fission.io/docs/languages/go/#custom-builds
var Config string

var srv *service.EvaluatorPluginService

func EvalHandler(w http.ResponseWriter, r *http.Request) {
	srv.Eval(w, r)
}

func init() {
	var err error
	srv, err = evltr_plg_cmd.NewEvaluatorPluginService(Config)
	if err != nil {
		panic(err)
	}
}
