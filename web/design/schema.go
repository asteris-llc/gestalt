package design

import (
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

var (
	// Schema is the base schema for the whole app
	Schema = Type("schema", func() {
		Attribute("name", String, func() { Pattern(`[a-zA-Z0-9\-]+`) })
		Attribute("description", String, "human readable description")
		Attribute("root", String, "root for this schema (backend prefix + name if not set)")
		Attribute("backend", String, "a registered backend")

		Attribute("fields", ArrayOf(Field))

		Required("backend", "name", "fields")
	})

	// Field is a single field in the schema
	Field = Type("field", func() {
		Attribute("name", String, func() { Pattern(`[a-zA-Z0-9\-/]+`) })
		Attribute("description", String, "human readable description")
		Attribute("root", String, "root for this key (backend prefix + schema name if not set)")

		Attribute("type", String, "type of value expected", func() {
			Enum("string", "integer", "float", "boolean")
		})

		// TODO: format attribute for email/url/ipv4/ipv6/etc

		Required("name", "type")
	})

	// SchemaMedia is the media type for Schema
	SchemaMedia = MediaType("application/vnd.asteris.gestalt.schema+json", func() {
		Reference(Schema)

		fields := func() {
			Attribute("name")
			Attribute("description")
			Attribute("root")
			Attribute("backend")
			Attribute("fields")

			Attribute("values", ArrayOf(String), "links to values")
		}

		Attributes(fields)
		View("default", fields)
	})
)
