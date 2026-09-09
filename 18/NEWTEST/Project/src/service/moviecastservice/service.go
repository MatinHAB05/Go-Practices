package moviecastservice

import (
	"fmt"
	"movieapp/api/param"
	"movieapp/service/castservice"
	"movieapp/service/movieservice"
)

type Repository interface {
	IsExist(int, int) (*bool, error)
	Create(int, int) (*int, *int, error)
}

type Service struct {
	moviecastRepo Repository
	MovieSrv      *movieservice.Service
	CastSrv       *castservice.Service
}

func (s *Service) New(movieService *movieservice.Service, castService *castservice.Service, moviecastRepo Repository) {
	s.moviecastRepo = moviecastRepo
	s.MovieSrv = movieService
	s.CastSrv = castService
}

func (s *Service) IsExist(MovieId, CastId int) (*bool, error) {
	ok, err := s.IsExist(MovieId, CastId)
	if err != nil {
		return nil, fmt.Errorf("is movie-cast exist : %w\n", err)
	}
	return ok, nil
}

func (s *Service) Create(link param.LinkMovie2CastRequest) (*int, *int, error) {
	ok, err := s.MovieSrv.IsExist(link.MovieId)
	if err != nil {
		return nil, nil, fmt.Errorf("is movie exist service : %w\n", err)
	} else if !*ok {
		return nil, nil, &movieservice.ErrNotFoundMovie
	}
	ok, err = s.CastSrv.IsExist(link.CastId)
	if err != nil {
		return nil, nil, fmt.Errorf("is cast exist service : %w\n", err)
	} else if !*ok {
		return nil, nil, &castservice.ErrCastNotFound
	}
	ok, err = s.IsExist(link.MovieId, link.CastId)
	if err != nil {
		return nil, nil, fmt.Errorf("is movie-cast exist : %w\n", err)
	} else if *ok {
		return nil, nil, &ErrAlreadyLinkedMovieCast
	}
	m, c, err := s.moviecastRepo.Create(link.MovieId, link.CastId)
	if err != nil {
		return nil, nil, fmt.Errorf("create movie-cast: %w\n", err)
	}
	return m, c, nil
}
