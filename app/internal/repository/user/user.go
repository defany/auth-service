package userrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/defany/auth-service/app/internal/model"
	"github.com/defany/auth-service/app/internal/repository/user/converter"
	storemodel "github.com/defany/auth-service/app/internal/repository/user/model"
	"github.com/jackc/pgx/v5"
)

func (r *repository) User(ctx context.Context, id uint64) (model.User, error) {
	q := r.qb.Select(idColumn, emailColumn, nameColumn, roleColumn, createdAtColumn, updatedAtColumn, passwordColumn).
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

func (r *repository) UserByNickname(ctx context.Context, nickname string) (model.User, error) {
	q := r.qb.Select(idColumn, emailColumn, nameColumn, roleColumn, createdAtColumn, updatedAtColumn, passwordColumn).
		From(table).
		Where(squirrel.Eq{
			nameColumn: nickname,
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

func (r *repository) DoesHaveAccess(ctx context.Context, userRole string, endpoint string) error {
	q := fmt.Sprintf("select exists(select endpoint from %s where role = $1 and endpoint = $2)", tableEndpointsPermissions)

	rows, err := r.db.Query(ctx, q, userRole, endpoint)
	if err != nil {
		return err
	}

	doesHaveAccess, err := pgx.CollectOneRow(rows, pgx.RowTo[bool])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}

		return err
	}

	if !doesHaveAccess {
		return errors.New("user doesn't have access to this resource")
	}

	return nil
}
