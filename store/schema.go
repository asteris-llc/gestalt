package store

import (
	"encoding/json"
	"github.com/asteris-llc/gestalt/web/app"
	"github.com/docker/libkv/store"
)

// StoreSchema stores or updates a schema as a JSON blob and makes sure that the
// schema is valid before storage.
func (s *Store) StoreSchema(name string, schema *app.Schema) error {
	// make sure before we store this schema that the name matches
	schema.Name = name

	schemaBlob, err := json.Marshal(schema)
	if err != nil {
		return err
	}

	// validate that the schema is valid
	return s.schemaStore.Put(
		ensurePrefix(s.schemaStore.Prefix, name),
		schemaBlob,
		&store.WriteOptions{},
	)
}

// RetrieveSchema gets a schema as a JSON blob
func (s *Store) RetrieveSchema(name string) (*app.Schema, error) {
	pair, err := s.schemaStore.Get(ensurePrefix(s.schemaStore.Prefix, name))

	if err != nil {
		return nil, err
	} else if pair == nil || len(pair.Value) == 0 {
		return nil, ErrMissingKey
	}

	schema := &app.Schema{}
	err = json.Unmarshal(pair.Value, schema)
	if err != nil {
		return nil, err
	}

	return schema, nil
}

// DeleteSchema removes a schema from the K/V tree
func (s *Store) DeleteSchema(name string) error {
	return s.schemaStore.Delete(ensurePrefix(s.schemaStore.Prefix, name))
}
