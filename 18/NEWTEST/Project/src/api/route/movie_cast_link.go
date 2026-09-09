package route

import (
	"movieapp/api/handler"

	"github.com/gin-gonic/gin"
)

func Linkingroute(rg *gin.RouterGroup) *handler.MovieCastLinking_Handler {
	mc_l_h := handler.MovieCastLinking_Handler{}

	mc_l := rg.Group("/movie-cast-link")

	mc_l.POST("/", mc_l_h.Movie_Cast_Linking)
	return &mc_l_h
}
