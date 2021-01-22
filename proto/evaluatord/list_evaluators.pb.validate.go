// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: list_evaluators.proto

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
var _list_evaluators_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ListEvaluatorsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListEvaluatorsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetEvaluator() == nil {
		return ListEvaluatorsRequestValidationError{
			field:  "Evaluator",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetEvaluator()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListEvaluatorsRequestValidationError{
				field:  "Evaluator",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ListEvaluatorsRequestValidationError is the validation error returned by
// ListEvaluatorsRequest.Validate if the designated constraints aren't met.
type ListEvaluatorsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListEvaluatorsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListEvaluatorsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListEvaluatorsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListEvaluatorsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListEvaluatorsRequestValidationError) ErrorName() string {
	return "ListEvaluatorsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListEvaluatorsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListEvaluatorsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListEvaluatorsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListEvaluatorsRequestValidationError{}

// Validate checks the field values on ListEvaluatorsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListEvaluatorsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetEvaluators() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListEvaluatorsResponseValidationError{
					field:  fmt.Sprintf("Evaluators[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListEvaluatorsResponseValidationError is the validation error returned by
// ListEvaluatorsResponse.Validate if the designated constraints aren't met.
type ListEvaluatorsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListEvaluatorsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListEvaluatorsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListEvaluatorsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListEvaluatorsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListEvaluatorsResponseValidationError) ErrorName() string {
	return "ListEvaluatorsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListEvaluatorsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListEvaluatorsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListEvaluatorsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListEvaluatorsResponseValidationError{}