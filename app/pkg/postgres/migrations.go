package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"os"
)

type Migrator struct {
	provider *goose.Provider
}

func NewMigrator(db *pgxpool.Pool, migrationsDir string) (*Migrator, error) {
	provider, err := goose.NewProvider(goose.DialectPostgres, stdlib.OpenDBFromPool(db), os.DirFS(migrationsDir))
	if err != nil {
		return nil, err
	}

	return &Migrator{
		provider: provider,
	}, nil
}

func (u *Migrator) Up(ctx context.Context) ([]*goose.MigrationResult, error) {
	res, err := u.provider.Up(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}
