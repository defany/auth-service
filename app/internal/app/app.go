package app

import (
	"context"
	"fmt"
	"net"

	"github.com/defany/auth-service/app/pkg/closer"
	userv1 "github.com/defany/auth-service/app/pkg/gen/proto/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	di         *DI
	grpcServer *grpc.Server
}

func NewApp() *App {
	a := &App{}

	a.setupDI()

	return a
}

func (a *App) Run(ctx context.Context) error {
	defer func() {
		a.di.Log(ctx).Info("closing application... :(")

		closer.Close()

		a.di.Log(ctx).Info("application closed")
	}()

	a.setupDI()

	a.registerUserService(ctx)

	return a.runGRPCServer(ctx)
}

func (a *App) DI() *DI {
	return a.di
}

func (a *App) setupDI() {
	a.di = newDI()
}

func (a *App) runGRPCServer(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.di.Config(ctx).Server.Port))
	if err != nil {
		return err
	}

	a.di.Log(ctx).Info("Go!")

	if err := a.grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (a *App) registerUserService(ctx context.Context) {
	a.grpcServer = grpc.NewServer()
	reflection.Register(a.grpcServer)

	userv1.RegisterUserServiceServer(a.grpcServer, a.di.UserImpl(ctx))

	return
}
