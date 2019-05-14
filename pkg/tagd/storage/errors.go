package metathings_tagd_storage

import "errors"

var (
	ErrUnknownDriver = errors.New("unknown driver")
	ErrNotFound      = errors.New("not found")
)
