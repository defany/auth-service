package postgres

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/defany/auth-service/app/pkg/logger/sl"
	"github.com/defany/auth-service/app/pkg/retry"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func (c *Config) dsn() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		c.Username, c.Password,
		c.Host, c.Port, c.Database,
	)
}

func NewConfig(username string, password string, host string, port string, database string) *Config {
	return &Config{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
	}
}

func NewClient(ctx context.Context, log *slog.Logger, maxAttempts int, maxDelay time.Duration, cfg *Config) (pool *pgxpool.Pool, err error) {
	dsn := cfg.dsn()

	err = retry.WithAttempts(maxAttempts, maxDelay, func() error {
		log.Info("connecting to postgresql database...")

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pgxCfg, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			log.Error("Unable to parse configs", sl.ErrAttr(err))

			return err
		}

		pool, err = pgxpool.NewWithConfig(ctx, pgxCfg)
		if err != nil {
			log.Error("failed to connect to postgres...", sl.ErrAttr(err))

			return err
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Error("ping to postgres failed...", sl.ErrAttr(err))
		}

		return err
	})

	if err != nil {
		log.Error("all attempts are exceeded. unable to connect to postgres database")

		return nil, err
	}

	log.Info("connected to postgresql")

	return
}
