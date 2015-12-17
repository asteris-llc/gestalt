package validator

import (
	"fmt"
	"github.com/asteris-llc/gestalt/web/app"
)

// Validator validates values
type Validator struct {
	schema *app.Schema
}

// New returns a new Validator
func New(s *app.Schema) *Validator {
	return &Validator{s}
}

// Field returns the named field
func (v *Validator) Field(name string) (*app.Field, error) {
	for _, field := range v.schema.Fields {
		if field.Name == name {
			return field, nil
		}
	}
	return nil, ErrNoField
}

// ValidateField validates a single field against the schema contained therein
func (v *Validator) ValidateField(name string, value interface{}) error {
	field, err := v.Field(name)
	if err != nil {
		return err
	}

	valid := true

	switch field.Type {
	case "string":
		_, valid = value.(string)

	case "integer":
		_, valid = value.(int)

	case "float":
		_, valid = value.(float64)

	case "boolean":
		_, valid = value.(bool)
	}

	if !valid {
		return fmt.Errorf(`"%v" is not a valid %s`, value, field.Type)
	}

	return err
}

// ValidateAll validates all the values in the fields, erroring on any extra
func (v *Validator) ValidateAll(values map[string]interface{}) map[string]error {
	// copy the keys into a new map so we can delete them as we go to detect extra fields
	mValues := map[string]interface{}{}
	for k, v := range values {
		mValues[k] = v
	}

	errors := map[string]error{}

	for _, field := range v.schema.Fields {
		value, ok := mValues[field.Name]
		if ok {
			err := v.ValidateField(field.Name, value)
			if err != nil {
				errors[field.Name] = err
			}
		} else if field.Required {
			errors[field.Name] = ErrFieldRequired
		}

		delete(mValues, field.Name)
	}

	for field := range mValues {
		errors[field] = ErrNoField
	}

	return errors
}
