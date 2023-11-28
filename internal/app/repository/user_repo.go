package repository

import (
	"grpctest/internal/app/model"
	"grpctest/utils"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user model.User) (model.User, error)
	GetUser(id string) (model.User, error)
}

type UserRepoImpl struct {
	Db *gorm.DB
}

func NewUserRepoImpl(db *gorm.DB) UserRepo {
	return &UserRepoImpl{Db: db}
}
func (u *UserRepoImpl) CreateUser(user model.User) (model.User, error) {
	var userModel model.User
	result := u.Db.Create(&user)
	utils.ErrorPanic(result.Error)

	result = u.Db.Where("id =?", user.ID).First(&userModel)
	utils.ErrorPanic(result.Error)
	return userModel, nil
}
func (u *UserRepoImpl) GetUser(id string) (model.User, error) {
	var user model.User
	result := u.Db.Where("id =?", id).First(&user)
	utils.ErrorPanic(result.Error)
	return user, nil
}
