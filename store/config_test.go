package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromConfigValid(t *testing.T) {
	t.Parallel()

	config := []byte(`schemaBackend="consul"
defaultBackend="consul"

[[backend]]
type="consul"
name="consul"
prefix="/config"
endpoints=["http://127.0.0.1:8500"]`)

	store, err := FromConfig(config)
	assert.Nil(t, err)
	assert.NotNil(t, store)
}

func TestFromConfigMissingBackend(t *testing.T) {
	t.Parallel()

	config := []byte(`schemaBackend="consul"
defaultBackend="consul"`)

	store, err := FromConfig(config)
	assert.Nil(t, store)
	if assert.NotNil(t, err) {
		assert.Equal(t, err.Error(), `could not find backend named "consul"`)
	}
}

func TestFromConfigMissingSchemaBackend(t *testing.T) {
	t.Parallel()

	config := []byte(`defaultBackend="consul"

[[backend]]
type="consul"
name="consul"
prefix="/config"
endpoints=["http://127.0.0.1:8500"]`)

	store, err := FromConfig(config)
	assert.Nil(t, store)
	if assert.NotNil(t, err) {
		assert.Equal(t, err.Error(), `schema backend is required (schemaBackend key)`)
	}
}

func TestFromConfigMissingDefaultBackend(t *testing.T) {
	t.Parallel()

	config := []byte(`schemaBackend="consul"

[[backend]]
type="consul"
name="consul"
prefix="/config"
endpoints=["http://127.0.0.1:8500"]`)

	store, err := FromConfig(config)
	assert.Nil(t, store)
	if assert.NotNil(t, err) {
		assert.Equal(t, err.Error(), `default backend is required (defaultBackend key)`)
	}
}

func TestFromConfigMissingEndpoints(t *testing.T) {
	t.Parallel()

	config := []byte(`schemaBackend="consul"
defaultBackend="consul"

[[backend]]
type="consul"
name="consul"
prefix="/config"`)

	store, err := FromConfig(config)
	assert.Nil(t, store)
	if assert.NotNil(t, err) {
		assert.Equal(t, err.Error(), `need at least one endpoint in backend "consul"`)
	}
}

func TestFromConfigMissingName(t *testing.T) {
	t.Parallel()

	config := []byte(`schemaBackend="consul"
defaultBackend="consul"

[[backend]]
type="consul"
prefix="/config"
endpoints=["http://127.0.0.1:8500"]`)

	store, err := FromConfig(config)
	assert.Nil(t, store)
	if assert.NotNil(t, err) {
		assert.Equal(t, err.Error(), `name is required in backend 0`)
	}
}

func TestFromConfigMissingTypeValidName(t *testing.T) {
	t.Parallel()

	config := []byte(`schemaBackend="consul"
defaultBackend="consul"

[[backend]]
name="consul"
prefix="/config"
endpoints=["http://127.0.0.1:8500"]`)

	store, err := FromConfig(config)
	assert.Nil(t, err)
	assert.NotNil(t, store)
}

func TestFromConfigInvalidType(t *testing.T) {
	t.Parallel()

	config := []byte(`schemaBackend="consul"
defaultBackend="consul"

[[backend]]
name="consul"
type="not-a-backend"
prefix="/config"
endpoints=["http://127.0.0.1:8500"]`)

	store, err := FromConfig(config)
	assert.Nil(t, store)
	if assert.NotNil(t, err) {
		assert.Equal(t, err.Error(), `invalid type "not-a-backend" in backend "consul"`)
	}
}
