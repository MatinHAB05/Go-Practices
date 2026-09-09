package handler

import (
	"log"
	"movieapp/api/param"
	baseerrorresponse "movieapp/common/baseresponse"
	"movieapp/common/serviceerror"
	"movieapp/entity"
	"movieapp/service/castservice"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type CastHandler struct {
	castSrv *castservice.Service
}

func (ch *CastHandler) New(s *castservice.Service) {
	ch.castSrv = s
}

// ShowAllCast godoc
//
//	@Summary		List all casts
//	@Description	Get all casts
//	@Tags			Casts
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	baseerrorresponse.Response
//	@Failure		500	{object}	baseerrorresponse.Response
//
//	@Router			/v0/cast/show-all [get]
func (ch *CastHandler) ShowAllCast(c *gin.Context) {
	all_cast, err := ch.castSrv.ListAllCasts()
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, baseerrorresponse.New(all_cast))
}

// ShowCastById godoc
//
//	@Summary		Get cast by ID
//	@Description	Retrieve a cast by its ID
//	@Tags			Casts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Cast ID"
//	@Success		200	{object}	baseerrorresponse.Response
//	@Failure		400	{object}	baseerrorresponse.Response
//	@Failure		404	{object}	baseerrorresponse.Response
//	@Failure		500	{object}	baseerrorresponse.Response
//
//	@Router			/v0/cast/show/{id} [get]
func (ch *CastHandler) ShowCastById(c *gin.Context) {
	temp := param.GetCastByIdRequest{}
	err := c.ShouldBindUri(&temp)
	if err != nil {
		log.Printf("bind error : %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseerrorresponse.NewError(err))
		return
	}
	i := temp.ID
	cast, err := ch.castSrv.GetCastById(i)
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, baseerrorresponse.New(cast.Cast))
}

// AddCast godoc
//
//	@Summary		Create a new cast
//	@Description	Add a new cast
//	@Tags			Casts
//	@Accept			json
//	@Produce		json
//	@Param			cast	body		param.CreateCastRequest	true	"Cast Data"
//	@Success		201		{object}	baseerrorresponse.Response	
// @Failure		400		{object}	baseerrorresponse.Response
//	@Failure		500		{object}	baseerrorresponse.Response
//
//	@Router			/v0/cast/add-cast [post]
func (ch *CastHandler) AddCast(c *gin.Context) {
	var temp param.CreateCastRequest
	err := c.ShouldBindJSON(&temp)
	if err != nil {
		log.Printf("bind error : %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseerrorresponse.NewError(err))
		return
	}
	cast, err := ch.castSrv.Create(temp)
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}
	// cast.Movies = []int{} // CONSIDER !!!!!!
	c.IndentedJSON(http.StatusCreated, baseerrorresponse.New(*cast))
}

// RemoveCast godoc
//
//	@Summary		Delete a cast
//	@Description	Remove a cast by ID
//	@Tags			Casts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Cast ID"
//	@Success		200	{object}	baseerrorresponse.Response	
// @Failure		400	{object}	baseerrorresponse.Response
//	@Failure		404	{object}	baseerrorresponse.Response
//	@Failure		500	{object}	baseerrorresponse.Response
//
//	@Router			/v0/cast/rem-cast [delete]
func (ch *CastHandler) RemoveCast(c *gin.Context) {
	temp := param.GetCastByIdRequest{}
	err := c.ShouldBindJSON(&temp)
	if err != nil {
		log.Printf("bind error : %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseerrorresponse.NewError(err))
		return
	}
	cast, err := ch.castSrv.Delete(temp.ID)
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, baseerrorresponse.New(cast.Cast))

}

// ShowCastWithFilter godoc
//
//	@Summary		Filter casts
//	@Description	Filter casts by name
//	@Tags			Casts
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string	false	"Cast name prefix"
//	@Success		200		{object}	baseerrorresponse.Response
//	@Failure		400		{object}	baseerrorresponse.Response
//	@Failure		500		{object}	baseerrorresponse.Response
//
//	@Router			/v0/cast/show [get]
func (ch *CastHandler) ShowCastWithFilter(c *gin.Context) {
	temp := param.GetCastWithFilterRequest{}
	err := c.ShouldBindQuery(&temp)
	if err != nil {
		log.Printf("bind error : %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseerrorresponse.NewError(err))
		return
	}

	all_cast, err := ch.castSrv.ListAllCasts()
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}

	_, exist_name := c.GetQuery("name")
	res := make([]entity.Cast, 0)
	for i := range all_cast {
		if exist_name && !strings.HasPrefix(all_cast[i].Name, temp.Name) {
			continue
		}

		res = append(res, all_cast[i])
	}
	c.IndentedJSON(http.StatusOK, baseerrorresponse.New(res))

}
