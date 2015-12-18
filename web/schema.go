package web

import (
	"github.com/asteris-llc/gestalt/store"
	"github.com/asteris-llc/gestalt/web/app"
	"github.com/raphael/goa"
)

// TODO: more realistic error handling

// SchemaController implements the schema resource.
type SchemaController struct {
	goa.Controller

	store *store.Store
}

// NewSchemaController creates a schema controller.
func NewSchemaController(service goa.Service, store *store.Store) app.SchemaController {
	return &SchemaController{
		Controller: service.NewController("SchemaController"),
		store:      store,
	}
}

// Create runs the create action.
func (c *SchemaController) Create(ctx *app.CreateSchemaContext) error {
	schema := app.Schema(*ctx.Payload)
	err := c.store.StoreSchema(schema.Name, &schema)
	if err != nil {
		ctx.Logger.Error(err.Error())
		return ctx.InternalServerError()
	}

	if ctx.HasSetDefaults && ctx.SetDefaults {
		err = c.store.StoreDefaultValues(schema.Name)
		if err != nil {
			ctx.Logger.Error(err.Error())
			return ctx.InternalServerError()
		}
	}

	return ctx.Created(&schema)
}

// Delete runs the delete action.
func (c *SchemaController) Delete(ctx *app.DeleteSchemaContext) error {
	if ctx.HasDeleteKeys && ctx.DeleteKeys {
		err := c.store.DeleteValues(ctx.Name)
		if err == store.ErrMissingKey {
			return ctx.NotFound()
		} else if err != nil {
			ctx.Logger.Error(err.Error())
			return ctx.InternalServerError()
		}
	}

	err := c.store.DeleteSchema(ctx.Name)
	if err == store.ErrMissingKey {
		return ctx.NotFound()
	} else if err != nil {
		ctx.Logger.Error(err.Error())
		return ctx.InternalServerError()
	}

	return ctx.OK([]byte{})
}

// Get runs the get action.
func (c *SchemaController) Get(ctx *app.GetSchemaContext) error {
	schema, err := c.store.RetrieveSchema(ctx.Name)
	if err == store.ErrMissingKey {
		return ctx.NotFound()
	} else if err != nil {
		ctx.Logger.Error(err.Error())
		return ctx.InternalServerError()
	}

	return ctx.OK(schema)
}

// List runs the list action.
func (c *SchemaController) List(ctx *app.ListSchemaContext) error {
	schemas, err := c.store.ListSchemas()
	if err != nil {
		ctx.Logger.Error(err.Error())
		return ctx.InternalServerError()
	}

	return ctx.OK(schemas)
}

// SetDefaults runs the setDefaults action.
func (c *SchemaController) SetDefaults(ctx *app.SetDefaultsSchemaContext) error {
	err := c.store.StoreDefaultValues(ctx.Name)
	if err == store.ErrMissingKey {
		return ctx.NotFound()
	} else if err != nil {
		ctx.Logger.Error(err.Error())
		return ctx.InternalServerError()
	}

	return ctx.OK([]byte{})
}

// Update runs the update action.
func (c *SchemaController) Update(ctx *app.UpdateSchemaContext) error {
	schema := app.Schema(*ctx.Payload)
	err := c.store.StoreSchema(schema.Name, &schema)
	if err == store.ErrMissingKey {
		return ctx.NotFound()
	} else if err != nil {
		ctx.Logger.Error(err.Error())
		return ctx.InternalServerError()
	}

	if ctx.HasSetDefaults && ctx.SetDefaults {
		err = c.store.StoreDefaultValues(schema.Name)
		if err == store.ErrMissingKey {
			return ctx.NotFound()
		} else if err != nil {
			ctx.Logger.Error(err.Error())
			return ctx.InternalServerError()
		}
	}

	return ctx.OK(&schema)
}
