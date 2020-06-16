package main

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sync"

	opentracing_helper "github.com/nayotta/metathings/pkg/common/opentracing"
	evltr_plg_cmd "github.com/nayotta/metathings/pkg/plugin/evaluator/cmd"
	service "github.com/nayotta/metathings/pkg/plugin/evaluator/service"
)

// initialized by go build flags
// https://docs.fission.io/docs/languages/go/#custom-builds
var FissionConfigName string

func getFissionConfigPath() string {
	var ret string
	root := "/configs"

	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && path != root {
			ret = path
			return io.EOF
		}
		return nil
	}); err != io.EOF {
		panic(err)
	}

	return path.Join(ret, FissionConfigName)
}

var srv *service.EvaluatorPluginService

var evalHandler http.HandlerFunc
var evalHandlerOnce sync.Once

func EvalHandler(w http.ResponseWriter, r *http.Request) {
	evalHandlerOnce.Do(func() {
		evalHandler = opentracing_helper.Middleware(srv, "Eval")
	})
	evalHandler(w, r)
}

var receiveDataHandler http.HandlerFunc
var receiveDataHandlerOnce sync.Once

func ReceiveDataHandler(w http.ResponseWriter, r *http.Request) {
	receiveDataHandlerOnce.Do(func() {
		receiveDataHandler = opentracing_helper.Middleware(srv, "ReceiveData")
	})
	receiveDataHandler(w, r)
}

var timerWebhookHandler http.HandlerFunc
var timerWebhookHandlerOnce sync.Once

func TimerWebhookHandler(w http.ResponseWriter, r *http.Request) {
	timerWebhookHandlerOnce.Do(func() {
		timerWebhookHandler = opentracing_helper.Middleware(srv, "TimerWebhook")
	})
	timerWebhookHandler(w, r)
}

func init() {
	var err error
	srv, err = evltr_plg_cmd.NewEvaluatorPluginService(getFissionConfigPath())
	if err != nil {
		panic(err)
	}
}
