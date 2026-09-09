package movieservice

import (
	"fmt"
	"movieapp/api/param"
	"movieapp/entity"
)

type Repository interface {
	GetAllMovies() ([]entity.Movie, error)
	GetMovieById(int) (*entity.Movie, error)
	Create(entity.Movie) (*entity.Movie, error)
	Delete(int) (*entity.Movie, error)
	IsExist(int) (*bool, error)
}
type Service struct {
	movieRepo Repository
}

func (s *Service) New(movieRepo Repository) {
	s.movieRepo = movieRepo
}

func (s *Service) ListAllMovies() ([]entity.Movie, error) {
	all_m, err := s.movieRepo.GetAllMovies()
	if err != nil {
		return nil, fmt.Errorf("get all movies : %w\n", err)
	}
	return all_m, nil
}

func (s *Service) GetMovieById(index int) (param.GetMovieByIdResponse, error) {
	movie, err := s.movieRepo.GetMovieById(index)
	if err != nil {
		return param.GetMovieByIdResponse{}, fmt.Errorf("get movie by id : %w\n", err)
	} else if movie == nil {
		return param.GetMovieByIdResponse{}, &ErrNotFoundMovie
	}
	return param.GetMovieByIdResponse{Movie: movie}, nil
}

func (s *Service) Create(cmr param.CreateMovieRequest) (*entity.Movie, error) {
	new_m := entity.Movie{
		Id:          -1,
		Title:       cmr.Title,
		ReleaseYear: cmr.ReleaseYear,
		Quality:     cmr.Quality,
		Casts:       []int{},
	}
	m, err := s.movieRepo.Create(new_m)
	if err != nil {
		return &entity.Movie{}, fmt.Errorf("create movie : %w\n", err)
	}
	return m, nil
}

func (s *Service) Delete(id int) (param.DeleteMovieResponse, error) {
	ok, err := s.movieRepo.IsExist(id)
	if err != nil {
		return param.DeleteMovieResponse{}, fmt.Errorf("is exist movie : %w\n", err)
	}
	if !*ok {
		return param.DeleteMovieResponse{}, &ErrNotFoundMovie
	}
	m, err := s.movieRepo.Delete(id)
	if err != nil {
		return param.DeleteMovieResponse{}, fmt.Errorf("delete movie : %w\n", err)
	}
	return param.DeleteMovieResponse{Movie: m}, nil
}

func (s *Service) IsExist(id int) (*bool, error) {
	ok, err := s.movieRepo.IsExist(id)
	if err != nil {
		return nil, fmt.Errorf("is exist movie : %w\n", err)
	}
	return ok, nil
}
