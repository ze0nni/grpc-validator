package server

import (
	"context"

	"github.com/ze0nni/grpc-validator/api"
)

// NewValidatorServer returns new instanc of 'api.ValidatorServer
func NewValidatorServer() api.ValidatorServer {
	return &server{}
}

type server struct {
}

func (server) A(context.Context, *api.ValidateRequest) (*api.ValidateResult, error) {
	panic("not implements")
}
