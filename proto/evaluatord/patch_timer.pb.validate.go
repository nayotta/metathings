// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: patch_timer.proto

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
var _patch_timer_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on PatchTimerRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *PatchTimerRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetTimer() == nil {
		return PatchTimerRequestValidationError{
			field:  "Timer",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetTimer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PatchTimerRequestValidationError{
				field:  "Timer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PatchTimerRequestValidationError is the validation error returned by
// PatchTimerRequest.Validate if the designated constraints aren't met.
type PatchTimerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PatchTimerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PatchTimerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PatchTimerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PatchTimerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PatchTimerRequestValidationError) ErrorName() string {
	return "PatchTimerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PatchTimerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPatchTimerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PatchTimerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PatchTimerRequestValidationError{}

// Validate checks the field values on PatchTimerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PatchTimerResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTimer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PatchTimerResponseValidationError{
				field:  "Timer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PatchTimerResponseValidationError is the validation error returned by
// PatchTimerResponse.Validate if the designated constraints aren't met.
type PatchTimerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PatchTimerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PatchTimerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PatchTimerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PatchTimerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PatchTimerResponseValidationError) ErrorName() string {
	return "PatchTimerResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PatchTimerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPatchTimerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PatchTimerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PatchTimerResponseValidationError{}
