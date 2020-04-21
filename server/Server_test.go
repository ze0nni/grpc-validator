package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ze0nni/grpc-validator/api"
)

func TestNewValidatorServer(t *testing.T) {
	s := NewValidatorServer()

	assert.NotNil(t, s)
}

func TestValidateSuccesForEmptyMap(t *testing.T) {
	s := NewValidatorServer()
	req := &api.ValidateRequest{
		Struct: &api.Struct{
			Foo: "foo",
			Bar: "bar",
			Baz: "baz",
		},
		Filter: make(map[string]string),
	}
	res, err := s.A(nil, req)

	assert.NoError(t, err)
	assert.True(t, res.GetSuccess())
}

func TestSuccesForFiltersKeyNotContaintsStructKeys(t *testing.T) {
	s := NewValidatorServer()

	strt := &api.Struct{
		Foo: "foo",
		Bar: "bar",
		Baz: "baz",
	}
	filter := make(map[string]string)
	filter["bat"] = "bat"

	req := &api.ValidateRequest{
		Struct: strt,
		Filter: filter,
	}
	res, err := s.A(nil, req)

	assert.NoError(t, err)
	assert.True(t, res.GetSuccess())
}

func TestNotSuccesWhenFooValueNotMatch(t *testing.T) {
	s := NewValidatorServer()

	strt := &api.Struct{
		Foo: "foo",
		Bar: "bar",
		Baz: "baz",
	}
	filter := make(map[string]string)
	filter["foo"] = "bar"

	req := &api.ValidateRequest{
		Struct: strt,
		Filter: filter,
	}
	res, err := s.A(nil, req)

	assert.NoError(t, err)
	assert.False(t, res.GetSuccess())
}

func TestSuccesWhenMapContaintsAllKeys(t *testing.T) {
	s := NewValidatorServer()

	strt := &api.Struct{
		Foo: "foo",
		Bar: "bar",
		Baz: "baz",
	}
	filter := make(map[string]string)
	filter["foo"] = "foo"
	filter["bar"] = "bar"
	filter["baz"] = "baz"

	req := &api.ValidateRequest{
		Struct: strt,
		Filter: filter,
	}
	res, err := s.A(nil, req)

	assert.NoError(t, err)
	assert.True(t, res.GetSuccess())
}

func TestNotSuccesWhenOneKeyNotMatch(t *testing.T) {
	s := NewValidatorServer()

	strt := &api.Struct{
		Foo: "foo",
		Bar: "bar",
		Baz: "baz",
	}
	filter := make(map[string]string)
	filter["foo"] = "foo"
	filter["bar"] = "error"
	filter["baz"] = "baz"

	req := &api.ValidateRequest{
		Struct: strt,
		Filter: filter,
	}
	res, err := s.A(nil, req)

	assert.NoError(t, err)
	assert.False(t, res.GetSuccess())
}
