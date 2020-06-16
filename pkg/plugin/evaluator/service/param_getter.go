package metathings_plugin_evaluator_service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type GetParamFunc func(*http.Request, string) string

func getFissionParam(r *http.Request, k string) string {
	return r.Header.Get(fmt.Sprintf("x-fission-params-%v", k))
}

func getGorillaParam(r *http.Request, k string) string {
	return mux.Vars(r)[k]
}
