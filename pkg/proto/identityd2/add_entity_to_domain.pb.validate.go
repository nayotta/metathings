// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: add_entity_to_domain.proto

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
var _add_entity_to_domain_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on AddEntityToDomainRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AddEntityToDomainRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetDomain() == nil {
		return AddEntityToDomainRequestValidationError{
			field:  "Domain",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetDomain()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddEntityToDomainRequestValidationError{
				field:  "Domain",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetEntity() == nil {
		return AddEntityToDomainRequestValidationError{
			field:  "Entity",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetEntity()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddEntityToDomainRequestValidationError{
				field:  "Entity",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// AddEntityToDomainRequestValidationError is the validation error returned by
// AddEntityToDomainRequest.Validate if the designated constraints aren't met.
type AddEntityToDomainRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddEntityToDomainRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddEntityToDomainRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddEntityToDomainRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddEntityToDomainRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddEntityToDomainRequestValidationError) ErrorName() string {
	return "AddEntityToDomainRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddEntityToDomainRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddEntityToDomainRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddEntityToDomainRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddEntityToDomainRequestValidationError{}
