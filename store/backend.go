package store

import (
	"github.com/docker/libkv/store"
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
