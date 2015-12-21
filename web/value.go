package web

import (
	"encoding/json"
	"errors"
	"github.com/asteris-llc/gestalt/store"
	"github.com/asteris-llc/gestalt/web/app"
	"github.com/raphael/goa"
	"strings"
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
	err := c.store.DeleteValue(ctx.Name, strings.TrimLeft(ctx.Value, "/"))
	if err != nil {
		ctx.Logger.Error(err.Error())
		return ctx.InternalServerError()
	}

	return ctx.NoContent()
}

// List runs the list action.
func (c *ValueController) List(ctx *app.ListValueContext) error {
	values, err := c.store.RetrieveValues(ctx.Name)
	if err == store.ErrMissingKey {
		return ctx.NotFound()
	} else if err != nil {
		ctx.Logger.Error(err.Error())
		return ctx.InternalServerError()
	}

	resp, err := json.Marshal(values)
	if err != nil {
		ctx.Logger.Error(err.Error())
		return ctx.InternalServerError()
	}

	return ctx.OK(resp)
}

// Show runs the show action.
func (c *ValueController) Show(ctx *app.ShowValueContext) error {
	// because of a little bug in httprouter and how we consume values, requests
	// that should go to List go here instead. The following snippet redirects
	// them.
	if ctx.Name == "/" {
		newCtx, err := app.NewListValueContext(ctx.Context)
		if err != nil {
			return err
		}
		return c.List(newCtx)
	}

	value, err := c.store.RetrieveValue(ctx.Name, strings.TrimLeft(ctx.Value, "/"))
	if err == store.ErrMissingKey {
		return ctx.NotFound()
	} else if err != nil {
		ctx.Logger.Error(err.Error())
		return ctx.InternalServerError()
	}

	resp, err := json.Marshal(value)
	if err != nil {
		ctx.Logger.Error(err.Error())
		return ctx.InternalServerError()
	}

	return ctx.OK(resp)
}

// Write runs the write action.
func (c *ValueController) Write(ctx *app.WriteValueContext) error {
	// because of a little bug in httprouter and how we consume values, requests
	// that should go to WriteAll go here instead. The following snippet redirects
	// them.
	if ctx.Name == "/" {
		newCtx, err := app.NewWriteAllValueContext(ctx.Context)
		if err != nil {
			return err
		}
		return c.WriteAll(newCtx)
	}

	err := c.store.StoreValue(ctx.Name, strings.TrimLeft(ctx.Value, "/"), ctx.Payload())
	if err == store.ErrMissingKey {
		return ctx.NotFound()
	} else if err != nil {
		ctx.Logger.Error(err.Error())
		return ctx.InternalServerError()
	}

	resp, err := json.Marshal(ctx.Payload())
	if err != nil {
		ctx.Logger.Error(err.Error())
		return ctx.InternalServerError()
	}
	return ctx.OK(resp)
}

// WriteAll runs the writeAll action.
func (c *ValueController) WriteAll(ctx *app.WriteAllValueContext) error {
	vals, ok := ctx.Payload().(map[string]interface{})
	if !ok {
		return ctx.BadRequest(goa.NewBadRequestError(errors.New("could not convert to hash of any")))
	}

	for k, v := range vals {
		delete(vals, k)
		vals[strings.TrimLeft(k, "/")] = v
	}

	err := c.store.StoreValues(ctx.Name, vals)
	if err == store.ErrMissingKey {
		return ctx.NotFound()
	} else if err != nil {
		ctx.Logger.Error(err.Error())
		return ctx.InternalServerError()
	}

	resp, err := json.Marshal(ctx.Payload())
	if err != nil {
		ctx.Logger.Error(err.Error())
		return ctx.InternalServerError()
	}
	return ctx.OK(resp)
}
