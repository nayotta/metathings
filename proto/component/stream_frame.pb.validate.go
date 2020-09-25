// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: stream_frame.proto

package component

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
var _stream_frame_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on OpStreamCallConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *OpStreamCallConfig) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSession()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpStreamCallConfigValidationError{
				field:  "Session",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMethod()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpStreamCallConfigValidationError{
				field:  "Method",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAck()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpStreamCallConfigValidationError{
				field:  "Ack",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// OpStreamCallConfigValidationError is the validation error returned by
// OpStreamCallConfig.Validate if the designated constraints aren't met.
type OpStreamCallConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpStreamCallConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpStreamCallConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpStreamCallConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpStreamCallConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpStreamCallConfigValidationError) ErrorName() string {
	return "OpStreamCallConfigValidationError"
}

// Error satisfies the builtin error interface
func (e OpStreamCallConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpStreamCallConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpStreamCallConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpStreamCallConfigValidationError{}

// Validate checks the field values on StreamCallConfig with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *StreamCallConfig) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Session

	// no validation rules for Method

	// no validation rules for Ack

	return nil
}

// StreamCallConfigValidationError is the validation error returned by
// StreamCallConfig.Validate if the designated constraints aren't met.
type StreamCallConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamCallConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamCallConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamCallConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamCallConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamCallConfigValidationError) ErrorName() string { return "StreamCallConfigValidationError" }

// Error satisfies the builtin error interface
func (e StreamCallConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamCallConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamCallConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamCallConfigValidationError{}

// Validate checks the field values on OpStreamCallAck with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *OpStreamCallAck) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpStreamCallAckValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// OpStreamCallAckValidationError is the validation error returned by
// OpStreamCallAck.Validate if the designated constraints aren't met.
type OpStreamCallAckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpStreamCallAckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpStreamCallAckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpStreamCallAckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpStreamCallAckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpStreamCallAckValidationError) ErrorName() string { return "OpStreamCallAckValidationError" }

// Error satisfies the builtin error interface
func (e OpStreamCallAckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpStreamCallAck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpStreamCallAckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpStreamCallAckValidationError{}

// Validate checks the field values on StreamCallAck with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *StreamCallAck) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Value

	return nil
}

// StreamCallAckValidationError is the validation error returned by
// StreamCallAck.Validate if the designated constraints aren't met.
type StreamCallAckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamCallAckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamCallAckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamCallAckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamCallAckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamCallAckValidationError) ErrorName() string { return "StreamCallAckValidationError" }

// Error satisfies the builtin error interface
func (e StreamCallAckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamCallAck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamCallAckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamCallAckValidationError{}

// Validate checks the field values on OpStreamCallExit with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *OpStreamCallExit) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// OpStreamCallExitValidationError is the validation error returned by
// OpStreamCallExit.Validate if the designated constraints aren't met.
type OpStreamCallExitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpStreamCallExitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpStreamCallExitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpStreamCallExitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpStreamCallExitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpStreamCallExitValidationError) ErrorName() string { return "OpStreamCallExitValidationError" }

// Error satisfies the builtin error interface
func (e OpStreamCallExitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpStreamCallExit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpStreamCallExitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpStreamCallExitValidationError{}

// Validate checks the field values on StreamCallExit with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *StreamCallExit) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// StreamCallExitValidationError is the validation error returned by
// StreamCallExit.Validate if the designated constraints aren't met.
type StreamCallExitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamCallExitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamCallExitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamCallExitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamCallExitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamCallExitValidationError) ErrorName() string { return "StreamCallExitValidationError" }

