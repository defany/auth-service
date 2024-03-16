package interceptor

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

type validator interface {
	Validate() error
	ValidateAll() error
}

func GRPCValidate(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	if v, ok := req.(validator); ok {
		if err := v.Validate(); err != nil {
			log.Println("VALIDATE ERR", err)

			return nil, err
		}
	}

	log.Println(req)

	return handler(ctx, req)
}
