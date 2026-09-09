package handlers

import (
	"movieapp/db"
	"movieapp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct{}

// ShowAllMovie godoc
//
//	@Summary		Get all movies
//	@Description	Retrieve a list of all movies
//	@Tags			movies
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	models.Movie
//	@Router			/v0/movie/show-all [get]
func (mh *MovieHandler) ShowAllMovie(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, db.ShowAllMovie())
}

// ShowMovieById godoc
//
//	@Summary		Get movie by ID
//	@Description	Retrieve a single movie using its ID
//	@Tags			movies
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Movie ID"
//	@Success		200	{object}	models.Movie
//	@Failure		400	{object}	map[string]string
//	@Failure		404	{object}	map[string]string
//	@Router			/v0/movie/show/{id} [get]
func (mh *MovieHandler) ShowMovieById(c *gin.Context) {
	var temp struct {
		ID int `uri:"id" binding:"required"`
	}
	err := c.ShouldBindUri(&temp)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	i := temp.ID
	m, ind := db.GetMovieById(i, db.ShowAllMovie())
	if ind < 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Not Found Movie"})
		return
	}
	c.IndentedJSON(http.StatusOK, m)
}

// ShowMovieWithFilter godoc
//
//	@Summary		Filter movies
//	@Description	Retrieve movies using optional filters such as title, release year comparison, and quality
//	@Tags			movies
//	@Accept			json
//	@Produce		json
//	@Param			title			query		string	false	"Movie title"
//	@Param			release			query		int		false	"Release year"
//	@Param			release-action	query		string	false	"Release year comparator (e.g. >, <, =, <=, >=)"
//	@Param			quality			query		string	false	"Movie quality"
//	@Success		200				{array}		models.Movie
//	@Failure		400				{object}	map[string]string
//	@Router			/v0/movie/show [get]
func (mh *MovieHandler) ShowMovieWithFilter(c *gin.Context) {
	// TODO : VALIDATE.....
	var temp struct {
		Title         string `form:"title" binding:"omitempty,movieTitle"`
		ReleaseYear   int    `form:"release" binding:"omitempty,releaseYear"`
		ReleaseAction string `form:"release-action" binding:"omitempty,comparatorOperator"`
		Quality       string `form:"quality" binding:"omitempty,movieQuality"`
	}
	err := c.ShouldBindQuery(&temp)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, exist_t := c.GetQuery("title")
	_, exist_r := c.GetQuery("release")
	_, exist_ac := c.GetQuery("release-action")
	_, exist_q := c.GetQuery("quality")

	if exist_ac != exist_r {
		if !exist_ac {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": "For Filtering By Release Year Must Enter Both [action] and [release]! Not Found Action",
			})
		} else if !exist_r {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": "For Filtering By Release Year Must Enter Both [action] and [release]! Not Found release",
			})
		}
		return
	}
	res := db.ShowAllMovie()
	if exist_t {
		res = db.FilterMoviesByTitle(temp.Title, res)
	}
	if exist_q {
		res = db.FilterMoviesByQuality(temp.Quality, res)
	}
	if exist_r {
		res = db.FilterMoviesByDate(temp.ReleaseAction, temp.ReleaseYear, res)
	}
	c.IndentedJSON(http.StatusOK, res)

}

// AddMovie godoc
//
//	@Summary		Add a new movie
//	@Description	Create a new movie using the provided data
//	@Tags			movies
//	@Accept			json
//	@Produce		json
//	@Param			movie	body		models.Movie	true	"Movie data"
//	@Success		200		{object}	models.Movie
//	@Failure		400		{object}	string
//	@Router			/v0/movie/add-movie [post]
func (mh *MovieHandler) AddMovie(c *gin.Context) {
	var temp models.Movie
	err := c.ShouldBindJSON(&temp)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	m := db.AddMovie(temp.Title, strconv.Itoa(temp.ReleaseYear), string(temp.Quality))
	m.Actors = []int{} // CONSIDER !!!!!!
	c.IndentedJSON(http.StatusOK, *m)

}

// RemoveMovie godoc
//
//	@Summary		Remove a movie
//	@Description	Delete a movie using its ID
//	@Tags			movies
//	@Accept			json
//	@Produce		json
//	@Param			movie	body		object	true	"Movie ID"
//	@Success		200		{object}	models.Movie
//	@Failure		400		{object}	string
//	@Router			/movie/rem-movie [delete]
func (mh *MovieHandler) RemoveMovie(c *gin.Context) {
	var temp struct {
		ID int `json:"id"`
	}
	err := c.ShouldBindJSON(&temp)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	m := db.RemoveMovie(temp.ID)
	c.IndentedJSON(http.StatusOK, *m)

}
