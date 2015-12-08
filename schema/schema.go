package schema

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"strings"
)

var (
	// ErrNoField indicates that a name could not be found in the spec
	ErrNoField = errors.New("field not found")
)

// Schema does validation and storage on values
type Schema struct {
	raw map[string]interface{}
}

// New returns a new Schema instance from the specified JSON
func New(input []byte) (*Schema, error) {
	raw := map[string]interface{}{}
	err := json.Unmarshal(input, &raw)
	if err != nil {
		return nil, err
	}

	schema := &Schema{raw: raw}

	return schema, schema.validateInput()
}

func (s *Schema) validateInput() error {
	loader := gojsonschema.NewGoLoader(s.raw)
	_, err := gojsonschema.NewSchema(loader)
	return err
}

// BackendName returns the backend name that this schema expects
func (s *Schema) BackendName() (string, bool) {
	backend, ok := s.raw["backend"]

	if ok {
		return backend.(string), ok
	}
	return "", ok
}

// ValidateAll validates a given JSON input for the entire schema
func (s *Schema) ValidateAll(input []byte) (valid bool, errs []error) {
	loader := gojsonschema.NewStringLoader(string(input))
	result, err := gojsonschema.Validate(gojsonschema.NewGoLoader(s.raw), loader)
	if err != nil {
		return false, []error{err}
	}

	errs = []error{}
	for _, err := range result.Errors() {
		errs = append(errs, fmt.Errorf("%v", err))
	}

	return result.Valid(), errs
}

// ValidateField validates a single named path in the JSON
func (s *Schema) ValidateField(name string, input []byte) (valid bool, errs []error) {
	field, err := s.getField(name)
	if err != nil {
		return false, []error{ErrNoField}
	}

	loader := gojsonschema.NewStringLoader(string(input))
	result, err := gojsonschema.Validate(gojsonschema.NewGoLoader(field), loader)
	if err != nil {
		return false, []error{err}
	}

	errs = []error{}
	for _, err := range result.Errors() {
		errs = append(errs, fmt.Errorf("%v", err))
	}

	return result.Valid(), errs
}

func (s *Schema) getField(name string) (map[string]interface{}, error) {
	field := s.raw
	for _, part := range strings.Split(name, "/") {
		if part == "" {
			continue
		}

		rawProps, ok := field["properties"]
		if !ok {
			return nil, ErrNoField
		}

		props, ok := rawProps.(map[string]interface{})
		if !ok {
			return nil, ErrNoField
		}

		property, ok := props[part]
		if !ok {
			return nil, ErrNoField
		}

		field, ok = property.(map[string]interface{})
		if !ok {
			return nil, ErrNoField
		}
	}

	return field, nil
}
