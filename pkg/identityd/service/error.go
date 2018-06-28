package metathings_identity_service

import (
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"
)

var (
	Unauthenticated = errors.New("unauthenticated")
)

func mapCode(code int) codes.Code {
	switch code {
	case http.StatusBadRequest:
		return codes.InvalidArgument
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	case http.StatusForbidden:
		return codes.PermissionDenied
	case http.StatusNotFound:
		return codes.NotFound
	case http.StatusConflict:
		return codes.InvalidArgument
	default:
		return codes.Internal
	}
}
