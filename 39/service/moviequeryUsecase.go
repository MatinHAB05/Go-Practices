package service

import "movieapp/api/param/serviceparam"

type MovieQuery interface {
	GetMovieWithCast(MovieId int) (*serviceparam.GetMovieWithCastsByIdResponse, error)
	GetAllMoviesWithCast() ([]serviceparam.GetMovieWithCastsByIdResponse, error)
}
