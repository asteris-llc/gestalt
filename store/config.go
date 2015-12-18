package store

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/boltdb"
	"github.com/docker/libkv/store/consul"
	"github.com/docker/libkv/store/etcd"
	"github.com/docker/libkv/store/zookeeper"
	"time"
)

type config struct {
	SchemaBackend  string `toml:"schemaBackend"`
	DefaultBackend string `toml:"defaultBackend"`

	Backends []backendConfig `toml:"backend"`
}

type backendConfig struct {
	Name      string   `toml:"name"`
	Type      string   `toml:"type"`
	Prefix    string   `toml:"prefix"`
	Endpoints []string `toml:"endpoints"`

	// Used for the boltdb backend
	Bucket string `toml:"bucket"`

	// Specified in seconds
	ConnectionTimeout time.Duration `toml:"connectionTimeout"`
}

// FromConfig loads a Store from config bytes
func FromConfig(src []byte) (*Store, error) {
	var dest config

	err := toml.Unmarshal(src, &dest)
	if err != nil {
		return nil, err
	}

	// make real backends
	backends := []*Backend{}
	for i, backend := range dest.Backends {
		if backend.Name == "" {
			return nil, fmt.Errorf(`name is required in backend %d`, i)
		}

		if backend.Type == "" {
			backend.Type = backend.Name
		}

		backendConfig := &store.Config{
			Bucket:            backend.Bucket,
			ConnectionTimeout: backend.ConnectionTimeout * time.Second,
		}

		if len(backend.Endpoints) == 0 {
			return nil, fmt.Errorf(`need at least one endpoint in backend "%s"`, backend.Name)
		}

		var (
			inner store.Store
			err   error
		)

		switch backend.Type {
		case "consul":
			inner, err = consul.New(backend.Endpoints, backendConfig)
		case "etcd":
			inner, err = etcd.New(backend.Endpoints, backendConfig)
		case "boltdb":
			inner, err = boltdb.New(backend.Endpoints, backendConfig)
		case "zookeeper":
			inner, err = zookeeper.New(backend.Endpoints, backendConfig)
		default:
			err = fmt.Errorf(`invalid type "%s" in backend "%s"`, backend.Type, backend.Name)
		}

		if err != nil {
			return nil, err
		}

		backends = append(backends, NewBackend(inner, backend.Name, backend.Prefix))
	}

	if dest.SchemaBackend == "" {
		return nil, errors.New("schema backend is required (schemaBackend key)")
	}
	var schemaBackend *Backend
	for _, backend := range backends {
		if backend.Name == dest.SchemaBackend {
			schemaBackend = backend
		}
	}
	if schemaBackend == nil {
		return nil, fmt.Errorf(`could not find backend named "%s"`, dest.SchemaBackend)
	}

	if dest.DefaultBackend == "" {
		return nil, errors.New("default backend is required (defaultBackend key)")
	}
	var defaultBackend *Backend
	for _, backend := range backends {
		if backend.Name == dest.DefaultBackend {
			defaultBackend = backend
		}
	}
	if defaultBackend == nil {
		return nil, fmt.Errorf(`could not find backend named "%s"`, dest.DefaultBackend)
	}

	return New(backends, schemaBackend, defaultBackend)
}
