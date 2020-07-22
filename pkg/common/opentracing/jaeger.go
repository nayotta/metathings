package opentracing_helper

import (
	"io"
	"sync"

	"github.com/opentracing/opentracing-go"
	jaeger_config "github.com/uber/jaeger-client-go/config"

	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

type JaegerTracerFactory struct{}

var set_jaeger_as_global_tracer_once sync.Once

func (jtf *JaegerTracerFactory) New(args ...interface{}) (opentracing.Tracer, io.Closer, error) {
	var ok bool
	var service_name string
	var disabled string
	var agent_host string
	var agent_port string
	var agent_user string
	var agent_password string
	var endpoint string
	var sampler_type string
	var sampler_param string
	var err error

	if err = opt_helper.Setopt(map[string]func(string, interface {
	}) error{
		"option": func(key string, val interface{}) error {
			if _, ok = val.(*TracerOption); !ok {
				return opt_helper.InvalidArgument("option")
			}
			return nil
		},
		"service_name":   opt_helper.ToString(&service_name),
		"disabled":       opt_helper.ToString(&disabled),
		"agent_host":     opt_helper.ToString(&agent_host),
		"agent_port":     opt_helper.ToString(&agent_port),
		"agent_user":     opt_helper.ToString(&agent_user),
		"agent_password": opt_helper.ToString(&agent_password),
		"endpoint":       opt_helper.ToString(&endpoint),
		"sampler_type":   opt_helper.ToString(&sampler_type),
		"sampler_param":  opt_helper.ToString(&sampler_param),
	})(args...); err != nil {
		return nil, nil, err
	}

	opt_helper.SetenvIfNotExists("JAEGER_SERVICE_NAME", service_name)
	opt_helper.SetenvIfNotExists("JAEGER_DISABLED", disabled)
	opt_helper.SetenvIfNotExists("JAEGER_AGENT_HOST", agent_host)
	opt_helper.SetenvIfNotExists("JAEGER_AGENT_PORT", agent_port)
	opt_helper.SetenvIfNotExists("JAEGER_AGENT_USER", agent_user)
	opt_helper.SetenvIfNotExists("JAEGER_AGENT_PASSWORD", agent_password)
	opt_helper.SetenvIfNotExists("JAEGER_ENDPOINT", endpoint)
	opt_helper.SetenvIfNotExists("JAEGER_SAMPLER_TYPE", sampler_type)
	opt_helper.SetenvIfNotExists("JAEGER_SAMPLER_PARAM", sampler_param)

	cfg, err := jaeger_config.FromEnv()
	if err != nil {
		return nil, nil, err
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}

	set_jaeger_as_global_tracer_once.Do(func() {
		opentracing.SetGlobalTracer(tracer)
	})

	return tracer, closer, nil
}

func init() {
	registry_tracer_factory("jaeger", new(JaegerTracerFactory))
}
