package schema

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
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
	field, ok := s.fields()[name]
	if !ok {
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

// Defaults retrieves the defaults from the schema in a map
func (s *Schema) Defaults() map[string]string {
	defaults := map[string]string{}

	for key, field := range s.fields() {
		def, ok := field["default"]
		if ok {
			defaults[key] = fmt.Sprintf("%v", def)
		}
	}

	return defaults
}

func (s *Schema) fields() map[string]map[string]interface{} {
	fields := map[string]map[string]interface{}{}

	type Item struct {
		path  string
		value map[string]interface{}
	}
	queue := []Item{}

	props, ok := s.raw["properties"]
	if !ok {
		return fields
	}

	for k, v := range props.(map[string]interface{}) {
		queue = append(queue, Item{k, v.(map[string]interface{})})
	}

	for len(queue) != 0 {
		item := queue[0]
		queue = queue[1:]

		fields[item.path] = item.value

		props, ok = item.value["properties"]
		if ok {
			for k, v := range props.(map[string]interface{}) {
				queue = append(
					queue,
					Item{
						item.path + "/" + k,
						v.(map[string]interface{})},
				)
			}
		}
	}

	return fields
}
