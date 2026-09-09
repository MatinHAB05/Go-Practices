package serviceparam

import "movieapp/entity"

type GetMovieByIdInpuyt struct {
	ID int `uri:"id"`
}

type RemoveMovieByIdInput struct {
	ID int `json:"id"`
}

type GetMovieWithFilterInput struct {
	Title         string `form:"title"`
	ReleaseYear   int    `form:"release"`
	ReleaseAction string `form:"release-action"`
	Quality       string `form:"quality"`
}

type CreateMovieInput struct {
	Title       string              `json:"title"`
	ReleaseYear int                 `json:"releaseYear"`
	Quality     entity.VideoQuality `json:"quality"`
}

type DeleteMovieResponse struct {
	Movie *entity.Movie `json:"movies"`
}
