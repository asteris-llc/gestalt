package schema

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"path"
)

var (
	// ErrNoField indicates that a name could not be found in the spec
	ErrNoField = errors.New("field not found")
)

// Schema does validation and storage on values
type Schema struct {
	Field
	Required []string `json:"required,omitempty"`

	// extended values for gestalt
	Backend string `json:"backend,omitempty"`

	loader gojsonschema.JSONLoader
}

// New returns a new Schema instance from the specified JSON
func New(raw []byte) (*Schema, error) {
	schema := new(Schema)
	err := json.Unmarshal(raw, schema)
	if err != nil {
		return nil, err
	}

	mapped := map[string]interface{}{}
	err = json.Unmarshal(raw, &mapped)
	if err != nil {
		return nil, err
	}

	schema.loader = gojsonschema.NewGoLoader(mapped)

	return schema, schema.validateInput()
}

func (s *Schema) validateInput() error {
	_, err := gojsonschema.NewSchema(s.loader)
	return err
}

// ValidateAll validates a given JSON input for the entire schema
func (s *Schema) ValidateAll(input []byte) (valid bool, errs []error) {
	loader := gojsonschema.NewStringLoader(string(input))
	result, err := gojsonschema.Validate(s.loader, loader)
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
	field, ok := s.Fields()[name]
	if !ok {
		return false, []error{ErrNoField}
	}

	loader := gojsonschema.NewStringLoader(string(input))
	result, err := gojsonschema.Validate(gojsonschema.NewGoLoader(field.Map()), loader)
	if err != nil {
		return false, []error{err}
	}

	errs = []error{}
	for _, err := range result.Errors() {
		errs = append(errs, fmt.Errorf("%v", err))
	}

	return result.Valid(), errs
}

// Defaults retrieves the defaults from the schema in a map
func (s *Schema) Defaults() map[string]interface{} {
	defaults := map[string]interface{}{}

	for key, field := range s.Fields() {
		if field.Default != nil {
			defaults[key] = *field.Default
		}
	}

	return defaults
}

// Fields returns a flattened list of fields
func (s *Schema) Fields() map[string]*Field {
	fields := map[string]*Field{}

	type Item struct {
		path  string
		value *Field
	}
	queue := []Item{}

	for name := range s.Properties {
		field := s.Properties[name]
		queue = append(queue, Item{name, &field})
	}

	for len(queue) != 0 {
		item := queue[0]
		queue = queue[1:]

		fields[item.path] = item.value

		for name := range item.value.Properties {
			field := item.value.Properties[name]
			queue = append(queue, Item{path.Join(item.path, name), &field})
		}
	}

	return fields
}

// FlatRequired returns a list of fields that are required
func (s *Schema) FlatRequired() []string {
	required := map[string]bool{}

	for _, name := range s.Required {
		required[name] = true
	}

	for name, field := range s.Fields() {
		if field.Required != nil && *field.Required {
			required[name] = true
		}
	}

	out := []string{}
	for field := range required {
		out = append(out, field)
	}
	return out
}
