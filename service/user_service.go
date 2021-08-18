package service

import (
	"game-api/constant"
	"game-api/entity"
	"game-api/repository/postgres"
	"game-api/utils"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type userSvc struct {
	userRepo postgres.UserRepository
}

func NewUserService(us postgres.UserRepository) UserService {
	return &userSvc{
		userRepo: us,
	}
}

type UserService interface {
	SaveUser(*entity.User) (*entity.User, *entity.Error)
	GetUsers() ([]entity.User, *entity.Error)
	GetUser(uint64) (*entity.User, *entity.Error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, *entity.Error)
}

func (u *userSvc) SaveUser(user *entity.User) (*entity.User, *entity.Error) {
	result, err := u.userRepo.SaveUser(user)
	if err != nil {
		if utils.IsDuplicateRecord(err) {
			return nil, entity.NewError(http.StatusBadRequest, constant.MsgEmailDuplicated)
		}
		return nil, entity.NewError(http.StatusInternalServerError, constant.MsgServerError)
	}
	return result, nil
}

func (u *userSvc) GetUser(userId uint64) (*entity.User, *entity.Error) {
	result, err := u.userRepo.GetUser(userId)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, entity.NewError(constant.UserNotFound, constant.MsgUserNotFound)
		}
		return nil, entity.NewError(http.StatusInternalServerError, constant.MsgServerError)
	}
	return result, nil
}

func (u *userSvc) GetUsers() ([]entity.User, *entity.Error) {
	result, err := u.userRepo.GetUsers()
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, entity.NewError(http.StatusInternalServerError, constant.MsgServerError)
	}
	return result, nil
}

func (u *userSvc) GetUserByEmailAndPassword(user *entity.User) (*entity.User, *entity.Error) {
	result, err := u.userRepo.GetUserByEmailAndPassword(user)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, entity.NewError(constant.UserNotFound, constant.MsgUserNotFound)
		}
		return nil, entity.NewError(http.StatusInternalServerError, constant.MsgServerError)
	}
	//Verify the password
	err = utils.VerifyPassword(user.Password, result.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, entity.NewError(http.StatusBadRequest, constant.MsgInvalidPassword)
	}
	return result, nil
}
