// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: create_config.proto

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
var _create_config_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CreateConfigRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateConfigRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetConfig() == nil {
		return CreateConfigRequestValidationError{
			field:  "Config",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateConfigRequestValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateConfigRequestValidationError is the validation error returned by
// CreateConfigRequest.Validate if the designated constraints aren't met.
type CreateConfigRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateConfigRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateConfigRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateConfigRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateConfigRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateConfigRequestValidationError) ErrorName() string {
	return "CreateConfigRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateConfigRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateConfigRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateConfigRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateConfigRequestValidationError{}

// Validate checks the field values on CreateConfigResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateConfigResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateConfigResponseValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateConfigResponseValidationError is the validation error returned by
// CreateConfigResponse.Validate if the designated constraints aren't met.
type CreateConfigResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateConfigResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateConfigResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateConfigResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateConfigResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateConfigResponseValidationError) ErrorName() string {
	return "CreateConfigResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateConfigResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateConfigResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateConfigResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateConfigResponseValidationError{}
