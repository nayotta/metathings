package opentracing_helper

import "errors"

var (
	ErrInvalidTracerDriver = errors.New("invalid tracer driver")
)
