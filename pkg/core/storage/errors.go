package metathings_core_storage

import "errors"

var (
	ErrNotFound    = errors.New("Not Found")
	NothingChanged = errors.New("Nothing Changed")
)
