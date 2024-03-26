package repository

import (
	"context"

	"github.com/defany/auth-service/app/internal/model"
)

type UserRepository interface {
	User(ctx context.Context, id uint64) (model.User, error)
	UserByNickname(ctx context.Context, nickname string) (model.User, error)
	Create(ctx context.Context, user model.UserCreate) (uint64, error)
	Update(ctx context.Context, user model.UserUpdate) error
	Delete(ctx context.Context, id uint64) error
	DoesHaveAccess(ctx context.Context, userRole string, endpoint string) error
}

type LogRepository interface {
	Log(ctx context.Context, log model.Log) error
}
