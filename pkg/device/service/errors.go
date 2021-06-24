package metathings_device_service

import (
	"errors"

	"google.golang.org/grpc/codes"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
)

var (
	ErrModuleNotFound              = errors.New("module not found")
	ErrInitializeConnectionTimeout = errors.New("initialize connection timeout")
	ErrConnectToSameNode           = errors.New("connect to same node")
	ErrExistedConnection           = errors.New("existed connection")
	ErrNoConnection                = errors.New("no connection")
)

var em = grpc_helper.ErrorMapping{
	ErrModuleNotFound: codes.NotFound,
}
