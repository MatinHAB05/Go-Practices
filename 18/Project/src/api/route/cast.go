package route

import (
	"movieapp/api/handler"

	"github.com/gin-gonic/gin"
)

func Castroute(rg *gin.RouterGroup) *handler.CastHandler {
	castHandler := handler.CastHandler{}

	cast := rg.Group("/cast")

	cast.GET("/show-all", castHandler.ShowAllCast)
	cast.GET("/show/:id", castHandler.ShowCastById)
	cast.GET("/show", castHandler.ShowCastWithFilter)
	cast.POST("/add-cast", castHandler.AddCast)
	cast.DELETE("/rem-cast", castHandler.RemoveCast)
	return &castHandler
}
