package userservice

import (
	"context"
	"github.com/defany/auth-service/app/internal/model"
)

func (s *service) Update(ctx context.Context, user model.UserUpdate) error {
	return s.repo.Update(ctx, user)
}
