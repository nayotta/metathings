package metathings_device_cloud_service

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
	grpc_helper "github.com/nayotta/metathings/pkg/common/grpc"
)

var (
	ErrUnexpectedContentType = errors.New("unexpected content type")
)

func ParseHttpRequestBody(r *http.Request, v proto.Message) error {
	if !strings.HasPrefix(strings.ToLower(r.Header.Get("Content-Type")), "application/json") {
		return ErrUnexpectedContentType
	}

	if err := grpc_helper.JSONPBUnmarshaler.Unmarshal(r.Body, v); err != nil {
		return err
	}

	return nil
}

func ParseHttpResponseBody(v proto.Message) ([]byte, error) {
	s, err := grpc_helper.JSONPBMarshaler.MarshalToString(v)
	if err != nil {
		return nil, err
	}
	return []byte(s), nil
}

func GetTokenFromHeader(r *http.Request) string {
	return strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
}

func GetSessionFromHeader(r *http.Request) int64 {
	sess, _ := strconv.ParseInt(r.Header.Get("MT-Module-Session"), 10, 64)
	return sess
}
