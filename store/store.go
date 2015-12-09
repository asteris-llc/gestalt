package store

import (
	"errors"
	"github.com/asteris-llc/gestalt/schema"
)

var (
	// ErrNotFound is returned when a given schema is not found
	ErrNotFound = errors.New("not found")

	// ErrBackendNotFound is returned with a given backend is not found
	ErrBackendNotFound = errors.New("backend not found")
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

func (s *Store) getBackend(name string) (*Backend, error) {
	for _, backend := range s.backends {
		if backend.Name == name {
			return backend, nil
		}
	}

	return nil, ErrBackendNotFound
}

func (s *Store) getBackendForSchema(target *schema.Schema) (*Backend, error) {
	if target.Backend == "" {
		return s.defaultStore, nil
	}

	return s.getBackend(target.Backend)
}
