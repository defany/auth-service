package storemodel

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint64              `db:"id"`
	Nickname  string              `db:"name"`
	Email     string              `db:"email"`
	Password  string              `db:"password"`
	Role      string              `db:"role"`
	CreatedAt time.Time           `db:"created_at"`
	UpdatedAt sql.Null[time.Time] `db:"updated_at"`
}
