//************************************************************************//
// gestalt: Application Contexts
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
	"fmt"
	"strconv"

	"github.com/raphael/goa"
)

// CreateSchemaContext provides the schema create action context.
type CreateSchemaContext struct {
	*goa.Context
	SetDefaults bool

	HasSetDefaults bool
	Payload        *CreateSchemaPayload
}

// NewCreateSchemaContext parses the incoming request URL and body, performs validations and creates the
// context used by the schema controller create action.
func NewCreateSchemaContext(c *goa.Context) (*CreateSchemaContext, error) {
	var err error
	ctx := CreateSchemaContext{Context: c}
	rawSetDefaults, ok := c.Get("setDefaults")
	if ok {
		if setDefaults, err2 := strconv.ParseBool(rawSetDefaults); err2 == nil {
			ctx.SetDefaults = setDefaults
		} else {
			err = goa.InvalidParamTypeError("setDefaults", rawSetDefaults, "boolean", err)
		}
		ctx.HasSetDefaults = true
	}
	p, err := NewCreateSchemaPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// CreateSchemaPayload is the schema create action payload.
type CreateSchemaPayload struct {
	// a registered backend
	Backend string
	// human readable description
	Description string
	Fields      []*Field
	Name        string
	// root for this schema (backend prefix + name if not set)
	Root string
}

// NewCreateSchemaPayload instantiates a CreateSchemaPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewCreateSchemaPayload(raw interface{}) (p *CreateSchemaPayload, err error) {
	p, err = UnmarshalCreateSchemaPayload(raw, err)
	return
}

