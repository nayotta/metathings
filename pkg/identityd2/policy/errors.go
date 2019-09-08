package metathings_identityd2_policy

import "errors"

var (
	ErrInvalidArguments          = errors.New("invalid arguments")
	ErrPermissionDenied          = errors.New("permission denied")
	ErrUnauthenticated           = errors.New("unauthenticated")
	ErrInvalidBackendDriver      = errors.New("invalid backend driver")
	ErrInvalidBackendCacheDriver = errors.New("invalid backend cache driver")
	ErrNoCached                  = errors.New("no cached")
)
