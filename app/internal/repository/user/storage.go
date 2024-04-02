package userrepo

import (
	"github.com/Masterminds/squirrel"
	repo "github.com/defany/auth-service/app/internal/repository"
	"github.com/defany/db/pkg/postgres"
)

const (
	table                     = "users"
	tableEndpointsPermissions = "roles_endpoints_permissions"
)

const (
	idColumn        = "id"
	nicknameColumn  = "nickname"
	emailColumn     = "email"
	passwordColumn  = "password"
	roleColumn      = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repository struct {
	db postgres.Postgres
	qb squirrel.StatementBuilderType
}

func NewRepository(db postgres.Postgres) repo.UserRepository {
	return &repository{
		db: db,
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