// UnmarshalCreateSchemaPayload unmarshals and validates a raw interface{} into an instance of CreateSchemaPayload
func UnmarshalCreateSchemaPayload(source interface{}, inErr error) (target *CreateSchemaPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(CreateSchemaPayload)
		if v, ok := val["backend"]; ok {
			var tmp1 string
			if val, ok := v.(string); ok {
				tmp1 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Backend`, v, "string", err)
			}
			target.Backend = tmp1
		} else {
			err = goa.MissingAttributeError(`payload`, "backend", err)
		}
		if v, ok := val["description"]; ok {
			var tmp2 string
			if val, ok := v.(string); ok {
				tmp2 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Description`, v, "string", err)
			}
			target.Description = tmp2
		}
		if v, ok := val["fields"]; ok {
			var tmp3 []*Field
			if val, ok := v.([]interface{}); ok {
				tmp3 = make([]*Field, len(val))
				for tmp4, v := range val {
					tmp3[tmp4], err = UnmarshalField(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Fields`, v, "array", err)
			}
			target.Fields = tmp3
		} else {
			err = goa.MissingAttributeError(`payload`, "fields", err)
		}
		if v, ok := val["name"]; ok {
			var tmp5 string
			if val, ok := v.(string); ok {
				tmp5 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Name`, v, "string", err)
			}
			if err == nil {
				if tmp5 != "" {
					if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, tmp5); !ok {
						err = goa.InvalidPatternError(`payload.Name`, tmp5, `[a-zA-Z0-9\-]+`, err)
					}
				}
			}
			target.Name = tmp5
		} else {
			err = goa.MissingAttributeError(`payload`, "name", err)
		}
		if v, ok := val["root"]; ok {
			var tmp6 string
			if val, ok := v.(string); ok {
				tmp6 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Root`, v, "string", err)
			}
			target.Root = tmp6
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateSchemaContext) Created(resp *Schema) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.schema+json; charset=utf-8")
	return ctx.JSON(201, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateSchemaContext) InternalServerError() error {
	return ctx.Respond(500, nil)
}

// DeleteSchemaContext provides the schema delete action context.
type DeleteSchemaContext struct {
	*goa.Context
	DeleteKeys bool

	HasDeleteKeys bool
	Name          string
}

// NewDeleteSchemaContext parses the incoming request URL and body, performs validations and creates the
// context used by the schema controller delete action.
func NewDeleteSchemaContext(c *goa.Context) (*DeleteSchemaContext, error) {
	var err error
	ctx := DeleteSchemaContext{Context: c}
	rawDeleteKeys, ok := c.Get("deleteKeys")
	if ok {
		if deleteKeys, err2 := strconv.ParseBool(rawDeleteKeys); err2 == nil {
			ctx.DeleteKeys = deleteKeys
		} else {
			err = goa.InvalidParamTypeError("deleteKeys", rawDeleteKeys, "boolean", err)
		}
		ctx.HasDeleteKeys = true
	}
	rawName, ok := c.Get("name")
	if ok {
		ctx.Name = rawName
		if ctx.Name != "" {
			if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, ctx.Name); !ok {
				err = goa.InvalidPatternError(`name`, ctx.Name, `[a-zA-Z0-9\-]+`, err)
			}
		}
	}
	return &ctx, err
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *DeleteSchemaContext) InternalServerError() error {
	return ctx.Respond(500, nil)
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteSchemaContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteSchemaContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// GetSchemaContext provides the schema get action context.
type GetSchemaContext struct {
	*goa.Context
	Name string
}

// NewGetSchemaContext parses the incoming request URL and body, performs validations and creates the
// context used by the schema controller get action.
func NewGetSchemaContext(c *goa.Context) (*GetSchemaContext, error) {
	var err error
	ctx := GetSchemaContext{Context: c}
	rawName, ok := c.Get("name")
	if ok {
		ctx.Name = rawName
		if ctx.Name != "" {
			if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, ctx.Name); !ok {
				err = goa.InvalidPatternError(`name`, ctx.Name, `[a-zA-Z0-9\-]+`, err)
			}
		}
	}
	return &ctx, err
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *GetSchemaContext) InternalServerError() error {
	return ctx.Respond(500, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetSchemaContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *GetSchemaContext) OK(resp *Schema) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.schema+json; charset=utf-8")
	return ctx.JSON(200, r)
}

// ListSchemaContext provides the schema list action context.
type ListSchemaContext struct {
	*goa.Context
}

// NewListSchemaContext parses the incoming request URL and body, performs validations and creates the
// context used by the schema controller list action.
func NewListSchemaContext(c *goa.Context) (*ListSchemaContext, error) {
	var err error
	ctx := ListSchemaContext{Context: c}
	return &ctx, err
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListSchemaContext) InternalServerError() error {
	return ctx.Respond(500, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ListSchemaContext) OK(resp SchemaCollection) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.schema+json; type=collection; charset=utf-8")
	return ctx.JSON(200, r)
}

// SetDefaultsSchemaContext provides the schema setDefaults action context.
type SetDefaultsSchemaContext struct {
	*goa.Context
	Name string
}

// NewSetDefaultsSchemaContext parses the incoming request URL and body, performs validations and creates the
// context used by the schema controller setDefaults action.
func NewSetDefaultsSchemaContext(c *goa.Context) (*SetDefaultsSchemaContext, error) {
	var err error
	ctx := SetDefaultsSchemaContext{Context: c}
	rawName, ok := c.Get("name")
	if ok {
		ctx.Name = rawName
		if ctx.Name != "" {
			if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, ctx.Name); !ok {
				err = goa.InvalidPatternError(`name`, ctx.Name, `[a-zA-Z0-9\-]+`, err)
			}
		}
	}
	return &ctx, err
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *SetDefaultsSchemaContext) InternalServerError() error {
	return ctx.Respond(500, nil)
}

// NoContent sends a HTTP response with status code 204.
func (ctx *SetDefaultsSchemaContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *SetDefaultsSchemaContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// UpdateSchemaContext provides the schema update action context.
type UpdateSchemaContext struct {
	*goa.Context
	Name        string
	SetDefaults bool

	HasSetDefaults bool
	Payload        *UpdateSchemaPayload
}

// NewUpdateSchemaContext parses the incoming request URL and body, performs validations and creates the
// context used by the schema controller update action.
func NewUpdateSchemaContext(c *goa.Context) (*UpdateSchemaContext, error) {
	var err error
	ctx := UpdateSchemaContext{Context: c}
	rawName, ok := c.Get("name")
	if ok {
		ctx.Name = rawName
		if ctx.Name != "" {
			if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, ctx.Name); !ok {
				err = goa.InvalidPatternError(`name`, ctx.Name, `[a-zA-Z0-9\-]+`, err)
			}
		}
	}
	rawSetDefaults, ok := c.Get("setDefaults")
	if ok {
		if setDefaults, err2 := strconv.ParseBool(rawSetDefaults); err2 == nil {
			ctx.SetDefaults = setDefaults
		} else {
			err = goa.InvalidParamTypeError("setDefaults", rawSetDefaults, "boolean", err)
		}
		ctx.HasSetDefaults = true
	}
	p, err := NewUpdateSchemaPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// UpdateSchemaPayload is the schema update action payload.
type UpdateSchemaPayload struct {
	// a registered backend
	Backend string
	// human readable description
	Description string
	Fields      []*Field
	Name        string
	// root for this schema (backend prefix + name if not set)
	Root string
}

// NewUpdateSchemaPayload instantiates a UpdateSchemaPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewUpdateSchemaPayload(raw interface{}) (p *UpdateSchemaPayload, err error) {
	p, err = UnmarshalUpdateSchemaPayload(raw, err)
	return
}

// UnmarshalUpdateSchemaPayload unmarshals and validates a raw interface{} into an instance of UpdateSchemaPayload
func UnmarshalUpdateSchemaPayload(source interface{}, inErr error) (target *UpdateSchemaPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(UpdateSchemaPayload)
		if v, ok := val["backend"]; ok {
			var tmp7 string
			if val, ok := v.(string); ok {
				tmp7 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Backend`, v, "string", err)
			}
			target.Backend = tmp7
		} else {
			err = goa.MissingAttributeError(`payload`, "backend", err)
		}
		if v, ok := val["description"]; ok {
			var tmp8 string
			if val, ok := v.(string); ok {
				tmp8 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Description`, v, "string", err)
			}
			target.Description = tmp8
		}
		if v, ok := val["fields"]; ok {
			var tmp9 []*Field
			if val, ok := v.([]interface{}); ok {
				tmp9 = make([]*Field, len(val))
				for tmp10, v := range val {
					tmp9[tmp10], err = UnmarshalField(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Fields`, v, "array", err)
			}
			target.Fields = tmp9
		} else {
			err = goa.MissingAttributeError(`payload`, "fields", err)
		}
		if v, ok := val["name"]; ok {
			var tmp11 string
			if val, ok := v.(string); ok {
				tmp11 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Name`, v, "string", err)
			}
			if err == nil {
				if tmp11 != "" {
					if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, tmp11); !ok {
						err = goa.InvalidPatternError(`payload.Name`, tmp11, `[a-zA-Z0-9\-]+`, err)
					}
				}
			}
			target.Name = tmp11
		} else {
			err = goa.MissingAttributeError(`payload`, "name", err)
		}
		if v, ok := val["root"]; ok {
			var tmp12 string
			if val, ok := v.(string); ok {
				tmp12 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Root`, v, "string", err)
			}
			target.Root = tmp12
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *UpdateSchemaContext) InternalServerError() error {
	return ctx.Respond(500, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateSchemaContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateSchemaContext) OK(resp *Schema) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.schema+json; charset=utf-8")
	return ctx.JSON(200, r)
}

