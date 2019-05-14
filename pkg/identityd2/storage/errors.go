package metathings_identityd2_storage

import "errors"

var (
	ErrInitialized  = errors.New("system initialized")
	InvalidArgument = errors.New("invalid argument")
)
