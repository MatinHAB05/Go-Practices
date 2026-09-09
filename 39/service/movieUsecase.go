package service

import (
	"movieapp/api/param/serviceparam"
	"movieapp/entity"
)

type Movie interface {
	ListAllMovies() ([]entity.Movie, error)
	GetMovieById(int) (*entity.Movie, error)

	Create(serviceparam.CreateMovieInput) (*entity.Movie, error)

	Delete(int) (*serviceparam.DeleteMovieResponse, error)

	IsExist(int) (*bool, error)
}
