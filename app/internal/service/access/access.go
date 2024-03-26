package accessservice

import (
	"time"

	"github.com/defany/auth-service/app/internal/repository"
)

type JWTConfig interface {
	RefreshSecretKey() []byte
	AccessSecretKey() []byte
	RefreshTokenDuration() time.Duration
	AccessTokenDuration() time.Duration
}

type Service struct {
	userRepo repository.UserRepository

	jwtConfig JWTConfig
}

func NewService(userRepo repository.UserRepository, jwtConfig JWTConfig) *Service {
	return &Service{
		userRepo:  userRepo,
		jwtConfig: jwtConfig,
	}
}
