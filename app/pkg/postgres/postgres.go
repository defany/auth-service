package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type Querier interface {
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, query string, args ...interface{}) (commandTag pgconn.CommandTag, err error)

	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)

	Close()
}

type postgres struct {
	log *slog.Logger

	db *pgxpool.Pool
}

func NewPostgres(ctx context.Context, log *slog.Logger, cfg *Config) (Querier, error) {
	p := &postgres{
		log: log,
	}

	pool, err := NewClient(ctx, log, cfg)
	if err != nil {
		return nil, err
	}

	p.db = pool

	return p, nil
}

func (p *postgres) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	tx, ok := ExtractTX(ctx)
	if ok {
		return tx.Query(ctx, query, args...)
	}

	return p.db.Query(ctx, query, args...)
}

func (p *postgres) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	tx, ok := ExtractTX(ctx)
	if ok {
		return tx.QueryRow(ctx, query, args...)
	}

	return p.db.QueryRow(ctx, query, args...)
}

func (p *postgres) Exec(ctx context.Context, query string, args ...interface{}) (commandTag pgconn.CommandTag, err error) {
	tx, ok := ExtractTX(ctx)
	if ok {
		return tx.Exec(ctx, query, args...)
	}

	return p.db.Exec(ctx, query, args...)
}

func (p *postgres) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.db.BeginTx(ctx, txOptions)
}

func (p *postgres) Close() {
	p.db.Close()
}
