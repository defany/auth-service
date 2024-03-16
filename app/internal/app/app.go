package app

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net"
	"net/http"

	"github.com/defany/auth-service/app/internal/interceptor"
	"github.com/defany/auth-service/app/pkg/closer"
	userv1 "github.com/defany/auth-service/app/pkg/gen/proto/user/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	di         *DI
	grpcServer *grpc.Server
	httpServer *http.Server
}

func NewApp() *App {
	a := &App{}

	return a
}

func (a *App) Run(ctx context.Context) error {
	defer func() {
		closer.Close()

		a.DI().Log(ctx).Info("application closed :9")
	}()

	a.registerUserService(ctx)

	wg, ctx := errgroup.WithContext(ctx)

	wg.Go(func() error {
		return a.runHTTPServer(ctx)
	})

	wg.Go(func() error {
		return a.runGRPCServer(ctx)
	})

	return wg.Wait()
}

func (a *App) DI() *DI {
	if a.di != nil {
		return a.di
	}

	a.di = newDI()

	return a.di
}

func (a *App) runGRPCServer(ctx context.Context) error {
	lis, err := net.Listen("tcp", a.DI().Config(ctx).Server.GRPC.Addr)
	if err != nil {
		return err
	}

	a.DI().Log(ctx).Info("Go grpc!", slog.String("addr", a.DI().Config(ctx).Server.GRPC.Addr))

	closer.Add(func() error {
		a.DI().Log(ctx).Info("closing grpc server")

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

func (a *App) runHTTPServer(ctx context.Context) error {
	if a.httpServer == nil {
		if err := a.setupHTTPServer(ctx); err != nil {
			return err
		}
	}

	a.DI().Log(ctx).Info("Go http!", slog.String("addr", a.DI().Config(ctx).Server.HTTP.Addr))

	closer.Add(func() error {
		a.DI().Log(ctx).Info("closing http server")

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
	mux := runtime.NewServeMux(runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
		log.Println(err)
	}))

	serverConfig := a.DI().Config(ctx).Server

	grpcAddr := serverConfig.GRPC.Addr
	httpAddr := serverConfig.HTTP.Addr

	err := userv1.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcAddr, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		return err
	}

	a.httpServer = &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	return nil
}

func (a *App) registerUserService(ctx context.Context) {
	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(interceptor.GRPCValidate),
	)

	reflection.Register(a.grpcServer)

	userv1.RegisterUserServiceServer(a.grpcServer, a.DI().UserImpl(ctx))

	return
}
