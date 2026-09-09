package repository

import "movieapp/entity"

type MovieRepository interface {
	GetAllMovies() ([]entity.Movie, error)
	GetMovieById(int) (*entity.Movie, error)
	GetMovieByIds([]int) ([]entity.Movie, error)
	Create(entity.Movie) (*entity.Movie, error)
	Delete(int) (*entity.Movie, error)
	IsExist(int) (*bool, error)
}
