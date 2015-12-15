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
			tmp48 := map[string]interface{}{
				"description": source.Description,
				"name":        source.Name,
				"root":        source.Root,
				"type":        source.Type,
			}
			target = tmp48
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
			var tmp49 string
			if val, ok := v.(string); ok {
				tmp49 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Description`, v, "string", err)
			}
			target.Description = tmp49
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
					if ok := goa.ValidatePattern(`[a-zA-Z0-9\-/]+`, tmp50); !ok {
						err = goa.InvalidPatternError(`load.Name`, tmp50, `[a-zA-Z0-9\-/]+`, err)
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
		if v, ok := val["type"]; ok {
			var tmp52 string
			if val, ok := v.(string); ok {
				tmp52 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Type`, v, "string", err)
			}
			if err == nil {
				if tmp52 != "" {
					if !(tmp52 == "string" || tmp52 == "integer" || tmp52 == "float" || tmp52 == "boolean") {
						err = goa.InvalidEnumValueError(`load.Type`, tmp52, []interface{}{"string", "integer", "float", "boolean"}, err)
					}
				}
			}
			target.Type = tmp52
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
			tmp53 := map[string]interface{}{
				"backend":     source.Backend,
				"description": source.Description,
				"name":        source.Name,
				"root":        source.Root,
			}
			if source.Fields != nil {
				tmp54 := make([]map[string]interface{}, len(source.Fields))
				for tmp55, tmp56 := range source.Fields {
					tmp54[tmp55], err = MarshalField(tmp56, err)
				}
				tmp53["fields"] = tmp54
			}
			target = tmp53
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
			var tmp57 string
			if val, ok := v.(string); ok {
				tmp57 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Backend`, v, "string", err)
			}
			target.Backend = tmp57
		} else {
			err = goa.MissingAttributeError(`load`, "backend", err)
		}
		if v, ok := val["description"]; ok {
			var tmp58 string
			if val, ok := v.(string); ok {
				tmp58 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Description`, v, "string", err)
			}
			target.Description = tmp58
		}
		if v, ok := val["fields"]; ok {
			var tmp59 []*Field
			if val, ok := v.([]interface{}); ok {
				tmp59 = make([]*Field, len(val))
				for tmp60, v := range val {
					tmp59[tmp60], err = UnmarshalField(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`load.Fields`, v, "array", err)
			}
			target.Fields = tmp59
		} else {
			err = goa.MissingAttributeError(`load`, "fields", err)
		}
		if v, ok := val["name"]; ok {
			var tmp61 string
			if val, ok := v.(string); ok {
				tmp61 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			if err == nil {
				if tmp61 != "" {
					if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, tmp61); !ok {
						err = goa.InvalidPatternError(`load.Name`, tmp61, `[a-zA-Z0-9\-]+`, err)
					}
				}
			}
			target.Name = tmp61
		} else {
			err = goa.MissingAttributeError(`load`, "name", err)
		}
		if v, ok := val["root"]; ok {
			var tmp62 string
			if val, ok := v.(string); ok {
				tmp62 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Root`, v, "string", err)
			}
			target.Root = tmp62
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}
