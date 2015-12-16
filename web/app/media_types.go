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
	tmp17 := map[string]interface{}{
		"backend":     source.Backend,
		"description": source.Description,
		"name":        source.Name,
		"root":        source.Root,
	}
	if source.Fields != nil {
		tmp18 := make([]map[string]interface{}, len(source.Fields))
		for tmp19, tmp20 := range source.Fields {
			tmp18[tmp19], err = MarshalField(tmp20, err)
		}
		tmp17["fields"] = tmp18
	}
	if source.Values != nil {
		tmp21 := make([]string, len(source.Values))
		for tmp22, tmp23 := range source.Values {
			tmp21[tmp22] = tmp23
		}
		tmp17["values"] = tmp21
	}
	target = tmp17
	return
}

// UnmarshalAsterisGestaltSchema unmarshals and validates a raw interface{} into an instance of AsterisGestaltSchema
func UnmarshalAsterisGestaltSchema(source interface{}, inErr error) (target *AsterisGestaltSchema, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(AsterisGestaltSchema)
		if v, ok := val["backend"]; ok {
			var tmp24 string
			if val, ok := v.(string); ok {
				tmp24 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Backend`, v, "string", err)
			}
			target.Backend = tmp24
		}
		if v, ok := val["description"]; ok {
			var tmp25 string
			if val, ok := v.(string); ok {
				tmp25 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Description`, v, "string", err)
			}
			target.Description = tmp25
		}
		if v, ok := val["fields"]; ok {
			var tmp26 []*Field
			if val, ok := v.([]interface{}); ok {
				tmp26 = make([]*Field, len(val))
				for tmp27, v := range val {
					tmp26[tmp27], err = UnmarshalField(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`load.Fields`, v, "array", err)
			}
			target.Fields = tmp26
		}
		if v, ok := val["name"]; ok {
			var tmp28 string
			if val, ok := v.(string); ok {
				tmp28 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			if err == nil {
				if tmp28 != "" {
					if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, tmp28); !ok {
						err = goa.InvalidPatternError(`load.Name`, tmp28, `[a-zA-Z0-9\-]+`, err)
					}
				}
			}
			target.Name = tmp28
		}
		if v, ok := val["root"]; ok {
			var tmp29 string
			if val, ok := v.(string); ok {
				tmp29 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Root`, v, "string", err)
			}
			target.Root = tmp29
		}
		if v, ok := val["values"]; ok {
			var tmp30 []string
			if val, ok := v.([]interface{}); ok {
				tmp30 = make([]string, len(val))
				for tmp31, v := range val {
					if val, ok := v.(string); ok {
						tmp30[tmp31] = val
					} else {
						err = goa.InvalidAttributeTypeError(`load.Values[*]`, v, "string", err)
					}
				}
			} else {
				err = goa.InvalidAttributeTypeError(`load.Values`, v, "array", err)
			}
			target.Values = tmp30
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// AsterisGestaltSchemaCollection media type
// Identifier: application/vnd.asteris.gestalt.schema+json; type=collection
type AsterisGestaltSchemaCollection []*AsterisGestaltSchema

// LoadAsterisGestaltSchemaCollection loads raw data into an instance of AsterisGestaltSchemaCollection running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadAsterisGestaltSchemaCollection(raw interface{}) (res AsterisGestaltSchemaCollection, err error) {
	res, err = UnmarshalAsterisGestaltSchemaCollection(raw, err)
	return
}

// Dump produces raw data from an instance of AsterisGestaltSchemaCollection running all the
// validations. See LoadAsterisGestaltSchemaCollection for the definition of raw data.
func (mt AsterisGestaltSchemaCollection) Dump() (res []map[string]interface{}, err error) {
	res = make([]map[string]interface{}, len(mt))
	for i, tmp32 := range mt {
		var tmp33 map[string]interface{}
		tmp33, err = MarshalAsterisGestaltSchema(tmp32, err)
		res[i] = tmp33
	}
	return
}

// Validate validates the media type instance.
func (mt AsterisGestaltSchemaCollection) Validate() (err error) {
	for _, e := range mt {
		for _, e := range e.Fields {
			if e.Name != "" {
				if ok := goa.ValidatePattern(`[a-zA-Z0-9\-/]+`, e.Name); !ok {
					err = goa.InvalidPatternError(`response[*].fields[*].name`, e.Name, `[a-zA-Z0-9\-/]+`, err)
				}
			}
			if e.Type != "" {
				if !(e.Type == "string" || e.Type == "integer" || e.Type == "float" || e.Type == "boolean") {
					err = goa.InvalidEnumValueError(`response[*].fields[*].type`, e.Type, []interface{}{"string", "integer", "float", "boolean"}, err)
				}
			}
		}
		if e.Name != "" {
			if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, e.Name); !ok {
				err = goa.InvalidPatternError(`response[*].name`, e.Name, `[a-zA-Z0-9\-]+`, err)
			}
		}
	}
	return
}

// MarshalAsterisGestaltSchemaCollection validates and renders an instance of AsterisGestaltSchemaCollection into a interface{}
// using view "default".
func MarshalAsterisGestaltSchemaCollection(source AsterisGestaltSchemaCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalAsterisGestaltSchema(res, err)
	}
	return
}

// UnmarshalAsterisGestaltSchemaCollection unmarshals and validates a raw interface{} into an instance of AsterisGestaltSchemaCollection
func UnmarshalAsterisGestaltSchemaCollection(source interface{}, inErr error) (target AsterisGestaltSchemaCollection, err error) {
	err = inErr
	if val, ok := source.([]interface{}); ok {
		target = make([]*AsterisGestaltSchema, len(val))
		for tmp34, v := range val {
			target[tmp34], err = UnmarshalAsterisGestaltSchema(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	return
}
