// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: get_object.proto

package deviced

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
var _get_object_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetObjectRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetObjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetObject() == nil {
		return GetObjectRequestValidationError{
			field:  "Object",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetObject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetObjectRequestValidationError{
				field:  "Object",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetObjectRequestValidationError is the validation error returned by
// GetObjectRequest.Validate if the designated constraints aren't met.
type GetObjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetObjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetObjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetObjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetObjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetObjectRequestValidationError) ErrorName() string { return "GetObjectRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetObjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetObjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetObjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetObjectRequestValidationError{}

// Validate checks the field values on GetObjectResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetObjectResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetObject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetObjectResponseValidationError{
				field:  "Object",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetObjectResponseValidationError is the validation error returned by
// GetObjectResponse.Validate if the designated constraints aren't met.
type GetObjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetObjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetObjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetObjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetObjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetObjectResponseValidationError) ErrorName() string {
	return "GetObjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetObjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetObjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetObjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetObjectResponseValidationError{}
