// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: patch_domain.proto

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
var _patch_domain_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on PatchDomainRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PatchDomainRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetDomain() == nil {
		return PatchDomainRequestValidationError{
			field:  "Domain",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetDomain()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PatchDomainRequestValidationError{
				field:  "Domain",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PatchDomainRequestValidationError is the validation error returned by
// PatchDomainRequest.Validate if the designated constraints aren't met.
type PatchDomainRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PatchDomainRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PatchDomainRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PatchDomainRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PatchDomainRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PatchDomainRequestValidationError) ErrorName() string {
	return "PatchDomainRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PatchDomainRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPatchDomainRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PatchDomainRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PatchDomainRequestValidationError{}

// Validate checks the field values on PatchDomainResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PatchDomainResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetDomain()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PatchDomainResponseValidationError{
				field:  "Domain",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PatchDomainResponseValidationError is the validation error returned by
// PatchDomainResponse.Validate if the designated constraints aren't met.
type PatchDomainResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PatchDomainResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PatchDomainResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PatchDomainResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PatchDomainResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PatchDomainResponseValidationError) ErrorName() string {
	return "PatchDomainResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PatchDomainResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPatchDomainResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PatchDomainResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PatchDomainResponseValidationError{}