package handler

import (
	"log"
	"movieapp/api/param"
	baseerrorresponse "movieapp/common/baseresponse"
	"movieapp/common/serviceerror"
	"movieapp/service/moviecastservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieCastLinking_Handler struct {
	moviecastSrv *moviecastservice.Service
}

func (mch *MovieCastLinking_Handler) New(s *moviecastservice.Service) {
	mch.moviecastSrv = s
}

// MovieCastLinking godoc
//
//	@Summary		Link movie to cast
//	@Description	Create relation between a movie and a cast
//	@Tags			MovieCast
//	@Accept			json
//	@Produce		json
//	@Param			data	body		param.LinkMovie2CastRequest	true	"Movie Cast Link"
//	@Success		201		{object}	baseerrorresponse.Response
//	@Failure		400		{object}	baseerrorresponse.Response
//	@Failure		500		{object}	baseerrorresponse.Response
//
//	@Router			/v0/movie-cast-link/ [post]
func (mc_l_h *MovieCastLinking_Handler) Movie_Cast_Linking(c *gin.Context) {
	temp := param.LinkMovie2CastRequest{}
	err := c.ShouldBindJSON(&temp)
	if err != nil {
		log.Printf("bind error : %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseerrorresponse.NewError(err))
		return
	}

	mov_id, cast_id, err := mc_l_h.moviecastSrv.Create(temp)
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}
	mov, err := mc_l_h.moviecastSrv.MovieSrv.GetMovieById(*mov_id)
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}
	cast, err := mc_l_h.moviecastSrv.CastSrv.GetCastById(*cast_id)
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, baseerrorresponse.New(gin.H{
		"Movie": *mov.Movie,
		"Cast":  *cast.Cast,
	}))
}
