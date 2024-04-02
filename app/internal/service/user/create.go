package userservice

import (
	"context"

	"github.com/defany/auth-service/app/internal/model"
	"github.com/defany/slogger/pkg/logger/sl"
)

func (s *service) Create(ctx context.Context, user model.UserCreate) (uint64, error) {
	var userID uint64

	password, err := s.passHasher.GenerateFromPassword([]byte(user.Password))
	if err != nil {
		return 0, err
	}

	user.Password = string(password)

	err = s.tx.ReadCommitted(ctx, func(ctx context.Context) error {
		id, err := s.repo.Create(ctx, user)
		if err != nil {
			return err
		}

		userID = id

		err = s.log.Log(ctx, model.Log{
			Action: model.LogCreateUser,
			UserID: userID,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return 0, sl.Err(sl.FnName(), err)
	}

	return userID, nil
}
