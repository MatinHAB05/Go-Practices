package serviceparam

import "movieapp/entity"

type GetMovieWithCastsByIdResponse struct {
	Id          int                 `json:"id"`
	Title       string              `json:"title"`
	ReleaseYear int                 `json:"releaseYear"`
	Quality     entity.VideoQuality `json:"quality"`
	Casts       []entity.Cast       `json:"casts"`
}

type GetCastWithMoviesByIdResponse struct {
	Name   string         `json:"name"`
	Id     int            `json:"id"`
	Movies []entity.Movie `json:"movies"`
}
