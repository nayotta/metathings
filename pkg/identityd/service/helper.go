package metathings_identityd_service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	codec "github.com/nayotta/metathings/pkg/identityd/service/encode_decode"
)

func encodeError(err error) error {
	switch err {
	case codec.Unimplemented:
		return status.Errorf(codes.Unimplemented, "unimplemented")
	default:
		return status.Errorf(codes.Internal, err.Error())
	}
}
