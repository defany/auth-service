package interceptor

import (
	"context"
	"time"

	"github.com/defany/auth-service/app/pkg/metrics"
	"google.golang.org/grpc"
)

type Interceptor struct{}

func NewInterceptor() *Interceptor {
	return &Interceptor{}
}

func (i *Interceptor) Interceptor(ctx context.Context, req any, server *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	startAt := time.Now()

	metrics.IncRequestCounter()

	status := "success"

	res, err := handler(ctx, req)
	if err != nil {
		status = "error"
	}

	metrics.HistogramResponseTimeObserve(status, server.FullMethod, time.Since(startAt).Seconds())
	metrics.IncResponseCounter(status, server.FullMethod)

	return res, err
}
