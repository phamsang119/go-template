package service

import (
	"food-app/repository/postgres"
)

type Service struct {
	repo        *postgres.Repositories
	UserService *UserService
}

func NewService(repo *postgres.Repositories) *Service {
	userSvc := NewUserService(repo.User)
	return &Service{
		repo:        repo,
		UserService: &userSvc,
	}
}