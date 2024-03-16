package app

import (
	"context"
	"log/slog"
	"os"

	"github.com/defany/auth-service/app/internal/api/user"
	"github.com/defany/auth-service/app/internal/config"
	"github.com/defany/auth-service/app/internal/repository"
	logrepo "github.com/defany/auth-service/app/internal/repository/log"
	userrepo "github.com/defany/auth-service/app/internal/repository/user"
	defserv "github.com/defany/auth-service/app/internal/service"
	userservice "github.com/defany/auth-service/app/internal/service/user"
	"github.com/defany/auth-service/app/pkg/closer"
	"github.com/defany/db/pkg/postgres"
	"github.com/defany/slogger/pkg/logger/sl"
)

type DI struct {
	log *slog.Logger

	cfg *config.Config

	repositories struct {
		user repository.UserRepository
		log  repository.LogRepository
	}

	services struct {
		user defserv.UserService
	}

	implementations struct {
		user *user.Implementation
	}

	txManager postgres.TxManager
	db        postgres.Postgres
}

func newDI() *DI {
	return &DI{}
}

func (d *DI) Log(ctx context.Context) *slog.Logger {
	if d.log != nil {
		return d.log
	}

	d.log = sl.NewSlogLogger(d.Config(ctx).Logger)

	return d.log
}

func (d *DI) Config(_ context.Context) *config.Config {
	if d.cfg != nil {
		return d.cfg
	}

	d.cfg = config.MustLoad()

	return d.cfg
}

func (d *DI) Database(ctx context.Context) postgres.Postgres {
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
		d.Log(ctx).Info("closing db connection pool")

		db.Close()

		return nil
	})

	d.db = db

	return d.db
}

func (d *DI) TxManager(ctx context.Context) postgres.TxManager {
	if d.txManager != nil {
		return d.txManager
	}

	d.txManager = postgres.NewTxManager(d.Database(ctx))

	return d.txManager
}

func (d *DI) UserRepo(ctx context.Context) repository.UserRepository {
	if d.repositories.user != nil {
		return d.repositories.user
	}

	d.repositories.user = userrepo.NewRepository(d.Database(ctx))

	return d.repositories.user
}

func (d *DI) LogRepo(ctx context.Context) repository.LogRepository {
	if d.repositories.log != nil {
		return d.repositories.log
	}

	d.repositories.log = logrepo.NewRepository(d.Database(ctx))

	return d.repositories.log
}

func (d *DI) UserService(ctx context.Context) defserv.UserService {
	if d.services.user != nil {
		return d.services.user
	}

	d.services.user = userservice.NewService(d.TxManager(ctx), d.UserRepo(ctx), d.LogRepo(ctx))

	return d.services.user
}

func (d *DI) UserImpl(ctx context.Context) *user.Implementation {
	if d.implementations.user != nil {
		return d.implementations.user
	}

	d.implementations.user = user.NewImplementation(d.Log(ctx), d.UserService(ctx))

	return d.implementations.user
}
