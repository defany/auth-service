package userservice

import (
	"context"
	"github.com/defany/auth-service/app/internal/model"
	"github.com/defany/slogger/pkg/logger/sl"
)

func (s *service) Update(ctx context.Context, user model.UserUpdate) error {
	op := sl.FnName()

	err := s.tx.ReadCommitted(ctx, func(ctx context.Context) error {
		if err := s.repo.Update(ctx, user); err != nil {
			return err
		}

		err := s.log.Log(ctx, model.Log{
			Action: model.LogUpdateUser,
			UserID: user.ID,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return sl.Err(op, err)
	}

	return nil
}
