//************************************************************************//
// gestalt: Application User Types
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

import (
	"github.com/raphael/goa"
)

// SchemaPayload type
type SchemaPayload struct {
	// a registered backend
	Backend string
	// human readable description
	Description string
	Fields      []*Field
	Name        string
	// root for this schema (backend prefix + name if not set)
	Root string
}

// MarshalSchemaPayload validates and renders an instance of SchemaPayload into a interface{}
func MarshalSchemaPayload(source *SchemaPayload, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if source.Backend == "" {
		err = goa.MissingAttributeError(``, "backend", err)
	}
	if source.Name == "" {
		err = goa.MissingAttributeError(``, "name", err)
	}
	if source.Fields == nil {
		err = goa.MissingAttributeError(``, "fields", err)
	}

	if err == nil {
		if source.Backend == "" {
			err = goa.MissingAttributeError(``, "backend", err)
		}
		if source.Name == "" {
			err = goa.MissingAttributeError(``, "name", err)
		}
		if source.Fields == nil {
			err = goa.MissingAttributeError(``, "fields", err)
		}
		if err == nil {
			if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, source.Name); !ok {
				err = goa.InvalidPatternError(`.name`, source.Name, `[a-zA-Z0-9\-]+`, err)
			}
			tmp26 := map[string]interface{}{
				"backend":     source.Backend,
				"description": source.Description,
				"name":        source.Name,
				"root":        source.Root,
			}
			if source.Fields != nil {
				tmp27 := make([]map[string]interface{}, len(source.Fields))
				for tmp28, tmp29 := range source.Fields {
					tmp27[tmp28], err = MarshalField(tmp29, err)
				}
				tmp26["fields"] = tmp27
			}
			target = tmp26
		}
	}
	return
}

// UnmarshalSchemaPayload unmarshals and validates a raw interface{} into an instance of SchemaPayload
func UnmarshalSchemaPayload(source interface{}, inErr error) (target *SchemaPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(SchemaPayload)
		if v, ok := val["backend"]; ok {
			var tmp30 string
			if val, ok := v.(string); ok {
				tmp30 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Backend`, v, "string", err)
			}
			target.Backend = tmp30
		} else {
			err = goa.MissingAttributeError(`load`, "backend", err)
		}
		if v, ok := val["description"]; ok {
			var tmp31 string
			if val, ok := v.(string); ok {
				tmp31 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Description`, v, "string", err)
			}
			target.Description = tmp31
		}
		if v, ok := val["fields"]; ok {
			var tmp32 []*Field
			if val, ok := v.([]interface{}); ok {
				tmp32 = make([]*Field, len(val))
				for tmp33, v := range val {
					tmp32[tmp33], err = UnmarshalField(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`load.Fields`, v, "array", err)
			}
			target.Fields = tmp32
		} else {
			err = goa.MissingAttributeError(`load`, "fields", err)
		}
		if v, ok := val["name"]; ok {
			var tmp34 string
			if val, ok := v.(string); ok {
				tmp34 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			if err == nil {
				if tmp34 != "" {
					if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, tmp34); !ok {
						err = goa.InvalidPatternError(`load.Name`, tmp34, `[a-zA-Z0-9\-]+`, err)
					}
				}
			}
			target.Name = tmp34
		} else {
			err = goa.MissingAttributeError(`load`, "name", err)
		}
		if v, ok := val["root"]; ok {
			var tmp35 string
			if val, ok := v.(string); ok {
				tmp35 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Root`, v, "string", err)
			}
			target.Root = tmp35
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// Field type
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

// MarshalField validates and renders an instance of Field into a interface{}
func MarshalField(source *Field, inErr error) (target map[string]interface{}, err error) {
	err = inErr
	if source.Name == "" {
		err = goa.MissingAttributeError(``, "name", err)
	}
	if source.Type == "" {
		err = goa.MissingAttributeError(``, "type", err)
	}

	if err == nil {
		if source.Name == "" {
			err = goa.MissingAttributeError(``, "name", err)
		}
		if source.Type == "" {
			err = goa.MissingAttributeError(``, "type", err)
		}
		if err == nil {
			if ok := goa.ValidatePattern(`[a-zA-Z0-9\-/]+`, source.Name); !ok {
				err = goa.InvalidPatternError(`.name`, source.Name, `[a-zA-Z0-9\-/]+`, err)
			}
			if !(source.Type == "string" || source.Type == "integer" || source.Type == "float" || source.Type == "boolean") {
				err = goa.InvalidEnumValueError(`.type`, source.Type, []interface{}{"string", "integer", "float", "boolean"}, err)
			}
			tmp36 := map[string]interface{}{
				"default":     source.Default,
				"description": source.Description,
				"name":        source.Name,
				"required":    source.Required,
				"root":        source.Root,
				"type":        source.Type,
			}
			target = tmp36
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
			var tmp37 interface{}
			tmp37 = v
			target.Default = tmp37
		}
		if v, ok := val["description"]; ok {
			var tmp38 string
			if val, ok := v.(string); ok {
				tmp38 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Description`, v, "string", err)
			}
			target.Description = tmp38
		}
		if v, ok := val["name"]; ok {
			var tmp39 string
			if val, ok := v.(string); ok {
				tmp39 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			if err == nil {
				if tmp39 != "" {
					if ok := goa.ValidatePattern(`[a-zA-Z0-9\-/]+`, tmp39); !ok {
						err = goa.InvalidPatternError(`load.Name`, tmp39, `[a-zA-Z0-9\-/]+`, err)
					}
				}
			}
			target.Name = tmp39
		} else {
			err = goa.MissingAttributeError(`load`, "name", err)
		}
		if v, ok := val["required"]; ok {
			var tmp40 bool
			if val, ok := v.(bool); ok {
				tmp40 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Required`, v, "bool", err)
			}
			target.Required = tmp40
		}
		if v, ok := val["root"]; ok {
			var tmp41 string
			if val, ok := v.(string); ok {
				tmp41 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Root`, v, "string", err)
			}
			target.Root = tmp41
		}
		if v, ok := val["type"]; ok {
			var tmp42 string
			if val, ok := v.(string); ok {
				tmp42 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Type`, v, "string", err)
			}
			if err == nil {
				if tmp42 != "" {
					if !(tmp42 == "string" || tmp42 == "integer" || tmp42 == "float" || tmp42 == "boolean") {
						err = goa.InvalidEnumValueError(`load.Type`, tmp42, []interface{}{"string", "integer", "float", "boolean"}, err)
					}
				}
			}
			target.Type = tmp42
		} else {
			err = goa.MissingAttributeError(`load`, "type", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}
