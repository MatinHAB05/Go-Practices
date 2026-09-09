package main

import (
	"log"

	"movieapp/api/route"
	"movieapp/api/validation"
	_ "movieapp/docs"
	"movieapp/template/NoRoute"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("castName", validation.CastNameValid)
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

// @title			Movie App API
// @version		0.0
// @termsOfService	https://No-Where-Be-Allah-Ghasam-1
// @contact.name	Matin Hasanali Baki ( Matin HAB )
// @contact.url	https://No-Where-Be-Allah-Ghasam-2
// @contact.email	m9652973@gmail.com
// @license.name	Apache 9797.0
// @license.url	https://No-Where-Be-Allah-Ghasam-3
// @description	REST API for managing movies and casts
// @host			localhost:8097
// @BasePath		/api/v0
func main() {

	gin_server := gin.Default()
	NoRoute.InitNoRoutePage(gin_server)

	api := gin_server.Group("/api", func(ctx *gin.Context) {
		ctx.Header("creator", "MatinHAB")
		ctx.Next()
	})

	v0 := api.Group("/v0", func(ctx *gin.Context) {
		ctx.Header("api-version", "v0.0.0")
		ctx.Next()
	})
	v0.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	mh := route.Movieroute(v0)
	ch := route.Castroute(v0)
	mch := route.Linkingroute(v0)
	ih := route.Inforoute(v0)

	InitServices(mh, ch, mch, ih)

	log.Println("Start...")
	if err := gin_server.Run(":8097"); err != nil {
		log.Fatal(err)
	}

}
