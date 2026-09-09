package bindparam

import "movieapp/entity"

type GetMovieByIdRequest struct {
	ID int `uri:"id" binding:"required" json:"id"`
}
type GetMovieWithFilterRequest struct {
	Title         string `form:"title" binding:"omitempty,movieTitle"`
	ReleaseYear   int    `form:"release" binding:"omitempty,releaseYear"`
	ReleaseAction string `form:"release-action" binding:"omitempty,comparatorOperator"`
	Quality       string `form:"quality" binding:"omitempty,movieQuality"`
}
type CreateMovieRequest struct {
	Title       string              `json:"title" binding:"omitempty,movieTitle"`
	ReleaseYear int                 `json:"releaseYear" binding:"omitempty,releaseYear"`
	Quality     entity.VideoQuality `json:"quality" binding:"omitempty,movieQuality"`
}

type RemoveMovieByIdRequest struct {
	ID int `json:"id" binding:"required"`
}

// type GetMovieByIdResponse struct {
// 	Id          int                 `json:"id"`
// 	Title       string              `json:"title"`
// 	ReleaseYear int                 `json:"releaseYear"`
// 	Quality     entity.VideoQuality `json:"quality"`
// 	Casts       []entity.Cast       `json:"casts"`
// }

// type DeleteMovieResponse struct {
// 	Movie *entity.Movie
// }
