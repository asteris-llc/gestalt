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

// Schema media type
// Identifier: application/vnd.schema+json
type Schema struct {
	// a registered backend
	Backend string
	// human readable description
	Description string
	Fields      []*Field
	Name        string
	// root for this schema (backend prefix + name if not set)
	Root string
}

// LoadSchema loads raw data into an instance of Schema running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadSchema(raw interface{}) (res *Schema, err error) {
	res, err = UnmarshalSchema(raw, err)
	return
}

// Dump produces raw data from an instance of Schema running all the
// validations. See LoadSchema for the definition of raw data.
func (mt *Schema) Dump() (res map[string]interface{}, err error) {
	res, err = MarshalSchema(mt, err)
	return
}

// Validate validates the media type instance.
func (mt *Schema) Validate() (err error) {
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

// MarshalSchema validates and renders an instance of Schema into a interface{}
// using view "default".
func MarshalSchema(source *Schema, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if source.Name != "" {
		if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, source.Name); !ok {
			err = goa.InvalidPatternError(`.name`, source.Name, `[a-zA-Z0-9\-]+`, err)
		}
	}
	tmp13 := map[string]interface{}{
		"backend":     source.Backend,
		"description": source.Description,
		"name":        source.Name,
		"root":        source.Root,
	}
	if source.Fields != nil {
		tmp14 := make([]map[string]interface{}, len(source.Fields))
		for tmp15, tmp16 := range source.Fields {
			tmp14[tmp15], err = MarshalField(tmp16, err)
		}
		tmp13["fields"] = tmp14
	}
	target = tmp13
	return
}

// UnmarshalSchema unmarshals and validates a raw interface{} into an instance of Schema
func UnmarshalSchema(source interface{}, inErr error) (target *Schema, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Schema)
		if v, ok := val["backend"]; ok {
			var tmp17 string
			if val, ok := v.(string); ok {
				tmp17 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Backend`, v, "string", err)
			}
			target.Backend = tmp17
		}
		if v, ok := val["description"]; ok {
			var tmp18 string
			if val, ok := v.(string); ok {
				tmp18 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Description`, v, "string", err)
			}
			target.Description = tmp18
		}
		if v, ok := val["fields"]; ok {
			var tmp19 []*Field
			if val, ok := v.([]interface{}); ok {
				tmp19 = make([]*Field, len(val))
				for tmp20, v := range val {
					tmp19[tmp20], err = UnmarshalField(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`load.Fields`, v, "array", err)
			}
			target.Fields = tmp19
		}
		if v, ok := val["name"]; ok {
			var tmp21 string
			if val, ok := v.(string); ok {
				tmp21 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			if err == nil {
				if tmp21 != "" {
					if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, tmp21); !ok {
						err = goa.InvalidPatternError(`load.Name`, tmp21, `[a-zA-Z0-9\-]+`, err)
					}
				}
			}
			target.Name = tmp21
		}
		if v, ok := val["root"]; ok {
			var tmp22 string
			if val, ok := v.(string); ok {
				tmp22 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Root`, v, "string", err)
			}
			target.Root = tmp22
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// SchemaCollection media type
// Identifier: application/vnd.schema+json; type=collection
type SchemaCollection []*Schema

// LoadSchemaCollection loads raw data into an instance of SchemaCollection running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadSchemaCollection(raw interface{}) (res SchemaCollection, err error) {
	res, err = UnmarshalSchemaCollection(raw, err)
	return
}

// Dump produces raw data from an instance of SchemaCollection running all the
// validations. See LoadSchemaCollection for the definition of raw data.
func (mt SchemaCollection) Dump() (res []map[string]interface{}, err error) {
	res = make([]map[string]interface{}, len(mt))
	for i, tmp23 := range mt {
		var tmp24 map[string]interface{}
		tmp24, err = MarshalSchema(tmp23, err)
		res[i] = tmp24
	}
	return
}

// Validate validates the media type instance.
func (mt SchemaCollection) Validate() (err error) {
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

// MarshalSchemaCollection validates and renders an instance of SchemaCollection into a interface{}
// using view "default".
func MarshalSchemaCollection(source SchemaCollection, inErr error) (target []map[string]interface{}, err error) {
	err = inErr
	target = make([]map[string]interface{}, len(source))
	for i, res := range source {
		target[i], err = MarshalSchema(res, err)
	}
	return
}

// UnmarshalSchemaCollection unmarshals and validates a raw interface{} into an instance of SchemaCollection
func UnmarshalSchemaCollection(source interface{}, inErr error) (target SchemaCollection, err error) {
	err = inErr
	if val, ok := source.([]interface{}); ok {
		target = make([]*Schema, len(val))
		for tmp25, v := range val {
			target[tmp25], err = UnmarshalSchema(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	return
}
