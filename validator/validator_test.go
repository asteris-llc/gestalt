package validator

import (
	"errors"
	"fmt"
	"github.com/asteris-llc/gestalt/web/app"
	"github.com/stretchr/testify/assert"
	"testing"
)

var schema = &app.Schema{
	Name: "test",
	Fields: []*app.Field{
		{Name: "string", Type: "string"},
		{Name: "integer", Type: "integer"},
		{Name: "float", Type: "float"},
		{Name: "boolean", Type: "boolean"},

		{Name: "required", Type: "string", Required: true},
	},
}

func TestField(t *testing.T) {
	t.Parallel()

	v := New(schema)

	// exists
	field, err := v.Field("string")
	assert.Nil(t, err)
	assert.NotNil(t, field)

	// doesn't exist
	field, err = v.Field("bad-field")
	assert.Equal(t, ErrNoField, err)
	assert.Nil(t, field)
}

func TestValidateField(t *testing.T) {
	t.Parallel()

	v := New(schema)

	assert.Nil(t, v.ValidateField("string", "test"))
	assert.Nil(t, v.ValidateField("integer", 3))
	assert.Nil(t, v.ValidateField("float", 3.14))
	assert.Nil(t, v.ValidateField("boolean", true))

	assert.NotNil(t, v.ValidateField("string", 1))
	assert.NotNil(t, v.ValidateField("integer", 3.14))
	assert.NotNil(t, v.ValidateField("float", 3))
	assert.NotNil(t, v.ValidateField("boolean", "banana"))
}

func TestValidateFields(t *testing.T) {
	t.Parallel()

	v := New(schema)

	values := map[string]interface{}{
		"string":  "test",
		"boolean": "banana",
		"extra":   "some invalid value",
	}

	errs := v.ValidateAll(values)

	assert.Equal(t, errs["boolean"], errors.New(`"banana" is not a valid boolean`))
	delete(errs, "boolean")

	assert.Equal(t, errs["extra"], ErrNoField)
	delete(errs, "extra")

	assert.Equal(t, errs["required"], ErrFieldRequired)
	delete(errs, "required")

	assert.Equal(t, 0, len(errs), fmt.Sprintf("%v", errs))
}
