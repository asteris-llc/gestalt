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
			tmp26 := map[string]interface{}{
				"default":     source.Default,
				"description": source.Description,
				"name":        source.Name,
				"required":    source.Required,
				"root":        source.Root,
				"type":        source.Type,
			}
			target = tmp26
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
			var tmp27 interface{}
			tmp27 = v
			target.Default = tmp27
		}
		if v, ok := val["description"]; ok {
			var tmp28 string
			if val, ok := v.(string); ok {
				tmp28 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Description`, v, "string", err)
			}
			target.Description = tmp28
		}
		if v, ok := val["name"]; ok {
			var tmp29 string
			if val, ok := v.(string); ok {
				tmp29 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Name`, v, "string", err)
			}
			if err == nil {
				if tmp29 != "" {
					if ok := goa.ValidatePattern(`[a-zA-Z0-9\-/]+`, tmp29); !ok {
						err = goa.InvalidPatternError(`load.Name`, tmp29, `[a-zA-Z0-9\-/]+`, err)
					}
				}
			}
			target.Name = tmp29
		} else {
			err = goa.MissingAttributeError(`load`, "name", err)
		}
		if v, ok := val["required"]; ok {
			var tmp30 bool
			if val, ok := v.(bool); ok {
				tmp30 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Required`, v, "bool", err)
			}
			target.Required = tmp30
		}
		if v, ok := val["root"]; ok {
			var tmp31 string
			if val, ok := v.(string); ok {
				tmp31 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Root`, v, "string", err)
			}
			target.Root = tmp31
		}
		if v, ok := val["type"]; ok {
			var tmp32 string
			if val, ok := v.(string); ok {
				tmp32 = val
			} else {
				err = goa.InvalidAttributeTypeError(`load.Type`, v, "string", err)
			}
			if err == nil {
				if tmp32 != "" {
					if !(tmp32 == "string" || tmp32 == "integer" || tmp32 == "float" || tmp32 == "boolean") {
						err = goa.InvalidEnumValueError(`load.Type`, tmp32, []interface{}{"string", "integer", "float", "boolean"}, err)
					}
				}
			}
			target.Type = tmp32
		} else {
			err = goa.MissingAttributeError(`load`, "type", err)
		}
	} else {
		err = goa.InvalidAttributeTypeError(`load`, source, "dictionary", err)
	}
	return
}
