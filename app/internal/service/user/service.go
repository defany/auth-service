package userservice

import (
	"github.com/defany/auth-service/app/internal/repository"
	defserv "github.com/defany/auth-service/app/internal/service"
	"github.com/defany/auth-service/app/pkg/postgres"
)

type service struct {
	tx   postgres.TxManager
	repo repository.UserRepository
	log  repository.LogRepository
}

func NewService(tx postgres.TxManager, repo repository.UserRepository, log repository.LogRepository) defserv.UserService {
	return &service{
		tx:   tx,
		repo: repo,
		log:  log,
	}
}
