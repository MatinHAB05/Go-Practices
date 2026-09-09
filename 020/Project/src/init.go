package main

import (
	"movieapp/api/handler"
	"movieapp/appconstant"
	mysqldb "movieapp/repository/mysql"
	mysqlcast "movieapp/repository/mysql/cast"
	mysqlmovie "movieapp/repository/mysql/movie"
	mysqlmoviecast "movieapp/repository/mysql/moviecast"
	"movieapp/service/castqueryservice"
	"movieapp/service/castservice"
	"movieapp/service/moviecastservice"
	"movieapp/service/moviequeryservice"
	"movieapp/service/movieservice"
)

func InitServices(mh *handler.MovieHandler,
	ch *handler.CastHandler,
	mch *handler.MovieCastLinkingHandler,
	ih *handler.InfoHandler,
) {
	mysqlConfig := mysqldb.Config{
		Username: appconstant.MYSQLUSERNAME,
		Password: appconstant.MYSQLPASSWORD,
		Port:     appconstant.MYSQLPORT,
		Host:     appconstant.MYSQLHOST,
		DBName:   appconstant.MYSQLDBNAME,
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
	moviecastService.New(moviecastRepo)

	moviequeryService := moviequeryservice.Service{}
	moviequeryService.New(movieRepo, castRepo, moviecastRepo)

	castqueryService := castqueryservice.Service{}
	castqueryService.New(movieRepo, castRepo, moviecastRepo)

	mh.New(&movieService, &moviecastService, &moviequeryService)
	ch.New(&castService, &moviecastService, &castqueryService)
	mch.New(&movieService, &castService, &moviecastService)
	ih.New()
}
