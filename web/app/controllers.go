//************************************************************************//
// gestalt: Application Controllers
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
	"github.com/julienschmidt/httprouter"
	"github.com/raphael/goa"
)

// SchemaController is the controller interface for the Schema actions.
type SchemaController interface {
	goa.Controller
	Create(*CreateSchemaContext) error
	Delete(*DeleteSchemaContext) error
	Get(*GetSchemaContext) error
	List(*ListSchemaContext) error
	SetDefaults(*SetDefaultsSchemaContext) error
	Update(*UpdateSchemaContext) error
}

// MountSchemaController "mounts" a Schema resource controller on the given service.
func MountSchemaController(service goa.Service, ctrl SchemaController) {
	router := service.HTTPHandler().(*httprouter.Router)
	var h goa.Handler
	h = func(c *goa.Context) error {
		ctx, err := NewCreateSchemaContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Create(ctx)
	}
	router.Handle("POST", "/v1/schemas", ctrl.NewHTTPRouterHandle("Create", h))
	service.Info("mount", "ctrl", "Schema", "action", "Create", "route", "POST /v1/schemas")
	h = func(c *goa.Context) error {
		ctx, err := NewDeleteSchemaContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(ctx)
	}
	router.Handle("DELETE", "/v1/schemas/:name", ctrl.NewHTTPRouterHandle("Delete", h))
	service.Info("mount", "ctrl", "Schema", "action", "Delete", "route", "DELETE /v1/schemas/:name")
	h = func(c *goa.Context) error {
		ctx, err := NewGetSchemaContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Get(ctx)
	}
	router.Handle("GET", "/v1/schemas/:name", ctrl.NewHTTPRouterHandle("Get", h))
	service.Info("mount", "ctrl", "Schema", "action", "Get", "route", "GET /v1/schemas/:name")
	h = func(c *goa.Context) error {
		ctx, err := NewListSchemaContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(ctx)
	}
	router.Handle("GET", "/v1/schemas", ctrl.NewHTTPRouterHandle("List", h))
	service.Info("mount", "ctrl", "Schema", "action", "List", "route", "GET /v1/schemas")
	h = func(c *goa.Context) error {
		ctx, err := NewSetDefaultsSchemaContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.SetDefaults(ctx)
	}
	router.Handle("POST", "/v1/schemas/:name/setDefaults", ctrl.NewHTTPRouterHandle("SetDefaults", h))
	service.Info("mount", "ctrl", "Schema", "action", "SetDefaults", "route", "POST /v1/schemas/:name/setDefaults")
	h = func(c *goa.Context) error {
		ctx, err := NewUpdateSchemaContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Update(ctx)
	}
	router.Handle("PUT", "/v1/schemas/:name", ctrl.NewHTTPRouterHandle("Update", h))
	service.Info("mount", "ctrl", "Schema", "action", "Update", "route", "PUT /v1/schemas/:name")
}

// ValueController is the controller interface for the Value actions.
type ValueController interface {
	goa.Controller
	Delete(*DeleteValueContext) error
	List(*ListValueContext) error
	Show(*ShowValueContext) error
	Write(*WriteValueContext) error
}

// MountValueController "mounts" a Value resource controller on the given service.
func MountValueController(service goa.Service, ctrl ValueController) {
	router := service.HTTPHandler().(*httprouter.Router)
	var h goa.Handler
	h = func(c *goa.Context) error {
		ctx, err := NewDeleteValueContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Delete(ctx)
	}
	router.Handle("DELETE", "/v1/schemas/:name/values/*value", ctrl.NewHTTPRouterHandle("Delete", h))
	service.Info("mount", "ctrl", "Value", "action", "Delete", "route", "DELETE /v1/schemas/:name/values/*value")
	h = func(c *goa.Context) error {
		ctx, err := NewListValueContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.List(ctx)
	}
	router.Handle("GET", "/v1/schemas/:name/values", ctrl.NewHTTPRouterHandle("List", h))
	service.Info("mount", "ctrl", "Value", "action", "List", "route", "GET /v1/schemas/:name/values")
	h = func(c *goa.Context) error {
		ctx, err := NewShowValueContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	router.Handle("GET", "/v1/schemas/:name/values/*value", ctrl.NewHTTPRouterHandle("Show", h))
	service.Info("mount", "ctrl", "Value", "action", "Show", "route", "GET /v1/schemas/:name/values/*value")
	h = func(c *goa.Context) error {
		ctx, err := NewWriteValueContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Write(ctx)
	}
	router.Handle("PUT", "/v1/schemas/:name/values/*value", ctrl.NewHTTPRouterHandle("Write", h))
	service.Info("mount", "ctrl", "Value", "action", "Write", "route", "PUT /v1/schemas/:name/values/*value")
}
