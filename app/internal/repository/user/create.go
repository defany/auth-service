package userrepo

import (
	"context"
	"fmt"

	"github.com/defany/auth-service/app/internal/model"
	"github.com/jackc/pgx/v5"
)

func (r *repository) Create(ctx context.Context, user model.UserCreate) (uint64, error) {
	q := r.qb.Insert(table).
		Columns(nicknameColumn, emailColumn, passwordColumn, roleColumn).
		Values(user.Name, user.Email, user.Password, user.Role).
		Suffix(fmt.Sprintf("returning %s", idColumn))

	sql, args, err := q.ToSql()
	if err != nil {
		return 0, err
	}

	rows, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		return 0, err
	}

	id, err := pgx.CollectOneRow(rows, pgx.RowTo[uint64])
	if err != nil {
		return 0, err
	}

	return id, nil
}
