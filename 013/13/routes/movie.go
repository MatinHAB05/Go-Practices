package routes

import (
	"movieapp/handlers"

	"github.com/gin-gonic/gin"
)

func MovieRoutes(rg *gin.RouterGroup) {
	movieHandler := handlers.MovieHandler{}

	movie := rg.Group("/movie")

	movie.GET("/show-all", movieHandler.ShowAllMovie)
	movie.GET("/show/:id", movieHandler.ShowMovieById)
	movie.GET("/show", movieHandler.ShowMovieWithFilter)
	movie.POST("/add-movie", movieHandler.AddMovie)
	movie.DELETE("/rem-movie", movieHandler.RemoveMovie)
}
