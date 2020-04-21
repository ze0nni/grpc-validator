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

func (server) A(ctx context.Context, in *api.ValidateRequest) (*api.ValidateResult, error) {
	return &api.ValidateResult{
		Success: true,
	}, nil
}
