package validator

import (
	"errors"
)

var (
	// ErrNoField is returned when a field could not be found
	ErrNoField = errors.New("no field found")
)
