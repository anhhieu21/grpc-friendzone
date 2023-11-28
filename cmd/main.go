package main

import (
	"flag"
	"fmt"
	"grpctest/api/handler"
	"grpctest/api/pb"
	"grpctest/helper/middleware"
	"grpctest/internal/app/model"
	"grpctest/internal/app/repository"
	"grpctest/internal/app/service"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	DatabaseConnection()
}

var DB *gorm.DB
var err error

func DatabaseConnection() {
	host := "nonsense.ddns.net"
	port := "5432"
	dbName := "friendzone"
	dbUser := "postgres"
	password := "1111"
	dns := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		host,
		port,
		dbName,
		dbUser,
		password)
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	DB.AutoMigrate(model.Movie{})
	DB.AutoMigrate(model.User{})
	if err != nil {
		log.Fatal("Error connectiong to the database...", err)
	}
	fmt.Println("Databse connection successful")

}

var (
	port = flag.Int("port", 8080, "gRPC server port")
)

func main() {
	fmt.Println("gRPC server running ...")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	movieRepoImpl := repository.NewMovieRepoImpl(DB)
	movieServiceImpl := service.NewMovieServiceImpl(movieRepoImpl)
	movieHandler := handler.NewMovieHandler(movieServiceImpl)
	//user
	userRepoImpl := repository.NewUserRepoImpl(DB)
	userServiceImpl := service.NewUserServiceImpl(userRepoImpl)
	userHandler := handler.NewUserhandler(userServiceImpl)

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(middleware.ValidateUserInterceptor()))
	pb.RegisterMovieServiceServer(s, movieHandler)
	pb.RegisterUserServiceServer(s, userHandler)

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
