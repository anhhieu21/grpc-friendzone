package service

import (
	"fmt"
	"grpctest/internal/app/model"
	"grpctest/internal/app/model/req"
	"grpctest/internal/app/repository"

	"github.com/google/uuid"
)

type MovieService interface {
	CreateMovie(movie req.MovieRequest) model.Movie
	GetMovie(id string) model.Movie
	GetMovies() []model.Movie
	UpdateMovie(movie req.MovieRequest) (model.Movie, error)
	DeleteMovie(id string) (bool, error)
}
type MovieServiceImpl struct {
	MovieRepo repository.MovieRepo
}

func NewMovieServiceImpl(repo repository.MovieRepo) MovieService {
	return &MovieServiceImpl{
		MovieRepo: repo,
	}
}

// CreateMovie implements MovieService.
func (m *MovieServiceImpl) CreateMovie(movie req.MovieRequest) model.Movie {
	fmt.Println("Create Movie")
	data := model.Movie{
		ID:    uuid.New().String(),
		Title: movie.Title,
		Genre: movie.Genre,
	}
	m.MovieRepo.CreateMovie(data)
	return data
}

// DeleteMovie implements MovieService.
func (m *MovieServiceImpl) DeleteMovie(id string) (bool, error) {
	fmt.Println("Delete Movie")
	return m.MovieRepo.DeleteMovie(id)
}

// GetMovie implements MovieService.
func (m *MovieServiceImpl) GetMovie(id string) model.Movie {
	fmt.Println("Delete Movie")
	return m.MovieRepo.GetMovie(id)
}

// GetMovies implements MovieService.
func (m *MovieServiceImpl) GetMovies() []model.Movie {
	fmt.Println("Get Movies")
	return m.MovieRepo.GetMovies()
}

// UpdateMovie implements MovieService.
func (m *MovieServiceImpl) UpdateMovie(movie req.MovieRequest) (model.Movie, error) {
	fmt.Println("Update Movie")
	var data = model.Movie{
		ID:    movie.ID,
		Title: movie.Title,
		Genre: movie.Genre,
	}
	result, err := m.MovieRepo.UpdateMovie(data)
	return result, err
}
