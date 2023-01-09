// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: note.proto

package note_v1

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

// Validate checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Empty) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EmptyMultiError, or nil if none found.
func (m *Empty) ValidateAll() error {
	return m.validate(true)
}

func (m *Empty) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EmptyMultiError(errors)
	}

	return nil
}

// EmptyMultiError is an error wrapping multiple validation errors returned by
// Empty.ValidateAll() if the designated constraints aren't met.
type EmptyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmptyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmptyMultiError) AllErrors() []error { return m }

// EmptyValidationError is the validation error returned by Empty.Validate if
// the designated constraints aren't met.
type EmptyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyValidationError) ErrorName() string { return "EmptyValidationError" }

// Error satisfies the builtin error interface
func (e EmptyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmpty.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyValidationError{}

// Validate checks the field values on NoteInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NoteInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NoteInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NoteInfoMultiError, or nil
// if none found.
func (m *NoteInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *NoteInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Title

	// no validation rules for Text

	// no validation rules for Author

	if len(errors) > 0 {
		return NoteInfoMultiError(errors)
	}

	return nil
}

// NoteInfoMultiError is an error wrapping multiple validation errors returned
// by NoteInfo.ValidateAll() if the designated constraints aren't met.
type NoteInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NoteInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NoteInfoMultiError) AllErrors() []error { return m }

// NoteInfoValidationError is the validation error returned by
// NoteInfo.Validate if the designated constraints aren't met.
type NoteInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NoteInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NoteInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NoteInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NoteInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NoteInfoValidationError) ErrorName() string { return "NoteInfoValidationError" }

// Error satisfies the builtin error interface
func (e NoteInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNoteInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NoteInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NoteInfoValidationError{}

// Validate checks the field values on Note with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Note) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Note with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in NoteMultiError, or nil if none found.
func (m *Note) ValidateAll() error {
	return m.validate(true)
}

func (m *Note) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NoteValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NoteValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NoteValidationError{
				field:  "Info",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NoteValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NoteValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NoteValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NoteValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NoteValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NoteValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return NoteMultiError(errors)
	}

	return nil
}

// NoteMultiError is an error wrapping multiple validation errors returned by
// Note.ValidateAll() if the designated constraints aren't met.
type NoteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NoteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NoteMultiError) AllErrors() []error { return m }

// NoteValidationError is the validation error returned by Note.Validate if the
// designated constraints aren't met.
type NoteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NoteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NoteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NoteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NoteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NoteValidationError) ErrorName() string { return "NoteValidationError" }

