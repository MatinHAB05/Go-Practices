package main

import (
	"log"

	"movieapp/api/route"
	"movieapp/api/template/NoRoute"
	"movieapp/api/validation"
	_ "movieapp/docs"
	mysqldb "movieapp/repository/mysql"
	mysqlcast "movieapp/repository/mysql/cast"
	mysqlmovie "movieapp/repository/mysql/movie"
	mysqlmoviecast "movieapp/repository/mysql/moviecast"
	"movieapp/service/castservice"
	"movieapp/service/moviecastservice"
	"movieapp/service/movieservice"

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

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:8097
//	@BasePath		/

func main() {

	gin_server := gin.Default()

	NoRoute.InitNoRoutePage(gin_server)

	v0 := gin_server.Group("/v0", func(ctx *gin.Context) {
		ctx.Header("api-version", "v0.0.0")
		ctx.Next()
	})

	mh := route.Movieroute(v0)
	ch := route.Castroute(v0)
	mch := route.Linkingroute(v0)
	ih := route.Inforoute(v0)
	gin_server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	mysqlConfig := mysqldb.Config{
		Username: "movieapp",
		Password: "moviepass",
		Port:     3308,
		Host:     "localhost",
		DBName:   "movieapp",
	}
	mysqlDB := mysqldb.New(mysqlConfig)
	movieRepo := mysqlmovie.New(mysqlDB)
	castRepo := mysqlcast.New(mysqlDB)
	moviecastRepo := mysqlmoviecast.New(mysqlDB)

	movieService := movieservice.Service{}
	movieService.New(movieRepo)

	castService := castservice.Service{}
	castService.New(castRepo)

	moviecastService := moviecastservice.Service{}
	moviecastService.New(&movieService, &castService, moviecastRepo)

	mh.New(&movieService)
	ch.New(&castService)
	mch.New(&moviecastService)
	ih.New()

	log.Println("Start...")
	if err := gin_server.Run(":8097"); err != nil {
		log.Fatal(err)
	}

}
