package errorservice

import (
	"fmt"
	"movieapp/common/serviceerror"
	"net/http"
)

var (
	ErrCastNotFound = serviceerror.AppError{
		HTTPStatus: http.StatusNotFound,
		Err:        fmt.Errorf("cast not found"),
		Message:    "cast not found",
	}
)
