package opentracing_helper

import (
	"io"
	"sync"

	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
)

type TracerOption struct {
	*viper.Viper
}

type TracerFactory interface {
	New(args ...interface{}) (opentracing.Tracer, io.Closer, error)
}

var tracer_factories map[string]TracerFactory
var tracer_factories_once sync.Once

func registry_tracer_factory(name string, fty TracerFactory) {
	tracer_factories[name] = fty
}

func NewTracer(name string, args ...interface{}) (opentracing.Tracer, io.Closer, error) {
	fty, ok := tracer_factories[name]

	if !ok {
		return nil, nil, ErrInvalidTracerDriver
	}

	return fty.New(args...)
}

func init() {
	tracer_factories_once.Do(func() {
		tracer_factories = make(map[string]TracerFactory)
	})
}
