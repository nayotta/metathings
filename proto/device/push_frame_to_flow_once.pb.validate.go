// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: push_frame_to_flow_once.proto

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
var _push_frame_to_flow_once_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on PushFrameToFlowOnceRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PushFrameToFlowOnceRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() == nil {
		return PushFrameToFlowOnceRequestValidationError{
			field:  "Id",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PushFrameToFlowOnceRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetFlow() == nil {
		return PushFrameToFlowOnceRequestValidationError{
			field:  "Flow",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetFlow()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PushFrameToFlowOnceRequestValidationError{
				field:  "Flow",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetFrame() == nil {
		return PushFrameToFlowOnceRequestValidationError{
			field:  "Frame",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetFrame()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PushFrameToFlowOnceRequestValidationError{
				field:  "Frame",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PushFrameToFlowOnceRequestValidationError is the validation error returned
// by PushFrameToFlowOnceRequest.Validate if the designated constraints aren't met.
type PushFrameToFlowOnceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PushFrameToFlowOnceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PushFrameToFlowOnceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PushFrameToFlowOnceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PushFrameToFlowOnceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PushFrameToFlowOnceRequestValidationError) ErrorName() string {
	return "PushFrameToFlowOnceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PushFrameToFlowOnceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPushFrameToFlowOnceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PushFrameToFlowOnceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PushFrameToFlowOnceRequestValidationError{}
