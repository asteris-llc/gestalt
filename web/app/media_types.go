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

// Field media type
// Identifier: application/vnd.asteris.gestalt.field+json
type Field struct {
	// the default for this field
	Default interface{}
	// human readable description
	Description string
	Name        string
	// this field is required
	Required bool
	// root for this key (backend prefix + schema name if not set)
	Root string
	// type of value expected
	Type string
}

// LoadField loads raw data into an instance of Field running all the
// validations. Raw data is defined by data that the JSON unmarshaler would create when unmarshaling
// into a variable of type interface{}. See https://golang.org/pkg/encoding/json/#Unmarshal for the
// complete list of supported data types.
func LoadField(raw interface{}) (res *Field, err error) {
	res, err = UnmarshalField(raw, err)
	return
}

// Dump produces raw data from an instance of Field running all the
// validations. See LoadField for the definition of raw data.
func (mt *Field) Dump() (res map[string]interface{}, err error) {
	res, err = MarshalField(mt, err)
	return
}

// Validate validates the media type instance.
func (mt *Field) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name", err)
	}
	if mt.Type == "" {
		err = goa.MissingAttributeError(`response`, "type", err)
	}

	if ok := goa.ValidatePattern(`[a-zA-Z0-9\-/]+`, mt.Name); !ok {
		err = goa.InvalidPatternError(`response.name`, mt.Name, `[a-zA-Z0-9\-/]+`, err)
	}
	if !(mt.Type == "string" || mt.Type == "integer" || mt.Type == "float" || mt.Type == "boolean") {
		err = goa.InvalidEnumValueError(`response.type`, mt.Type, []interface{}{"string", "integer", "float", "boolean"}, err)
	}
	return
}

// MarshalField validates and renders an instance of Field into a interface{}
// using view "default".
func MarshalField(source *Field, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if source.Name == "" {
		err = goa.MissingAttributeError(``, "name", err)
	}

	if err == nil {
		if source.Name == "" {
			err = goa.MissingAttributeError(``, "name", err)
		}
		if err == nil {
			if ok := goa.ValidatePattern(`[a-zA-Z0-9\-/]+`, source.Name); !ok {
				err = goa.InvalidPatternError(`.name`, source.Name, `[a-zA-Z0-9\-/]+`, err)
			}
			if source.Type != "" {
				if !(source.Type == "string" || source.Type == "integer" || source.Type == "float" || source.Type == "boolean") {
					err = goa.InvalidEnumValueError(`.type`, source.Type, []interface{}{"string", "integer", "float", "boolean"}, err)
				}
			}
			tmp13 := map[string]interface{}{
				"default":     source.Default,
				"description": source.Description,
				"name":        source.Name,
				"required":    source.Required,
				"root":        source.Root,
				"type":        source.Type,
			}
			target = tmp13
		}
	}
	return
}

// UnmarshalField unmarshals and validates a raw interface{} into an instance of Field
func UnmarshalField(source interface{}, inErr error) (target *Field, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Field)
		if v, ok := val["default"]; ok {
			var tmp14 interface{}
			tmp14 = v
			target.Default = tmp14
		}
		if v, ok := val["description"]; ok {
			var tmp15 string
			if val, ok := v.(string); ok {
				tmp15 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Description`, v, "string", err)
			}
			target.Description = tmp15
		}
		if v, ok := val["name"]; ok {
			var tmp16 string
			if val, ok := v.(string); ok {
				tmp16 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			if err == nil {
				if tmp16 != "" {
					if ok := goa.ValidatePattern(`[a-zA-Z0-9\-/]+`, tmp16); !ok {
						err = goa.InvalidPatternError(`load.Name`, tmp16, `[a-zA-Z0-9\-/]+`, err)
					}
				}
			}
			target.Name = tmp16
		} else {
			err = goa.MissingAttributeError(`load`, "name", err)
		}
		if v, ok := val["required"]; ok {
			var tmp17 bool
			if val, ok := v.(bool); ok {
				tmp17 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Required`, v, "bool", err)
			}
			target.Required = tmp17
		}
		if v, ok := val["root"]; ok {
			var tmp18 string
			if val, ok := v.(string); ok {
				tmp18 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Root`, v, "string", err)
			}
			target.Root = tmp18
		}
		if v, ok := val["type"]; ok {
			var tmp19 string
			if val, ok := v.(string); ok {
				tmp19 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Type`, v, "string", err)
			}
			if err == nil {
				if tmp19 != "" {
					if !(tmp19 == "string" || tmp19 == "integer" || tmp19 == "float" || tmp19 == "boolean") {
						err = goa.InvalidEnumValueError(`load.Type`, tmp19, []interface{}{"string", "integer", "float", "boolean"}, err)
					}
				}
			}
			target.Type = tmp19
		} else {
			err = goa.MissingAttributeError(`load`, "type", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// Schema media type
// Identifier: application/vnd.asteris.gestalt.schema+json
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
	if mt.Backend == "" {
		err = goa.MissingAttributeError(`response`, "backend", err)
	}
	if mt.Name == "" {
		err = goa.MissingAttributeError(`response`, "name", err)
	}
	if mt.Fields == nil {
		err = goa.MissingAttributeError(`response`, "fields", err)
	}

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
	if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, mt.Name); !ok {
		err = goa.InvalidPatternError(`response.name`, mt.Name, `[a-zA-Z0-9\-]+`, err)
	}
	return
}

// MarshalSchema validates and renders an instance of Schema into a interface{}
// using view "default".
func MarshalSchema(source *Schema, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if source.Backend == "" {
		err = goa.MissingAttributeError(``, "backend", err)
	}

	if err == nil {
		if source.Backend == "" {
			err = goa.MissingAttributeError(``, "backend", err)
		}
		if err == nil {
			if source.Name != "" {
				if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, source.Name); !ok {
					err = goa.InvalidPatternError(`.name`, source.Name, `[a-zA-Z0-9\-]+`, err)
				}
			}
			tmp20 := map[string]interface{}{
				"backend":     source.Backend,
				"description": source.Description,
				"name":        source.Name,
				"root":        source.Root,
			}
			if source.Fields != nil {
				tmp21 := make([]map[string]interface{}, len(source.Fields))
				for tmp22, tmp23 := range source.Fields {
					tmp21[tmp22], err = MarshalField(tmp23, err)
				}
				tmp20["fields"] = tmp21
			}
			target = tmp20
		}
	}
	return
}

// UnmarshalSchema unmarshals and validates a raw interface{} into an instance of Schema
func UnmarshalSchema(source interface{}, inErr error) (target *Schema, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(Schema)
		if v, ok := val["backend"]; ok {
			var tmp24 string
			if val, ok := v.(string); ok {
				tmp24 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Backend`, v, "string", err)
			}
			target.Backend = tmp24
		} else {
			err = goa.MissingAttributeError(`load`, "backend", err)
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
		} else {
			err = goa.MissingAttributeError(`load`, "fields", err)
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
		} else {
			err = goa.MissingAttributeError(`load`, "name", err)
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
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// SchemaCollection media type
// Identifier: application/vnd.asteris.gestalt.schema+json; type=collection
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
	for i, tmp30 := range mt {
		var tmp31 map[string]interface{}
		tmp31, err = MarshalSchema(tmp30, err)
		res[i] = tmp31
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
		for tmp32, v := range val {
			target[tmp32], err = UnmarshalSchema(v, err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "array", err)
	}
	return
}
