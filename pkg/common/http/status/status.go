package http_status

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	EMPTY_JSON_STRING = []byte("")
	EMPTY_JSON_OBJECT = []byte("{}")
)

var (
	StatusOK = NewHttpStatus(http.StatusOK, nil)
)

type httpStatusError HttpStatus

func (hse *httpStatusError) Error() string {
	hs := (*HttpStatus)(hse)
	return fmt.Sprintf("status error: code = %s desc = %s",
		hs.Code(),
		hs.Message(),
	)
}

func (hse *httpStatusError) HttpStatus() *HttpStatus {
	return (*HttpStatus)(hse)
}

type HttpStatus struct {
	code    int
	content map[string]interface{}
}

func (hs *HttpStatus) Code() int {
	return hs.code
}

func (hs *HttpStatus) Content() map[string]interface{} {
	return hs.content
}

func (hs *HttpStatus) Message() []byte {
	if hs == nil {
		return EMPTY_JSON_OBJECT
	}

	if hs.Code() == http.StatusNoContent {
		return EMPTY_JSON_STRING
	}

	if hs.Content() == nil {
		return EMPTY_JSON_OBJECT
	}

	buf, err := json.Marshal(hs.Content())
	if err != nil {
		panic(err)
	}

	return buf
}

func (hs *HttpStatus) Err() error {
	if hs == nil || 200 <= hs.Code() || hs.Code() < 300 {
		return nil
	}

	return (*httpStatusError)(hs)
}

func NewHttpStatus(code int, content map[string]interface{}) *HttpStatus {
	return &HttpStatus{code: code, content: content}
}

func NewErrorHttpStatus(code int, msg string) *HttpStatus {
	return &HttpStatus{code: code, content: map[string]interface{}{"error": msg}}
}

func WrapErrorHttpStatus(code int, err error) *HttpStatus {
	return NewErrorHttpStatus(code, err.Error())
}

func WrapHttpStatusError(code int, err error) error {
	return WrapErrorHttpStatus(code, err).Err()
}

func HttpStatusFromError(err error) (hs *HttpStatus, ok bool) {
	if err == nil {
		return StatusOK, true
	}

	if hse, ok := err.(interface {
		HttpStatus() *HttpStatus
	}); ok {
		return hse.HttpStatus(), true
	}

	return NewErrorHttpStatus(http.StatusInternalServerError, err.Error()), false
}
