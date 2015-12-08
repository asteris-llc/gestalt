package store

import (
	"errors"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type StoreSuite struct {
	suite.Suite

	prefix      string
	mock        *mock.Mock
	backend     *Backend
	store       *Store
	schemaBytes []byte
}

func (s *StoreSuite) SetupSuite() {
	s.prefix = "mock/"
	s.schemaBytes = []byte(`{"type": "string"}`)
}

func (s *StoreSuite) SetupTest() {
	s.prefix = "mock/"

	s.mock = &mock.Mock{}
	s.backend = NewBackend(s.mock, "mock", s.prefix)

	var err error
	s.store, err = New([]*Backend{s.backend}, s.backend, s.backend)
	s.Require().Nil(err)
}

// STORE

func (s *StoreSuite) TestStoreInvalid() {
	err := s.store.Store("invalid", []byte{})
	s.Assert().NotNil(err)

	s.mock.AssertNotCalled(s.T(), "Put", s.prefix+"invalid", "")
}

func (s *StoreSuite) TestStoreValid() {
	s.mock.On("Put", s.prefix+"valid", s.schemaBytes, &store.WriteOptions{}).Return(nil)

	err := s.store.Store("valid", s.schemaBytes)
	s.Assert().Nil(err)

	s.mock.AssertExpectations(s.T())
}

func (s *StoreSuite) TestStoreError() {
	err := errors.New("test")
	s.mock.On("Put", s.prefix+"valid", s.schemaBytes, &store.WriteOptions{}).Return(err)

	err2 := s.store.Store("valid", s.schemaBytes)
	s.Assert().Equal(err, err2)

	s.mock.AssertExpectations(s.T())
}

// RETRIEVE

func (s *StoreSuite) TestRetrievePresent() {
	s.mock.On("Get", s.prefix+"present").Return(&store.KVPair{Key: s.prefix + "present", Value: s.schemaBytes}, nil)

	schema, err := s.store.Retrieve("present")
	s.Assert().Nil(err)
	s.Assert().Equal(s.schemaBytes, schema)

	s.mock.AssertExpectations(s.T())
}

func (s *StoreSuite) TestRetrieveAbsent() {
	s.mock.On("Get", s.prefix+"absent").Return(&store.KVPair{}, nil)

	schema, err := s.store.Retrieve("absent")
	s.Assert().Equal(err, ErrNotFound)
	s.Assert().Equal(schema, []byte{})

	s.mock.AssertExpectations(s.T())
}

func (s *StoreSuite) TestRetrieveError() {
	err := errors.New("test")
	s.mock.On("Get", s.prefix+"error").Return(&store.KVPair{}, err)

	schema, err2 := s.store.Retrieve("error")
	s.Assert().Equal(err2, err)
	s.Assert().Equal(schema, []byte{})

	s.mock.AssertExpectations(s.T())
}

// DELETE

func (s *StoreSuite) TestDeletePresent() {
	s.mock.On("Delete", s.prefix+"present").Return(nil)

	err := s.store.Delete("present")
	s.Assert().Nil(err)

	s.mock.AssertExpectations(s.T())
}

func (s *StoreSuite) TestDeleteError() {
	err := errors.New("test")
	s.mock.On("Delete", s.prefix+"error").Return(err)

	err2 := s.store.Delete("error")
	s.Assert().Equal(err2, err)

	s.mock.AssertExpectations(s.T())
}

func TestStoreSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(StoreSuite))
}
