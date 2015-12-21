package store

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	// ErrFieldRequired is returned when a field is required and receives a
	// request to be deleted.
	ErrFieldRequired = errors.New("field is required")

	// ErrMissingKey is returned when a key is missing
	ErrMissingKey = errors.New("key not found")

	// ErrMissingField is returned when a field is missing
	ErrMissingField = errors.New("no such field")

	// ErrMissingBackend is returned with a given backend is not found
	ErrMissingBackend = errors.New("backend not found")
)

// DecodeError is returned for errors in decoding values
type DecodeError struct {
	Field string
	Err   error
}

func (d *DecodeError) Error() string {
	if err, ok := d.Err.(*strconv.NumError); ok {
		return fmt.Sprintf(`%s: parsing "%s": %s`, d.Field, err.Num, err.Err)
	}

	return fmt.Sprintf("%s: %s", d.Field, d.Err)
}

// MultiError is multiple errors rolled into one
type MultiError struct {
	errs []error
}

// NewMultiError wraps a number of errors in a MutliError
func NewMultiError(errs ...error) *MultiError {
	return &MultiError{errs}
}

// Append adds a new error onto the end of the MultiError
func (m *MultiError) Append(err error) {
	m.errs = append(m.errs, err)
}

func (m *MultiError) Error() string {
	errs := []string{}
	for _, err := range m.errs {
		errs = append(errs, err.Error())
	}

	return strings.Join(errs, "\n")
}
