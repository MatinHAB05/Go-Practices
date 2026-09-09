package movieservice

import (
	"fmt"
	"movieapp/api/param/serviceparam"
	"movieapp/entity"
	"movieapp/repository"
	"movieapp/service/errorservice"
)

type Service struct {
	movieRepo repository.MovieRepository
}

func (s *Service) New(movieRepo repository.MovieRepository) {
	s.movieRepo = movieRepo
}

func (s *Service) ListAllMovies() ([]entity.Movie, error) {
	all_m, err := s.movieRepo.GetAllMovies()
	if err != nil {
		return nil, fmt.Errorf("get all movies : %w\n", err)
	}
	return all_m, nil
}

func (s *Service) GetMovieById(index int) (*entity.Movie, error) {
	movie, err := s.movieRepo.GetMovieById(index)
	if err != nil {
		return nil, fmt.Errorf("get movie by id-%d : %w\n", index, err)
	} else if movie == nil {
		return nil, &errorservice.ErrMovieNotFound
	}
	return movie, nil
}

func (s *Service) Create(cmr serviceparam.CreateMovieInput) (*entity.Movie, error) {
	new_m := entity.Movie{
		Id:          -1,
		Title:       cmr.Title,
		ReleaseYear: cmr.ReleaseYear,
		Quality:     cmr.Quality,
	}
	m, err := s.movieRepo.Create(new_m)
	if err != nil {
		return &entity.Movie{}, fmt.Errorf("create movie : %w\n", err)
	}
	return m, nil
}

func (s *Service) Delete(id int) (*serviceparam.DeleteMovieResponse, error) {
	ok, err := s.movieRepo.IsExist(id)
	if err != nil {
		return nil, fmt.Errorf("is exist movie id-%d : %w\n", id, err)
	} else if !*ok {
		return nil, &errorservice.ErrMovieNotFound
	}

	m, err := s.movieRepo.Delete(id)
	if err != nil {
		return nil, fmt.Errorf("delete movie id-%d : %w\n", id, err)
	}
	return &serviceparam.DeleteMovieResponse{Movie: m}, nil
}

func (s *Service) IsExist(id int) (*bool, error) {
	ok, err := s.movieRepo.IsExist(id)
	if err != nil {
		return nil, fmt.Errorf("is exist movie id-%d : %w\n", id, err)
	}
	return ok, nil
}
