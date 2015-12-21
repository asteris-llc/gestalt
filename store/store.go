package store

import (
	"github.com/asteris-llc/gestalt/web/app"
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

	return nil, ErrMissingBackend
}

func (s *Store) getBackendForSchema(target *app.Schema) (*Backend, error) {
	if target.Backend == "" {
		return s.defaultStore, nil
	}

	return s.getBackend(target.Backend)
}
