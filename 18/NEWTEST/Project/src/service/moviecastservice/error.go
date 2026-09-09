package moviecastservice

import (
	"errors"
	"movieapp/common/serviceerror"
	"net/http"
)

var (
	//	ErrNotFoundMovie = serviceerror.AppError{
	//		Err:        errors.New("movie not found"),
	//		HTTPStatus: http.StatusNotFound,
	//		Message:    "movie not found",
	//	}
	//
	//	ErrDuplicatedMovieId = AppError{
	//		Err:        errors.New("movie id is duplicated"),
	//		HTTPStatus: http.StatusNotAcceptable,
	//		Message:    "movie id is duplicated",
	//	}
	ErrAlreadyLinkedMovieCast = serviceerror.AppError{
		Err:        errors.New("movie-cast already linked"),
		HTTPStatus: http.StatusNotAcceptable,
		Message:    "movie-cast already linked",
	}
)
