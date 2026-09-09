package repositoryparam

type MovieCast struct {
	Id      int `json:"id" db:"id"`
	MovieId int `json:"movie_id" db:"movie_id"`
	CastId  int `json:"cast_id" db:"cast_id"`
}

type CastsOfMovieInput struct {
	Id     int `json:"id" db:"id"`
	CastId int `json:"cast_id" db:"cast_id"`
}

type MoviesOfCastInput struct {
	Id      int `json:"id" db:"id"`
	MovieId int `json:"movie_id" db:"movie_id"`
}
type GetAllCastsOfMovieResponse struct {
	Id     int `json:"id"`
	CastId int `json:"cast_id"`
}

type GetAllMoviesOfCastResponse struct {
	Id      int `json:"id"`
	MovieId int `json:"movie_id"`
}
