package handler

import (
	"context"
	"grpctest/api/pb"
	"grpctest/internal/app/model/req"
	"grpctest/internal/app/service"
)

type Userhandler struct {
	pb.UnimplementedUserServiceServer
	userService service.UserService
}

func NewUserhandler(userService service.UserService) *Userhandler {
	return &Userhandler{userService: userService}
}

func (handler *Userhandler) RegisterUser(ctx context.Context, rq *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := req.UserRequest{
		Name:     rq.User.GetName(),
		Age:      rq.User.GetAge(),
		Email:    rq.User.GetEmail(),
		Phone:    rq.User.GetPhone(),
		Password: rq.User.GetPassword(),
	}
	result, err := handler.userService.Register(user)

	return &pb.CreateUserResponse{
		User: &pb.User{
			Id:    result.ID,
			Name:  result.Name,
			Email: result.Email,
			Phone: result.Phone,
			Age:   result.Age,
		},
	}, err
}
