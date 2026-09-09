package handler

import (
	"log"
	"movieapp/api/param/bindparam"
	"movieapp/api/param/serviceparam"
	"movieapp/common/baseresponse"
	"movieapp/common/serviceerror"
	"movieapp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieCastLinkingHandler struct {
	castSrv      service.Cast
	movieSrv     service.Movie
	moviecastSrv service.MovieCast
}

func (mclh *MovieCastLinkingHandler) New(movieSrv service.Movie, castSrv service.Cast, moviecastSrv service.MovieCast) {
	mclh.moviecastSrv = moviecastSrv
	mclh.castSrv = castSrv
	mclh.movieSrv = movieSrv
}

// MovieCastLinking godoc
//
//	@Summary		Link movie to cast
//	@Description	Create relation between a movie and a cast
//	@Tags			MovieCast
//	@Accept			json
//	@Produce		json
//	@Param			data	body		bindparam.LinkMovie2CastRequest	true	"Movie Cast Link"
//	@Success		201		{object}	baseresponse.Response
//	@Failure		400		{object}	baseresponse.Response
//	@Failure		500		{object}	baseresponse.Response
//
//	@Router			/movie-cast/ [post]
func (mclh *MovieCastLinkingHandler) MovieCastLinking(c *gin.Context) {
	temp := bindparam.LinkMovie2CastRequest{}
	err := c.ShouldBindJSON(&temp)
	if err != nil {
		log.Printf("bind error : %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseresponse.NewError(err))
		return
	}

	mc, err := mclh.moviecastSrv.Create(serviceparam.LinkMovie2CastInput{
		CastId:  temp.CastId,
		MovieId: temp.MovieId,
	})
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}
	mov, err := mclh.movieSrv.GetMovieById(mc.MovieId)
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}
	cast, err := mclh.castSrv.GetCastById(mc.CastId)
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, baseresponse.New(gin.H{
		"Movie": *mov,
		"Cast":  *cast,
	}))
}

// ShowAllMovieCasts godoc
//
//	@Summary		Get all movie-cast relations
//	@Description	Returns the list of all movie and cast relationships
//	@Tags			MovieCast
//	@Produce		json
//	@Success		200	{object}	baseresponse.Response
//	@Failure		400	{object}	baseresponse.Response
//	@Failure		500	{object}	baseresponse.Response
//	@Router			/movie-cast/ [get]
func (mclh *MovieCastLinkingHandler) ShowAllMovieCasts(c *gin.Context) {
	all_mc, err := mclh.moviecastSrv.GetAllMovieCasts()
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, baseresponse.New(all_mc))
}

// ShowMovieCastById godoc
//
//	@Summary		Get movie-cast relation by ID
//	@Description	Returns a specific movie-cast relationship using its ID
//	@Tags			MovieCast
//	@Produce		json
//	@Param			id	path		int	true	"MovieCast ID"
//	@Success		200	{object}	baseresponse.Response
//	@Failure		400	{object}	baseresponse.Response
//	@Failure		404	{object}	baseresponse.Response
//	@Failure		500	{object}	baseresponse.Response
//	@Router			/movie-cast/{id} [get]
func (mclh *MovieCastLinkingHandler) ShowMovieCastById(c *gin.Context) {
	var temp bindparam.GetMovieCastByIdRequest
	err := c.ShouldBindUri(&temp)
	if err != nil {
		log.Printf("bind error : %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseresponse.NewError(err))
		return
	}
	mc, err := mclh.moviecastSrv.GetMovieCastById(temp.ID)
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, baseresponse.New(*mc))
}

// ShowMovieCastByMovieIdCastId godoc
//
//	@Summary		Get movie-cast relation by movie ID and cast ID
//	@Description	Returns the relationship between a specific movie and cast member
//	@Tags			MovieCast
//	@Produce		json
//	@Param			movie_id	path		int	true	"Movie ID"
//	@Param			cast_id		path		int	true	"Cast ID"
//	@Success		200			{object}	baseresponse.Response
//	@Failure		400			{object}	baseresponse.Response
//	@Failure		404			{object}	baseresponse.Response
//	@Failure		500			{object}	baseresponse.Response
//	@Router			/movie-cast/movie/{movie_id}/cast/{cast_id} [get]
func (mclh *MovieCastLinkingHandler) ShowMovieCastByMovieIdCastId(c *gin.Context) {
	var temp bindparam.GetMovieCastByMovieIdCastIdRequest
	err := c.ShouldBindUri(&temp)
	if err != nil {
		log.Printf("bind error : %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseresponse.NewError(err))
		return
	}
	mc, err := mclh.moviecastSrv.GetMovieCastByMovieIdCastId(temp.MovieId, temp.CastId)
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, baseresponse.New(*mc))
}

// RemoveMovieCastById godoc
//
//	@Summary		Delete movie-cast relation by ID
//	@Description	Deletes a movie-cast relationship using its ID
//	@Tags			MovieCast
//	@Produce		json
//	@Param			id	path		int	true	"MovieCast ID"
//	@Success		200	{object}	baseresponse.Response
//	@Failure		400	{object}	baseresponse.Response
//	@Failure		404	{object}	baseresponse.Response
//	@Failure		500	{object}	baseresponse.Response
//	@Router			/movie-cast/{id} [delete]
func (mclh *MovieCastLinkingHandler) RemoveMovieCastById(c *gin.Context) {
	var temp bindparam.DeletetMovieCastByIdRequest
	err := c.ShouldBindUri(&temp)
	if err != nil {
		log.Printf("bind error : %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseresponse.NewError(err))
		return
	}
	mc, err := mclh.moviecastSrv.Delete(temp.ID)
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, baseresponse.New(*mc))
}

// RemoveMovieCastByMovieIdCastId godoc
//
//	@Summary		Delete movie-cast relation by movie ID and cast ID
//	@Description	Deletes the relationship between a specific movie and cast member
//	@Tags			MovieCast
//	@Produce		json
//	@Param			movie_id	path		int	true	"Movie ID"
//	@Param			cast_id		path		int	true	"Cast ID"
//	@Success		200			{object}	baseresponse.Response
//	@Failure		400			{object}	baseresponse.Response
//	@Failure		404			{object}	baseresponse.Response
//	@Failure		500			{object}	baseresponse.Response
//	@Router			/movie-cast/movie/{movie_id}/cast/{cast_id} [delete]
func (mclh *MovieCastLinkingHandler) RemoveMovieCastByMovieIdCastId(c *gin.Context) {
	var temp bindparam.DeleteMovieCastByMovieIdCastIdRequest
	err := c.ShouldBindUri(&temp)
	if err != nil {
		log.Printf("bind error : %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseresponse.NewError(err))
		return
	}
	mc, err := mclh.moviecastSrv.DeleteByMovieIdCastId(temp.MovieId, temp.CastId)
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, baseresponse.New(*mc))
}
