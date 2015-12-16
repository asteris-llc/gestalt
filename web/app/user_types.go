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
	// human readable description
	Description string
	Name        string
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
				"description": source.Description,
				"name":        source.Name,
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
		if v, ok := val["description"]; ok {
			var tmp36 string
			if val, ok := v.(string); ok {
				tmp36 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Description`, v, "string", err)
			}
			target.Description = tmp36
		}
		if v, ok := val["name"]; ok {
			var tmp37 string
			if val, ok := v.(string); ok {
				tmp37 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			if err == nil {
				if tmp37 != "" {
					if ok := goa.ValidatePattern(`[a-zA-Z0-9\-/]+`, tmp37); !ok {
						err = goa.InvalidPatternError(`load.Name`, tmp37, `[a-zA-Z0-9\-/]+`, err)
					}
				}
			}
			target.Name = tmp37
		} else {
			err = goa.MissingAttributeError(`load`, "name", err)
		}
		if v, ok := val["root"]; ok {
			var tmp38 string
			if val, ok := v.(string); ok {
				tmp38 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Root`, v, "string", err)
			}
			target.Root = tmp38
		}
		if v, ok := val["type"]; ok {
			var tmp39 string
			if val, ok := v.(string); ok {
				tmp39 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Type`, v, "string", err)
			}
			if err == nil {
				if tmp39 != "" {
					if !(tmp39 == "string" || tmp39 == "integer" || tmp39 == "float" || tmp39 == "boolean") {
						err = goa.InvalidEnumValueError(`load.Type`, tmp39, []interface{}{"string", "integer", "float", "boolean"}, err)
					}
				}
			}
			target.Type = tmp39
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
			tmp40 := map[string]interface{}{
				"backend":     source.Backend,
				"description": source.Description,
				"name":        source.Name,
				"root":        source.Root,
			}
			if source.Fields != nil {
				tmp41 := make([]map[string]interface{}, len(source.Fields))
				for tmp42, tmp43 := range source.Fields {
					tmp41[tmp42], err = MarshalField(tmp43, err)
				}
				tmp40["fields"] = tmp41
			}
			target = tmp40
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
			var tmp44 string
			if val, ok := v.(string); ok {
				tmp44 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Backend`, v, "string", err)
			}
			target.Backend = tmp44
		} else {
			err = goa.MissingAttributeError(`load`, "backend", err)
		}
		if v, ok := val["description"]; ok {
			var tmp45 string
			if val, ok := v.(string); ok {
				tmp45 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Description`, v, "string", err)
			}
			target.Description = tmp45
		}
		if v, ok := val["fields"]; ok {
			var tmp46 []*Field
			if val, ok := v.([]interface{}); ok {
				tmp46 = make([]*Field, len(val))
				for tmp47, v := range val {
					tmp46[tmp47], err = UnmarshalField(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`load.Fields`, v, "array", err)
			}
			target.Fields = tmp46
		} else {
			err = goa.MissingAttributeError(`load`, "fields", err)
		}
		if v, ok := val["name"]; ok {
			var tmp48 string
			if val, ok := v.(string); ok {
				tmp48 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			if err == nil {
				if tmp48 != "" {
					if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, tmp48); !ok {
						err = goa.InvalidPatternError(`load.Name`, tmp48, `[a-zA-Z0-9\-]+`, err)
					}
				}
			}
			target.Name = tmp48
		} else {
			err = goa.MissingAttributeError(`load`, "name", err)
		}
		if v, ok := val["root"]; ok {
			var tmp49 string
			if val, ok := v.(string); ok {
				tmp49 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Root`, v, "string", err)
			}
			target.Root = tmp49
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}
