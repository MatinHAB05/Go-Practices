package handler

import (
	"fmt"
	"log"
	filtermovie "movieapp/api/helper/filter/movie"
	"movieapp/api/param"
	baseerrorresponse "movieapp/common/baseresponse"
	"movieapp/common/serviceerror"
	"movieapp/entity"
	"movieapp/service/movieservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	movieSrv *movieservice.Service
}

func (mh *MovieHandler) New(s *movieservice.Service) {
	mh.movieSrv = s
}

// ShowAllMovie godoc
//
//	@Summary		Get all movies
//	@Description	Get list of all movies
//	@Tags			Movies
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	baseerrorresponse.Response
//	@Failure		500	{object}	baseerrorresponse.Response
//
//	@Router			/v0/movie/show-all [get]
func (mh *MovieHandler) ShowAllMovie(c *gin.Context) {
	all_movies, err := mh.movieSrv.ListAllMovies()
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, baseerrorresponse.New(all_movies))
}

// ShowMovieById godoc
//
//	@Summary		Get movie by ID
//	@Description	Retrieve a single movie using its ID from URI
//	@Tags			Movies
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Movie ID"
//	@Success		200	{object}	baseerrorresponse.Response
//	@Failure		400	{object}	baseerrorresponse.Response
//	@Failure		404	{object}	baseerrorresponse.Response
//	@Failure		500	{object}	baseerrorresponse.Response
//
//	@Router			/v0/movie/show/{id} [get]
func (mh *MovieHandler) ShowMovieById(c *gin.Context) {
	temp := param.GetMovieByIdRequest{}
	err := c.ShouldBindUri(&temp)
	if err != nil {
		log.Printf("bind error : %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseerrorresponse.NewError(err))
		return
	}
	i := temp.ID
	m_res, err := mh.movieSrv.GetMovieById(i)
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, baseerrorresponse.New(*m_res.Movie))
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
//	@Param			release-action	query		string	false	"Release filter action (gt, lt, eq)"
//	@Success		200				{object}	baseerrorresponse.Response
//	@Failure		400				{object}	baseerrorresponse.Response
//	@Failure		500				{object}	baseerrorresponse.Response
//
//	@Router			/v0/movie/show [get]
func (mh *MovieHandler) ShowMovieWithFilter(c *gin.Context) {
	temp := param.GetMovieWithFilterRequest{}
	err := c.ShouldBindQuery(&temp)
	if err != nil {
		log.Printf("bind error : %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseerrorresponse.NewError(err))
		return
	}

	_, exist_t := c.GetQuery("title")
	_, exist_r := c.GetQuery("release")
	_, exist_ac := c.GetQuery("release-action")
	_, exist_q := c.GetQuery("quality")

	if exist_ac != exist_r {
		if !exist_ac {
			log.Printf("business error : %v\n", fmt.Errorf("for filtering by release year must enter both [action] and [release]! not found action"))
			c.AbortWithStatusJSON(http.StatusBadRequest, baseerrorresponse.NewSErrorf("not action found"))
		} else if !exist_r {
			log.Printf("business error : %v\n", fmt.Errorf("for filtering by release year must enter both [action] and [release]! not found release"))
			c.AbortWithStatusJSON(http.StatusBadRequest, baseerrorresponse.NewSErrorf("not release found"))
		}
		return
	}
	res, err := mh.movieSrv.ListAllMovies()
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}

	final := make([]entity.Movie, 0)

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
	c.IndentedJSON(http.StatusOK, baseerrorresponse.New(final))

}

// AddMovie godoc
//
//	@Summary		Create a new movie
//	@Description	Add a new movie to the system
//	@Tags			Movies
//	@Accept			json
//	@Produce		json
//	@Param			movie	body		param.CreateMovieRequest	true	"Movie data"
//	@Success		201		{object}	baseerrorresponse.Response
//	@Failure		400		{object}	baseerrorresponse.Response
//	@Failure		500		{object}	baseerrorresponse.Response
//
//	@Router			/v0/movie/add-movie [post]
func (mh *MovieHandler) AddMovie(c *gin.Context) {
	var temp param.CreateMovieRequest
	err := c.ShouldBindJSON(&temp)
	if err != nil {
		log.Printf("bind error : %v\n", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseerrorresponse.NewError(err))
		return
	}

	m, err := mh.movieSrv.Create(temp)
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}
	// m.Casts = []int{} // CONSIDER !!!!!!
	c.IndentedJSON(http.StatusCreated, baseerrorresponse.New(*m))

}

// RemoveMovie godoc
//
//	@Summary		Delete a movie
//	@Description	Remove a movie by its ID
//	@Tags			Movies
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Movie ID"
//	@Success		200	{object}	baseerrorresponse.Response
//	@Failure		400	{object}	baseerrorresponse.Response
//	@Failure		404	{object}	baseerrorresponse.Response
//	@Failure		500	{object}	baseerrorresponse.Response
//
//	@Router			/v0/movie/rem-movie [delete]
func (mh *MovieHandler) RemoveMovie(c *gin.Context) {
	temp := param.GetMovieByIdRequest{}
	err := c.ShouldBindJSON(&temp)
	if err != nil {
		log.Printf("bind error : %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, baseerrorresponse.NewError(err))
		return
	}
	m, err := mh.movieSrv.Delete(temp.ID)
	if err != nil {
		serviceerror.RespondError(c, err)
		return
	}

	c.IndentedJSON(http.StatusOK, baseerrorresponse.New(*m.Movie))

}
