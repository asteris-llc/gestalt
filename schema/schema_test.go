package schema

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var sampleSchema = []byte(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "id": "testobj",
  "type": "object",
  "properties": {
    "num": {
      "type": "integer"
    },
    "str": {
      "type": "string"
    }
  },
  "required": [
    "num",
    "str"
  ]
}`)

// New

func TestNewValid(t *testing.T) {
	t.Parallel()

	schema, err := New([]byte(`{"type": "string"}`))
	assert.Nil(t, err)
	assert.NotNil(t, schema)
}

func TestNewInvalid(t *testing.T) {
	t.Parallel()

	schema, err := New([]byte{})
	assert.NotNil(t, err)
	assert.Nil(t, schema)
}

// BackendName

func TestBackendNamePresent(t *testing.T) {
	t.Parallel()

	schema, err := New([]byte(`{"backend": "test"}`))
	require.Nil(t, err)

	backend, ok := schema.BackendName()
	assert.Equal(t, backend, "test")
	assert.True(t, ok)
}

func TestBackendNameAbsent(t *testing.T) {
	t.Parallel()

	schema, err := New([]byte(`{}`))
	require.Nil(t, err)

	backend, ok := schema.BackendName()
	assert.Equal(t, backend, "")
	assert.False(t, ok)
}

// ValidateAll

func TestValidateAllValid(t *testing.T) {
	t.Parallel()

	schema, err := New(sampleSchema)
	require.Nil(t, err)

	valid, errs := schema.ValidateAll([]byte(`{"num":1,"str":"banana"}`))
	assert.True(t, valid)
	assert.Equal(t, len(errs), 0)
}

func TestValidateAllInvalid(t *testing.T) {
	t.Parallel()

	schema, err := New(sampleSchema)
	require.Nil(t, err)

	valid, errs := schema.ValidateAll([]byte(`{}`))
	assert.False(t, valid)
	assert.Equal(t, len(errs), 2)
}

// ValidateField

func TestValidateFieldValid(t *testing.T) {
	t.Parallel()

	schema, err := New(sampleSchema)
	require.Nil(t, err)

	valid, errs := schema.ValidateField("num", []byte("1"))
	assert.True(t, valid)
	assert.Equal(t, len(errs), 0)
}

func TestValidateFieldInvalid(t *testing.T) {
	t.Parallel()

	schema, err := New(sampleSchema)
	require.Nil(t, err)

	valid, errs := schema.ValidateField("num", []byte(`"banana"`))
	assert.False(t, valid)
	assert.Equal(t, len(errs), 1)
}

func TestValidateFieldBadName(t *testing.T) {
	t.Parallel()

	schema, err := New(sampleSchema)
	require.Nil(t, err)

	valid, errs := schema.ValidateField("num/b", []byte("1"))
	assert.False(t, valid)
	assert.Equal(t, errs, []error{ErrNoField})
}
