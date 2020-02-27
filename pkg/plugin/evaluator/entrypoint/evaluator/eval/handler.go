package main

import (
	"net/http"

	evltr_plg_cmd "github.com/nayotta/metathings/pkg/plugin/evaluator/cmd"
	service "github.com/nayotta/metathings/pkg/plugin/evaluator/service"
)

var srv *service.EvaluatorPluginService

func Handler(w http.ResponseWriter, r *http.Request) {
	srv.Eval(w, r)
}

func init() {
	var err error
	srv, err = evltr_plg_cmd.NewEvaluatorPluginServiceFromEnv()
	if err != nil {
		panic(err)
	}
}
