package handler

import (
	"context"
	"grpctest/api/pb"
	"grpctest/internal/app/model/req"
	"grpctest/internal/app/service"
	"grpctest/utils"
)

type MovieHandler struct {
	pb.UnimplementedMovieServiceServer
	movieService service.MovieService
}

func NewMovieHandler(movieService service.MovieService) *MovieHandler {
	return &MovieHandler{movieService: movieService}
}

func (m *MovieHandler) CreateMovie(c context.Context, rq *pb.CreateMovieRequest) (*pb.CreateMovieResponse, error) {
	var movie = req.MovieRequest{
		Title: rq.Movie.GetTitle(),
		Genre: rq.Movie.GetGenre(),
	}
	result := m.movieService.CreateMovie(movie)

	return &pb.CreateMovieResponse{Movie: &pb.Movie{
		Id:    result.ID,
		Title: result.Title,
		Genre: result.Genre,
	}}, nil
}
func (m *MovieHandler) DeleteMovie(ctx context.Context, rq *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {
	id := rq.GetId()
	result, err := m.movieService.DeleteMovie(id)
	utils.ErrorPanic(err)
	return &pb.DeleteMovieResponse{Success: result}, nil
}
func (m *MovieHandler) GetMovie(ctx context.Context, rq *pb.ReadMovieRequest) (*pb.ReadMovieResponse, error) {
	id := rq.GetId()
	result := m.movieService.GetMovie(id)
	return &pb.ReadMovieResponse{Movie: &pb.Movie{
		Id:    result.ID,
		Title: result.Title,
		Genre: result.Genre,
	}}, nil
}
func (m *MovieHandler) GetMovies(ctx context.Context, rq *pb.ReadMoviesRequest) (*pb.ReadMoviesResponse, error) {
	var movies []*pb.Movie
	result := m.movieService.GetMovies()
	for _, e := range result {
		newMovie := &pb.Movie{
			Id:    e.ID,
			Title: e.Title,
			Genre: e.Genre,
		}
		movies = append(movies, newMovie)
	}
	return &pb.ReadMoviesResponse{Movies: movies}, nil
}
func (m *MovieHandler) UpdateMovie(ctx context.Context, rq *pb.UpdateMovieRequest) (*pb.UpdateMovieResponse, error) {
	result, err := m.movieService.UpdateMovie(req.MovieRequest{
		ID: rq.Movie.GetId(),
		Title: rq.Movie.GetTitle(),
		Genre: rq.Movie.GetGenre(),
	})
	utils.ErrorPanic(err)
	return &pb.UpdateMovieResponse{Movie: &pb.Movie{
		Id:    result.ID,
		Title: result.Title,
		Genre: result.Genre,
	}}, nil
}
