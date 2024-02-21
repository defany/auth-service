package userservice

import (
	"context"
	"github.com/defany/auth-service/app/internal/model"
)

func (s *service) Create(ctx context.Context, user model.UserCreate) (int, error) {
	return s.repo.Create(ctx, user)
}
