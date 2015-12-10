package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/asteris-llc/gestalt/schema"
	"github.com/docker/libkv/store"
	"path"
	"strconv"
)

var (
	// ErrFieldRequired is returned when a field is required and receives a
	// request to be deleted.
	ErrFieldRequired = errors.New("field is required")

	// ErrMissingKey is returned when a key is missing
	ErrMissingKey = errors.New("no such key") // TODO: consistent missing key across project

	// ErrMissingField is returned when a field is missing
	ErrMissingField = errors.New("no such field") // TODO: probably move to schema?
)

// DecodeError is returned for errors in decoding values
type DecodeError struct {
	Field string
	Err   error
}

func (d *DecodeError) Error() string {
	return fmt.Sprintf("%s: %s", d.Field, d.Err)
}

// RetrieveValues gets all the values from the backend in a map
func (s *Store) RetrieveValues(app string) (map[string]interface{}, error) {
	out := map[string]interface{}{}

	schemaBytes, err := s.RetrieveSchema(app)
	if err != nil {
		return out, err
	}

	target, err := schema.New(schemaBytes)
	if err != nil {
		return out, err
	}

	backend, err := s.getBackendForSchema(target)
	if err != nil {
		return out, err
	}

	//  object, array, null, any
	for name, field := range target.Fields() {
		value, err := backend.Get(ensurePrefix(backend.Prefix, path.Join(app, name)))

		if err != nil {
			return out, err
		} else if value == nil || len(value.Value) == 0 {
			continue
		}

		decoded, err := s.decodeValue(value.Value, field.Type)
		if err != nil {
			return out, &DecodeError{name, err}
		}

		if decoded != nil {
			out[name] = decoded
		}
	}

	return out, nil
}

// RetrieveValue retrieves a single designated value
func (s *Store) RetrieveValue(app, key string) (interface{}, error) {
	schemaBytes, err := s.RetrieveSchema(app)
	if err != nil {
		return nil, err
	}

	target, err := schema.New(schemaBytes)
	if err != nil {
		return nil, err
	}

	backend, err := s.getBackendForSchema(target)
	if err != nil {
		return nil, err
	}

	field, ok := target.Fields()[key]
	if !ok {
		return nil, ErrMissingField
	}

	raw, err := backend.Get(ensurePrefix(backend.Prefix, path.Join(app, key)))
	if err != nil {
		return nil, err
	} else if raw == nil || len(raw.Value) == 0 {
		return nil, ErrMissingKey
	}

	// TODO: remember to check for strconv.NumError in the HTTP parts and convert
	// it to a more friendly error type
	out, err := s.decodeValue(raw.Value, field.Type)
	if err != nil {
		return nil, &DecodeError{key, err}
	}

	return out, nil
}

func (s *Store) decodeValue(value []byte, typ string) (interface{}, error) {
	stringed := string(value)

	switch typ {
	case "string", "any":
		return stringed, nil

	case "number":
		return strconv.ParseFloat(stringed, 64)

	case "integer":
		return strconv.Atoi(stringed)

	case "boolean":
		return strconv.ParseBool(stringed)

		// unsupported / ignored types
	case "object", "array", "null":
		return nil, nil

	default:
		return nil, fmt.Errorf(`don't know how to decode "%s"`, typ)
	}
}

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
