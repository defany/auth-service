package repository

import (
	"context"
	"github.com/defany/auth-service/app/internal/model"
)

type UserRepository interface {
	User(ctx context.Context, id int) (model.User, error)
	Create(ctx context.Context, user model.UserCreate) (int, error)
	Update(ctx context.Context, user model.UserUpdate) error
	Delete(ctx context.Context, id int) error
}
