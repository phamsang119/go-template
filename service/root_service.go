package service

import (
	"game-api/repository/postgres"
)

type Service struct {
	repo        *postgres.Repositories
	UserService UserService
}

func NewService(repo *postgres.Repositories) *Service {
	return &Service{
		repo:        repo,
		UserService: NewUserService(repo.User),
	}
}
