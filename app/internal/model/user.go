package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID        uint64
	Nickname  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserCreate struct {
	Nickname string `validate:"required|min_len:6"`
	Email    string `validate:"required|email"`
	Password string `validate:"required|min_len:8"`
	Role     string `validate:"required"`
}

type UserUpdate struct {
	ID    uint64
	Name  *string
	Email *string
	Role  *string
}

type UserLoginInput struct {
	Nickname string
	Password string
}

type UserLoginOutput struct {
	RefreshToken string
}

type UserClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}
