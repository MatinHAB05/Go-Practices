package moviequeryservice

import (
	"fmt"
	"movieapp/api/param/serviceparam"
	"movieapp/entity"
	"movieapp/repository"
	"movieapp/service/errorservice"
)

type Service struct {
	movieRepo     repository.MovieRepository
	castRepo      repository.CastRepository
	moviecastRepo repository.MovieCastRepository
}

func (s *Service) New(moveiRepo repository.MovieRepository, castRepo repository.CastRepository, moviecastRepo repository.MovieCastRepository) {
	s.movieRepo = moveiRepo
	s.castRepo = castRepo
	s.moviecastRepo = moviecastRepo
}

func (s *Service) GetMovieWithCast(MovieId int) (*serviceparam.GetMovieWithCastsByIdResponse, error) {
	movie, err := s.movieRepo.GetMovieById(MovieId)
	if err != nil {
		return nil, fmt.Errorf("get movie by id-%d : %w", MovieId, err)
	} else if movie == nil {
		return nil, &errorservice.ErrMovieNotFound
	}

	casts_temp, err := s.moviecastRepo.GetAllCastsOfMovie(MovieId)
	if err != nil {
		return nil, fmt.Errorf("get all cast of movie id-%d : %w", MovieId, err)
	}
	casts := make([]entity.Cast, 0)

	for i := range casts_temp {
		var c *entity.Cast
		c, err := s.castRepo.GetCastById(casts_temp[i].CastId)
		if err != nil {
			return nil, fmt.Errorf("get cast by id-%d : %w\n", casts_temp[i].CastId, err)
		} else if c == nil {
			return nil, &errorservice.ErrCastNotFound
		}
		casts = append(casts, *c)
	}

	return &serviceparam.GetMovieWithCastsByIdResponse{
		Id:          movie.Id,
		Title:       movie.Title,
		ReleaseYear: movie.ReleaseYear,
		Quality:     movie.Quality,
		Casts:       casts,
	}, nil
}

func (s *Service) GetAllMoviesWithCast() ([]serviceparam.GetMovieWithCastsByIdResponse, error) {
	all_movies, err := s.movieRepo.GetAllMovies()
	if err != nil {
		return nil, fmt.Errorf("get all movies : %w", err)
	}
	final := make([]serviceparam.GetMovieWithCastsByIdResponse, 0)
	for i := range all_movies {
		var movie *entity.Movie = &all_movies[i]
		var MovieId int = movie.Id

		obj, err := s.GetMovieWithCast(MovieId)
		if err != nil {
			return nil, fmt.Errorf("get movie with id-%d with its casts : %w", MovieId, err)
		}

		final = append(final, *obj)
	}
	return final, nil
}
