package userrepo

import (
	"context"
	"github.com/Masterminds/squirrel"
)

func (r *repository) Delete(ctx context.Context, id uint64) error {
	q := r.qb.Delete(table).Where(squirrel.Eq{
		idColumn: id,
	})

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
