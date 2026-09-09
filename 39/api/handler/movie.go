package handler

import (
	"fmt"
	"log"
	"movieapp/api/helper/filter/filtermovie"
	"movieapp/api/param/bindparam"
	"movieapp/api/param/serviceparam"
	"movieapp/common/baseresponse"
	"movieapp/common/serviceerror"
	"movieapp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	movieSrv      service.Movie
	moviecastSrv  service.MovieCast
	moviequerySrv service.MovieQuery
}

func (mh *MovieHandler) New(movieSrv service.Movie, moviecastSrv service.MovieCast, moviequerySrv service.MovieQuery) {
	mh.movieSrv = movieSrv
	mh.moviecastSrv = moviecastSrv
	mh.moviequerySrv = moviequerySrv
}

// ShowAllMovie godoc
//
//	@Summary		Get all movies
//	@Description	Get list of all movies
//	@Tags			Movies
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	baseresponse.Response
//	@Failure		500	{object}	baseresponse.Response
//
//	@Router			/movie/show-all [get]
func (mh *MovieHandler) ShowAllMovie(c *gin.Context) {

	log.Println(c.Request.URL.Port())
	log.Println(c.Request.Host)
	log.Println(c.Request.URL.Host)

	all_movies, err := mh.moviequerySrv.GetAllMoviesWithCast()
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, baseresponse.New(all_movies))
}

// ShowMovieById godoc
//
//	@Summary		Get movie by ID
//	@Description	Retrieve a single movie using its ID from URI
//	@Tags			Movies
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Movie ID"
//	@Success		200	{object}	baseresponse.Response
//	@Failure		400	{object}	baseresponse.Response
//	@Failure		404	{object}	baseresponse.Response
//	@Failure		500	{object}	baseresponse.Response
//
//	@Router			/movie/show/{id} [get]
func (mh *MovieHandler) ShowMovieById(c *gin.Context) {
	temp := bindparam.GetMovieByIdRequest{}
	err := c.ShouldBindUri(&temp)
	if err != nil {
		log.Printf("bind error : %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseresponse.NewError(err))
		return
	}
	i := temp.ID
	m_res, err := mh.moviequerySrv.GetMovieWithCast(i)
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, baseresponse.New(*m_res))
}

// ShowMovieWithFilter godoc
//
//	@Summary		Filter movies
//	@Description	Get movies with optional filters
//	@Tags			Movies
//	@Accept			json
//	@Produce		json
//	@Param			title			query		string	false	"Movie title"
//	@Param			quality			query		string	false	"Movie quality"
//	@Param			release			query		int		false	"Release year"
//	@Param			release-action	query		string	false	"Release filter action (gr, ls, eq, grq, lse, neq)"
//	@Success		200				{object}	baseresponse.Response
//	@Failure		400				{object}	baseresponse.Response
//	@Failure		500				{object}	baseresponse.Response
//
//	@Router			/movie/show [get]
func (mh *MovieHandler) ShowMovieWithFilter(c *gin.Context) {
	temp := bindparam.GetMovieWithFilterRequest{}
	err := c.ShouldBindQuery(&temp)
	if err != nil {
		log.Printf("bind error : %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseresponse.NewError(err))
		return
	}

	_, exist_t := c.GetQuery("title")
	_, exist_r := c.GetQuery("release")
	_, exist_ac := c.GetQuery("release-action")
	_, exist_q := c.GetQuery("quality")

	if exist_ac != exist_r {
		if !exist_ac {
			log.Printf("business error : %v\n", fmt.Errorf("for filtering by release year must enter both [action] and [release]! not found action"))
			c.AbortWithStatusJSON(http.StatusBadRequest, baseresponse.NewSErrorf("not action found"))
		} else if !exist_r {
			log.Printf("business error : %v\n", fmt.Errorf("for filtering by release year must enter both [action] and [release]! not found release"))
			c.AbortWithStatusJSON(http.StatusBadRequest, baseresponse.NewSErrorf("not release found"))
		}
		return
	}
	res, err := mh.moviequerySrv.GetAllMoviesWithCast()
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}

	final := make([]serviceparam.GetMovieWithCastsByIdResponse, 0)

	for i := range res {
		if exist_t && !filtermovie.ByTitle(res[i].Title, temp.Title) {
			continue
		}

		if exist_q && !filtermovie.ByQuality(res[i].Quality.ToString(), temp.Quality) {
			continue
		}

		if exist_r && !filtermovie.ByDate(res[i].ReleaseYear, temp.ReleaseYear, temp.ReleaseAction) {
			continue
		}
		final = append(final, res[i])
	}
	c.IndentedJSON(http.StatusOK, baseresponse.New(final))

}

// AddMovie godoc
//
//	@Summary		Create a new movie
//	@Description	Add a new movie to the system
//	@Tags			Movies
//	@Accept			json
//	@Produce		json
//	@Param			movie	body		bindparam.CreateMovieRequest	true	"Movie data"
//	@Success		201		{object}	baseresponse.Response
//	@Failure		400		{object}	baseresponse.Response
//	@Failure		500		{object}	baseresponse.Response
//
//	@Router			/movie/add-movie [post]
func (mh *MovieHandler) AddMovie(c *gin.Context) {
	var temp bindparam.CreateMovieRequest
	err := c.ShouldBindJSON(&temp)
	if err != nil {
		log.Printf("bind error : %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseresponse.NewError(err))
		return
	}

	m, err := mh.movieSrv.Create(serviceparam.CreateMovieInput{
		Title:       temp.Title,
		ReleaseYear: temp.ReleaseYear,
		Quality:     temp.Quality,
	})
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, baseresponse.New(*m))

}

// RemoveMovie godoc
//
//	@Summary		Delete a movie
//	@Description	Remove a movie by its ID
//	@Tags			Movies
//	@Accept			json
//	@Produce		json
//	@Param			id	body		bindparam.RemoveMovieByIdRequest	true	"Movie ID"
//	@Success		200	{object}	baseresponse.Response
//	@Failure		400	{object}	baseresponse.Response
//	@Failure		404	{object}	baseresponse.Response
//	@Failure		500	{object}	baseresponse.Response
//	@Router			/movie/rem-movie [delete]
func (mh *MovieHandler) RemoveMovie(c *gin.Context) {
	temp := bindparam.RemoveMovieByIdRequest{}
	err := c.ShouldBindJSON(&temp)
	if err != nil {
		log.Printf("bind error : %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseresponse.NewError(err))
		return
	}
	m, err := mh.movieSrv.Delete(temp.ID)
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, baseresponse.New(*m.Movie))

}
