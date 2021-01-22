// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: get_object_content.proto

package device

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _get_object_content_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetObjectContentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetObjectContentRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetObject() == nil {
		return GetObjectContentRequestValidationError{
			field:  "Object",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetObject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetObjectContentRequestValidationError{
				field:  "Object",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetObjectContentRequestValidationError is the validation error returned by
// GetObjectContentRequest.Validate if the designated constraints aren't met.
type GetObjectContentRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetObjectContentRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetObjectContentRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetObjectContentRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetObjectContentRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetObjectContentRequestValidationError) ErrorName() string {
	return "GetObjectContentRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetObjectContentRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetObjectContentRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetObjectContentRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetObjectContentRequestValidationError{}

// Validate checks the field values on GetObjectContentResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetObjectContentResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Content

	return nil
}

// GetObjectContentResponseValidationError is the validation error returned by
// GetObjectContentResponse.Validate if the designated constraints aren't met.
type GetObjectContentResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetObjectContentResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetObjectContentResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetObjectContentResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetObjectContentResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetObjectContentResponseValidationError) ErrorName() string {
	return "GetObjectContentResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetObjectContentResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetObjectContentResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetObjectContentResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetObjectContentResponseValidationError{}