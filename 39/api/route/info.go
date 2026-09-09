package route

import (
	"movieapp/api/handler"

	"github.com/gin-gonic/gin"
)

func Inforoute(rg *gin.RouterGroup) *handler.InfoHandler {
	i_f := handler.InfoHandler{}
	rg.GET("/inspiration", i_f.Inspiration)
	rg.GET("/info", i_f.Info)
	return &i_f
}
