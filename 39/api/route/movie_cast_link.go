package route

import (
	"movieapp/api/handler"

	"github.com/gin-gonic/gin"
)

func Linkingroute(rg *gin.RouterGroup) *handler.MovieCastLinkingHandler {
	mc_l_h := handler.MovieCastLinkingHandler{}

	mc_l := rg.Group("/movie-cast")

	mc_l.POST("/", mc_l_h.MovieCastLinking)
	mc_l.DELETE("/:id", mc_l_h.RemoveMovieCastById)
	mc_l.DELETE("/movie/:movie_id/cast/:cast_id", mc_l_h.RemoveMovieCastByMovieIdCastId)
	mc_l.GET("/movie/:movie_id/cast/:cast_id", mc_l_h.ShowMovieCastByMovieIdCastId)
	mc_l.GET("/:id", mc_l_h.ShowMovieCastById)
	mc_l.GET("/", mc_l_h.ShowAllMovieCasts)

	return &mc_l_h
}
