package bindparam

type LinkMovie2CastRequest struct {
	CastId  int `json:"cast-id" binding:"required"`
	MovieId int `json:"movie-id" binding:"required"`
}

type GetMovieCastByIdRequest struct {
	ID int `uri:"id" json:"id" binding:"required"`
}
type GetMovieCastByMovieIdCastIdRequest struct {
	MovieId int `uri:"movie_id" json:"movie_id" binding:"required"`
	CastId  int `uri:"cast_id" json:"cast_id" binding:"required"`
}

type DeletetMovieCastByIdRequest struct {
	ID int `uri:"id" json:"id" binding:"required"`
}
type DeleteMovieCastByMovieIdCastIdRequest struct {
	MovieId int `uri:"movie_id" json:"movie_id" binding:"required"`
	CastId  int `uri:"cast_id" json:"cast_id" binding:"required"`
}
