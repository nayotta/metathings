package metathings_device_service

import (
	"errors"

	"google.golang.org/grpc/codes"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
)

var (
	ErrModuleNotFound = errors.New("module not found")
)

var em = grpc_helper.ErrorMapping{
	ErrModuleNotFound: codes.NotFound,
}
