// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: put_object.proto

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
var _put_object_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on PutObjectRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *PutObjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetObject() == nil {
		return PutObjectRequestValidationError{
			field:  "Object",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetObject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PutObjectRequestValidationError{
				field:  "Object",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetContent() == nil {
		return PutObjectRequestValidationError{
			field:  "Content",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetContent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PutObjectRequestValidationError{
				field:  "Content",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PutObjectRequestValidationError is the validation error returned by
// PutObjectRequest.Validate if the designated constraints aren't met.
type PutObjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PutObjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PutObjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PutObjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PutObjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PutObjectRequestValidationError) ErrorName() string { return "PutObjectRequestValidationError" }

// Error satisfies the builtin error interface
func (e PutObjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPutObjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PutObjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PutObjectRequestValidationError{}
