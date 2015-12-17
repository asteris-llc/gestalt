package web

import (
	"github.com/asteris-llc/gestalt/store"
	"github.com/asteris-llc/gestalt/web/app"
	"github.com/raphael/goa"
)

// ValueController implements the value resource.
type ValueController struct {
	goa.Controller

	store *store.Store
}

// NewValueController creates a value controller.
func NewValueController(service goa.Service, store *store.Store) app.ValueController {
	return &ValueController{
		Controller: service.NewController("ValueController"),
		store:      store,
	}
}

// Delete runs the delete action.
func (c *ValueController) Delete(ctx *app.DeleteValueContext) error {
	return nil
}

// List runs the list action.
func (c *ValueController) List(ctx *app.ListValueContext) error {
	return nil
}

// Show runs the show action.
func (c *ValueController) Show(ctx *app.ShowValueContext) error {
	return nil
}

// Write runs the write action.
func (c *ValueController) Write(ctx *app.WriteValueContext) error {
	return nil
}
