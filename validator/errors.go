package validator

import (
	"errors"
)

var (
	// ErrNoField is returned when a field could not be found
	ErrNoField = errors.New("no field found")

	// ErrFieldRequired is returned when a field is required
	ErrFieldRequired = errors.New("field is required")
)

// ValidationError is returned when a value is not valid. This does look a
// little unnecessary, but it's useful to be able to distinguish errors we
// should present to the user versus errors that should just be in the logs
type ValidationError struct {
	Err error
}

func (v *ValidationError) Error() string {
	return v.Err.Error()
}
