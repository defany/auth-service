package main

import (
	context "context"
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/defany/auth-service/app/pkg/gen/auth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const port = 50001

type server struct {
	authv1.UnimplementedAuthServer
}

func (s *server) Create(ctx context.Context, request *authv1.CreateUserRequest) (*authv1.CreateUserResponse, error) {
	log := slog.With(
		slog.String("name", request.GetName()),
		slog.String("email", request.GetEmail()),
		slog.String("password", request.GetPassword()),
		slog.String("password_confirm", request.GetPasswordConfirm()),
		slog.String("role", request.GetRole().String()),
	)

	log.Info("create user request")

	return &authv1.CreateUserResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Get(ctx context.Context, request *authv1.GetUserRequest) (*authv1.GetUserResponse, error) {
	log := slog.With(
		slog.Int64("id", request.GetId()),
	)

	log.Info("get user request")

	resData := authv1.GetUserResponse{
		Id:              request.GetId(),
		Name:            gofakeit.Name(),
		Email:           gofakeit.Email(),
		Password:        gofakeit.Password(false, true, true, true, false, 6),
		PasswordConfirm: gofakeit.BeerName(),
		Role:            authv1.UserRole_ADMIN,
		CreatedAt:       timestamppb.Now(),
		UpdatedAt:       nil,
	}

	log.Info("user info", slog.Any("info", &resData))

	return &resData, nil
}

func (s *server) Update(ctx context.Context, request *authv1.UpdateUserRequest) (*emptypb.Empty, error) {
	log := slog.With(
		slog.Int64("id", request.GetId()),
		slog.String("email", request.GetEmail().GetValue()),
		slog.String("name", request.GetName().GetValue()),
		slog.String("role", request.GetRole().String()),
	)

	log.Info("update user request")

	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, request *authv1.DeleteUserRequest) (*emptypb.Empty, error) {
	log := slog.With(
		slog.Int64("id", request.GetId()),
	)

	log.Info("delete user request")

	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		slog.Error("failed to listen: %v", err)
		os.Exit(1)
	}

	s := grpc.NewServer()

	reflection.Register(s)

	authv1.RegisterAuthServer(s, &server{})

	slog.Info("listening", slog.String("port", lis.Addr().String()))

	if err := s.Serve(lis); err != nil {
		slog.Error("failed to serve: %v", err)
		os.Exit(1)
	}
}
