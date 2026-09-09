package main

import (
	"log"
	"movieapp/routes"
	"movieapp/validation"

	_ "movieapp/docs"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("actorName", validation.ActorNameValid)
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("movieTitle", validation.MovieTitleValid)
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("comparatorOperator", validation.ComparatorOperatorValid)
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("movieQuality", validation.MovieQualityValid)
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("releaseYear", validation.ReleaseYearValid)
	}

}

// @title			Swagger Example API
// @version		1.0
// @description	This is a sample server celler server.
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8080
// @BasePath		/api/v
func main() {
	// TODO TEST VALIDATIONS........ pass+
	gin_server := gin.Default()

	gin_server.LoadHTMLGlob("template/*")

	gin_server.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{
			"title": "Page Not Found",
		})
	})

	v0 := gin_server.Group("/v0", func(ctx *gin.Context) {
		ctx.Header("api-version", "v0.0.0")
		ctx.Next()
	})

	routes.MovieRoutes(v0)
	routes.CastRoutes(v0)
	routes.LinkingRoutes(v0)
	routes.InfoRoutes(v0)
	gin_server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := gin_server.Run(":8097"); err != nil {
		log.Fatal(err)
	}

}
