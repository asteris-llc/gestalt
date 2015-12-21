package store

import (
	"encoding/json"
	"github.com/asteris-llc/gestalt/web/app"
	"github.com/docker/libkv/store"
	"path"
)

func (s *Store) schemaPath(name string) string {
	return ensurePrefix(s.schemaStore.Prefix, path.Join("schemas", name))
}

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
		s.schemaPath(name),
		schemaBlob,
		&store.WriteOptions{},
	)
}

// RetrieveSchema gets a schema
func (s *Store) RetrieveSchema(name string) (*app.Schema, error) {
	pair, err := s.schemaStore.Get(s.schemaPath(name))

	if err == store.ErrKeyNotFound {
		return nil, ErrMissingKey
	} else if err != nil {
		return nil, err
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
	schemas := []*app.Schema{}
	raws, err := s.schemaStore.List(ensurePrefix(s.schemaStore.Prefix, "schemas"))

	if err == store.ErrKeyNotFound {
		return schemas, nil
	} else if err != nil {
		return nil, err
	}

	for _, raw := range raws {
		schema := &app.Schema{}
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
	return s.schemaStore.Delete(s.schemaPath(name))
}
