package postgres

import (
	"game-api/entity"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, error)
	GetUser(uint64) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) SaveUser(user *entity.User) (*entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) GetUser(id uint64) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", id).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetUsers() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepo) GetUserByEmailAndPassword(u *entity.User) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", u.Email).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
