package userservice

import (
	"context"
	"github.com/defany/auth-service/app/internal/model"
	"github.com/defany/slogger/pkg/logger/sl"
)

func (s *service) Get(ctx context.Context, id uint64) (model.User, error) {
	op := sl.FnName()

	var user model.User

	err := s.tx.ReadCommitted(ctx, func(ctx context.Context) error {
		u, err := s.repo.User(ctx, id)
		if err != nil {
			return err
		}

		user = u

		err = s.log.Log(ctx, model.Log{
			Action: model.LogGetUser,
			UserID: id,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return model.User{}, sl.Err(op, err)
	}

	return user, nil
}
