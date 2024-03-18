package logrepo

import (
	"context"

	"github.com/defany/auth-service/app/internal/model"
)

func (r *repository) Log(ctx context.Context, log model.Log) error {
	q := r.qb.Insert(logs).
		Columns(logsAction, logsUserID).
		Values(log.Action, log.UserID)

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
