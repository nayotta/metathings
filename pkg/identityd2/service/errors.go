package metathings_identityd2_service

import "errors"

var (
	ErrPermissionDenied = errors.New("permission denied")
	ErrUnauthenticated  = errors.New("unauthenticated")
)
