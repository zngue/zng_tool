// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/test/v1/test.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on UpdateUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateUserRequestMultiError, or nil if none found.
func (m *UpdateUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Age

	// no validation rules for Sex

	// no validation rules for Image

	if len(errors) > 0 {
		return UpdateUserRequestMultiError(errors)
	}

	return nil
}

// UpdateUserRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateUserRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateUserRequestMultiError) AllErrors() []error { return m }

// UpdateUserRequestValidationError is the validation error returned by
// UpdateUserRequest.Validate if the designated constraints aren't met.
type UpdateUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserRequestValidationError) ErrorName() string {
	return "UpdateUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUserRequestValidationError) Error() string {
	if strings.Contains(e.reason, "syMsg") {
		return strings.Trim(e.Reason(), "syMsg")
	}
	return e.field + e.reason
}

var _ error = UpdateUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserRequestValidationError{}

// Validate checks the field values on GetList2TestRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetList2TestRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetList2TestRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetList2TestRequestMultiError, or nil if none found.
func (m *GetList2TestRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetList2TestRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if _, ok := _GetList2TestRequest_Status_InLookup[m.GetStatus()]; !ok {
		err := GetList2TestRequestValidationError{
			field:  "status",
			reason: "的值必须在列表中 [0 1 2 3 4]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Age

	// no validation rules for Page

	// no validation rules for PageSize

	// no validation rules for Order

	if len(errors) > 0 {
		return GetList2TestRequestMultiError(errors)
	}

	return nil
}

// GetList2TestRequestMultiError is an error wrapping multiple validation
// errors returned by GetList2TestRequest.ValidateAll() if the designated
// constraints aren't met.
type GetList2TestRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetList2TestRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetList2TestRequestMultiError) AllErrors() []error { return m }

// GetList2TestRequestValidationError is the validation error returned by
// GetList2TestRequest.Validate if the designated constraints aren't met.
type GetList2TestRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetList2TestRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetList2TestRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetList2TestRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetList2TestRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetList2TestRequestValidationError) ErrorName() string {
	return "GetList2TestRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetList2TestRequestValidationError) Error() string {
	if strings.Contains(e.reason, "syMsg") {
		return strings.Trim(e.Reason(), "syMsg")
	}
	return e.field + e.reason
}

var _ error = GetList2TestRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetList2TestRequestValidationError{}

var _GetList2TestRequest_Status_InLookup = map[int32]struct{}{
	0: {},
	1: {},
	2: {},
	3: {},
	4: {},
}

// Validate checks the field values on UserList2Rely with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserList2Rely) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserList2Rely with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserList2RelyMultiError, or
// nil if none found.
func (m *UserList2Rely) ValidateAll() error {
	return m.validate(true)
}

func (m *UserList2Rely) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetUser() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UserList2RelyValidationError{
						field:  fmt.Sprintf("User[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UserList2RelyValidationError{
						field:  fmt.Sprintf("User[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserList2RelyValidationError{
					field:  fmt.Sprintf("User[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	// no validation rules for Page

	// no validation rules for PageSize

	// no validation rules for Order

	// no validation rules for TotalPage

	// no validation rules for Status

	// no validation rules for Age

	if len(errors) > 0 {
		return UserList2RelyMultiError(errors)
	}

	return nil
}

// UserList2RelyMultiError is an error wrapping multiple validation errors
// returned by UserList2Rely.ValidateAll() if the designated constraints
// aren't met.
type UserList2RelyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserList2RelyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserList2RelyMultiError) AllErrors() []error { return m }

// UserList2RelyValidationError is the validation error returned by
// UserList2Rely.Validate if the designated constraints aren't met.
type UserList2RelyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserList2RelyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserList2RelyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserList2RelyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserList2RelyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserList2RelyValidationError) ErrorName() string { return "UserList2RelyValidationError" }

// Error satisfies the builtin error interface
func (e UserList2RelyValidationError) Error() string {
	if strings.Contains(e.reason, "syMsg") {
		return strings.Trim(e.Reason(), "syMsg")
	}
	return e.field + e.reason
}

var _ error = UserList2RelyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserList2RelyValidationError{}

// Validate checks the field values on GetListTestRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetListTestRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetListTestRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetListTestRequestMultiError, or nil if none found.
func (m *GetListTestRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetListTestRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return GetListTestRequestMultiError(errors)
	}

	return nil
}

// GetListTestRequestMultiError is an error wrapping multiple validation errors
// returned by GetListTestRequest.ValidateAll() if the designated constraints
// aren't met.
type GetListTestRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetListTestRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetListTestRequestMultiError) AllErrors() []error { return m }

// GetListTestRequestValidationError is the validation error returned by
// GetListTestRequest.Validate if the designated constraints aren't met.
type GetListTestRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetListTestRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetListTestRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetListTestRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetListTestRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetListTestRequestValidationError) ErrorName() string {
	return "GetListTestRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetListTestRequestValidationError) Error() string {
	if strings.Contains(e.reason, "syMsg") {
		return strings.Trim(e.Reason(), "syMsg")
	}
	return e.field + e.reason
}

var _ error = GetListTestRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetListTestRequestValidationError{}

