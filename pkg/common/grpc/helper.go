package grpc_helper

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	InvalidFullMethodName = errors.New("invalid full method name")
)

type MethodDescription struct {
	Package string
	Service string
	Method  string
}

func HttpStatusCode2GrpcStatusCode(code int) codes.Code {
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

func ParseMethodDescription(fullMethodName string) (*MethodDescription, error) {
	pack_serv_meth := strings.Split(fullMethodName, "/")
	if len(pack_serv_meth) != 3 {
		return nil, InvalidFullMethodName
	}

	pack_serv := strings.SplitAfter(pack_serv_meth[1], ".")
	serv := pack_serv[len(pack_serv)-1]
	pack := pack_serv_meth[1][0 : len(pack_serv_meth[1])-len(serv)-1]
	meth := pack_serv_meth[2]

	return &MethodDescription{
		Package: pack,
		Service: serv,
		Method:  meth,
	}, nil
}

// github.com/grpc-ecosystem/go-grpc-middleware/auth/metadata.go:AuthFromMD
func AuthFromMD(ctx context.Context, expectedScheme string, headerAuthorize ...string) (string, error) {
	var authorize string
	if len(headerAuthorize) > 0 {
		authorize = headerAuthorize[0]
	} else {
		authorize = "authorization"
	}

	val := metautils.ExtractIncoming(ctx).Get(authorize)
	if val == "" {
		return "", status.Errorf(codes.Unauthenticated, fmt.Sprintf("request unauthenticated with %s, got empty string", expectedScheme))
	}

	splits := strings.SplitN(val, " ", 2)
	if len(splits) < 2 {
		return "", status.Errorf(codes.Unauthenticated, "Bad authorization string")
	}
	if strings.ToLower(splits[0]) != strings.ToLower(expectedScheme) {
		return "", status.Errorf(codes.Unauthenticated, fmt.Sprintf("request unauthenticated with %s, but got %s", expectedScheme, val))
	}
	return splits[1], nil
}

type AuthorizationTokenParser struct{}

func (p AuthorizationTokenParser) GetTokenFromContext(ctx context.Context) (string, error) {
	return AuthFromMD(ctx, "mt", "authorization")
}

func (p AuthorizationTokenParser) GetSubjectTokenFromContext(ctx context.Context) (string, error) {
	return AuthFromMD(ctx, "mt", "authorization-subject")
}

func HandleGRPCError(logger log.FieldLogger, err error, format string, args ...interface{}) error {
	if err == io.EOF {
		return nil
	}

	if status.Code(err) == codes.Canceled {
		return nil
	}

	logger.WithError(err).Errorf(format, args...)

	return err
}
