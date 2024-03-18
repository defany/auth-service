package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

type validator interface {
	Validate() error
	ValidateAll() error
}

func GRPCValidate(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	if v, ok := req.(validator); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	return handler(ctx, req)
}
