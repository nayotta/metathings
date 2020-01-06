package cmd_contrib

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/fx"

	cfg_helper "github.com/nayotta/metathings/pkg/common/config"
	opentracing_helper "github.com/nayotta/metathings/pkg/common/opentracing"
)

type OpentracingOptioner interface {
	GetTracer() string
	SetTracer(string)

	GetData() map[string]interface{}
}

type OpentracingOption struct {
	Opentracing map[string]interface{} `mapstructure:"opentracing"`
}

func (o *OpentracingOption) GetTracer() string {
	tr, ok := o.Opentracing["tracer"]
	if !ok {
		return "unknown"
	}
	return tr.(string)
}

func (o *OpentracingOption) SetTracer(tr string) {
	o.Opentracing["tracer"] = tr
}

func (o *OpentracingOption) GetData() map[string]interface{} {
	return o.Opentracing
}

type OpentracingResult struct {
	fx.Out

	Tracer opentracing.Tracer `name:"opentracing_tracer"`
	Closer io.Closer          `name:"opentracing_closer"`
}

func NewOpentracing(srv_opt ServiceOptioner, opt OpentracingOptioner) (OpentracingResult, error) {
	name, args, err := cfg_helper.ParseConfigOption("tracer", opt.GetData())
	if err != nil {
		return OpentracingResult{}, err
	}

	// get service name from other config
	args = append(args, "service_name", srv_opt.GetServiceName())

	tracer, closer, err := opentracing_helper.NewTracer(name, args...)
	if err != nil {
		return OpentracingResult{}, err
	}

	return OpentracingResult{
		Tracer: tracer,
		Closer: closer,
	}, nil
}
