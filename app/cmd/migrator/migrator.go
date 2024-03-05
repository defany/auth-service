package main

import (
	"context"
	"github.com/defany/auth-service/app/internal/app"
	"github.com/defany/auth-service/app/pkg/logger/sl"
	"github.com/defany/auth-service/app/pkg/postgres"
	"log/slog"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a := app.NewApp()

	db := a.DI().Database(ctx)

	log := a.DI().Log(ctx)

	migrator, err := postgres.NewMigrator(db.Pool(), a.DI().Config(ctx).Database.MigrationsDir)
	if err != nil {
		log.Error("failed to setup migrator", sl.ErrAttr(err))

		return
	}

	log.Info("upping the migrations")

	upped, err := migrator.Up(ctx)
	if err != nil {
		log.Error("failed to up migrations", sl.ErrAttr(err))

		return
	}

	if len(upped) == 0 {
		log.Info("there is no migrations to up")

		return
	}

	for _, migration := range upped {
		log.Info("migration upped!", slog.String("name", migration.Source.Path))
	}
}
