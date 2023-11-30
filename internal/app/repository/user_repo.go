package repository

import (
	"fmt"
	"grpctest/helper/auth"
	"grpctest/internal/app/model"
	"grpctest/utils"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user model.User) (model.User, error)
	GetUser(id string) (model.User, error)
	Login(email, password string) (string, error)
}

type UserRepoImpl struct {
	Db *gorm.DB
}

// Login implements UserRepo.
func (u *UserRepoImpl) Login(email string, password string) (string, error) {
	var user model.User
	if err := u.Db.Where("email = ? ", email).First(&user).Error; err != nil {
		return "", err
	}
	fmt.Print(user.Email)
	if err := auth.ComparePassword(password, user.Password); err != nil {
		return "", err
	}
	token, err := auth.GenToken(user)
	return token, err
}

func NewUserRepoImpl(db *gorm.DB) UserRepo {
	return &UserRepoImpl{Db: db}
}

func (u *UserRepoImpl) CreateUser(user model.User) (model.User, error) {
	var userModel model.User

	result := u.Db.Create(&user)

	if result.Error != nil {
		return model.User{}, result.Error
	}

	result = u.Db.Where("id =?", user.ID).First(&userModel)
	return userModel, result.Error
}
func (u *UserRepoImpl) GetUser(id string) (model.User, error) {
	var user model.User
	result := u.Db.Where("id =?", id).First(&user)
	utils.ErrorPanic(result.Error)
	return user, nil
}
