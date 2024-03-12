package logrepo

import (
	"github.com/Masterminds/squirrel"
	repo "github.com/defany/auth-service/app/internal/repository"
	"github.com/defany/db/pkg/postgres"
)

const (
	logs = "logs"
)

const (
	logsAction = "action"
	logsUserID = "user_id"
)

type repository struct {
	db postgres.Postgres
	qb squirrel.StatementBuilderType
}

func NewRepository(db postgres.Postgres) repo.LogRepository {
	return &repository{
		db: db,
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
