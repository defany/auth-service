package userservice

import (
	"context"

	"github.com/defany/auth-service/app/internal/model"
	"github.com/defany/slogger/pkg/logger/sl"
)

func (s *service) Delete(ctx context.Context, id uint64) error {
	op := sl.FnName()

	err := s.tx.ReadCommitted(ctx, func(ctx context.Context) error {
		if err := s.repo.Delete(ctx, id); err != nil {
			return err
		}

		err := s.log.Log(ctx, model.Log{
			Action: model.LogDeleteUser,
			UserID: id,
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
