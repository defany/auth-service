package userrepo

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/defany/auth-service/app/internal/model"
	"github.com/defany/auth-service/app/internal/repository/user/converter"
	storemodel "github.com/defany/auth-service/app/internal/repository/user/model"
	"github.com/jackc/pgx/v5"
)

func (r *repository) User(ctx context.Context, id uint64) (model.User, error) {
	q := r.qb.Select(idColumn, emailColumn, nameColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(table).
		Where(squirrel.Eq{
			idColumn: id,
		})

	sql, args, err := q.ToSql()
	if err != nil {
		return model.User{}, err
	}

	rows, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		return model.User{}, err
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[storemodel.User])
	if err != nil {
		return model.User{}, err
	}

	return converter.UserToModel(user), nil
}