// Validate checks the field values on GetTestRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetTestRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTestRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetTestRequestMultiError,
// or nil if none found.
func (m *GetTestRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTestRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		var err error

		err = GetTestRequestValidationError{
			field:  "name",
			reason: "的长度必须最少为 1",
		}

		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _GetTestRequest_Status_InLookup[m.GetStatus()]; !ok {
		err := GetTestRequestValidationError{
			field:  "status",
			reason: "的值必须在列表中 [0 1 2 3 4]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetTestRequestMultiError(errors)
	}

	return nil
}

// GetTestRequestMultiError is an error wrapping multiple validation errors
// returned by GetTestRequest.ValidateAll() if the designated constraints
// aren't met.
type GetTestRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTestRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTestRequestMultiError) AllErrors() []error { return m }

// GetTestRequestValidationError is the validation error returned by
// GetTestRequest.Validate if the designated constraints aren't met.
type GetTestRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTestRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTestRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTestRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTestRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTestRequestValidationError) ErrorName() string { return "GetTestRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetTestRequestValidationError) Error() string {
	if strings.Contains(e.reason, "syMsg") {
		return strings.Trim(e.Reason(), "syMsg")
	}
	return e.field + e.reason
}

var _ error = GetTestRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTestRequestValidationError{}

var _GetTestRequest_Status_InLookup = map[int32]struct{}{
	0: {},
	1: {},
	2: {},
	3: {},
	4: {},
}

// Validate checks the field values on GetTestReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetTestReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTestReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetTestReplyMultiError, or
// nil if none found.
func (m *GetTestReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTestReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return GetTestReplyMultiError(errors)
	}

	return nil
}

// GetTestReplyMultiError is an error wrapping multiple validation errors
// returned by GetTestReply.ValidateAll() if the designated constraints aren't met.
type GetTestReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTestReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTestReplyMultiError) AllErrors() []error { return m }

// GetTestReplyValidationError is the validation error returned by
// GetTestReply.Validate if the designated constraints aren't met.
type GetTestReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTestReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTestReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTestReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTestReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTestReplyValidationError) ErrorName() string { return "GetTestReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetTestReplyValidationError) Error() string {
	if strings.Contains(e.reason, "syMsg") {
		return strings.Trim(e.Reason(), "syMsg")
	}
	return e.field + e.reason
}

var _ error = GetTestReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTestReplyValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Age

	if all {
		switch v := interface{}(m.GetUserExit()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "userExit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "userExit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUserExit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "userExit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UserType

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	if strings.Contains(e.reason, "syMsg") {
		return strings.Trim(e.Reason(), "syMsg")
	}
	return e.field + e.reason
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on Data with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Data) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Data with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in DataMultiError, or nil if none found.
func (m *Data) ValidateAll() error {
	return m.validate(true)
}

func (m *Data) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return DataMultiError(errors)
	}

	return nil
}

// DataMultiError is an error wrapping multiple validation errors returned by
// Data.ValidateAll() if the designated constraints aren't met.
type DataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DataMultiError) AllErrors() []error { return m }

// DataValidationError is the validation error returned by Data.Validate if the
// designated constraints aren't met.
type DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataValidationError) ErrorName() string { return "DataValidationError" }

// Error satisfies the builtin error interface
func (e DataValidationError) Error() string {
	if strings.Contains(e.reason, "syMsg") {
		return strings.Trim(e.Reason(), "syMsg")
	}
	return e.field + e.reason
}

var _ error = DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataValidationError{}

// Validate checks the field values on UserExit with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserExit) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserExit with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserExitMultiError, or nil
// if none found.
func (m *UserExit) ValidateAll() error {
	return m.validate(true)
}

func (m *UserExit) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Image

	// no validation rules for Height

	// no validation rules for Weight

	if len(errors) > 0 {
		return UserExitMultiError(errors)
	}

	return nil
}

// UserExitMultiError is an error wrapping multiple validation errors returned
// by UserExit.ValidateAll() if the designated constraints aren't met.
type UserExitMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserExitMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserExitMultiError) AllErrors() []error { return m }

// UserExitValidationError is the validation error returned by
// UserExit.Validate if the designated constraints aren't met.
type UserExitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserExitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserExitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserExitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserExitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserExitValidationError) ErrorName() string { return "UserExitValidationError" }

// Error satisfies the builtin error interface
func (e UserExitValidationError) Error() string {
	if strings.Contains(e.reason, "syMsg") {
		return strings.Trim(e.Reason(), "syMsg")
	}
	return e.field + e.reason
}

var _ error = UserExitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserExitValidationError{}

// Validate checks the field values on UserList with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserList) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserList with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserListMultiError, or nil
// if none found.
func (m *UserList) ValidateAll() error {
	return m.validate(true)
}

func (m *UserList) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetUser() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UserListValidationError{
						field:  fmt.Sprintf("User[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UserListValidationError{
						field:  fmt.Sprintf("User[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserListValidationError{
					field:  fmt.Sprintf("User[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return UserListMultiError(errors)
	}

	return nil
}

// UserListMultiError is an error wrapping multiple validation errors returned
// by UserList.ValidateAll() if the designated constraints aren't met.
type UserListMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserListMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserListMultiError) AllErrors() []error { return m }

// UserListValidationError is the validation error returned by
// UserList.Validate if the designated constraints aren't met.
type UserListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserListValidationError) ErrorName() string { return "UserListValidationError" }

// Error satisfies the builtin error interface
func (e UserListValidationError) Error() string {
	if strings.Contains(e.reason, "syMsg") {
		return strings.Trim(e.Reason(), "syMsg")
	}
	return e.field + e.reason
}

var _ error = UserListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserListValidationError{}
