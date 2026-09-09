package repositoryparam

type GetAllCastsOfMovieRespnse struct {
	Id     int `json:"id"`
	CastId int `json:"cast_id"`
}

type GetAllMoviesOfCastRespnse struct {
	Id      int `json:"id"`
	MovieId int `json:"movie_id"`
}
