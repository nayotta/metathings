// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: list_devices.proto

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
var _list_devices_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ListDevicesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListDevicesRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetDevice()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListDevicesRequestValidationError{
				field:  "Device",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ListDevicesRequestValidationError is the validation error returned by
// ListDevicesRequest.Validate if the designated constraints aren't met.
type ListDevicesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListDevicesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListDevicesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListDevicesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListDevicesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListDevicesRequestValidationError) ErrorName() string {
	return "ListDevicesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListDevicesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListDevicesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListDevicesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListDevicesRequestValidationError{}

// Validate checks the field values on ListDevicesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListDevicesResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetDevices() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListDevicesResponseValidationError{
					field:  fmt.Sprintf("Devices[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListDevicesResponseValidationError is the validation error returned by
// ListDevicesResponse.Validate if the designated constraints aren't met.
type ListDevicesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListDevicesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListDevicesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListDevicesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListDevicesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListDevicesResponseValidationError) ErrorName() string {
	return "ListDevicesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListDevicesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListDevicesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListDevicesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListDevicesResponseValidationError{}
