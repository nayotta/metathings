package metathings_evaluatord_service

import (
	"io"

	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
	"google.golang.org/grpc/codes"
)

var em = grpc_helper.ErrorMapping{
	io.EOF: codes.OutOfRange,
}
