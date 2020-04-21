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
	strt := in.GetStruct()
	for k, v := range in.GetFilter() {
		if ("foo" == k && strt.Foo != v) ||
			("bar" == k && strt.Bar != v) ||
			("baz" == k && strt.Baz != v) {

			return &api.ValidateResult{
				Success: false,
			}, nil
		}
	}

	return &api.ValidateResult{
		Success: true,
	}, nil
}
