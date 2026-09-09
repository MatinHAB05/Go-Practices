package handlers

import (
	"movieapp/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieCastLinking_Handlers struct{}

// Movie_Cast_Linking godoc
//
//	@Summary		Link actor to movie
//	@Description	Create a relation between an actor and a movie using their IDs
//	@Tags			linking
//	@Accept			json
//	@Produce		json
//	@Param			link	body		object					true	"Actor-Movie linking payload (actor-id, movie-id)"
//	@Success		200		{object}	map[string]interface{}	"Movie and Actor if linked successfully"
//	@Failure		400		{string}	string					"Binding error"
//	@Router			/v0/movie-cast-link/ [post]
func (mc_l_h *MovieCastLinking_Handlers) Movie_Cast_Linking(c *gin.Context) {
	var temp struct {
		ActorId int `json:"actor-id" binding:"required"`
		MovieId int `json:"movie-id" binding:"required"`
	}
	err := c.ShouldBindJSON(&temp)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	m, ac := db.LinkActorToMovie(temp.ActorId, temp.MovieId)

	if m == nil && ac == nil {
		c.IndentedJSON(http.StatusOK, "Already Linked")
		return
	} else if ac == nil {
		c.IndentedJSON(http.StatusOK, "Not Found Actor")
		return
	} else if m == nil {
		c.IndentedJSON(http.StatusOK, "Not Found Movie")
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"Movie": m,
		"Actor": ac,
	})
}
