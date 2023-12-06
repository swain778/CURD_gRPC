package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"gRPC_CURD/models"
	pb "gRPC_CURD/proto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	port = flag.Int("port", 50051, "gRPC server port")
)

type MovieServer struct {
	pb.UnimplementedMovieServiceServer
}

func (*MovieServer) CreateMovie(ctx context.Context, req *pb.CreateMovieRequest) (*pb.CreateMovieResponse, error) {
	fmt.Println("Create Movie")
	movie := req.GetMovie()
	movie.Id = uuid.New().String()

	data := models.Movie{
		ID:    movie.GetId(),
		Title: movie.GetTitle(),
		Genre: movie.GetGenre(),
	}

	res := DB.Create(&data)
	if res.RowsAffected == 0 {
		return nil, errors.New("movie creation successfully")
	}
	return &pb.CreateMovieResponse{
		Movie: &pb.Movie{
			Id:    movie.GetId(),
			Title: movie.GetTitle(),
			Genre: movie.GetGenre(),
		},
	}, nil
}

func (*MovieServer) GetMovie(ctx context.Context, req *pb.ReadMovieRequest) (*pb.ReadMovieResponse, error) {
	fmt.Println("Read Movie", req.GetId())
	var movie models.Movie
	res := DB.Find(&movie, "id = ?", req.GetId())
	if res.RowsAffected == 0 {
		return nil, errors.New("movie not found")
	}
	return &pb.ReadMovieResponse{
		Movie: &pb.Movie{
			Id:    movie.ID,
			Title: movie.Title,
			Genre: movie.Genre,
		},
	}, nil
}

func (*MovieServer) GetMovies(ctx context.Context, req *pb.ReadMoviesRequest) (*pb.ReadMoviesResponse, error) {
	fmt.Println("Read Movies")
	movies := []*pb.Movie{}
	res := DB.Find(&movies)
	if res.RowsAffected == 0 {
		return nil, errors.New("movie not found")
	}
	return &pb.ReadMoviesResponse{
		Movies: movies,
	}, nil
}

func (*MovieServer) UpdateMovie(ctx context.Context, req *pb.UpdateMovieRequest) (*pb.UpdateMovieResponse, error) {
	fmt.Println("Update Movie")
	var movie models.Movie
	reqMovie := req.GetMovie()

	res := DB.Model(&movie).Where("id=?", reqMovie.Id).Updates(models.Movie{Title: reqMovie.Title, Genre: reqMovie.Genre})

	if res.RowsAffected == 0 {
		return nil, errors.New("movie not found")
	}

	return &pb.UpdateMovieResponse{
		Movie: &pb.Movie{
			Id:    movie.ID,
			Title: movie.Title,
			Genre: movie.Genre,
		},
	}, nil
}

func (*MovieServer) DeleteMovie(ctx context.Context, req *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {
	fmt.Println("Delete Movie")
	var movie models.Movie
	res := DB.Where("id=?", req.GetId()).Delete(&movie)
	if res.RowsAffected == 0 {
		return nil, errors.New("movie not found")
	}
	return &pb.DeleteMovieResponse{
		Success: true,
	}, nil
}
