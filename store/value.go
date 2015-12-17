package store

import (
	"github.com/asteris-llc/gestalt/validator"
	"github.com/asteris-llc/gestalt/web/app"
	"github.com/docker/libkv/store"
	"reflect"
)

// setup is a convienence method for getting a valid schema and backend
func (s *Store) setup(schemaName string) (*app.Schema, *Backend, error) {
	schema, err := s.RetrieveSchema(schemaName)
	if err != nil {
		return nil, nil, err
	}

	backend, err := s.getBackendForSchema(schema)
	if err != nil {
		return nil, nil, err
	}

	return schema, backend, nil
}

// RetrieveValues gets all the values from the backend in a map
func (s *Store) RetrieveValues(schemaName string) (map[string]interface{}, error) {
	out := map[string]interface{}{}

	schema, backend, err := s.setup(schemaName)
	if err != nil {
		return out, err
	}

	for _, field := range schema.Fields {
		value, err := backend.Get(backend.FieldKey(schema, field))
		if err != nil {
			return out, err
		} else if value == nil || len(value.Value) == 0 {
			continue
		}

		decoded, err := unmarshal(value.Value, field.Type)
		if err != nil {
			return out, &DecodeError{field.Name, err}
		}

		if decoded != nil {
			out[field.Name] = decoded
		}
	}

	return out, nil
}

// RetrieveValue retrieves a single designated value
func (s *Store) RetrieveValue(schemaName, fieldName string) (interface{}, error) {
	schema, backend, err := s.setup(schemaName)
	if err != nil {
		return nil, err
	}

	v := validator.New(schema)

	field, err := v.Field(fieldName)
	if err == validator.ErrNoField {
		return nil, ErrMissingField
	} else if err != nil {
		return nil, err
	}

	raw, err := backend.Get(backend.FieldKey(schema, field))
	if err != nil {
		return nil, err
	} else if raw == nil || len(raw.Value) == 0 {
		return nil, ErrMissingKey
	}

	out, err := unmarshal(raw.Value, field.Type)
	if err != nil {
		return nil, &DecodeError{field.Name, err}
	}

	return out, nil
}

// retrieve one value

// StoreValues stores all the values specified
func (s *Store) StoreValues(schemaName string, values map[string]interface{}) []error {
	schema, backend, err := s.setup(schemaName)
	if err != nil {
		return []error{err}
	}

	v := validator.New(schema)

	errors := v.ValidateAll(values)
	if len(errors) != 0 {
		outErrors := []error{}
		for field, error := range errors {
			outErrors = append(outErrors, &DecodeError{field, error})
		}
		return outErrors
	}

	for fieldName, value := range values {
		field, err := v.Field(fieldName)
		if err != nil {
			return []error{err}
		}

		err = backend.Put(
			backend.FieldKey(schema, field),
			marshal(value),
			&store.WriteOptions{},
		)
		if err != nil {
			return []error{err}
		}
	}

	return []error{}
}

// StoreDefaultValues stores the default values for an app
func (s *Store) StoreDefaultValues(schemaName string) error {
	schema, backend, err := s.setup(schemaName)
	if err != nil {
		return err
	}

	for _, field := range schema.Fields {
		if reflect.ValueOf(field.Default).IsValid() {
			err = backend.Put(
				backend.FieldKey(schema, field),
				marshal(field.Default),
				&store.WriteOptions{},
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// StoreValue stores a single value
func (s *Store) StoreValue(schemaName, fieldName string, value interface{}) error {
	schema, backend, err := s.setup(schemaName)
	if err != nil {
		return err
	}

	v := validator.New(schema)
	field, err := v.Field(fieldName)
	if err != nil {
		return err
	}

	err = v.ValidateField(field.Name, value)
	if err != nil {
		return err
	}

	return backend.Put(
		backend.FieldKey(schema, field),
		marshal(value),
		&store.WriteOptions{},
	)
}

// DeleteValues deletes all the values
func (s *Store) DeleteValues(schemaName string) error {
	schema, backend, err := s.setup(schemaName)
	if err != nil {
		return err
	}

	return backend.DeleteTree(backend.SchemaKey(schema))
}

// DeleteValue deletes a single value or sets it back to the default. If the
// value is required and does not have a default, this method will return an error.
func (s *Store) DeleteValue(schemaName, fieldName string) error {
	schema, backend, err := s.setup(schemaName)
	if err != nil {
		return err
	}

	v := validator.New(schema)

	field, err := v.Field(fieldName)
	if err != nil {
		return err
	}

	if field.Required {
		return ErrFieldRequired
	}

	if reflect.ValueOf(field.Default).IsValid() {
		backend.Put(
			backend.FieldKey(schema, field),
			marshal(field.Default),
			&store.WriteOptions{},
		)
	} else {
		backend.Delete(backend.FieldKey(schema, field))
	}

	return nil
}
