package store

import (
	"fmt"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type StoreValueSuite struct {
	suite.Suite

	prefix      string
	mock        *mock.Mock
	backend     *Backend
	store       *Store
	schemaBytes []byte
}

func (s *StoreValueSuite) SetupSuite() {
	s.prefix = "mock/"
	s.schemaBytes = []byte(`{
	  "properties": {
	    "required": {"type": "string"},
	    "optional": {"type": "string"},
	    "integer": {"type": "integer"},
	    "boolean": {"type": "boolean"},
	    "number": {"type": "number"},
	    "default": {"type": "string", "default": "default"},
	    "nested": {
	      "type": "object",
	      "properties": {
	        "inner": {"type": "string", "default": "nested/inner"}
	      }
	    }
	  },
	  "required": ["required"]
	}`)
}

func (s *StoreValueSuite) SetupTest() {
	s.prefix = "mock/"

	s.mock = &mock.Mock{}
	s.mock.On("Get", s.prefix+"test").Return(&store.KVPair{Key: s.prefix + "test", Value: s.schemaBytes}, nil)
	s.backend = NewBackend(s.mock, "mock", s.prefix)

	var err error
	s.store, err = New([]*Backend{s.backend}, s.backend, s.backend)
	s.Require().Nil(err)
}

// RetrieveValues

func (s *StoreValueSuite) TestRetrieveValues() {
	// set
	s.mock.On("Get", s.prefix+"test/required").Return(&store.KVPair{Key: s.prefix + "test/required", Value: []byte("required")}, nil)
	s.mock.On("Get", s.prefix+"test/number").Return(&store.KVPair{Key: s.prefix + "test/number", Value: []byte("3.14")}, nil)
	s.mock.On("Get", s.prefix+"test/integer").Return(&store.KVPair{Key: s.prefix + "test/integer", Value: []byte("3")}, nil)
	s.mock.On("Get", s.prefix+"test/boolean").Return(&store.KVPair{Key: s.prefix + "test/boolean", Value: []byte("true")}, nil)

	// not set
	for _, name := range []string{"optional", "default", "nested", "nested/inner"} {
		s.mock.On("Get", s.prefix+"test/"+name).Return(&store.KVPair{}, nil)
	}

	values, err := s.store.RetrieveValues("test")
	s.Require().Nil(err)

	s.mock.AssertExpectations(s.T())

	s.Assert().Equal("required", values["required"].(string))
	s.Assert().Equal(3.14, values["number"].(float64))
	s.Assert().Equal(3, values["integer"].(int))
	s.Assert().True(values["boolean"].(bool))
}

// RetrieveValues

func (s *StoreValueSuite) TestRetrieveValueValid() {
	s.mock.On("Get", s.prefix+"test/integer").Return(&store.KVPair{Key: s.prefix + "test/integer", Value: []byte("3")}, nil)

	value, err := s.store.RetrieveValue("test", "integer")
	s.Require().Nil(err)

	s.mock.AssertExpectations(s.T())

	s.Assert().Equal(value, 3)
}

func (s *StoreValueSuite) TestRetrieveValueInvalid() {
	s.mock.On("Get", s.prefix+"test/integer").Return(&store.KVPair{Key: s.prefix + "test/integer", Value: []byte("x")}, nil)

	_, err := s.store.RetrieveValue("test", "integer")
	s.Require().NotNil(err)
	s.Assert().Equal(`integer: strconv.ParseInt: parsing "x": invalid syntax`, err.Error())

	s.mock.AssertExpectations(s.T())
}

func (s *StoreValueSuite) TestRetrieveValueMissing() {
	s.mock.On("Get", s.prefix+"test/integer").Return(&store.KVPair{}, nil)

	_, err := s.store.RetrieveValue("test", "integer")
	s.Assert().Equal(ErrMissingKey, err)

	s.mock.AssertExpectations(s.T())
}

func (s *StoreValueSuite) TestRetrieveValueBadKey() {
	_, err := s.store.RetrieveValue("test", "blah")
	s.Assert().Equal(ErrMissingField, err)

	s.mock.AssertExpectations(s.T())
}

// StoreValues

func (s *StoreValueSuite) TestStoreValuesValid() {
	s.mock.On("Put", s.prefix+"test/required", []byte("a"), &store.WriteOptions{}).Return(nil)

	errors := s.store.StoreValues("test", []byte(`{"required": "a"}`))
	s.Assert().Equal(0, len(errors), fmt.Sprintf("%v", errors))

	s.mock.AssertExpectations(s.T())
}

func (s *StoreValueSuite) TestStoreValuesInvalid() {
	errors := s.store.StoreValues("test", []byte(`{"required": 1}`))
	s.Assert().Equal(1, len(errors), fmt.Sprintf("%v", errors))

	s.mock.AssertExpectations(s.T())
}

// StoreDefaultValues

func (s *StoreValueSuite) TestStoreDefaultValues() {
	s.mock.On("Put", s.prefix+"test/default", []byte("default"), &store.WriteOptions{}).Return(nil)
	s.mock.On("Put", s.prefix+"test/nested/inner", []byte("nested/inner"), &store.WriteOptions{}).Return(nil)

	err := s.store.StoreDefaultValues("test")
	s.Assert().Nil(err)

	s.mock.AssertExpectations(s.T())
}

// StoreValue

func (s *StoreValueSuite) TestStoreValueValid() {
	s.mock.On("Put", s.prefix+"test/required", []byte("a"), &store.WriteOptions{}).Return(nil)

	errors := s.store.StoreValue("test", "required", []byte(`"a"`))
	s.Assert().Equal(0, len(errors), fmt.Sprintf("%v", errors))

	s.mock.AssertExpectations(s.T())
}

func (s *StoreValueSuite) TestStoreValueInvalid() {
	errors := s.store.StoreValue("test", "required", []byte("1"))
	s.Assert().Equal(1, len(errors), fmt.Sprintf("%v", errors))

	s.mock.AssertExpectations(s.T())
}

func (s *StoreValueSuite) TestStoreValueNoKey() {
	errors := s.store.StoreValue("test", "blah", []byte{})
	s.Assert().Equal(1, len(errors), fmt.Sprintf("%v", errors))

	s.mock.AssertExpectations(s.T())
}

// DeleteValues

func (s *StoreValueSuite) TestDeleteValues() {
	s.mock.On("DeleteTree", s.prefix+"test").Return(nil)
	err := s.store.DeleteValues("test")
	s.Assert().Nil(err)

	s.mock.AssertExpectations(s.T())
}

// DeleteValue

func (s *StoreValueSuite) TestDeleteValueDeletable() {
	s.mock.On("Delete", s.prefix+"test/optional").Return(nil)

	err := s.store.DeleteValue("test", "optional")
	s.Assert().Nil(err)

	s.mock.AssertExpectations(s.T())
}

func (s *StoreValueSuite) TestDeleteValueRequired() {
	err := s.store.DeleteValue("test", "required")
	s.Assert().Equal(ErrFieldRequired, err)

	s.mock.AssertExpectations(s.T())
}

func (s *StoreValueSuite) TestDeleteValueDefault() {
	s.mock.On("Put", s.prefix+"test/default", []byte("default"), &store.WriteOptions{}).Return(nil)

	err := s.store.DeleteValue("test", "default")
	s.Assert().Nil(err)

	s.mock.AssertExpectations(s.T())
}

func TestStoreValueSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(StoreValueSuite))
}
