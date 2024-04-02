package userservice

import (
	"github.com/defany/auth-service/app/internal/repository"
	defserv "github.com/defany/auth-service/app/internal/service"
	"github.com/defany/db/pkg/postgres"
)

type PasswordHasher interface {
	GenerateFromPassword(password []byte) ([]byte, error)
}

type service struct {
	tx         postgres.TxManager
	repo       repository.UserRepository
	log        repository.LogRepository
	passHasher PasswordHasher
}

func NewService(tx postgres.TxManager, repo repository.UserRepository, log repository.LogRepository, passHasher PasswordHasher) defserv.UserService {
	return &service{
		tx:         tx,
		repo:       repo,
		log:        log,
		passHasher: passHasher,
	}
}
