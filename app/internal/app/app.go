package app

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"

	"github.com/defany/auth-service/app/internal/config"
	accessv1 "github.com/defany/auth-service/app/pkg/gen/proto/access/v1"
	authv1 "github.com/defany/auth-service/app/pkg/gen/proto/auth/v1"
	"github.com/defany/platcom/pkg/closer"
	"github.com/defany/platcom/pkg/swagger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"

	"github.com/defany/auth-service/app/internal/interceptor"
	_ "github.com/defany/auth-service/app/pkg/gen/gen-swagger/statik"
	userv1 "github.com/defany/auth-service/app/pkg/gen/proto/user/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	di               *DI
	grpcServer       *grpc.Server
	httpServer       *http.Server
	swaggerServer    *http.Server
	prometheusServer *http.Server
}

func NewApp() *App {
	a := &App{}

	return a
}

func (a *App) Run(ctx context.Context) error {
	defer func() {
		closer.Close()

		a.DI(ctx).Log(ctx).Info("application closed :9")
	}()

	wg, ctx := errgroup.WithContext(ctx)

	a.DI(ctx).SetupMetrics(ctx)

	wg.Go(func() error {
		return a.runSwaggerHTTPServer(ctx)
	})

	wg.Go(func() error {
		return a.runHTTPServer(ctx)
	})

	wg.Go(func() error {
		return a.runGRPCServer(ctx)
	})

	wg.Go(func() error {
		return a.runPrometheusServer(ctx)
	})

	return wg.Wait()
}

func (a *App) DI(ctx context.Context) *DI {
	if a.di != nil {
		return a.di
	}

	a.di = newDI(ctx)

	return a.di
}

func (a *App) runGRPCServer(ctx context.Context) error {
	if a.grpcServer == nil {
		a.setupGRPCServer(ctx)
	}

	lis, err := net.Listen("tcp", a.DI(ctx).Config(ctx).Server.GRPC.Addr())
	if err != nil {
		return err
	}

	a.DI(ctx).Log(ctx).Info("go grpc!", slog.String("addr", a.DI(ctx).Config(ctx).Server.GRPC.Addr()))

	closer.Add(func() error {
		a.DI(ctx).Log(ctx).Info("closing grpc server")

		a.grpcServer.GracefulStop()

		return nil
	})

	if err := a.grpcServer.Serve(lis); err != nil {
		if errors.Is(err, grpc.ErrServerStopped) {
			return nil
		}

		return err
	}

	return nil
}

func (a *App) setupGRPCServer(ctx context.Context) {
	itcp := interceptor.NewInterceptor()

	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(itcp.MetricsInterceptor),
	)

	reflection.Register(a.grpcServer)

	authv1.RegisterAuthServiceServer(a.grpcServer, a.DI(ctx).AuthImpl(ctx))
	userv1.RegisterUserServiceServer(a.grpcServer, a.DI(ctx).UserImpl(ctx))
	accessv1.RegisterAccessServiceServer(a.grpcServer, a.DI(ctx).AccessImpl(ctx))

	return
}

func (a *App) runHTTPServer(ctx context.Context) error {
	if a.httpServer == nil {
		if err := a.setupHTTPServer(ctx); err != nil {
			return err
		}
	}

	a.DI(ctx).Log(ctx).Info("go http!", slog.String("addr", a.DI(ctx).Config(ctx).Server.HTTP.Addr()))

	closer.Add(func() error {
		a.DI(ctx).Log(ctx).Info("closing http server")

		return a.httpServer.Shutdown(ctx)
	})

	if err := a.httpServer.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}

func (a *App) setupHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	serverConfig := a.DI(ctx).Config(ctx).Server

	grpcAddr := serverConfig.GRPC.Addr()
	httpAddr := serverConfig.HTTP.Addr()

	err := userv1.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcAddr, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		return err
	}

	corsMidd := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	a.httpServer = &http.Server{
		Addr:    httpAddr,
		Handler: corsMidd.Handler(mux),
	}

	return nil
}

func (a *App) runSwaggerHTTPServer(ctx context.Context) error {
	if a.swaggerServer == nil {
		if err := a.setupSwaggerHTTPServer(ctx); err != nil {
			return err
		}
	}

	a.DI(ctx).Log(ctx).Info("go swagger http!", slog.String("addr", a.DI(ctx).Config(ctx).Server.Swagger.Addr()))

	closer.Add(func() error {
		a.DI(ctx).Log(ctx).Info("closing swagger http server")

		return a.swaggerServer.Shutdown(ctx)
	})

	if err := a.swaggerServer.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}

func (a *App) setupSwaggerHTTPServer(ctx context.Context) error {
	sfs, err := fs.New()
	if err != nil {
		return err
	}

	serve := swagger.NewServe("/api.swagger.json")

	if a.DI(ctx).Config(ctx).Env == config.EnvDev {
		serve.WithHost(a.DI(ctx).Config(ctx).Server.HTTP.Addr())
	}

	serve.WithLogger(a.DI(ctx).Log(ctx))

	if err := serve.Setup(); err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(sfs)))
	mux.HandleFunc("/api.swagger.json", serve.Middleware())

	a.swaggerServer = &http.Server{
		Addr:    a.DI(ctx).Config(ctx).Server.Swagger.Addr(),
		Handler: mux,
	}

	return nil
}

func (a *App) runPrometheusServer(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	a.prometheusServer = &http.Server{
		Addr:    a.DI(ctx).Config(ctx).Server.Prometheus.Addr(),
		Handler: mux,
	}

	closer.Add(func() error {
		a.DI(ctx).Log(ctx).Info("closing prometheus http server")

		return a.prometheusServer.Shutdown(ctx)
	})

	a.DI(ctx).Log(ctx).Info("go prometheus!", slog.String("addr", a.DI(ctx).Config(ctx).Server.Prometheus.Addr()))

	err := a.prometheusServer.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}
