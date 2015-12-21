// !!! automatically generated !!!
// Use "make web/web.go" instead of editing this file.

package web

import (
	"github.com/asteris-llc/gestalt/store"
	"github.com/asteris-llc/gestalt/web/app"
	"github.com/asteris-llc/gestalt/web/swagger"
	"github.com/raphael/goa"
)

// Run starts the server
func Run(addr string, store *store.Store) {
	// Create service
	service := goa.New("API")

	// Setup middleware
	service.Use(goa.RequestID())
	service.Use(goa.LogRequest())
	service.Use(goa.Recover())

	// Mount "schema" controller
	c := NewSchemaController(service, store)
	app.MountSchemaController(service, c)
	// Mount "value" controller
	c2 := NewValueController(service, store)
	app.MountValueController(service, c2)

	// Mount Swagger spec provider controller
	swagger.MountController(service)

	// Start service, listen on given addr
	service.ListenAndServe(addr)
}
