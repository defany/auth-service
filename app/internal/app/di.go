package app

import (
	"context"
	"github.com/defany/auth-service/app/internal/api/user"
	"github.com/defany/auth-service/app/internal/config"
	"github.com/defany/auth-service/app/internal/repository"
	userrepo "github.com/defany/auth-service/app/internal/repository/user"
	defserv "github.com/defany/auth-service/app/internal/service"
	userservice "github.com/defany/auth-service/app/internal/service/user"
	"github.com/defany/auth-service/app/pkg/closer"
	"github.com/defany/auth-service/app/pkg/logger/sl"
	"github.com/defany/auth-service/app/pkg/postgres"
	"log/slog"
	"os"
)

type di struct {
	log *slog.Logger

	cfg *config.Config

	repositories struct {
		user repository.UserRepository
	}

	services struct {
		user defserv.UserService
	}

	implementations struct {
		user *user.Implementation
	}

	txManager postgres.TxManager
	db        postgres.Querier
}

func newDI() *di {
	return &di{}
}

func (d *di) Log(ctx context.Context) *slog.Logger {
	if d.log != nil {
		return d.log
	}

	d.log = sl.NewSlogLogger(d.Config(ctx).Logger)

	return d.log
}

func (d *di) Config(_ context.Context) *config.Config {
	if d.cfg != nil {
		return d.cfg
	}

	d.cfg = config.MustLoad()

	return d.cfg
}

func (d *di) Database(ctx context.Context) postgres.Querier {
	if d.db != nil {
		return d.db
	}

	cfg := d.Config(ctx)

	dbConfig := postgres.NewConfig(cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Database)

	dbConfig.WithRetryConnDelay(cfg.Database.ConnectAttemptsDelay)
	dbConfig.WithMaxConnAttempts(cfg.Database.ConnectAttempts)

	db, err := postgres.NewPostgres(ctx, d.Log(ctx), dbConfig)
	if err != nil {
		d.Log(ctx).Error("failed to connect to database", sl.ErrAttr(err))

		os.Exit(1)
	}

	closer.Add(func() error {
		db.Close()

		return nil
	})

	d.db = db

	return d.db
}

func (d *di) TxManager(ctx context.Context) postgres.TxManager {
	if d.txManager != nil {
		return d.txManager
	}

	d.txManager = postgres.NewTxManager(d.Database(ctx))

	return d.txManager
}

func (d *di) UserRepo(ctx context.Context) repository.UserRepository {
	if d.repositories.user != nil {
		return d.repositories.user
	}

	d.repositories.user = userrepo.NewRepository(d.Database(ctx))

	return d.repositories.user
}

func (d *di) UserService(ctx context.Context) defserv.UserService {
	if d.services.user != nil {
		return d.services.user
	}

	d.services.user = userservice.NewService(d.TxManager(ctx), d.UserRepo(ctx))

	return d.services.user
}

func (d *di) UserImpl(ctx context.Context) *user.Implementation {
	if d.implementations.user != nil {
		return d.implementations.user
	}

	d.implementations.user = user.NewImplementation(d.Log(ctx), d.UserService(ctx))

	return d.implementations.user
}
