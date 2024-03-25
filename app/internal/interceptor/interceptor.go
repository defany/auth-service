package interceptor

import (
	"context"

	"github.com/defany/auth-service/app/pkg/metrics"
	"google.golang.org/grpc"
)

type interceptor func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) error
type deferInterceptor func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler, res any, reqErr error) error

type validator interface {
	Validate() error
	ValidateAll() error
}

func Interceptor(metrics *metrics.Metrics) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if err := metrics.IncRequestsCount(); err != nil {
			return nil, err
		}

		if v, ok := req.(validator); ok {
			if err := v.Validate(); err != nil {
				return nil, err
			}
		}

		res, err := handler(ctx, req)
		if err != nil {
			if err := metrics.IncResponsesCount("error", info.FullMethod); err != nil {
				return nil, err
			}
		} else {
			if err := metrics.IncResponsesCount("success", info.FullMethod); err != nil {
				return nil, err
			}
		}

		return res, err
	}
}
