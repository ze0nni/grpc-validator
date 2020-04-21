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
