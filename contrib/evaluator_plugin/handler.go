package main

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

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

func EvalHandler(w http.ResponseWriter, r *http.Request) {
	srv.Eval(w, r)
}

func ReceiveDataHandler(w http.ResponseWriter, r *http.Request) {
	srv.ReceiveData(w, r)
}

func init() {
	var err error
	srv, err = evltr_plg_cmd.NewEvaluatorPluginService(getFissionConfigPath())
	if err != nil {
		panic(err)
	}
}
