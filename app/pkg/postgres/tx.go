package postgres

import (
	"context"
	errs "errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

type txKey struct{}

type TxManager interface {
	ReadCommitted(ctx context.Context, handler Handler) error
}

type Handler func(ctx context.Context) error

type txManager struct {
	db Postgres
}

func NewTxManager(db Postgres) TxManager {
	return &txManager{
		db: db,
	}
}

func (t *txManager) tx(ctx context.Context, opts pgx.TxOptions, handler Handler) (err error) {
	tx, ok := ExtractTX(ctx)
	if ok {
		return handler(ctx)
	}

	tx, err = t.db.BeginTx(ctx, opts)
	if err != nil {
		return errors.Wrap(err, "failed to start tx")
	}

	ctx = InjectTX(ctx, tx)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %w -> %v", err, r)
		}

		if err != nil {
			if txErr := tx.Rollback(ctx); txErr != nil {
				err = errs.Join(txErr, err)
			}

			return
		}

		if txErr := tx.Commit(ctx); txErr != nil {
			err = txErr
		}
	}()

	if err := handler(ctx); err != nil {
		return err
	}

	return nil
}

func (t *txManager) ReadCommitted(ctx context.Context, handler Handler) error {
	opts := pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}

	return t.tx(ctx, opts, handler)
}

func InjectTX(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func ExtractTX(ctx context.Context) (pgx.Tx, bool) {
	tx, ok := ctx.Value(txKey{}).(pgx.Tx)

	return tx, ok
}
