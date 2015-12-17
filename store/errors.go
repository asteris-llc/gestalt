package store

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	// ErrFieldRequired is returned when a field is required and receives a
	// request to be deleted.
	ErrFieldRequired = errors.New("field is required")

	// ErrMissingKey is returned when a key is missing
	ErrMissingKey = errors.New("key not found")

	// ErrMissingField is returned when a field is missing
	ErrMissingField = errors.New("no such field") // TODO: probably move to schema?

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
