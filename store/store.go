package store

import (
	"errors"
)

var (
	// ErrNotFound is returned when a given schema is not found
	ErrNotFound = errors.New("not found")
)

// Store stores and validates schemas.
type Store struct {
	backends     []*Backend
	schemaStore  *Backend
	defaultStore *Backend
}

// New returns a new Store with the given options
func New(backends []*Backend, schemaStore, defaultStore *Backend) (*Store, error) {
	store := &Store{
		backends:     backends,
		schemaStore:  schemaStore,
		defaultStore: defaultStore,
	}

	return store, nil
}
