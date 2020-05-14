package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/pflag"

	opentracing_helper "github.com/nayotta/metathings/pkg/common/opentracing"
	evltr_plg_cmd "github.com/nayotta/metathings/pkg/plugin/evaluator/cmd"
)

var (
	host   string
	port   int
	config string
)

func main() {
	pflag.StringVarP(&host, "host", "h", "127.0.0.1", "Evaluator Plugin TestServer listen host")
	pflag.IntVarP(&port, "port", "p", 8888, "Evaluator Plugin TestServer listen port")
	pflag.StringVarP(&config, "config", "c", "evaluator-plugin.yaml", "Evaluator Plugin TestServer config file")

	pflag.Parse()

	srv, err := evltr_plg_cmd.NewEvaluatorPluginService(config)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/eval", opentracing_helper.Middleware(srv, "Eval"))
	http.HandleFunc("/receive_data", opentracing_helper.Middleware(srv, "ReceiveData"))

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), nil))
}
