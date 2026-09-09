package routes

import (
	"movieapp/handlers"

	"github.com/gin-gonic/gin"
)

func LinkingRoutes(rg *gin.RouterGroup) {
	mc_l_h := handlers.MovieCastLinking_Handlers{}

	mc_l := rg.Group("/movie-cast-link")

	mc_l.POST("/", mc_l_h.Movie_Cast_Linking)
}
