package errorservice

import (
	"errors"
	"movieapp/common/serviceerror"
	"net/http"
)

var (
	ErrMovieNotFound = serviceerror.AppError{
		Err:        errors.New("movie not found"),
		HTTPStatus: http.StatusNotFound,
		Message:    "movie not found",
	}
	// ErrDuplicatedMovieId = AppError{
	// 	Err:        errors.New("movie id is duplicated"),
	// 	HTTPStatus: http.StatusNotAcceptable,
	// 	Message:    "movie id is duplicated",
	// }
)
