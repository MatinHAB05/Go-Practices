package serviceparam

import "movieapp/entity"

type LinkMovie2CastInput struct {
	CastId  int `json:"cast-id"`
	MovieId int `json:"movie-id"`
}

type DeleteMovieCastResponse struct {
	MovieCast *entity.MovieCast `json:"movie_cast"`
}

type GetAllCastsOfMovieRespnse struct {
	Id     int `json:"id"`
	CastId int `json:"cast_id"`
}

type GetAllMoviesOfCastRespnse struct {
	Id      int `json:"id"`
	MovieId int `json:"movie_id"`
}
