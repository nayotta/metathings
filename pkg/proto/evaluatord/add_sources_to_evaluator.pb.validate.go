// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: add_sources_to_evaluator.proto

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
var _add_sources_to_evaluator_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on AddSourcesToEvaluatorRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AddSourcesToEvaluatorRequest) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetSources() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AddSourcesToEvaluatorRequestValidationError{
					field:  fmt.Sprintf("Sources[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.GetEvaluator() == nil {
		return AddSourcesToEvaluatorRequestValidationError{
			field:  "Evaluator",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetEvaluator()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddSourcesToEvaluatorRequestValidationError{
				field:  "Evaluator",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// AddSourcesToEvaluatorRequestValidationError is the validation error returned
// by AddSourcesToEvaluatorRequest.Validate if the designated constraints
// aren't met.
type AddSourcesToEvaluatorRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddSourcesToEvaluatorRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddSourcesToEvaluatorRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddSourcesToEvaluatorRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddSourcesToEvaluatorRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddSourcesToEvaluatorRequestValidationError) ErrorName() string {
	return "AddSourcesToEvaluatorRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddSourcesToEvaluatorRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddSourcesToEvaluatorRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddSourcesToEvaluatorRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddSourcesToEvaluatorRequestValidationError{}
