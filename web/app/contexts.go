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
	// links to values
	Values []string
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
		if v, ok := val["values"]; ok {
			var tmp7 []string
			if val, ok := v.([]interface{}); ok {
				tmp7 = make([]string, len(val))
				for tmp8, v := range val {
					if val, ok := v.(string); ok {
						tmp7[tmp8] = val
					} else {
						err = goa.InvalidAttributeTypeError(`payload.Values[*]`, v, "string", err)
					}
				}
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Values`, v, "array", err)
			}
			target.Values = tmp7
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateSchemaContext) Created(resp *AsterisGestaltSchema) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.asteris.gestalt.schema+json; charset=utf-8")
	return ctx.JSON(201, r)
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

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteSchemaContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *DeleteSchemaContext) OK(resp []byte) error {
	return ctx.Respond(200, resp)
}

// GetSchemaContext provides the schema get action context.
type GetSchemaContext struct {
	*goa.Context
	Name    string
	Payload *GetSchemaPayload
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
	p, err := NewGetSchemaPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// GetSchemaPayload is the schema get action payload.
type GetSchemaPayload struct {
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

// NewGetSchemaPayload instantiates a GetSchemaPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewGetSchemaPayload(raw interface{}) (p *GetSchemaPayload, err error) {
	p, err = UnmarshalGetSchemaPayload(raw, err)
	return
}

// UnmarshalGetSchemaPayload unmarshals and validates a raw interface{} into an instance of GetSchemaPayload
func UnmarshalGetSchemaPayload(source interface{}, inErr error) (target *GetSchemaPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(GetSchemaPayload)
		if v, ok := val["backend"]; ok {
			var tmp9 string
			if val, ok := v.(string); ok {
				tmp9 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Backend`, v, "string", err)
			}
			target.Backend = tmp9
		}
		if v, ok := val["description"]; ok {
			var tmp10 string
			if val, ok := v.(string); ok {
				tmp10 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Description`, v, "string", err)
			}
			target.Description = tmp10
		}
		if v, ok := val["fields"]; ok {
			var tmp11 []*Field
			if val, ok := v.([]interface{}); ok {
				tmp11 = make([]*Field, len(val))
				for tmp12, v := range val {
					tmp11[tmp12], err = UnmarshalField(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Fields`, v, "array", err)
			}
			target.Fields = tmp11
		}
		if v, ok := val["name"]; ok {
			var tmp13 string
			if val, ok := v.(string); ok {
				tmp13 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Name`, v, "string", err)
			}
			if err == nil {
				if tmp13 != "" {
					if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, tmp13); !ok {
						err = goa.InvalidPatternError(`payload.Name`, tmp13, `[a-zA-Z0-9\-]+`, err)
					}
				}
			}
			target.Name = tmp13
		}
		if v, ok := val["root"]; ok {
			var tmp14 string
			if val, ok := v.(string); ok {
				tmp14 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Root`, v, "string", err)
			}
			target.Root = tmp14
		}
		if v, ok := val["values"]; ok {
			var tmp15 []string
			if val, ok := v.([]interface{}); ok {
				tmp15 = make([]string, len(val))
				for tmp16, v := range val {
					if val, ok := v.(string); ok {
						tmp15[tmp16] = val
					} else {
						err = goa.InvalidAttributeTypeError(`payload.Values[*]`, v, "string", err)
					}
				}
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Values`, v, "array", err)
			}
			target.Values = tmp15
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetSchemaContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *GetSchemaContext) OK(resp *AsterisGestaltSchema) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.asteris.gestalt.schema+json; charset=utf-8")
	return ctx.JSON(200, r)
}

// ListSchemaContext provides the schema list action context.
type ListSchemaContext struct {
	*goa.Context
	Payload *ListSchemaPayload
}

// NewListSchemaContext parses the incoming request URL and body, performs validations and creates the
// context used by the schema controller list action.
func NewListSchemaContext(c *goa.Context) (*ListSchemaContext, error) {
	var err error
	ctx := ListSchemaContext{Context: c}
	p, err := NewListSchemaPayload(c.Payload())
	if err != nil {
		return nil, err
	}
	ctx.Payload = p
	return &ctx, err
}

// ListSchemaPayload is the schema list action payload.
type ListSchemaPayload struct {
	// list of schemas
	Schemas []*Schema
}

// NewListSchemaPayload instantiates a ListSchemaPayload from a raw request body.
// It validates each field and returns an error if any validation fails.
func NewListSchemaPayload(raw interface{}) (p *ListSchemaPayload, err error) {
	p, err = UnmarshalListSchemaPayload(raw, err)
	return
}

// UnmarshalListSchemaPayload unmarshals and validates a raw interface{} into an instance of ListSchemaPayload
func UnmarshalListSchemaPayload(source interface{}, inErr error) (target *ListSchemaPayload, err error) {
	err = inErr
	if val, ok := source.(map[string]interface{}); ok {
		target = new(ListSchemaPayload)
		if v, ok := val["schemas"]; ok {
			var tmp17 []*Schema
			if val, ok := v.([]interface{}); ok {
				tmp17 = make([]*Schema, len(val))
				for tmp18, v := range val {
					tmp17[tmp18], err = UnmarshalSchema(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Schemas`, v, "array", err)
			}
			target.Schemas = tmp17
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *ListSchemaContext) OK(resp *AsterisGestaltSchemas) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.asteris.gestalt.schemas+json; charset=utf-8")
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

