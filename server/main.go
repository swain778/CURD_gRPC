package main

import (
	"fmt"
	"gRPC_CURD/models"
	"log"
	"net"

	pb "gRPC_CURD/proto"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	DatabaseConnection()
}

func DatabaseConnection() {
	host := "localhost"
	port := "5432"
	dbName := "postgres1"
	dbUser := "postgres1"
	password := "pass1234"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(models.Movie{})
	if err != nil {
		log.Fatalf("Error connecting to the database..", err)
	}
	fmt.Println("Database connection successful...")
}

func main() {
	fmt.Println("gRPC server running...")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	instance := MovieServer{
		UnimplementedMovieServiceServer: pb.UnimplementedMovieServiceServer{},
	}

	pb.RegisterMovieServiceServer(s, &instance)

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
