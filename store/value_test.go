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
	s.backend = NewBackend(s.mock, "mock", s.prefix)

	var err error
	s.store, err = New([]*Backend{s.backend}, s.backend, s.backend)
	s.Require().Nil(err)
}

// StoreValues

func (s *StoreValueSuite) TestStoreValuesValid() {
	s.mock.On("Get", s.prefix+"test").Return(&store.KVPair{Key: s.prefix + "test", Value: s.schemaBytes}, nil)
	s.mock.On("Put", s.prefix+"test/required", []byte("a"), &store.WriteOptions{}).Return(nil)

	errors := s.store.StoreValues("test", []byte(`{"required": "a"}`))
	s.Assert().Equal(0, len(errors), fmt.Sprintf("%v", errors))

	s.mock.AssertExpectations(s.T())
}

func (s *StoreValueSuite) TestStoreValuesInvalid() {
	s.mock.On("Get", s.prefix+"test").Return(&store.KVPair{Key: s.prefix + "test", Value: s.schemaBytes}, nil)

	errors := s.store.StoreValues("test", []byte(`{"required": 1}`))
	s.Assert().Equal(1, len(errors), fmt.Sprintf("%v", errors))

	s.mock.AssertExpectations(s.T())
}

func TestStoreValueSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(StoreValueSuite))
}
