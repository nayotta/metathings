package metathings_identityd2_policy

import "errors"

var (
	ErrPermissionDenied = errors.New("permission denied")
	ErrUnauthenticated  = errors.New("unauthenticated")
)
