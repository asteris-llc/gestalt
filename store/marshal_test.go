package store

import (
	"testing"
	"testing/quick"
)

func TestRoundTripString(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping round trip tests in short mode")
	}

	f := func(value string) bool {
		marshalled := marshal(value)
		unmarshalled, err := unmarshal(marshalled, "string")
		newValue, ok := unmarshalled.(string)

		return err == nil && ok && newValue == value
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestRoundtripInteger(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping round trip tests in short mode")
	}

	f := func(value int) bool {
		marshalled := marshal(value)
		unmarshalled, err := unmarshal(marshalled, "integer")
		newValue, ok := unmarshalled.(int)

		return err == nil && ok && newValue == value
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestRoundtripFloat(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping round trip tests in short mode")
	}

	f := func(value float64) bool {
		marshalled := marshal(value)
		unmarshalled, err := unmarshal(marshalled, "float")
		newValue, ok := unmarshalled.(float64)

		return err == nil && ok && newValue == value
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestRoundtripBool(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping round trip tests in short mode")
	}

	f := func(value bool) bool {
		marshalled := marshal(value)
		unmarshalled, err := unmarshal(marshalled, "boolean")
		newValue, ok := unmarshalled.(bool)

		return err == nil && ok && newValue == value
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
