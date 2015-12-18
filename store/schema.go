package store

import (
	"encoding/json"
	"github.com/asteris-llc/gestalt/web/app"
	"github.com/docker/libkv/store"
)

// StoreSchema stores or updates a schema and makes sure that the schema is
// valid before storage.
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

// RetrieveSchema gets a schema
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

// ListSchemas gets a list of schemas
func (s *Store) ListSchemas() ([]*app.Schema, error) {
	raws, err := s.schemaStore.List(s.schemaStore.Prefix)

	if err != nil {
		return nil, err
	}

	schemas := []*app.Schema{}
	for _, raw := range raws {
		var schema *app.Schema
		err = json.Unmarshal(raw.Value, schema)
		if err != nil {
			return schemas, err
		}

		schemas = append(schemas, schema)
	}

	return schemas, nil
}

// DeleteSchema removes a schema from the K/V tree
func (s *Store) DeleteSchema(name string) error {
	return s.schemaStore.Delete(ensurePrefix(s.schemaStore.Prefix, name))
}