// Error satisfies the builtin error interface
func (e StreamCallExitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamCallExit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamCallExitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamCallExitValidationError{}

// Validate checks the field values on OpUnaryCallValue with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *OpUnaryCallValue) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetSession() == nil {
		return OpUnaryCallValueValidationError{
			field:  "Session",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetSession()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpUnaryCallValueValidationError{
				field:  "Session",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetMethod() == nil {
		return OpUnaryCallValueValidationError{
			field:  "Method",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetMethod()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpUnaryCallValueValidationError{
				field:  "Method",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetValue() == nil {
		return OpUnaryCallValueValidationError{
			field:  "Value",
			reason: "value is required",
		}
	}

	if a := m.GetValue(); a != nil {

	}

	return nil
}

// OpUnaryCallValueValidationError is the validation error returned by
// OpUnaryCallValue.Validate if the designated constraints aren't met.
type OpUnaryCallValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpUnaryCallValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpUnaryCallValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpUnaryCallValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpUnaryCallValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpUnaryCallValueValidationError) ErrorName() string { return "OpUnaryCallValueValidationError" }

// Error satisfies the builtin error interface
func (e OpUnaryCallValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpUnaryCallValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpUnaryCallValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpUnaryCallValueValidationError{}

// Validate checks the field values on UnaryCallValue with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UnaryCallValue) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Session

	// no validation rules for Method

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UnaryCallValueValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UnaryCallValueValidationError is the validation error returned by
// UnaryCallValue.Validate if the designated constraints aren't met.
type UnaryCallValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnaryCallValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnaryCallValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnaryCallValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnaryCallValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnaryCallValueValidationError) ErrorName() string { return "UnaryCallValueValidationError" }

// Error satisfies the builtin error interface
func (e UnaryCallValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnaryCallValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnaryCallValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnaryCallValueValidationError{}

// Validate checks the field values on OpStreamCallValue with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *OpStreamCallValue) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Union.(type) {

	case *OpStreamCallValue_Value:

		if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OpStreamCallValueValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *OpStreamCallValue_Config:

		if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OpStreamCallValueValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *OpStreamCallValue_Ack:

		if v, ok := interface{}(m.GetAck()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OpStreamCallValueValidationError{
					field:  "Ack",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *OpStreamCallValue_Exit:

		if v, ok := interface{}(m.GetExit()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OpStreamCallValueValidationError{
					field:  "Exit",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// OpStreamCallValueValidationError is the validation error returned by
// OpStreamCallValue.Validate if the designated constraints aren't met.
type OpStreamCallValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpStreamCallValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpStreamCallValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpStreamCallValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpStreamCallValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpStreamCallValueValidationError) ErrorName() string {
	return "OpStreamCallValueValidationError"
}

// Error satisfies the builtin error interface
func (e OpStreamCallValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpStreamCallValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpStreamCallValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpStreamCallValueValidationError{}

// Validate checks the field values on StreamCallValue with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *StreamCallValue) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Union.(type) {

	case *StreamCallValue_Value:

		if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamCallValueValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *StreamCallValue_Config:

		if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamCallValueValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *StreamCallValue_Ack:

		if v, ok := interface{}(m.GetAck()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamCallValueValidationError{
					field:  "Ack",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *StreamCallValue_Exit:

		if v, ok := interface{}(m.GetExit()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamCallValueValidationError{
					field:  "Exit",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// StreamCallValueValidationError is the validation error returned by
// StreamCallValue.Validate if the designated constraints aren't met.
type StreamCallValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamCallValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamCallValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamCallValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamCallValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamCallValueValidationError) ErrorName() string { return "StreamCallValueValidationError" }

// Error satisfies the builtin error interface
func (e StreamCallValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamCallValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamCallValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamCallValueValidationError{}

// Validate checks the field values on ErrorValue with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ErrorValue) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Message

	return nil
}

// ErrorValueValidationError is the validation error returned by
// ErrorValue.Validate if the designated constraints aren't met.
type ErrorValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ErrorValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ErrorValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ErrorValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ErrorValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ErrorValueValidationError) ErrorName() string { return "ErrorValueValidationError" }

// Error satisfies the builtin error interface
func (e ErrorValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sErrorValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ErrorValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ErrorValueValidationError{}

// Validate checks the field values on UpStreamFrame with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UpStreamFrame) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Kind

	switch m.Union.(type) {

	case *UpStreamFrame_UnaryCall:

		if v, ok := interface{}(m.GetUnaryCall()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpStreamFrameValidationError{
					field:  "UnaryCall",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *UpStreamFrame_StreamCall:

		if v, ok := interface{}(m.GetStreamCall()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpStreamFrameValidationError{
					field:  "StreamCall",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *UpStreamFrame_Error:

		if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpStreamFrameValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// UpStreamFrameValidationError is the validation error returned by
// UpStreamFrame.Validate if the designated constraints aren't met.
type UpStreamFrameValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpStreamFrameValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpStreamFrameValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpStreamFrameValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpStreamFrameValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpStreamFrameValidationError) ErrorName() string { return "UpStreamFrameValidationError" }

// Error satisfies the builtin error interface
func (e UpStreamFrameValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpStreamFrame.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpStreamFrameValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpStreamFrameValidationError{}

// Validate checks the field values on DownStreamFrame with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DownStreamFrame) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := StreamFrameKind_name[int32(m.GetKind())]; !ok {
		return DownStreamFrameValidationError{
			field:  "Kind",
			reason: "value must be one of the defined enum values",
		}
	}

	switch m.Union.(type) {

	case *DownStreamFrame_UnaryCall:

		if v, ok := interface{}(m.GetUnaryCall()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DownStreamFrameValidationError{
					field:  "UnaryCall",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *DownStreamFrame_StreamCall:

		if v, ok := interface{}(m.GetStreamCall()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DownStreamFrameValidationError{
					field:  "StreamCall",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// DownStreamFrameValidationError is the validation error returned by
// DownStreamFrame.Validate if the designated constraints aren't met.
type DownStreamFrameValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownStreamFrameValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownStreamFrameValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownStreamFrameValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownStreamFrameValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownStreamFrameValidationError) ErrorName() string { return "DownStreamFrameValidationError" }

// Error satisfies the builtin error interface
func (e DownStreamFrameValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownStreamFrame.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownStreamFrameValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownStreamFrameValidationError{}
