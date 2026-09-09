package serviceerror

import (
	"errors"
	"log"
	baseerrorresponse "movieapp/common/baseresponse"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Err        error
	HTTPStatus int
	Message    string
}

func (e *AppError) Error() string { return e.Err.Error() }

func RespondError(c *gin.Context, err error) {
	var appErr *AppError

	if errors.As(err, &appErr) {
		log.Printf("error: %v\n", appErr.Err)
		c.AbortWithStatusJSON(appErr.HTTPStatus, baseerrorresponse.NewError(appErr.Err))
		return
	}

	log.Printf("unexpected error: %v\n", err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, baseerrorresponse.NewSErrorf("internal server error"))
}
