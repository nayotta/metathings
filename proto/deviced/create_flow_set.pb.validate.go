// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: create_flow_set.proto

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
var _create_flow_set_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CreateFlowSetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateFlowSetRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetFlowSet() == nil {
		return CreateFlowSetRequestValidationError{
			field:  "FlowSet",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetFlowSet()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateFlowSetRequestValidationError{
				field:  "FlowSet",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateFlowSetRequestValidationError is the validation error returned by
// CreateFlowSetRequest.Validate if the designated constraints aren't met.
type CreateFlowSetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateFlowSetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateFlowSetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateFlowSetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateFlowSetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateFlowSetRequestValidationError) ErrorName() string {
	return "CreateFlowSetRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateFlowSetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateFlowSetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateFlowSetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateFlowSetRequestValidationError{}

// Validate checks the field values on CreateFlowSetResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateFlowSetResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetFlowSet()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateFlowSetResponseValidationError{
				field:  "FlowSet",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateFlowSetResponseValidationError is the validation error returned by
// CreateFlowSetResponse.Validate if the designated constraints aren't met.
type CreateFlowSetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateFlowSetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateFlowSetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateFlowSetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateFlowSetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateFlowSetResponseValidationError) ErrorName() string {
	return "CreateFlowSetResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateFlowSetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateFlowSetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateFlowSetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateFlowSetResponseValidationError{}
