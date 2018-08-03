package metathings_identityd_service

import (
	"errors"
)

var (
	Unauthenticated               = errors.New("unauthenticated")
	ErrNotValidatedTokenInContext = errors.New("not validated token in context")
)
