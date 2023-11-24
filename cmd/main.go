package main

import (
	"flag"
	"fmt"
	"grpctest/api/pb"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Movie struct {
	ID        string `gorm:"primarykey"`
	Title     string
	Genre     string
	CreatedAt time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:false"`
}

func init() {
	DatabaseConnection()
}

var DB *gorm.DB
var err error

func DatabaseConnection() {
	host := "localhost"
	port := "5432"
	dbName := "postgres"
	dbUser := "postgres"
	password := "1111"
	dns := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password)
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	DB.AutoMigrate(Movie{})
	if err != nil {
		log.Fatal("Error connectiong to the database...", err)
	}
	fmt.Println("Databse connection successful")

}

var (
	port = flag.Int("port", 8080, "gRPC server port")
)

type server struct {
	pb.UnimplementedMovieServiceServer
}

func main() {
	fmt.Println("gRPC server running ...")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMovieServiceServer(s, &server{})

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
