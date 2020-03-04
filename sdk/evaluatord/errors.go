package metathings_evaluatord_sdk

import "errors"

var (
	ErrUnsupportedDataLauncherFactory = errors.New("unsupported data launcher factory")

	ErrUnsupportedDataEncoder = errors.New("unsupported data encoder")
	ErrUnsupportedDataDecoder = errors.New("unsupported data decoder")
)