// Error satisfies the builtin error interface
func (e NoteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNote.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NoteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NoteValidationError{}

// Validate checks the field values on UpdateNoteInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UpdateNoteInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateNoteInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UpdateNoteInfoMultiError,
// or nil if none found.
func (m *UpdateNoteInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateNoteInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTitle()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateNoteInfoValidationError{
					field:  "Title",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateNoteInfoValidationError{
					field:  "Title",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTitle()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateNoteInfoValidationError{
				field:  "Title",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetText()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateNoteInfoValidationError{
					field:  "Text",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateNoteInfoValidationError{
					field:  "Text",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetText()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateNoteInfoValidationError{
				field:  "Text",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetAuthor()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateNoteInfoValidationError{
					field:  "Author",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateNoteInfoValidationError{
					field:  "Author",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAuthor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateNoteInfoValidationError{
				field:  "Author",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateNoteInfoMultiError(errors)
	}

	return nil
}

// UpdateNoteInfoMultiError is an error wrapping multiple validation errors
// returned by UpdateNoteInfo.ValidateAll() if the designated constraints
// aren't met.
type UpdateNoteInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateNoteInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateNoteInfoMultiError) AllErrors() []error { return m }

// UpdateNoteInfoValidationError is the validation error returned by
// UpdateNoteInfo.Validate if the designated constraints aren't met.
type UpdateNoteInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateNoteInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateNoteInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateNoteInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateNoteInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateNoteInfoValidationError) ErrorName() string { return "UpdateNoteInfoValidationError" }

// Error satisfies the builtin error interface
func (e UpdateNoteInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateNoteInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateNoteInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateNoteInfoValidationError{}

// Validate checks the field values on CreateNoteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateNoteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateNoteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateNoteRequestMultiError, or nil if none found.
func (m *CreateNoteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateNoteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetNote()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateNoteRequestValidationError{
					field:  "Note",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateNoteRequestValidationError{
					field:  "Note",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNote()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateNoteRequestValidationError{
				field:  "Note",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateNoteRequestMultiError(errors)
	}

	return nil
}

// CreateNoteRequestMultiError is an error wrapping multiple validation errors
// returned by CreateNoteRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateNoteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateNoteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateNoteRequestMultiError) AllErrors() []error { return m }

// CreateNoteRequestValidationError is the validation error returned by
// CreateNoteRequest.Validate if the designated constraints aren't met.
type CreateNoteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateNoteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateNoteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateNoteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateNoteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateNoteRequestValidationError) ErrorName() string {
	return "CreateNoteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateNoteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateNoteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateNoteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateNoteRequestValidationError{}

// Validate checks the field values on CreateNoteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateNoteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateNoteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateNoteResponseMultiError, or nil if none found.
func (m *CreateNoteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateNoteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateNoteResponseMultiError(errors)
	}

	return nil
}

// CreateNoteResponseMultiError is an error wrapping multiple validation errors
// returned by CreateNoteResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateNoteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateNoteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateNoteResponseMultiError) AllErrors() []error { return m }

// CreateNoteResponseValidationError is the validation error returned by
// CreateNoteResponse.Validate if the designated constraints aren't met.
type CreateNoteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateNoteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateNoteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateNoteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateNoteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateNoteResponseValidationError) ErrorName() string {
	return "CreateNoteResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateNoteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateNoteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateNoteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateNoteResponseValidationError{}

// Validate checks the field values on GetNoteRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetNoteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNoteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetNoteRequestMultiError,
// or nil if none found.
func (m *GetNoteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNoteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() < 1 {
		err := GetNoteRequestValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetNoteRequestMultiError(errors)
	}

	return nil
}

// GetNoteRequestMultiError is an error wrapping multiple validation errors
// returned by GetNoteRequest.ValidateAll() if the designated constraints
// aren't met.
type GetNoteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNoteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNoteRequestMultiError) AllErrors() []error { return m }

// GetNoteRequestValidationError is the validation error returned by
// GetNoteRequest.Validate if the designated constraints aren't met.
type GetNoteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNoteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNoteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNoteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNoteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNoteRequestValidationError) ErrorName() string { return "GetNoteRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetNoteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNoteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNoteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNoteRequestValidationError{}

// Validate checks the field values on GetNoteResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetNoteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNoteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetNoteResponseMultiError, or nil if none found.
func (m *GetNoteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNoteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetNote()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetNoteResponseValidationError{
					field:  "Note",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetNoteResponseValidationError{
					field:  "Note",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNote()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetNoteResponseValidationError{
				field:  "Note",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetNoteResponseMultiError(errors)
	}

	return nil
}

// GetNoteResponseMultiError is an error wrapping multiple validation errors
// returned by GetNoteResponse.ValidateAll() if the designated constraints
// aren't met.
type GetNoteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNoteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNoteResponseMultiError) AllErrors() []error { return m }

// GetNoteResponseValidationError is the validation error returned by
// GetNoteResponse.Validate if the designated constraints aren't met.
type GetNoteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNoteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNoteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNoteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNoteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNoteResponseValidationError) ErrorName() string { return "GetNoteResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetNoteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNoteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNoteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNoteResponseValidationError{}

// Validate checks the field values on GetListNoteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetListNoteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetListNoteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetListNoteResponseMultiError, or nil if none found.
func (m *GetListNoteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetListNoteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetNote() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetListNoteResponseValidationError{
						field:  fmt.Sprintf("Note[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetListNoteResponseValidationError{
						field:  fmt.Sprintf("Note[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetListNoteResponseValidationError{
					field:  fmt.Sprintf("Note[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetListNoteResponseMultiError(errors)
	}

	return nil
}

// GetListNoteResponseMultiError is an error wrapping multiple validation
// errors returned by GetListNoteResponse.ValidateAll() if the designated
// constraints aren't met.
type GetListNoteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetListNoteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetListNoteResponseMultiError) AllErrors() []error { return m }

// GetListNoteResponseValidationError is the validation error returned by
// GetListNoteResponse.Validate if the designated constraints aren't met.
type GetListNoteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetListNoteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetListNoteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetListNoteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetListNoteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetListNoteResponseValidationError) ErrorName() string {
	return "GetListNoteResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetListNoteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetListNoteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetListNoteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetListNoteResponseValidationError{}

// Validate checks the field values on UpdateNoteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateNoteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateNoteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateNoteRequestMultiError, or nil if none found.
func (m *UpdateNoteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateNoteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetNote()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateNoteRequestValidationError{
					field:  "Note",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateNoteRequestValidationError{
					field:  "Note",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNote()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateNoteRequestValidationError{
				field:  "Note",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateNoteRequestMultiError(errors)
	}

	return nil
}

// UpdateNoteRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateNoteRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateNoteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateNoteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateNoteRequestMultiError) AllErrors() []error { return m }

// UpdateNoteRequestValidationError is the validation error returned by
// UpdateNoteRequest.Validate if the designated constraints aren't met.
type UpdateNoteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateNoteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateNoteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateNoteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateNoteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateNoteRequestValidationError) ErrorName() string {
	return "UpdateNoteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateNoteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateNoteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateNoteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateNoteRequestValidationError{}

// Validate checks the field values on DeleteNoteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteNoteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteNoteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteNoteRequestMultiError, or nil if none found.
func (m *DeleteNoteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteNoteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteNoteRequestMultiError(errors)
	}

	return nil
}

// DeleteNoteRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteNoteRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteNoteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteNoteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteNoteRequestMultiError) AllErrors() []error { return m }

// DeleteNoteRequestValidationError is the validation error returned by
// DeleteNoteRequest.Validate if the designated constraints aren't met.
type DeleteNoteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteNoteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteNoteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteNoteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteNoteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteNoteRequestValidationError) ErrorName() string {
	return "DeleteNoteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteNoteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteNoteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteNoteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteNoteRequestValidationError{}
