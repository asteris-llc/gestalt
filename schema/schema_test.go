package schema

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sort"
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
    },
    "three": {
      "type": "integer",
      "default": 3
    },
    "nested": {
      "type": "object",
      "properties": {
        "inner": {
          "type": "string",
          "default": "inner"
        }
      }
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
	assert.Equal(t, 0, len(errs))
}

func TestValidateFieldInvalid(t *testing.T) {
	t.Parallel()

	schema, err := New(sampleSchema)
	require.Nil(t, err)

	valid, errs := schema.ValidateField("num", []byte(`"banana"`))
	assert.False(t, valid)
	assert.Equal(t, 1, len(errs))
}

func TestValidateFieldBadName(t *testing.T) {
	t.Parallel()

	schema, err := New(sampleSchema)
	require.Nil(t, err)

	valid, errs := schema.ValidateField("num/b", []byte("1"))
	assert.False(t, valid)
	assert.Equal(t, errs, []error{ErrNoField})
}

// Defaults

func TestDefaults(t *testing.T) {
	t.Parallel()

	schema, err := New(sampleSchema)
	require.Nil(t, err)

	result := schema.Defaults()

	should := map[string]interface{}{
		"three":        float64(3),
		"nested/inner": "inner",
	}

	for k, v := range should {
		switch v.(type) {
		case float64:
			assert.Equal(t, v.(float64), result[k].(float64), k)
		default:
			assert.Equal(t, v, result[k], k)
		}
	}
}

// FlatRequired

func TestFlatRequired(t *testing.T) {
	t.Parallel()

	schema, err := New(sampleSchema)
	require.Nil(t, err)

	required := schema.FlatRequired()
	should := []string{"num", "str"}

	sort.Strings(required)
	sort.Strings(should)

	assert.Equal(t, should, required)
}
