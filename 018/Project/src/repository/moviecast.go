package repository

import (
	"movieapp/api/param/repositoryparam"
	"movieapp/entity"
)

type MovieCastRepository interface {
	GetAllCastsOfMovie(int) ([]repositoryparam.GetAllCastsOfMovieRespnse, error)
	GetAllMoviesOfCast(int) ([]repositoryparam.GetAllMoviesOfCastRespnse, error)
	GetMovieCastById(int) (*entity.MovieCast, error)
	GetMovieCastByMovieIdCastId(int, int) (*entity.MovieCast, error)
	GetAllMovieCast() ([]entity.MovieCast, error)
	IsExistMovieCast(int, int) (*bool, error)
	IsExist(int) (*bool, error)
	Create(entity.MovieCast) (*entity.MovieCast, error)
	Delete(int) (*entity.MovieCast, error)
	DeleteByMovieIdCastId(int, int) (*entity.MovieCast, error)
}
