package user

import (
	defserv "github.com/defany/auth-service/app/internal/service"
	userv1 "github.com/defany/auth-service/app/pkg/gen/user/v1"
	"log/slog"
)

type Implementation struct {
	userv1.UnimplementedUserServiceServer

	log *slog.Logger

	service defserv.UserService
}

func NewImplementation(log *slog.Logger, service defserv.UserService) *Implementation {
	return &Implementation{
		log:     log,
		service: service,
	}
}
