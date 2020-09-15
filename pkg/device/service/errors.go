package metathings_device_service

import (
	"errors"

	"google.golang.org/grpc/codes"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
)

var (
	ErrModuleNotFound             = errors.New("module not found")
	ErrSendRequestToStreamTimeout = errors.New("send request to stream timeout")
)

var em = grpc_helper.ErrorMapping{
	ErrModuleNotFound: codes.NotFound,
}
