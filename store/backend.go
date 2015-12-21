package store

import (
	"github.com/asteris-llc/gestalt/web/app"
	"github.com/docker/libkv/store"
	"path"
)

// Backend stores information about keys under backend, plus a name for the
// backend.
type Backend struct {
	store.Store

	Name   string
	Prefix string
}

// NewBackend wraps a backend for use in a store
func NewBackend(wraps store.Store, name, prefix string) *Backend {
	return &Backend{wraps, name, prefix}
}

// SchemaKey returns the root for a schema
func (b *Backend) SchemaKey(schema *app.Schema) string {
	if schema.Root != "" {
		return schema.Root
	}
	return path.Join(b.Prefix, schema.Name)
}

// FieldKey returns the root for a schema and field
func (b *Backend) FieldKey(schema *app.Schema, field *app.Field) string {
	if field.Root != "" {
		return field.Root
	}

	return path.Join(b.SchemaKey(schema), field.Name)
}
