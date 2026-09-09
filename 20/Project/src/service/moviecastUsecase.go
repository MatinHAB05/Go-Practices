package service

import (
	"movieapp/api/param/serviceparam"
	"movieapp/entity"
)

type MovieCast interface {
	GetAllMovieCasts() ([]entity.MovieCast, error)
	GetAllCastsOfMovie(movie_id int) ([]serviceparam.GetAllCastsOfMovieResponse, error)
	GetAllMoviesOfCast(cast_id int) ([]serviceparam.GetAllMoviesOfCastResponse, error)
	GetMovieCastById(id int) (*entity.MovieCast, error)
	GetMovieCastByMovieIdCastId(MovieId, CastId int) (*entity.MovieCast, error)

	Create(link serviceparam.LinkMovie2CastInput) (*entity.MovieCast, error)

	DeleteByMovieIdCastId(MovieId int, CastId int) (*serviceparam.DeleteMovieCastResponse, error)
	Delete(id int) (*serviceparam.DeleteMovieCastResponse, error)

	IsExistMovieCast(MovieId, CastId int) (*bool, error)
	IsExist(id int) (*bool, error)
}
