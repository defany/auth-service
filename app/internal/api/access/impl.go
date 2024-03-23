package access

import (
	"log/slog"

	defserv "github.com/defany/auth-service/app/internal/service"
	accessv1 "github.com/defany/auth-service/app/pkg/gen/proto/access/v1"
)

type Implementation struct {
	accessv1.UnimplementedAccessServiceServer

	log *slog.Logger

	service defserv.AccessService
}

func NewImplementation(log *slog.Logger, service defserv.AccessService) *Implementation {
	return &Implementation{
		log:     log,
		service: service,
	}
}
