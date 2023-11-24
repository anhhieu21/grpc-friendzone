package repository

import (
	"errors"
	"grpctest/internal/app/model"
	"grpctest/internal/app/model/req"
	"grpctest/utils"

	"gorm.io/gorm"
)

type MovieRepo interface {
	CreateMovie(movie model.Movie) model.Movie
	GetMovie(id string) model.Movie
	GetMovies() []model.Movie
	UpdateMovie(movie model.Movie) (model.Movie, error)
	DeleteMovie(id string) (bool, error)
}
type MovieRepoImpl struct {
	Db *gorm.DB
}

func NewMovieRepoImpl(db *gorm.DB) MovieRepo {
	return &MovieRepoImpl{Db: db}
}

func (m *MovieRepoImpl) GetMovie(id string) model.Movie {
	var movie model.Movie
	result := m.Db.Find(&movie, id)
	utils.ErrorPanic(result.Error)

	return movie
}

func (m *MovieRepoImpl) CreateMovie(movie model.Movie) model.Movie {
	var data model.Movie
	result := m.Db.Create(&movie)
	utils.ErrorPanic(result.Error)

	m.Db.Find(&data, movie.ID)

	return data
}
func (m *MovieRepoImpl) GetMovies() []model.Movie {
	var data []model.Movie
	result := m.Db.Find(&data)
	utils.ErrorPanic(result.Error)

	return data
}
func (m *MovieRepoImpl) UpdateMovie(movie model.Movie) (model.Movie, error) {
	var updateData = req.MovieRequest{
		Title: movie.Title,
		Genre: movie.Genre,
	}
	result := m.Db.Model(&movie).Where("id=?", movie.ID).Updates(&updateData)

	if result.RowsAffected == 0 {
		return movie, errors.New("movies not found")
	}
	utils.ErrorPanic(result.Error)
	return movie, nil
}
func (m *MovieRepoImpl) DeleteMovie(id string) (bool, error) {
	var data model.Movie
	result := m.Db.Where("id = ?", id).Delete(&data)
	if result.RowsAffected == 0 {
		return false, errors.New("movie not found")
	}
	utils.ErrorPanic(result.Error)
	return true, nil
}
