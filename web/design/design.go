package design

import (
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

func init() {
	API("gestalt", func() {
		Title("Gestalt API")
		BasePath("/v1")

		License(func() {
			Name("Apache 2.0")
			URL("http://www.apache.org/licenses/LICENSE-2.0.html")
		})
	})

	nameParam := func() { Param("name", String, func() { Pattern(`[a-zA-Z0-9\-]+`) }) }
	valueParam := func() { Param("value", String, func() { Pattern(`[a-zA-Z0-9\-/]+`) }) }

	Resource("schema", func() {
		Description("deal with schema values")
		BasePath("/schemas")

		Action("list", func() {
			Description("list the schemas present in the system")
			Routing(GET(""))

			Response(OK, func() {
				Description("a list of schemas")
				Media(CollectionOf(Schema, func() {
					View("default")
				}))
			})
			Response(InternalServerError)
		})

		Action("create", func() {
			Description("write a schema to the backend")
			Routing(POST(""))

			Payload(Schema)

			Params(func() {
				Param("setDefaults", Boolean, "set defaults when creating", func() { Default(true) })
			})

			Response(Created, func() {
				Description("schema was accepted")
				Media(Schema)
			})
			Response(InternalServerError)
		})

		Action("get", func() {
			Description("get a single schema")
			Routing(GET("/:name"))

			Params(nameParam)

			Response(OK, func() {
				Description("a single schema")
				Media(Schema)
			})

			Response(NotFound)
			Response(InternalServerError)
		})

		Action("update", func() {
			Description("update an existing schema")
			Routing(PUT("/:name"), POST("/:name"))

			Payload(Schema)

			Params(func() {
				nameParam()
				Param("setDefaults", Boolean, "set defaults when updating", func() { Default(false) })
			})

			Response(OK, func() {
				Description("update accepted")
				Media(Schema)
			})
			Response(NotFound)
			Response(InternalServerError)
		})

		Action("delete", func() {
			Description("delete an existing schema")
			Routing(DELETE("/:name"))

			Params(func() {
				nameParam()
				Param("deleteKeys", Boolean, "delete the keys for this app as well", func() { Default(false) })
			})

			Response(NoContent, func() { Description("deleted") })
			Response(NotFound)
			Response(InternalServerError)
		})

		Action("setDefaults", func() {
			Description("set defaults on an existing schema")
			Routing(POST("/:name/setDefaults"))

			Params(nameParam)

			Response(NoContent, func() { Description("defaults set") })
			Response(NotFound)
			Response(InternalServerError)
		})
	})

	Resource("value", func() {
		Description("deal with values as referenced by schemas")
		BasePath("/schemas/:name/values")

		Action("list", func() {
			Description("list the values present in the K/V store")
			Routing(GET(""))

			Params(nameParam)

			Response(OK, func() {
				Description("a list of values")
				Media("application/vnd.asteris.gestalt.values+json")
			})
			Response(NotFound)
			Response(InternalServerError)
		})

		Action("writeAll", func() {
			Description("write all the values at once")
			Routing(PUT(""), POST(""))

			Params(nameParam)

			Response(OK, func() {
				Description("the values which were written")
				Media("application/vnd.asteris.gestalt.values+json")
			})
			Response(NotFound)
			Response(InternalServerError)
		})

		Action("show", func() {
			Description("show a single value")
			Routing(GET("/*value"))

			Params(func() { nameParam(); valueParam() })

			Response(OK, func() {
				Description("a single value")
				Media("application/vnd.asteris.gestalt.value+json")
			})
			Response(NotFound)
			Response(InternalServerError)
		})

		Action("write", func() {
			Description("write a single value")
			Routing(PUT("/*value"), POST("/*value"))

			Params(func() { nameParam(); valueParam() })

			Response(OK, func() {
				Description("value was written")
				Media("application/vnd.asteris.gestalt.value+json")
			})
			Response(NotFound)
			Response(InternalServerError)
		})

		Action("delete", func() {
			Description("delete a single value")

			Routing(DELETE("/*value"))

			Params(func() { nameParam(); valueParam() })

			Response(NoContent, func() { Description("value was deleted") })
			Response(NotFound)
			Response(InternalServerError)
		})
	})
}