// DeleteValueContext provides the value delete action context.
type DeleteValueContext struct {
	*goa.Context
	Name  string
	Value string
}

// NewDeleteValueContext parses the incoming request URL and body, performs validations and creates the
// context used by the value controller delete action.
func NewDeleteValueContext(c *goa.Context) (*DeleteValueContext, error) {
	var err error
	ctx := DeleteValueContext{Context: c}
	rawName, ok := c.Get("name")
	if ok {
		ctx.Name = rawName
		if ctx.Name != "" {
			if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, ctx.Name); !ok {
				err = goa.InvalidPatternError(`name`, ctx.Name, `[a-zA-Z0-9\-]+`, err)
			}
		}
	}
	rawValue, ok := c.Get("value")
	if ok {
		ctx.Value = rawValue
		if ctx.Value != "" {
			if ok := goa.ValidatePattern(`[a-zA-Z0-9\-/]+`, ctx.Value); !ok {
				err = goa.InvalidPatternError(`value`, ctx.Value, `[a-zA-Z0-9\-/]+`, err)
			}
		}
	}
	return &ctx, err
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *DeleteValueContext) InternalServerError() error {
	return ctx.Respond(500, nil)
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteValueContext) NoContent() error {
	return ctx.Respond(204, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteValueContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// ListValueContext provides the value list action context.
type ListValueContext struct {
	*goa.Context
	Name string
}

// NewListValueContext parses the incoming request URL and body, performs validations and creates the
// context used by the value controller list action.
func NewListValueContext(c *goa.Context) (*ListValueContext, error) {
	var err error
	ctx := ListValueContext{Context: c}
	rawName, ok := c.Get("name")
	if ok {
		ctx.Name = rawName
		if ctx.Name != "" {
			if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, ctx.Name); !ok {
				err = goa.InvalidPatternError(`name`, ctx.Name, `[a-zA-Z0-9\-]+`, err)
			}
		}
	}
	return &ctx, err
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListValueContext) InternalServerError() error {
	return ctx.Respond(500, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListValueContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ListValueContext) OK(resp []byte) error {
	return ctx.Respond(200, resp)
}

// ShowValueContext provides the value show action context.
type ShowValueContext struct {
	*goa.Context
	Name  string
	Value string
}

// NewShowValueContext parses the incoming request URL and body, performs validations and creates the
// context used by the value controller show action.
func NewShowValueContext(c *goa.Context) (*ShowValueContext, error) {
	var err error
	ctx := ShowValueContext{Context: c}
	rawName, ok := c.Get("name")
	if ok {
		ctx.Name = rawName
		if ctx.Name != "" {
			if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, ctx.Name); !ok {
				err = goa.InvalidPatternError(`name`, ctx.Name, `[a-zA-Z0-9\-]+`, err)
			}
		}
	}
	rawValue, ok := c.Get("value")
	if ok {
		ctx.Value = rawValue
		if ctx.Value != "" {
			if ok := goa.ValidatePattern(`[a-zA-Z0-9\-/]+`, ctx.Value); !ok {
				err = goa.InvalidPatternError(`value`, ctx.Value, `[a-zA-Z0-9\-/]+`, err)
			}
		}
	}
	return &ctx, err
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowValueContext) InternalServerError() error {
	return ctx.Respond(500, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowValueContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowValueContext) OK(resp []byte) error {
	return ctx.Respond(200, resp)
}

// WriteValueContext provides the value write action context.
type WriteValueContext struct {
	*goa.Context
	Name  string
	Value string
}

// NewWriteValueContext parses the incoming request URL and body, performs validations and creates the
// context used by the value controller write action.
func NewWriteValueContext(c *goa.Context) (*WriteValueContext, error) {
	var err error
	ctx := WriteValueContext{Context: c}
	rawName, ok := c.Get("name")
	if ok {
		ctx.Name = rawName
		if ctx.Name != "" {
			if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, ctx.Name); !ok {
				err = goa.InvalidPatternError(`name`, ctx.Name, `[a-zA-Z0-9\-]+`, err)
			}
		}
	}
	rawValue, ok := c.Get("value")
	if ok {
		ctx.Value = rawValue
		if ctx.Value != "" {
			if ok := goa.ValidatePattern(`[a-zA-Z0-9\-/]+`, ctx.Value); !ok {
				err = goa.InvalidPatternError(`value`, ctx.Value, `[a-zA-Z0-9\-/]+`, err)
			}
		}
	}
	return &ctx, err
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *WriteValueContext) InternalServerError() error {
	return ctx.Respond(500, nil)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *WriteValueContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *WriteValueContext) OK(resp []byte) error {
	return ctx.Respond(200, resp)
}
