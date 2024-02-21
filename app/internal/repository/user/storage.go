package userrepo

import (
	"github.com/Masterminds/squirrel"
	repo "github.com/defany/auth-service/app/internal/repository"
	"github.com/defany/auth-service/app/pkg/postgres"
)

const (
	table = "users"
)

const (
	idColumn              = "id"
	nameColumn            = "name"
	emailColumn           = "email"
	passwordColumn        = "password"
	passwordConfirmColumn = "password_confirm"
	roleColumn            = "role"
	createdAtColumn       = "created_at"
	updatedAtColumn       = "updated_at"
)

type repository struct {
	db postgres.Querier
	qb squirrel.StatementBuilderType
}

func NewRepository(db postgres.Querier) repo.UserRepository {
	return &repository{
		db: db,
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
