package metathings_deviced_service

import (
	"errors"
	"os"

	"google.golang.org/grpc/codes"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	storage "github.com/nayotta/metathings/pkg/deviced/storage"
)

var (
	ErrUnexpectedMessage = errors.New("unexpected message")
	ErrDuplicatedDevice  = errors.New("duplicated device")
	ErrUnconnectedDevice = errors.New("unconnected device")
	ErrFlowNotFound      = errors.New("flow not found")
	ErrDeviceOffline     = errors.New("device offline")

	ErrPutObjectStreamingTimeout = errors.New("put object streaming timeout")
	ErrInvalidProtoset           = errors.New("invaild protoset")
)

var em = grpc_helper.ErrorMapping{
	storage.RecordNotFound: codes.NotFound,
	os.ErrNotExist:         codes.NotFound,
	ErrFlowNotFound:        codes.NotFound,
}
