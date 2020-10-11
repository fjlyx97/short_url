// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: short_url.proto

package ShortUrl

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
var _short_url_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on SetUrlReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SetUrlReq) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetUrl()) > 8182 {
		return SetUrlReqValidationError{
			field:  "Url",
			reason: "value length must be at most 8182 runes",
		}
	}

	if uri, err := url.Parse(m.GetUrl()); err != nil {
		return SetUrlReqValidationError{
			field:  "Url",
			reason: "value must be a valid URI",
			cause:  err,
		}
	} else if !uri.IsAbs() {
		return SetUrlReqValidationError{
			field:  "Url",
			reason: "value must be absolute",
		}
	}

	return nil
}

// SetUrlReqValidationError is the validation error returned by
// SetUrlReq.Validate if the designated constraints aren't met.
type SetUrlReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetUrlReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetUrlReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetUrlReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetUrlReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetUrlReqValidationError) ErrorName() string { return "SetUrlReqValidationError" }

// Error satisfies the builtin error interface
func (e SetUrlReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetUrlReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetUrlReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetUrlReqValidationError{}

// Validate checks the field values on SetUrlRsp with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SetUrlRsp) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for ShortUrl

	return nil
}

// SetUrlRspValidationError is the validation error returned by
// SetUrlRsp.Validate if the designated constraints aren't met.
type SetUrlRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetUrlRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetUrlRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetUrlRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetUrlRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetUrlRspValidationError) ErrorName() string { return "SetUrlRspValidationError" }

// Error satisfies the builtin error interface
func (e SetUrlRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetUrlRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetUrlRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetUrlRspValidationError{}

// Validate checks the field values on GetAfterForwardUrlReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetAfterForwardUrlReq) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetUrl()) > 20 {
		return GetAfterForwardUrlReqValidationError{
			field:  "Url",
			reason: "value length must be at most 20 runes",
		}
	}

	return nil
}

// GetAfterForwardUrlReqValidationError is the validation error returned by
// GetAfterForwardUrlReq.Validate if the designated constraints aren't met.
type GetAfterForwardUrlReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAfterForwardUrlReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAfterForwardUrlReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAfterForwardUrlReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAfterForwardUrlReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAfterForwardUrlReqValidationError) ErrorName() string {
	return "GetAfterForwardUrlReqValidationError"
}

// Error satisfies the builtin error interface
func (e GetAfterForwardUrlReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAfterForwardUrlReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAfterForwardUrlReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAfterForwardUrlReqValidationError{}

// Validate checks the field values on GetAfterForwardUrlRsp with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetAfterForwardUrlRsp) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for LongUrl

	return nil
}

// GetAfterForwardUrlRspValidationError is the validation error returned by
// GetAfterForwardUrlRsp.Validate if the designated constraints aren't met.
type GetAfterForwardUrlRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAfterForwardUrlRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAfterForwardUrlRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAfterForwardUrlRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAfterForwardUrlRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAfterForwardUrlRspValidationError) ErrorName() string {
	return "GetAfterForwardUrlRspValidationError"
}

// Error satisfies the builtin error interface
func (e GetAfterForwardUrlRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAfterForwardUrlRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAfterForwardUrlRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAfterForwardUrlRspValidationError{}
