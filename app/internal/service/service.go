package defserv

import (
	"context"
	"github.com/defany/auth-service/app/internal/model"
)

type UserService interface {
	Get(ctx context.Context, id uint64) (model.User, error)
	Create(ctx context.Context, user model.UserCreate) (uint64, error)
	Update(ctx context.Context, user model.UserUpdate) error
	Delete(ctx context.Context, id uint64) error
}
