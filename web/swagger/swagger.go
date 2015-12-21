//************************************************************************//
// gestalt Swagger Spec
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/asteris-llc/gestalt/web
// --design=github.com/asteris-llc/gestalt/web/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package swagger

import "github.com/raphael/goa"

// MountController mounts the swagger spec controller under "/swagger.json".
func MountController(service goa.Service) {
	service.ServeFiles("/swagger.json", "swagger/swagger.json")
}
