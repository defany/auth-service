package auth

import (
	"log/slog"

	defserv "github.com/defany/auth-service/app/internal/service"
	authv1 "github.com/defany/auth-service/app/pkg/gen/proto/auth/v1"
)

type Implementation struct {
	authv1.UnimplementedAuthServiceServer

	log *slog.Logger

	service defserv.AuthService
}

func NewImplementation(log *slog.Logger, service defserv.AuthService) *Implementation {
	return &Implementation{
		log:     log,
		service: service,
	}
}
