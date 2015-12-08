package store

import (
	"errors"
	"github.com/docker/libkv/store"
	"github.com/xeipuuv/gojsonschema"
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

// StoreSchema stores or updates a schema as a JSON blob and makes sure that the
// schema is valid before storage.
func (s *Store) StoreSchema(name string, schemaBlob []byte) error {
	// validate that the schema is valid
	_, err := gojsonschema.NewSchema(gojsonschema.NewStringLoader(string(schemaBlob)))
	if err != nil {
		return err
	}

	return s.schemaStore.Put(
		ensurePrefix(s.schemaStore.Prefix, name),
		schemaBlob,
		&store.WriteOptions{},
	)
}

// RetrieveSchema gets a schema as a JSON blob
func (s *Store) RetrieveSchema(name string) ([]byte, error) {
	pair, err := s.schemaStore.Get(ensurePrefix(s.schemaStore.Prefix, name))

	if err != nil {
		return []byte{}, err
	}

	if pair.Key == "" {
		return []byte{}, ErrNotFound
	}

	return pair.Value, err
}

// DeleteSchema removes a schema from the K/V tree
func (s *Store) DeleteSchema(name string) error {
	return s.schemaStore.Delete(ensurePrefix(s.schemaStore.Prefix, name))
}
