package service

import (
	"game-api/entity"
	"game-api/repository/postgres"
)

type userSvc struct {
	us postgres.UserRepository
}

func NewUserService(us postgres.UserRepository) UserService {
	return &userSvc{
		us: us,
	}
}

type UserService interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUsers() ([]entity.User, error)
	GetUser(uint64) (*entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string)
}

func (u *userSvc) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return u.us.SaveUser(user)
}

func (u *userSvc) GetUser(userId uint64) (*entity.User, error) {
	return u.us.GetUser(userId)
}

func (u *userSvc) GetUsers() ([]entity.User, error) {
	return u.us.GetUsers()
}

func (u *userSvc) GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string) {
	return u.us.GetUserByEmailAndPassword(user)
}
