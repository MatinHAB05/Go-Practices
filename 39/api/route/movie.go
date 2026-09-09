package route

import (
	"movieapp/api/handler"

	"github.com/gin-gonic/gin"
)

func Movieroute(rg *gin.RouterGroup) *handler.MovieHandler {
	movieHandler := handler.MovieHandler{}

	movie := rg.Group("/movie")

	movie.GET("/show-all", movieHandler.ShowAllMovie)
	movie.GET("/show/:id", movieHandler.ShowMovieById)
	movie.GET("/show", movieHandler.ShowMovieWithFilter)
	movie.POST("/add-movie", movieHandler.AddMovie)
	movie.DELETE("/rem-movie", movieHandler.RemoveMovie)
	return &movieHandler
}
