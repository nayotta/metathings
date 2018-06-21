package metathings_camerad_storage

import "errors"

var (
	ErrUnknownStorageDriver = errors.New("Unknown Storage Driver")
	ErrNotFound             = errors.New("Not Found")
	ErrNothingChanged       = errors.New("Nothing Changed")
)
