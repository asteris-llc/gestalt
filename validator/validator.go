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
