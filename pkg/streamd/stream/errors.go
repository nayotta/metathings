package stream_manager

import "errors"

var (
	// symbol errors
	ErrBadSymbolString = errors.New("bad symbol string")

	// stream errors
	ErrUnstartable  = errors.New("unstartable")
	ErrUnterminable = errors.New("unterminable")

	// input errors
	ErrInputDataCodec = errors.New("input data codec error")

	// output errors
	ErrOutputDataCodec = errors.New("output data codec error")
)
