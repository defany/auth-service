package userservice

import (
	"context"
	"github.com/defany/auth-service/app/internal/model"
)

func (s *service) Get(ctx context.Context, id int) (model.User, error) {
	return s.repo.User(ctx, id)
}
