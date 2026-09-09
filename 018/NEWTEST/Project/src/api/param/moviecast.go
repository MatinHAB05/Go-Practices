package param

type LinkMovie2CastRequest struct {
	CastId  int `json:"cast-id" binding:"required"`
	MovieId int `json:"movie-id" binding:"required"`
}
