package interceptor

import (
	"context"

	"github.com/defany/auth-service/app/pkg/metrics"
	"google.golang.org/grpc"
)

type interceptor func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) error

type validator interface {
	Validate() error
	ValidateAll() error
}

type Interceptor struct {
	interceptors []interceptor
}

func NewInterceptor(interceptors ...interceptor) *Interceptor {
	return &Interceptor{
		interceptors: interceptors,
	}
}

func (i *Interceptor) With(interceptors ...interceptor) *Interceptor {
	i.interceptors = append(i.interceptors, interceptors...)

	return i
}

func (i *Interceptor) WithGRPCValidate() *Interceptor {
	i.interceptors = append(i.interceptors, func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) error {
		if v, ok := req.(validator); ok {
			if err := v.Validate(); err != nil {
				return err
			}
		}

		return nil
	})

	return i
}

func (i *Interceptor) WithRequestsCounter(m *metrics.Metrics) *Interceptor {
	i.interceptors = append(i.interceptors, func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) error {
		return m.IncRequestsCount()
	})

	return i
}

func (i *Interceptor) Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, server *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		for _, interceptor := range i.interceptors {
			if err := interceptor(ctx, req, server, handler); err != nil {
				return nil, err
			}
		}

		return handler(ctx, req)
	}
}