// NotFound sends a HTTP response with status code 404.
func (ctx *SetDefaultsSchemaContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *SetDefaultsSchemaContext) OK(resp []byte) error {
	return ctx.Respond(200, resp)
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
	// links to values
	Values []string
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
			var tmp19 string
			if val, ok := v.(string); ok {
				tmp19 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Backend`, v, "string", err)
			}
			target.Backend = tmp19
		}
		if v, ok := val["description"]; ok {
			var tmp20 string
			if val, ok := v.(string); ok {
				tmp20 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Description`, v, "string", err)
			}
			target.Description = tmp20
		}
		if v, ok := val["fields"]; ok {
			var tmp21 []*Field
			if val, ok := v.([]interface{}); ok {
				tmp21 = make([]*Field, len(val))
				for tmp22, v := range val {
					tmp21[tmp22], err = UnmarshalField(v, err)
				}
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Fields`, v, "array", err)
			}
			target.Fields = tmp21
		}
		if v, ok := val["name"]; ok {
			var tmp23 string
			if val, ok := v.(string); ok {
				tmp23 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Name`, v, "string", err)
			}
			if err == nil {
				if tmp23 != "" {
					if ok := goa.ValidatePattern(`[a-zA-Z0-9\-]+`, tmp23); !ok {
						err = goa.InvalidPatternError(`payload.Name`, tmp23, `[a-zA-Z0-9\-]+`, err)
					}
				}
			}
			target.Name = tmp23
		}
		if v, ok := val["root"]; ok {
			var tmp24 string
			if val, ok := v.(string); ok {
				tmp24 = val
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Root`, v, "string", err)
			}
			target.Root = tmp24
		}
		if v, ok := val["values"]; ok {
			var tmp25 []string
			if val, ok := v.([]interface{}); ok {
				tmp25 = make([]string, len(val))
				for tmp26, v := range val {
					if val, ok := v.(string); ok {
						tmp25[tmp26] = val
					} else {
						err = goa.InvalidAttributeTypeError(`payload.Values[*]`, v, "string", err)
					}
				}
			} else {
				err = goa.InvalidAttributeTypeError(`payload.Values`, v, "array", err)
			}
			target.Values = tmp25
		}
	} else {
		err = goa.InvalidAttributeTypeError(`payload`, source, "dictionary", err)
	}
	return
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateSchemaContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateSchemaContext) OK(resp *AsterisGestaltSchema) error {
	r, err := resp.Dump()
	if err != nil {
		return fmt.Errorf("invalid response: %s", err)
	}
	ctx.Header().Set("Content-Type", "application/vnd.asteris.gestalt.schema+json; charset=utf-8")
	return ctx.JSON(200, r)
}

// DeleteValueContext provides the value delete action context.
type DeleteValueContext struct {
	*goa.Context
	Name       string
	SetDefault bool

	HasSetDefault bool
	Value         string
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
	rawSetDefault, ok := c.Get("setDefault")
	if ok {
		if setDefault, err2 := strconv.ParseBool(rawSetDefault); err2 == nil {
			ctx.SetDefault = setDefault
		} else {
			err = goa.InvalidParamTypeError("setDefault", rawSetDefault, "boolean", err)
		}
		ctx.HasSetDefault = true
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

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteValueContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *DeleteValueContext) OK(resp []byte) error {
	return ctx.Respond(200, resp)
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

// NotFound sends a HTTP response with status code 404.
func (ctx *WriteValueContext) NotFound() error {
	return ctx.Respond(404, nil)
}

// OK sends a HTTP response with status code 200.
func (ctx *WriteValueContext) OK(resp []byte) error {
	return ctx.Respond(200, resp)
}
