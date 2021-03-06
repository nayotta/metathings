// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: list_credentials_for_entity.proto

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
var _list_credentials_for_entity_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ListCredentialsForEntityRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListCredentialsForEntityRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetEntity()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListCredentialsForEntityRequestValidationError{
				field:  "Entity",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ListCredentialsForEntityRequestValidationError is the validation error
// returned by ListCredentialsForEntityRequest.Validate if the designated
// constraints aren't met.
type ListCredentialsForEntityRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCredentialsForEntityRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCredentialsForEntityRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCredentialsForEntityRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCredentialsForEntityRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCredentialsForEntityRequestValidationError) ErrorName() string {
	return "ListCredentialsForEntityRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListCredentialsForEntityRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCredentialsForEntityRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCredentialsForEntityRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCredentialsForEntityRequestValidationError{}

// Validate checks the field values on ListCredentialsForEntityResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ListCredentialsForEntityResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetCredentials() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListCredentialsForEntityResponseValidationError{
					field:  fmt.Sprintf("Credentials[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListCredentialsForEntityResponseValidationError is the validation error
// returned by ListCredentialsForEntityResponse.Validate if the designated
// constraints aren't met.
type ListCredentialsForEntityResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCredentialsForEntityResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCredentialsForEntityResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCredentialsForEntityResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCredentialsForEntityResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCredentialsForEntityResponseValidationError) ErrorName() string {
	return "ListCredentialsForEntityResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListCredentialsForEntityResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCredentialsForEntityResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCredentialsForEntityResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCredentialsForEntityResponseValidationError{}
