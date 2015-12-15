//************************************************************************//
// gestalt: Application Media Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/asteris-llc/gestalt/web
// --design=github.com/asteris-llc/gestalt/web/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/raphael/goa"

// AsterisGestaltSchema media type
// Identifier: application/vnd.asteris.gestalt.schema+json
type AsterisGestaltSchema struct {
	// a registered backend
	Backend string
	// human readable description
	Description string
	Fields      []*Field
	Name        string
	// root for this schema (backend prefix + name if not set)
	Root string
	// links to values
	Values []string
}

// LoadAsterisGestaltSchema loads raw data into an instance of AsterisGestaltSchema running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadAsterisGestaltSchema(raw interface{}) (res *AsterisGestaltSchema, err error) {
	res, err = UnmarshalAsterisGestaltSchema(raw, err)
	return
}

// Dump produces raw data from an instance of AsterisGestaltSchema running all the
// validations. See LoadAsterisGestaltSchema for the definition of raw data.
func (mt *AsterisGestaltSchema) Dump() (res map[string]interface{}, err error) {
	res, err = MarshalAsterisGestaltSchema(mt, err)
	return
}

// Validate validates the media type instance.
func (mt *AsterisGestaltSchema) Validate() (err error) {
	for _, e := range mt.Fields {
		if e.Name != "" {
			if ok := goa.ValidatePattern(`[a-zA-Z0-9\-/]+`, e.Name); !ok {
				err = goa.InvalidPatternError(`response.fields[*].name`, e.Name, `[a-zA-Z0-9\-/]+`, err)
			}
		}
		if e.Type != "" {
			if !(e.Type == "string" || e.Type == "integer" || e.Type == "float" || e.Type == "boolean") {
				err = goa.InvalidEnumValueError(`response.fields[*].type`, e.Type, []interface{}{"string", "integer", "float", "boolean"}, err)
			}
		}
	}
	if mt.Name != "" {
		if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, mt.Name); !ok {
			err = goa.InvalidPatternError(`response.name`, mt.Name, `[a-zA-Z0-9\-]+`, err)
		}
	}
	return
}

// MarshalAsterisGestaltSchema validates and renders an instance of AsterisGestaltSchema into a interface{}
// using view "default".
func MarshalAsterisGestaltSchema(source *AsterisGestaltSchema, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if source.Name != "" {
		if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, source.Name); !ok {
			err = goa.InvalidPatternError(`.name`, source.Name, `[a-zA-Z0-9\-]+`, err)
		}
	}
	tmp27 := map[string]interface{}{
		"backend":     source.Backend,
		"description": source.Description,
		"name":        source.Name,
		"root":        source.Root,
	}
	if source.Fields != nil {
		tmp28 := make([]map[string]interface{}, len(source.Fields))
		for tmp29, tmp30 := range source.Fields {
			tmp28[tmp29], err = MarshalField(tmp30, err)
		}
		tmp27["fields"] = tmp28
	}
	if source.Values != nil {
		tmp31 := make([]string, len(source.Values))
		for tmp32, tmp33 := range source.Values {
			tmp31[tmp32] = tmp33
		}
		tmp27["values"] = tmp31
	}
	target = tmp27
	return
}

// UnmarshalAsterisGestaltSchema unmarshals and validates a raw interface{} into an instance of AsterisGestaltSchema
func UnmarshalAsterisGestaltSchema(source interface{}, inErr error) (target *AsterisGestaltSchema, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(AsterisGestaltSchema)
		if v, ok := val["backend"]; ok {
			var tmp34 string
			if val, ok := v.(string); ok {
				tmp34 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Backend`, v, "string", err)
			}
			target.Backend = tmp34
		}
		if v, ok := val["description"]; ok {
			var tmp35 string
			if val, ok := v.(string); ok {
				tmp35 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Description`, v, "string", err)
			}
			target.Description = tmp35
		}
		if v, ok := val["fields"]; ok {
			var tmp36 []*Field
			if val, ok := v.([]interface{}); ok {
				tmp36 = make([]*Field, len(val))
				for tmp37, v := range val {
					tmp36[tmp37], err = UnmarshalField(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`load.Fields`, v, "array", err)
			}
			target.Fields = tmp36
		}
		if v, ok := val["name"]; ok {
			var tmp38 string
			if val, ok := v.(string); ok {
				tmp38 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			if err == nil {
				if tmp38 != "" {
					if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, tmp38); !ok {
						err = goa.InvalidPatternError(`load.Name`, tmp38, `[a-zA-Z0-9\-]+`, err)
					}
				}
			}
			target.Name = tmp38
		}
		if v, ok := val["root"]; ok {
			var tmp39 string
			if val, ok := v.(string); ok {
				tmp39 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Root`, v, "string", err)
			}
			target.Root = tmp39
		}
		if v, ok := val["values"]; ok {
			var tmp40 []string
			if val, ok := v.([]interface{}); ok {
				tmp40 = make([]string, len(val))
				for tmp41, v := range val {
					if val, ok := v.(string); ok {
						tmp40[tmp41] = val
					} else {
						err = goa.InvalidAttributeTypeError(`load.Values[*]`, v, "string", err)
					}
				}
			} else {
				err = goa.InvalidAttributeTypeError(`load.Values`, v, "array", err)
			}
			target.Values = tmp40
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// AsterisGestaltSchemas media type
// Identifier: application/vnd.asteris.gestalt.schemas+json
type AsterisGestaltSchemas struct {
	// list of schemas
	Schemas []*Schema
}

// LoadAsterisGestaltSchemas loads raw data into an instance of AsterisGestaltSchemas running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadAsterisGestaltSchemas(raw interface{}) (res *AsterisGestaltSchemas, err error) {
	res, err = UnmarshalAsterisGestaltSchemas(raw, err)
	return
}

// Dump produces raw data from an instance of AsterisGestaltSchemas running all the
// validations. See LoadAsterisGestaltSchemas for the definition of raw data.
func (mt *AsterisGestaltSchemas) Dump() (res map[string]interface{}, err error) {
	res, err = MarshalAsterisGestaltSchemas(mt, err)
	return
}

// Validate validates the media type instance.
func (mt *AsterisGestaltSchemas) Validate() (err error) {
	for _, e := range mt.Schemas {
		for _, e := range e.Fields {
			if e.Name != "" {
				if ok := goa.ValidatePattern(`[a-zA-Z0-9\-/]+`, e.Name); !ok {
					err = goa.InvalidPatternError(`response.schemas[*].fields[*].name`, e.Name, `[a-zA-Z0-9\-/]+`, err)
				}
			}
			if e.Type != "" {
				if !(e.Type == "string" || e.Type == "integer" || e.Type == "float" || e.Type == "boolean") {
					err = goa.InvalidEnumValueError(`response.schemas[*].fields[*].type`, e.Type, []interface{}{"string", "integer", "float", "boolean"}, err)
				}
			}
		}
		if e.Name != "" {
			if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, e.Name); !ok {
				err = goa.InvalidPatternError(`response.schemas[*].name`, e.Name, `[a-zA-Z0-9\-]+`, err)
			}
		}
	}
	return
}

// MarshalAsterisGestaltSchemas validates and renders an instance of AsterisGestaltSchemas into a interface{}
// using view "default".
func MarshalAsterisGestaltSchemas(source *AsterisGestaltSchemas, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	tmp42 := map[string]interface{}{}
	if source.Schemas != nil {
		tmp43 := make([]map[string]interface{}, len(source.Schemas))
		for tmp44, tmp45 := range source.Schemas {
			tmp43[tmp44], err = MarshalSchema(tmp45, err)
		}
		tmp42["schemas"] = tmp43
	}
	target = tmp42
	return
}

// UnmarshalAsterisGestaltSchemas unmarshals and validates a raw interface{} into an instance of AsterisGestaltSchemas
func UnmarshalAsterisGestaltSchemas(source interface{}, inErr error) (target *AsterisGestaltSchemas, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(AsterisGestaltSchemas)
		if v, ok := val["schemas"]; ok {
			var tmp46 []*Schema
			if val, ok := v.([]interface{}); ok {
				tmp46 = make([]*Schema, len(val))
				for tmp47, v := range val {
					tmp46[tmp47], err = UnmarshalSchema(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`load.Schemas`, v, "array", err)
			}
			target.Schemas = tmp46
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}
