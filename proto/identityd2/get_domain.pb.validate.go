// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: get_domain.proto

package identityd2

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
var _get_domain_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetDomainRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetDomainRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetDomain() == nil {
		return GetDomainRequestValidationError{
			field:  "Domain",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetDomain()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetDomainRequestValidationError{
				field:  "Domain",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetDomainRequestValidationError is the validation error returned by
// GetDomainRequest.Validate if the designated constraints aren't met.
type GetDomainRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDomainRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDomainRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDomainRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDomainRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDomainRequestValidationError) ErrorName() string { return "GetDomainRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetDomainRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDomainRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDomainRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDomainRequestValidationError{}

// Validate checks the field values on GetDomainResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetDomainResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetDomain()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetDomainResponseValidationError{
				field:  "Domain",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetDomainResponseValidationError is the validation error returned by
// GetDomainResponse.Validate if the designated constraints aren't met.
type GetDomainResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDomainResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDomainResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDomainResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDomainResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDomainResponseValidationError) ErrorName() string {
	return "GetDomainResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetDomainResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDomainResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDomainResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDomainResponseValidationError{}
