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

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raphael/goa"
)

// MountController mounts the swagger spec controller under "/swagger.json".
func MountController(service goa.Service) {
	ctrl := service.NewController("Swagger")
	service.Info("mount", "ctrl", "Swagger", "action", "Show", "route", "GET /swagger.json")
	h := ctrl.NewHTTPRouterHandle("Show", getSwagger)
	service.HTTPHandler().(*httprouter.Router).Handle("GET", "/swagger.json", h)
}

// getSwagger is the httprouter handle that returns the Swagger spec.
// func getSwagger(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
func getSwagger(ctx *goa.Context) error {
	ctx.Header().Set("Content-Type", "application/swagger+json")
	ctx.Header().Set("Cache-Control", "public, max-age=3600")
	return ctx.Respond(200, []byte(spec))
}

// Generated spec
const spec = `{"swagger":"2.0","info":{"title":"Gestalt API","license":{"name":"Apache 2.0","url":"http://www.apache.org/licenses/LICENSE-2.0.html"},"version":""},"basePath":"/v1","consumes":["application/json"],"produces":["application/json"],"paths":{"/schemas":{"get":{"description":"list the schemas present in the system","operationId":"schema#list","consumes":["application/json"],"produces":["application/json"],"responses":{"200":{"description":"a list of schemas"},"500":{"description":""}},"schemes":["https"]},"post":{"description":"write a schema to the backend","operationId":"schema#create","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"setDefaults","in":"query","description":"set defaults when creating","required":false,"type":"boolean","default":true},{"name":"payload","in":"body","required":true,"schema":{"$ref":"#/definitions/CreateSchemaPayload"}}],"responses":{"201":{"description":"schema was accepted"},"500":{"description":""}},"schemes":["https"]}},"/schemas/{name}":{"get":{"description":"get a single schema","operationId":"schema#get","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"name","in":"path","required":true,"type":"string","pattern":"[a-zA-Z0-9\\-]+"}],"responses":{"200":{"description":"a single schema"},"404":{"description":""},"500":{"description":""}},"schemes":["https"]},"put":{"description":"update an existing schema","operationId":"schema#update","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"name","in":"path","required":true,"type":"string","pattern":"[a-zA-Z0-9\\-]+"},{"name":"setDefaults","in":"query","description":"set defaults when updating","required":false,"type":"boolean","default":false},{"name":"payload","in":"body","required":true,"schema":{"$ref":"#/definitions/UpdateSchemaPayload"}}],"responses":{"200":{"description":"update accepted"},"404":{"description":""},"500":{"description":""}},"schemes":["https"]},"delete":{"description":"delete an existing schema","operationId":"schema#delete","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"deleteKeys","in":"query","description":"delete the keys for this app as well","required":false,"type":"boolean","default":false},{"name":"name","in":"path","required":true,"type":"string","pattern":"[a-zA-Z0-9\\-]+"}],"responses":{"200":{"description":"deleted"},"404":{"description":""},"500":{"description":""}},"schemes":["https"]}},"/schemas/{name}/setDefaults":{"post":{"description":"set defaults on an existing schema","operationId":"schema#setDefaults","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"name","in":"path","required":true,"type":"string","pattern":"[a-zA-Z0-9\\-]+"}],"responses":{"200":{"description":"defaults set"},"404":{"description":""},"500":{"description":""}},"schemes":["https"]}},"/schemas/{name}/values":{"get":{"description":"list the values present in the K/V store","operationId":"value#list","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"name","in":"path","required":true,"type":"string","pattern":"[a-zA-Z0-9\\-]+"}],"responses":{"200":{"description":"a list of values"},"404":{"description":""},"500":{"description":""}},"schemes":["https"]}},"/schemas/{name}/values/{value}":{"get":{"description":"show a single value","operationId":"value#show","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"name","in":"path","required":true,"type":"string","pattern":"[a-zA-Z0-9\\-]+"},{"name":"value","in":"path","required":true,"type":"string","pattern":"[a-zA-Z0-9\\-/]+"}],"responses":{"200":{"description":"a single value"},"404":{"description":""},"500":{"description":""}},"schemes":["https"]},"put":{"description":"write a single value","operationId":"value#write","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"name","in":"path","required":true,"type":"string","pattern":"[a-zA-Z0-9\\-]+"},{"name":"value","in":"path","required":true,"type":"string","pattern":"[a-zA-Z0-9\\-/]+"}],"responses":{"200":{"description":"value was written"},"404":{"description":""},"500":{"description":""}},"schemes":["https"]},"delete":{"description":"delete a single value","operationId":"value#delete","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"name","in":"path","required":true,"type":"string","pattern":"[a-zA-Z0-9\\-]+"},{"name":"value","in":"path","required":true,"type":"string","pattern":"[a-zA-Z0-9\\-/]+"}],"responses":{"200":{"description":"value was deleted"},"404":{"description":""},"500":{"description":""}},"schemes":["https"]}}},"definitions":{"CreateSchemaPayload":{"title":"CreateSchemaPayload","type":"object","properties":{"backend":{"type":"string","description":"a registered backend"},"description":{"type":"string","description":"human readable description"},"fields":{"type":"array","items":{"$ref":"#/definitions/field"}},"name":{"type":"string","pattern":"[a-zA-Z0-9\\-]+"},"root":{"type":"string","description":"root for this schema (backend prefix + name if not set)"}},"required":["backend","name","fields"]},"UpdateSchemaPayload":{"title":"UpdateSchemaPayload","type":"object","properties":{"backend":{"type":"string","description":"a registered backend"},"description":{"type":"string","description":"human readable description"},"fields":{"type":"array","items":{"$ref":"#/definitions/field"}},"name":{"type":"string","pattern":"[a-zA-Z0-9\\-]+"},"root":{"type":"string","description":"root for this schema (backend prefix + name if not set)"}},"required":["backend","name","fields"]},"field":{"title":"field","type":"object","properties":{"default":{"type":"any","description":"the default for this field"},"description":{"type":"string","description":"human readable description"},"name":{"type":"string","pattern":"[a-zA-Z0-9\\-/]+"},"required":{"type":"boolean","description":"this field is required","defaultValue":false},"root":{"type":"string","description":"root for this key (backend prefix + schema name if not set)"},"type":{"type":"string","description":"type of value expected","enum":["string","integer","float","boolean"]}},"required":["name","type"]}}} `
