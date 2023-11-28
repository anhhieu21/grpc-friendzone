package service

import (
	"grpctest/helper/auth"
	"grpctest/internal/app/model"
	"grpctest/internal/app/model/req"
	"grpctest/internal/app/repository"

	"github.com/google/uuid"
)

type UserService interface {
	Login(email string, password string) (model.User, error)
	Register(user req.UserRequest) (model.User, error)
}

type UserServiceImp struct {
	UserRepo repository.UserRepo
}

// Login implements UserService.
func (u *UserServiceImp) Login(email string, password string) (model.User, error) {
	panic("unimplemented")
}

// Register implements UserService.
func (u *UserServiceImp) Register(user req.UserRequest) (model.User, error) {
	hashedPassword := auth.HashPassword(user.Password)

	userModel := model.User{
		ID:       uuid.New().String(),
		Name:     user.Name,
		Age:      user.Age,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: hashedPassword,
	}

	result, err := u.UserRepo.CreateUser(userModel)
	return result, err
}

func NewUserServiceImpl(userRepo repository.UserRepo) UserService {
	return &UserServiceImp{
		UserRepo: userRepo,
	}
}
