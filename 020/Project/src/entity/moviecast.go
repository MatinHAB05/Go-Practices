package entity

import "fmt"

type MovieCast struct {
	Id      int `json:"id"`
	MovieId int `json:"movie_id"`
	CastId  int `json:"cast_id"`
}

func (mc *MovieCast) String() string {
	return fmt.Sprintf(`{id : "%d"  movie_id : "%d"  cast_id : "%d"}`, mc.Id, mc.MovieId, mc.CastId)
}
