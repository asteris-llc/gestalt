package web

import (
	"github.com/asteris-llc/gestalt/web/app"
	"github.com/raphael/goa"
)

// SchemaController implements the schema resource.
type SchemaController struct {
	goa.Controller
}

// NewSchemaController creates a schema controller.
func NewSchemaController(service goa.Service) app.SchemaController {
	return &SchemaController{Controller: service.NewController("SchemaController")}
}

// Create runs the create action.
func (c *SchemaController) Create(ctx *app.CreateSchemaContext) error {
	return nil
}

// Delete runs the delete action.
func (c *SchemaController) Delete(ctx *app.DeleteSchemaContext) error {
	return nil
}

// Get runs the get action.
func (c *SchemaController) Get(ctx *app.GetSchemaContext) error {
	return nil
}

// List runs the list action.
func (c *SchemaController) List(ctx *app.ListSchemaContext) error {
	return nil
}

// SetDefaults runs the setDefaults action.
func (c *SchemaController) SetDefaults(ctx *app.SetDefaultsSchemaContext) error {
	return nil
}

// Update runs the update action.
func (c *SchemaController) Update(ctx *app.UpdateSchemaContext) error {
	return nil
}
