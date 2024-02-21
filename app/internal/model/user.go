package model

import (
	"time"
)

type User struct {
	ID              uint64
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            string
	CreatedAt       time.Time
	UpdatedAt       *time.Time
}

type UserCreate struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            string
}

type UserUpdate struct {
	ID    uint64
	Name  *string
	Email *string
	Role  *string
}
