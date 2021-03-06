package design

import (
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

var (
	// Field is a single field in the schema
	Field = MediaType("application/vnd.asteris.gestalt.field+json", func() {
		TypeName("Field")

		fields := func() {
			Attribute("name", String, func() { Pattern(`[a-zA-Z0-9\-/]+`) })
			Attribute("description", String, "human readable description")
			Attribute("root", String, "root for this key (backend prefix + schema name if not set)")
			Attribute("required", Boolean, "this field is required", func() { Default(false) })
			Attribute("default", Any, "the default for this field")

			Attribute("type", String, "type of value expected", func() {
				Enum("string", "integer", "float", "boolean")
			})

			// TODO: format attribute for email/url/ipv4/ipv6/etc

			Required("name", "type")
		}

		Attributes(fields)
		View("default", fields)
	})

	// Schema is the media type for SchemaPayload
	Schema = MediaType("application/vnd.asteris.gestalt.schema+json", func() {
		TypeName("Schema")

		fields := func() {
			Attribute("name", String, func() { Pattern(`[a-zA-Z0-9\-]+`) })
			Attribute("description", String, "human readable description")
			Attribute("root", String, "root for this schema (backend prefix + name if not set)")
			Attribute("backend", String, "a registered backend")

			Attribute("fields", ArrayOf(Field), func() { MinLength(1) })

			Required("backend", "name", "fields")
		}

		Attributes(fields)
		View("default", fields)
	})
)
