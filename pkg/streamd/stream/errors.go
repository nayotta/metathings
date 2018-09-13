package stream_manager

import "errors"

var (
	// symbol errors
	ErrBadSymbolString = errors.New("bad symbol string")

	// stream manager errors
	ErrUnregisteredStreamManagerFactory = errors.New("unregistered stream manager factory")
	ErrStreamNotFound                   = errors.New("stream not found")

	// stream errors
	ErrUnregisteredStreamFactory = errors.New("unregistered stream factory")
	ErrUnstartable               = errors.New("unstartable")
	ErrUnterminable              = errors.New("unterminable")

	// upstream errors
	ErrUnregisteredUpstreamFactory = errors.New("unregistered upstream factory")

	// source errors
	ErrUnregisteredSourceFactory = errors.New("unregistered source factory")

	// input errors
	ErrUnregisteredInputFactory = errors.New("unregistered input factory")
	ErrInputDataCodec           = errors.New("input data codec error")

	// output errors
	ErrUnregisteredOutputFactory = errors.New("unregistered output factory")
	ErrOutputDataCodec           = errors.New("output data codec error")

	// group errors
	ErrUnregisteredGroupFactory = errors.New("unregistered group factory")

	// engine errors
	ErrUnexpectedResultType = errors.New("unexpected result type")
)
