package castqueryservice

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

func (s *Service) GetCastWithMovie(CastId int) (*serviceparam.GetCastWithMoviesByIdResponse, error) {
	cast, err := s.castRepo.GetCastById(CastId)
	if err != nil {
		return nil, fmt.Errorf("get cast by id-%d : %w", CastId, err)
	} else if cast == nil {
		return nil, &errorservice.ErrCastNotFound
	}

	movies_temp, err := s.moviecastRepo.GetAllMoviesOfCast(CastId)
	if err != nil {
		return nil, fmt.Errorf("get all movie of cast id-%d : %w", CastId, err)
	}
	movies := make([]entity.Movie, 0)
	for i := range movies_temp {
		var m *entity.Movie
		m, err := s.movieRepo.GetMovieById(movies_temp[i].MovieId)
		if err != nil {
			return nil, fmt.Errorf("get movie by id-%d : %w\n", movies_temp[i].MovieId, err)
		} else if m == nil {

			return nil, &errorservice.ErrMovieNotFound
		}
		movies = append(movies, *m)
	}

	return &serviceparam.GetCastWithMoviesByIdResponse{
		Name:   cast.Name,
		Id:     CastId,
		Movies: movies,
	}, nil
}

func (s *Service) GetAllCastWithMovie() ([]serviceparam.GetCastWithMoviesByIdResponse, error) {
	all_casts, err := s.castRepo.GetAllCasts()
	if err != nil {
		return nil, fmt.Errorf("get all casts : %w", err)
	}
	final := make([]serviceparam.GetCastWithMoviesByIdResponse, 0)
	for i := range all_casts {
		var cast *entity.Cast = &all_casts[i]
		var CastId int = cast.Id

		obj, err := s.GetCastWithMovie(CastId)
		if err != nil {
			return nil, fmt.Errorf("get cast with id-%d with its movies : %w", CastId, err)
		}

		final = append(final, *obj)
	}
	return final, nil
}
