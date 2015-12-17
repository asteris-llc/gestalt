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
			tmp35 := map[string]interface{}{
				"default":     source.Default,
				"description": source.Description,
				"name":        source.Name,
				"required":    source.Required,
				"root":        source.Root,
				"type":        source.Type,
			}
			target = tmp35
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
			var tmp36 interface{}
			tmp36 = v
			target.Default = tmp36
		}
		if v, ok := val["description"]; ok {
			var tmp37 string
			if val, ok := v.(string); ok {
				tmp37 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Description`, v, "string", err)
			}
			target.Description = tmp37
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
					if ok := goa.ValidatePattern(`[a-zA-Z0-9\-/]+`, tmp38); !ok {
						err = goa.InvalidPatternError(`load.Name`, tmp38, `[a-zA-Z0-9\-/]+`, err)
					}
				}
			}
			target.Name = tmp38
		} else {
			err = goa.MissingAttributeError(`load`, "name", err)
		}
		if v, ok := val["required"]; ok {
			var tmp39 bool
			if val, ok := v.(bool); ok {
				tmp39 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Required`, v, "bool", err)
			}
			target.Required = tmp39
		}
		if v, ok := val["root"]; ok {
			var tmp40 string
			if val, ok := v.(string); ok {
				tmp40 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Root`, v, "string", err)
			}
			target.Root = tmp40
		}
		if v, ok := val["type"]; ok {
			var tmp41 string
			if val, ok := v.(string); ok {
				tmp41 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Type`, v, "string", err)
			}
			if err == nil {
				if tmp41 != "" {
					if !(tmp41 == "string" || tmp41 == "integer" || tmp41 == "float" || tmp41 == "boolean") {
						err = goa.InvalidEnumValueError(`load.Type`, tmp41, []interface{}{"string", "integer", "float", "boolean"}, err)
					}
				}
			}
			target.Type = tmp41
		} else {
			err = goa.MissingAttributeError(`load`, "type", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}

// Schema type
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

// MarshalSchema validates and renders an instance of Schema into a interface{}
func MarshalSchema(source *Schema, inErr error) (target map[string]interface{}, err error) {
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
			tmp42 := map[string]interface{}{
				"backend":     source.Backend,
				"description": source.Description,
				"name":        source.Name,
				"root":        source.Root,
			}
			if source.Fields != nil {
				tmp43 := make([]map[string]interface{}, len(source.Fields))
				for tmp44, tmp45 := range source.Fields {
					tmp43[tmp44], err = MarshalField(tmp45, err)
				}
				tmp42["fields"] = tmp43
			}
			target = tmp42
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
			var tmp46 string
			if val, ok := v.(string); ok {
				tmp46 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Backend`, v, "string", err)
			}
			target.Backend = tmp46
		} else {
			err = goa.MissingAttributeError(`load`, "backend", err)
		}
		if v, ok := val["description"]; ok {
			var tmp47 string
			if val, ok := v.(string); ok {
				tmp47 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Description`, v, "string", err)
			}
			target.Description = tmp47
		}
		if v, ok := val["fields"]; ok {
			var tmp48 []*Field
			if val, ok := v.([]interface{}); ok {
				tmp48 = make([]*Field, len(val))
				for tmp49, v := range val {
					tmp48[tmp49], err = UnmarshalField(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`load.Fields`, v, "array", err)
			}
			target.Fields = tmp48
		} else {
			err = goa.MissingAttributeError(`load`, "fields", err)
		}
		if v, ok := val["name"]; ok {
			var tmp50 string
			if val, ok := v.(string); ok {
				tmp50 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			if err == nil {
				if tmp50 != "" {
					if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, tmp50); !ok {
						err = goa.InvalidPatternError(`load.Name`, tmp50, `[a-zA-Z0-9\-]+`, err)
					}
				}
			}
			target.Name = tmp50
		} else {
			err = goa.MissingAttributeError(`load`, "name", err)
		}
		if v, ok := val["root"]; ok {
			var tmp51 string
			if val, ok := v.(string); ok {
				tmp51 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Root`, v, "string", err)
			}
			target.Root = tmp51
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}
