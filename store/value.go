package store

import (
	"github.com/asteris-llc/gestalt/schema"
	"github.com/docker/libkv/store"
	"path"
)

// retrieve all the values

// retrieve one value

// StoreValues stores all the values specified
func (s *Store) StoreValues(app string, body []byte) []error {
	schemaBytes, err := s.RetrieveSchema(app)
	if err != nil {
		return []error{err}
	}

	target, err := schema.New(schemaBytes)
	if err != nil {
		return []error{err}
	}

	backend, err := s.getBackendForSchema(target)
	if err != nil {
		return []error{err}
	}

	valid, errors := target.ValidateAll(body)
	if !valid {
		return errors
	}

	kvs, err := flattenJSONForWriting(body)
	if err != nil {
		return []error{err}
	}

	for k, v := range kvs {
		err = backend.Put(ensurePrefix(backend.Prefix, path.Join(app, k)), v, &store.WriteOptions{})
		if err != nil {
			return []error{err}
		}
	}

	return []error{}
}

// store the default values

// store one value

// store a single default value

// delete a value (enforcing types, setting default if necessary)

// delete all the values
