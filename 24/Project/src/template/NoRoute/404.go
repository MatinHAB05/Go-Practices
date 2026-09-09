package NoRoute

import "github.com/gin-gonic/gin"

func InitNoRoutePage(gin_server *gin.Engine) {
	gin_server.LoadHTMLGlob("./template/NoRoute/*")

	gin_server.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{
			"title": "Page Not Found",
		})
	})

}
