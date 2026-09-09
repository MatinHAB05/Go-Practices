package errorservice

import (
	"errors"
	"movieapp/common/serviceerror"
	"net/http"
)

var (
	ErrMovieCastNotFound = serviceerror.AppError{
		Err:        errors.New("movie-cast not found"),
		HTTPStatus: http.StatusNotFound,
		Message:    "movie-cast not found",
	}

	ErrMovieCastForeignKeyFails = serviceerror.AppError{
		Err:        errors.New("Invalid movie_id or cast_id. The referenced record does not exist"),
		HTTPStatus: http.StatusNotAcceptable,
		Message:    "The movie or cast you are trying to reference does not exist.",
	}
	ErrMovieCastAlreadyLinked = serviceerror.AppError{
		Err:        errors.New("movie-cast already linked"),
		HTTPStatus: http.StatusNotAcceptable,
		Message:    "movie-cast already linked",
	}

	//	ErrDuplicatedMovieId = AppError{
	//		Err:        errors.New("movie id is duplicated"),
	//		HTTPStatus: http.StatusNotAcceptable,
	//		Message:    "movie id is duplicated",
	//	}
)
