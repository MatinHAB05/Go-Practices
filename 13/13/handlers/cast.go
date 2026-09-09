package handlers

import (
	"movieapp/db"
	"movieapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CastHandler struct{}

// ShowAllCast godoc
//
//	@Summary		Get all cast members
//	@Description	Retrieve a list of all actors/cast members
//	@Tags			cast
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	models.Actor
//	@Router			/v0/cast/show-all [get]
func (ch *CastHandler) ShowAllCast(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, db.ShowAllActor())
}

// ShowCastById godoc
//
//	@Summary		Get cast member by ID
//	@Description	Retrieve a single actor/cast member using its ID
//	@Tags			cast
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Cast ID"
//	@Success		200	{object}	models.Actor
//	@Failure		400	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Router			/v0/cast/show/{id} [get]
func (ch *CastHandler) ShowCastById(c *gin.Context) {
	var temp struct {
		ID int `uri:"id" binding:"required"`
	}
	err := c.ShouldBindUri(&temp)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	i := temp.ID
	a, ind := db.GetActorById(i, db.ShowAllActor())
	if ind < 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Not Found Actor"})
		return
	}
	c.IndentedJSON(http.StatusOK, a)
}

// AddCast godoc
//
//	@Summary		Add a new cast member
//	@Description	Create a new actor/cast member
//	@Tags			cast
//	@Accept			json
//	@Produce		json
//	@Param			cast	body		models.Actor	true	"Cast data"
//	@Success		200		{object}	models.Actor
//	@Failure		400		{object}	string
//	@Router			/v0/cast/add-cast [post]
func (ch *CastHandler) AddCast(c *gin.Context) {
	var temp models.Actor
	err := c.ShouldBindJSON(&temp)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	a := db.AddActor(temp.Name)
	a.Movies = []int{} // CONSIDER !!!!!!
	c.IndentedJSON(http.StatusOK, *a)
}

// RemoveCast godoc
//
//	@Summary		Remove a cast member
//	@Description	Delete a cast member using its ID
//	@Tags			cast
//	@Accept			json
//	@Produce		json
//	@Param			cast	body		object	true	"Cast ID"
//	@Success		200		{object}	models.Actor
//	@Failure		400		{object}	string
//	@Router			/v0/cast/rem-cast [delete]
func (ch *CastHandler) RemoveCast(c *gin.Context) {
	var temp struct {
		ID int `json:"id"`
	}
	err := c.ShouldBindJSON(&temp)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	a := db.RemoveActor(temp.ID)
	c.IndentedJSON(http.StatusOK, *a)

}

// ShowCastWithFilter godoc
//
//	@Summary		Filter cast members
//	@Description	Get cast members filtered by name
//	@Tags			cast
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string	false	"Actor name"
//	@Success		200		{array}		models.Actor
//	@Failure		400		{object}	map[string]string
//	@Router			/v0/cast/show [get]
func (ch *CastHandler) ShowCastWithFilter(c *gin.Context) {
	// TODO : VALIDATION....
	var temp struct {
		Name string `form:"name" binding:"omitempty,actorName"`
	}
	err := c.ShouldBindQuery(&temp)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := db.ShowAllActor()
	_, exist := c.GetQuery("name")
	if exist {
		res = db.FilterActorsByName(temp.Name, res)
	}
	c.IndentedJSON(http.StatusOK, res)

}
