package routes

import (
	"movieapp/handlers"

	"github.com/gin-gonic/gin"
)

func InfoRoutes(rg *gin.RouterGroup) {
	i_f := handlers.InfoHandler{}
	rg.GET("/inspiration", i_f.Inspiration)
	rg.GET("/info", i_f.Info)

}
