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

func TestStoreValueSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(StoreValueSuite))
}
