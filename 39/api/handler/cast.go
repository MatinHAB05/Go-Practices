package handler

import (
	"log"
	"movieapp/api/helper/filter/filtercast"
	"movieapp/api/param/bindparam"
	"movieapp/api/param/serviceparam"
	"movieapp/common/baseresponse"
	"movieapp/common/serviceerror"
	"movieapp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CastHandler struct {
	castSrv      service.Cast
	moviecastSrv service.MovieCast
	castquerySrv service.CastQuery
}

func (ch *CastHandler) New(castSrv service.Cast, moviecastSrv service.MovieCast, castquerySrv service.CastQuery) {
	ch.castSrv = castSrv
	ch.castquerySrv = castquerySrv
	ch.moviecastSrv = moviecastSrv
}

// ShowAllCast godoc
//
//	@Summary		List all casts
//	@Description	Get all casts
//	@Tags			Casts
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	baseresponse.Response
//	@Failure		500	{object}	baseresponse.Response
//
//	@Router			/cast/show-all [get]
func (ch *CastHandler) ShowAllCast(c *gin.Context) {
	all_cast, err := ch.castquerySrv.GetAllCastWithMovie()
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, baseresponse.New(all_cast))
}

// ShowCastById godoc
//
//	@Summary		Get cast by ID
//	@Description	Retrieve a cast by its ID
//	@Tags			Casts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Cast ID"
//	@Success		200	{object}	baseresponse.Response
//	@Failure		400	{object}	baseresponse.Response
//	@Failure		404	{object}	baseresponse.Response
//	@Failure		500	{object}	baseresponse.Response
//
//	@Router			/cast/show/{id} [get]
func (ch *CastHandler) ShowCastById(c *gin.Context) {
	temp := bindparam.GetCastByIdRequest{}
	err := c.ShouldBindUri(&temp)
	if err != nil {
		log.Printf("bind error : %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseresponse.NewError(err))
		return
	}
	i := temp.ID
	cast, err := ch.castquerySrv.GetCastWithMovie(i)
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, baseresponse.New(cast))
}

// AddCast godoc
//
//	@Summary		Create a new cast
//	@Description	Add a new cast
//	@Tags			Casts
//	@Accept			json
//	@Produce		json
//	@Param			cast	body		bindparam.CreateCastRequest	true	"Cast Data"
//	@Success		201		{object}	baseresponse.Response
//	@Failure		400		{object}	baseresponse.Response
//	@Failure		500		{object}	baseresponse.Response
//
//	@Router			/cast/add-cast [post]
func (ch *CastHandler) AddCast(c *gin.Context) {
	var temp bindparam.CreateCastRequest
	err := c.ShouldBindJSON(&temp)
	if err != nil {
		log.Printf("bind error : %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseresponse.NewError(err))
		return
	}
	cast, err := ch.castSrv.Create(serviceparam.CreateCastInput{Name: temp.Name})
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, baseresponse.New(*cast))
}

// RemoveCast godoc
//
//	@Summary		Delete a cast
//	@Description	Remove a cast by ID
//	@Tags			Casts
//	@Accept			json
//	@Produce		json
//	@Param			id	body		bindparam.RemoveCastByIdRequest	true	"Cast ID"
//	@Success		200	{object}	baseresponse.Response
//	@Failure		400	{object}	baseresponse.Response
//	@Failure		404	{object}	baseresponse.Response
//	@Failure		500	{object}	baseresponse.Response
//
//	@Router			/cast/rem-cast [delete]
func (ch *CastHandler) RemoveCast(c *gin.Context) {
	temp := bindparam.RemoveCastByIdRequest{}
	err := c.ShouldBindJSON(&temp)
	if err != nil {
		log.Printf("bind error : %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseresponse.NewError(err))
		return
	}
	cast, err := ch.castSrv.Delete(temp.ID)
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, baseresponse.New(cast.Cast))

}

// ShowCastWithFilter godoc
//
//	@Summary		Filter casts
//	@Description	Filter casts by name
//	@Tags			Casts
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string	false	"Cast name prefix"
//	@Success		200		{object}	baseresponse.Response
//	@Failure		400		{object}	baseresponse.Response
//	@Failure		500		{object}	baseresponse.Response
//	@Router			/cast/show [get]
func (ch *CastHandler) ShowCastWithFilter(c *gin.Context) {
	temp := bindparam.GetCastWithFilterRequest{}
	err := c.ShouldBindQuery(&temp)
	if err != nil {
		log.Printf("bind error : %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseresponse.NewError(err))
		return
	}

	all_cast, err := ch.castquerySrv.GetAllCastWithMovie()
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}

	_, exist_name := c.GetQuery("name")
	res := make([]serviceparam.GetCastWithMoviesByIdResponse, 0)
	for i := range all_cast {
		if exist_name && !filtercast.ByName(all_cast[i].Name, temp.Name) {
			continue
		}

		res = append(res, all_cast[i])
	}
	c.IndentedJSON(http.StatusOK, baseresponse.New(res))

}
