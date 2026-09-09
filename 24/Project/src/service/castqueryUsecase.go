package service

import "movieapp/api/param/serviceparam"

type CastQuery interface {
	GetCastWithMovie(CastId int) (*serviceparam.GetCastWithMoviesByIdResponse, error)
	GetAllCastWithMovie() ([]serviceparam.GetCastWithMoviesByIdResponse, error)
}
