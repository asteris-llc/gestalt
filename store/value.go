package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/asteris-llc/gestalt/schema"
	"github.com/docker/libkv/store"
	"path"
)

var (
	// ErrFieldRequired is returned when a field is required and receives a
	// request to be deleted.
	ErrFieldRequired = errors.New("field is required")
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
		err = backend.Put(
			ensurePrefix(backend.Prefix, path.Join(app, k)),
			v,
			&store.WriteOptions{},
		)
		if err != nil {
			return []error{err}
		}
	}

	return []error{}
}

// StoreDefaultValues stores the default values for an app
func (s *Store) StoreDefaultValues(app string) error {
	schemaBytes, err := s.RetrieveSchema(app)
	if err != nil {
		return err
	}

	target, err := schema.New(schemaBytes)
	if err != nil {
		return err
	}

	backend, err := s.getBackendForSchema(target)
	if err != nil {
		return err
	}

	for k, v := range target.Defaults() {
		err = backend.Put(
			ensurePrefix(backend.Prefix, path.Join(app, k)),
			[]byte(fmt.Sprintf("%v", v)),
			&store.WriteOptions{},
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// StoreValue stores a single value
func (s *Store) StoreValue(app, key string, jsonValue []byte) []error {
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

	valid, errors := target.ValidateField(key, jsonValue)
	if !valid {
		return errors
	}

	// Unmarshal the byte value before storage (to avoid storing quotes, etc)
	var value interface{}
	err = json.Unmarshal(jsonValue, &value)
	if err != nil {
		return []error{err}
	}
	byteValue := []byte(fmt.Sprintf("%v", value))

	err = backend.Put(
		ensurePrefix(backend.Prefix, path.Join(app, key)),
		byteValue,
		&store.WriteOptions{},
	)
	if err != nil {
		return []error{err}
	}

	return []error{}
}

// DeleteValues deletes all the values
func (s *Store) DeleteValues(app string) error {
	schemaBytes, err := s.RetrieveSchema(app)
	if err != nil {
		return err
	}

	target, err := schema.New(schemaBytes)
	if err != nil {
		return err
	}

	backend, err := s.getBackendForSchema(target)
	if err != nil {
		return err
	}

	return backend.DeleteTree(ensurePrefix(backend.Prefix, app))
}

// DeleteValue deletes a single value or sets it back to the default. If the
// value is required and does not have a default, this method will return an error.
func (s *Store) DeleteValue(app, key string) error {
	schemaBytes, err := s.RetrieveSchema(app)
	if err != nil {
		return err
	}

	target, err := schema.New(schemaBytes)
	if err != nil {
		return err
	}

	backend, err := s.getBackendForSchema(target)
	if err != nil {
		return err
	}

	for _, req := range target.FlatRequired() {
		if key == req {
			return ErrFieldRequired
		}
	}

	if field, ok := target.Defaults()[key]; ok {
		backend.Put(
			ensurePrefix(backend.Prefix, path.Join(app, key)),
			[]byte(fmt.Sprintf("%v", field)),
			&store.WriteOptions{},
		)
	} else {
		backend.Delete(ensurePrefix(backend.Prefix, path.Join(app, key)))
	}

	return nil
}
