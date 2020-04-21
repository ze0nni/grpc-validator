package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewValidatorServer(t *testing.T) {
	s := NewValidatorServer()

	assert.NotNil(t, s)
}
