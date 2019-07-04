package metathings_device_cloud_service

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

var (
	ErrUnexpectedContentType = errors.New("unexpected content type")
)

func ParseHttpRequestBody(r *http.Request, v interface{}) error {
	if !strings.HasPrefix(strings.ToLower(r.Header.Get("Content-Type")), "application/json") {
		return ErrUnexpectedContentType
	}

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(v); err != nil {
		return err
	}

	return nil
}

func GetTokenFromHeader(r *http.Request) string {
	return strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
}

func GetSessionFromHeader(r *http.Request) int64 {
	sess, _ := strconv.ParseInt(r.Header.Get("MT-Module-Session"), 10, 64)
	return sess
}
