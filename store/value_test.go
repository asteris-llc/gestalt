package store

import (
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
	  "name": "test",
	  "backend": "mock",
	  "fields": [
	    {"name": "required", "type": "string", "required": true},
	    {"name": "optional", "type": "string"},
	    {"name": "integer", "type": "integer", "default": 1}
	  ]
	}`)
}

func (s *StoreValueSuite) SetupTest() {
	s.prefix = "mock/"

	s.mock = &mock.Mock{}
	s.backend = NewBackend(s.mock, "mock", s.prefix)

	var err error
	s.store, err = New([]*Backend{s.backend}, s.backend, s.backend)
	s.Require().Nil(err)

	// this is down here because we need a fully initialized store to use schemaPath
	s.mock.On("Get", s.store.schemaPath("test")).Return(&store.KVPair{Key: s.prefix + "test", Value: s.schemaBytes}, nil)
}

// RetrieveValues

func (s *StoreValueSuite) TestRetrieveValues() {
	// set
	s.mock.On("Get", s.prefix+"test/required").Return(&store.KVPair{Key: s.prefix + "test/required", Value: []byte("required")}, nil)
	s.mock.On("Get", s.prefix+"test/optional").Return(&store.KVPair{Key: s.prefix + "test/optional", Value: []byte("optional")}, nil)
	s.mock.On("Get", s.prefix+"test/integer").Return(&store.KVPair{Key: s.prefix + "test/integer", Value: []byte("1")}, nil)

	values, err := s.store.RetrieveValues("test")
	s.Require().Nil(err)

	s.mock.AssertExpectations(s.T())

	s.Assert().Equal("required", values["required"].(string))
	s.Assert().Equal("optional", values["optional"].(string))
	s.Assert().Equal(1, values["integer"].(int))
}

// RetrieveValue

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
	s.Assert().Equal(`integer: parsing "x": invalid syntax`, err.Error())

	s.mock.AssertExpectations(s.T())
}

func (s *StoreValueSuite) TestRetrieveValueMissing() {
	s.mock.On("Get", s.prefix+"test/integer").Return(&store.KVPair{}, store.ErrKeyNotFound)

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

	err := s.store.StoreValues("test", map[string]interface{}{"required": "a"})
	s.Assert().Nil(err)

	s.mock.AssertExpectations(s.T())
}

func (s *StoreValueSuite) TestStoreValuesInvalid() {
	err := s.store.StoreValues("test", map[string]interface{}{"required": 1})
	s.Assert().Equal(`required: "1" is not a valid string`, err.Error())

	s.mock.AssertExpectations(s.T())
}

// StoreDefaultValues

func (s *StoreValueSuite) TestStoreDefaultValues() {
	s.mock.On("Put", s.prefix+"test/integer", []byte("1"), &store.WriteOptions{}).Return(nil)

	err := s.store.StoreDefaultValues("test")
	s.Assert().Nil(err)

	s.mock.AssertExpectations(s.T())
}

// StoreValue

func (s *StoreValueSuite) TestStoreValueValid() {
	s.mock.On("Put", s.prefix+"test/required", []byte("a"), &store.WriteOptions{}).Return(nil)

	err := s.store.StoreValue("test", "required", "a")
	s.Assert().Nil(err)

	s.mock.AssertExpectations(s.T())
}

func (s *StoreValueSuite) TestStoreValueInvalid() {
	err := s.store.StoreValue("test", "required", 1)
	s.Assert().NotNil(err)

	s.mock.AssertExpectations(s.T())
}

func (s *StoreValueSuite) TestStoreValueNoKey() {
	err := s.store.StoreValue("test", "blah", nil)
	s.Assert().NotNil(err)

	s.mock.AssertExpectations(s.T())
}

// DeleteValues

func (s *StoreValueSuite) TestDeleteValues() {
	s.mock.On("Delete", s.prefix+"test/required").Return(store.ErrKeyNotFound)
	s.mock.On("Delete", s.prefix+"test/integer").Return(store.ErrKeyNotFound)
	s.mock.On("Delete", s.prefix+"test/optional").Return(nil)
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
	s.mock.On("Put", s.prefix+"test/integer", []byte("1"), &store.WriteOptions{}).Return(nil)

	err := s.store.DeleteValue("test", "integer")
	s.Assert().Nil(err)

	s.mock.AssertExpectations(s.T())
}

func TestStoreValueSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(StoreValueSuite))
}
