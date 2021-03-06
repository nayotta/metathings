// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: create_evaluator.proto

package evaluatord

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
var _create_evaluator_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CreateEvaluatorRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateEvaluatorRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetEvaluator() == nil {
		return CreateEvaluatorRequestValidationError{
			field:  "Evaluator",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetEvaluator()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateEvaluatorRequestValidationError{
				field:  "Evaluator",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateEvaluatorRequestValidationError is the validation error returned by
// CreateEvaluatorRequest.Validate if the designated constraints aren't met.
type CreateEvaluatorRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateEvaluatorRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateEvaluatorRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateEvaluatorRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateEvaluatorRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateEvaluatorRequestValidationError) ErrorName() string {
	return "CreateEvaluatorRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateEvaluatorRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateEvaluatorRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateEvaluatorRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateEvaluatorRequestValidationError{}

// Validate checks the field values on CreateEvaluatorResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateEvaluatorResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetEvaluator()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateEvaluatorResponseValidationError{
				field:  "Evaluator",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateEvaluatorResponseValidationError is the validation error returned by
// CreateEvaluatorResponse.Validate if the designated constraints aren't met.
type CreateEvaluatorResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateEvaluatorResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateEvaluatorResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateEvaluatorResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateEvaluatorResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateEvaluatorResponseValidationError) ErrorName() string {
	return "CreateEvaluatorResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateEvaluatorResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateEvaluatorResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateEvaluatorResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateEvaluatorResponseValidationError{}
