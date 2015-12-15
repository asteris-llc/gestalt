// !!! automatically generated !!!
// Use "make web/web.go" instead of editing this file.

package web

import (
	"github.com/asteris-llc/gestalt/web/app"
	"github.com/asteris-llc/gestalt/web/swagger"
	"github.com/raphael/goa"
)

func Run(addr string) {
	// Create service
	service := goa.New("API")

	// Setup middleware
	service.Use(goa.RequestID())
	service.Use(goa.LogRequest())
	service.Use(goa.Recover())

	// Mount "schema" controller
	c := NewSchemaController(service)
	app.MountSchemaController(service, c)
	// Mount "value" controller
	c2 := NewValueController(service)
	app.MountValueController(service, c2)

	// Mount Swagger spec provider controller
	swagger.MountController(service)

	// Start service, listen on given addr
	service.ListenAndServe(addr)
}
