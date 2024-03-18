package userrepo

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/defany/auth-service/app/internal/model"
)

func (r *repository) Update(ctx context.Context, user model.UserUpdate) error {
	q := r.qb.Update(table).Where(squirrel.Eq{
		idColumn: user.ID,
	})

	if user.Name != nil {
		q = q.Set(nameColumn, user.Name)
	}

	if user.Email != nil {
		q = q.Set(emailColumn, user.Email)
	}

	if user.Role != nil {
		q = q.Set(roleColumn, user.Role)
	}

	sql, args, err := q.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
