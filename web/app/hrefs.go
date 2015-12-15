//************************************************************************//
// gestalt: Application Resource Href Factories
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

import "fmt"

// ValueHref returns the resource href.
func ValueHref(name, value interface{}) string {
	return fmt.Sprintf("/v1/schemas/%v/values/%v", name, value)
}
