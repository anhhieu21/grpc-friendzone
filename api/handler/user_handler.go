package handler

import (
	"context"
	"grpctest/api/pb"
	"grpctest/helper/auth"
	"grpctest/internal/app/model/req"
	"grpctest/internal/app/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Userhandler struct {
	pb.UnimplementedUserServiceServer
	userService service.UserService
}

func NewUserhandler(userService service.UserService) *Userhandler {
	return &Userhandler{userService: userService}
}

func (handler *Userhandler) Login(ctx context.Context, rq *pb.LoginRequest) (*pb.LoginResponse, error) {
	result, err := handler.userService.Login(rq.Email, rq.Password)
	return &pb.LoginResponse{
		Token: result.Token,
	}, err
}

func (handler *Userhandler) GetUser(ctx context.Context, rq *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	claims, err := auth.GetUserIdFromContext(ctx)
	if err != nil {
		return &pb.GetUserResponse{}, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}
	result, err := handler.userService.GetUser(claims.UserId)
	return &pb.GetUserResponse{User: &pb.User{
		Id:    result.ID,
		Name:  result.Name,
		Email: result.Email,
		Phone: result.Phone,
		Age:   result.Age,
	}}, err
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

	if err != nil {
		return &pb.CreateUserResponse{
			Error: &pb.Error{
				Code:    codes.AlreadyExists.String(),
				Message: "user with the same email or phone already exists",
			},
		}, status.Error(codes.AlreadyExists, "user with the same email or phone already exists")
	}

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
