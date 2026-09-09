package routes

import (
	"movieapp/handlers"

	"github.com/gin-gonic/gin"
)

func CastRoutes(rg *gin.RouterGroup) {
	castHandler := handlers.CastHandler{}

	cast := rg.Group("/cast")

	cast.GET("/show-all", castHandler.ShowAllCast)
	cast.GET("/show/:id", castHandler.ShowCastById)
	cast.GET("/show", castHandler.ShowCastWithFilter)
	cast.POST("/add-cast", castHandler.AddCast)
	cast.DELETE("/rem-cast", castHandler.RemoveCast)
}
