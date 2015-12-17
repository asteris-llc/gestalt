package store

import (
	"encoding/json"
	"errors"
	"github.com/asteris-llc/gestalt/web/app"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type StoreSchemaSuite struct {
	suite.Suite

	prefix      string
	mock        *mock.Mock
	backend     *Backend
	store       *Store
	schema      *app.Schema
	schemaBytes []byte
}

func (s *StoreSchemaSuite) SetupSuite() {
	s.prefix = "mock/"
}

func (s *StoreSchemaSuite) SetupTest() {
	s.mock = &mock.Mock{}
	s.backend = NewBackend(s.mock, "mock", s.prefix)

	s.schema = &app.Schema{
		Backend: "consul",
		Name:    "valid",
	}
	var err error
	s.schemaBytes, err = json.Marshal(s.schema)
	s.Require().Nil(err)

	s.store, err = New([]*Backend{s.backend}, s.backend, s.backend)
	s.Require().Nil(err)
}

// STORE

func (s *StoreSchemaSuite) TestStoreSchemaValid() {
	s.mock.On("Put", s.prefix+"valid", s.schemaBytes, &store.WriteOptions{}).Return(nil)

	err := s.store.StoreSchema("valid", s.schema)
	s.Assert().Nil(err)

	s.mock.AssertExpectations(s.T())
}

func (s *StoreSchemaSuite) TestStoreSchemaNameChange() {
	s.mock.On("Put", s.prefix+"valid", s.schemaBytes, &store.WriteOptions{}).Return(nil)

	s.schema.Name = "something-else"

	err := s.store.StoreSchema("valid", s.schema)
	s.Assert().Nil(err)

	s.mock.AssertExpectations(s.T())
}

func (s *StoreSchemaSuite) TestStoreSchemaError() {
	err := errors.New("test")
	s.mock.On("Put", s.prefix+"valid", s.schemaBytes, &store.WriteOptions{}).Return(err)

	err2 := s.store.StoreSchema("valid", s.schema)
	s.Assert().Equal(err, err2)

	s.mock.AssertExpectations(s.T())
}

// RETRIEVE

func (s *StoreSchemaSuite) TestRetrieveSchemaPresent() {
	s.mock.On("Get", s.prefix+"present").Return(&store.KVPair{Key: s.prefix + "present", Value: s.schemaBytes}, nil)

	schema, err := s.store.RetrieveSchema("present")
	s.Assert().Nil(err)
	s.Assert().Equal(s.schema, schema)

	s.mock.AssertExpectations(s.T())
}

func (s *StoreSchemaSuite) TestRetrieveSchemaAbsent() {
	s.mock.On("Get", s.prefix+"absent").Return(&store.KVPair{}, nil)

	schema, err := s.store.RetrieveSchema("absent")
	s.Assert().Equal(err, ErrMissingKey)
	s.Assert().Nil(schema)

	s.mock.AssertExpectations(s.T())
}

func (s *StoreSchemaSuite) TestRetrieveSchemaError() {
	err := errors.New("test")
	s.mock.On("Get", s.prefix+"error").Return(&store.KVPair{}, err)

	schema, err2 := s.store.RetrieveSchema("error")
	s.Assert().Equal(err2, err)
	s.Assert().Nil(schema)

	s.mock.AssertExpectations(s.T())
}

// DELETE

func (s *StoreSchemaSuite) TestDeleteSchemaPresent() {
	s.mock.On("Delete", s.prefix+"present").Return(nil)

	err := s.store.DeleteSchema("present")
	s.Assert().Nil(err)

	s.mock.AssertExpectations(s.T())
}

func (s *StoreSchemaSuite) TestDeleteSchemaError() {
	err := errors.New("test")
	s.mock.On("Delete", s.prefix+"error").Return(err)

	err2 := s.store.DeleteSchema("error")
	s.Assert().Equal(err2, err)

	s.mock.AssertExpectations(s.T())
}

func TestStoreSchemaSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(StoreSchemaSuite))
}
